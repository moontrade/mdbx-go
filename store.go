package mdbx

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

const (
	DataFileName = "mdbx.dat"
	LockFileName = "mdbx.lck"
)

var (
	storeMu sync.Mutex
	stores  = make(map[*Store]struct{})
)

type Store struct {
	id      int
	env     *Env
	closed  int64
	writeMu sync.Mutex
	mu      sync.Mutex
}

func (s *Store) Env() *Env {
	return s.env
}

func Open(
	path string,
	flags EnvFlags,
	mode os.FileMode,
	initEnv func(env *Env, create bool) error,
	init func(store *Store, create bool) error,
) (*Store, error) {
	store := &Store{}

	env, e := NewEnv()
	if e != ErrSuccess {
		return nil, e
	}
	store.env = env

	var err error
	var stat os.FileInfo
	stat, err = os.Stat(filepath.Join(path, DataFileName))
	var create bool
	if err != nil {
		create = true
	} else {
		// Init
		create = false
	}
	_ = stat

	if initEnv != nil {
		if err = initEnv(env, create); err != nil && err != ErrSuccess {
			store.Close()
			return nil, err
		}
	}

	if mode == 0 {
		mode = 0664
	}

	if err = store.env.Open(path, flags, mode); err != ErrSuccess {
		_ = store.Close()
		return nil, err
	}

	if init != nil {
		if err = init(store, create); err != nil && err != ErrSuccess {
			_ = store.Close()
			return nil, err
		}
	}

	if flags&EnvSafeNoSync != 0 || flags&EnvNoMetaSync != 0 {

	}
	storeMu.Lock()
	stores[store] = struct{}{}
	storeMu.Unlock()
	return store, nil
}

func (s *Store) IsClosed() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.closed > 0
}

func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.writeMu.Lock()
	defer s.writeMu.Unlock()

	if s.closed > 0 {
		return os.ErrClosed
	}
	s.closed = time.Now().UnixNano()
	if s.env != nil {
		_ = s.env.Close(true)
		s.env = nil
	}
	storeMu.Lock()
	delete(stores, s)
	storeMu.Unlock()
	return nil
}

func (s *Store) Update(fn func(tx *Tx) error) error {
	return s.UpdateLock(true, fn)
}

func (s *Store) UpdateLock(lockThread bool, fn func(tx *Tx) error) (err error) {
	if lockThread {
		// Write transactions must be bound to a single thread.
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
	}

	// Get exclusive write lock.
	s.writeMu.Lock()
	defer s.writeMu.Unlock()
	tx := Tx{}
	defer func() {
		// Abort if panic
		if !tx.IsCommitted() && !tx.IsAborted() && tx.txn != nil {
			if err = tx.Abort(); err != ErrSuccess {
				// Ignore
			}
		}

		// Propagate
		r := recover()
		if r != nil {
			if e, ok := r.(error); ok {
				err = e
			}
		}
	}()

	if err = s.env.Begin(&tx, TxReadWrite); err != ErrSuccess {
		return err
	}
	if err = fn(&tx); err != nil && err != ErrSuccess {
		// Abort if necessary
		if !tx.IsAborted() && !tx.IsCommitted() {
			if err = tx.Abort(); err != ErrSuccess {
				return err
			}
		}
		return err
	} else {
		if !tx.IsCommitted() {
			if err = tx.Commit(); err != ErrSuccess {
				return err
			}
		}
		return err
	}
}

func (s *Store) View(fn func(tx *Tx)) (err error) {
	tx := Tx{}
	defer func() {
		if !tx.IsAborted() {
			err = tx.Abort()
		}
		// recovery
		r := recover()
		if r != nil {
			if e, ok := r.(error); ok {
				err = e
			}
		}
	}()
	if err = s.env.Begin(&tx, TxReadOnly); err != ErrSuccess {
		return err
	}
	fn(&tx)
	return err
}

func (s *Store) ViewRenew(tx *Tx, fn func(tx *Tx)) (err error) {
	if tx == nil {
		return s.View(fn)
	}
	defer func() {
		if !tx.IsReset() {
			err = tx.Reset()
		}
		// recovery
		r := recover()
		if r != nil {
			if e, ok := r.(error); ok {
				err = e
			}
		}
	}()
	if err = tx.Renew(); err != ErrSuccess {
		return err
	}
	fn(tx)
	return err
}