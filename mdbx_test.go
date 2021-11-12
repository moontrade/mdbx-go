package mdbx

import (
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"testing"
	"unsafe"
)

func TestChk(t *testing.T) {
	_, out, err := Chk("-v", "-w", "testdata/1000000/mdbx.dat")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(out))
}

func TestEnv_Open(t *testing.T) {
	env, err := NewEnv()
	if err != ErrSuccess {
		t.Fatal(err)
	}
	if err = env.SetGeometry(Geometry{
		SizeLower:       1024 * 1024 * 16,
		SizeNow:         0,
		SizeUpper:       1024 * 1024 * 1024 * 16,
		GrowthStep:      1024 * 1024 * 16,
		ShrinkThreshold: 0,
		PageSize:        4096,
	}); err != ErrSuccess {
		t.Fatal(err)
	}
	if err = env.SetMaxDBS(4); err != ErrSuccess {
		t.Fatal(err)
	}
	err = env.Open(
		"testdata/74419951",
		EnvNoTLS|EnvNoReadAhead|EnvCoalesce|EnvLIFOReclaim|EnvSafeNoSync,
		0664,
	)
	if err != ErrSuccess {
		t.Fatal(err)
	}

	var txn Tx
	if err = env.Begin(&txn, TxReadWrite); err != ErrSuccess {
		t.Fatal(err)
	}

	var dbi DBI
	if dbi, err = txn.OpenDBI("", DBCreate); err != ErrSuccess {
		t.Fatal(err)
	}

	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, uint64(101))
	value := []byte("hello")

	keyVal := Bytes(&key)
	valueVal := Bytes(&value)

	if err = txn.Put(dbi, &keyVal, &valueVal, 0); err != ErrSuccess {
		t.Fatal(err)
	}

	var latency CommitLatency
	txn.CommitEx(&latency)

	err = env.Close(false)
	if err != ErrSuccess {
		t.Fatal(err)
	}
}

type Engine struct {
	env    *Env
	rootDB DBI
	write  Tx
	rd     Tx
}

func (engine *Engine) BeginWrite() (*Tx, Error) {
	engine.write.txn = nil
	engine.write.env = engine.env
	return &engine.write, engine.env.Begin(&engine.write, TxReadWrite)
}

func (engine *Engine) BeginRead() (*Tx, Error) {
	engine.rd.env = engine.env
	return &engine.rd, engine.rd.Renew()
}

func initDB(path string) (*Engine, Error) {
	engine := &Engine{}
	env, err := NewEnv()
	if err != ErrSuccess {
		return nil, err
	}
	engine.env = env
	if err = env.SetGeometry(Geometry{
		SizeLower:       1024 * 1024 * 16,
		SizeNow:         1024 * 1024 * 16,
		SizeUpper:       1024 * 1024 * 1024 * 16,
		GrowthStep:      1024 * 1024 * 16,
		ShrinkThreshold: 0,
		PageSize:        16384,
	}); err != ErrSuccess {
		return nil, err
	}
	if err = env.SetMaxDBS(1); err != ErrSuccess {
		return nil, err
	}

	os.Mkdir("testdata", 0755)
	err = env.Open(
		path,
		//EnvNoMemInit|EnvCoalesce|EnvLIFOReclaim|EnvSyncDurable,
		EnvNoMemInit|EnvCoalesce|EnvLIFOReclaim|EnvSafeNoSync|EnvWriteMap,
		0664,
	)
	if err != ErrSuccess {
		return nil, err
	}

	if err = env.Begin(&engine.write, TxReadWrite); err != ErrSuccess {
		return nil, err
	}

	if engine.rootDB, err = engine.write.OpenDBI("m", DBIntegerKey|DBCreate); err != ErrSuccess {
		return nil, err
	}

	if err = engine.write.Commit(); err != ErrSuccess {
		return nil, err
	}

	//if err = env.Begin(&engine.rd, TxReadOnly); err != ErrSuccess {
	//	return nil, err
	//}
	//if err = engine.rd.Reset(); err != ErrSuccess {
	//	return nil, err
	//}

	return engine, ErrSuccess
}

func BenchmarkTxn_Put(b *testing.B) {
	engine, err := initDB("./testdata/" + strconv.Itoa(b.N))
	if err != ErrSuccess {
		b.Fatal(err)
	}
	defer engine.env.Close(true)

	key := make([]byte, 8)
	data := []byte("hello")

	keyVal := Bytes(&key)
	dataVal := Bytes(&data)

	b.ResetTimer()
	b.ReportAllocs()

	txn, err := engine.BeginWrite()
	if err != ErrSuccess {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		//binary.BigEndian.PutUint64(key, uint64(20))
		//binary.LittleEndian.PutUint64(key, uint64(i))
		*(*uint64)(unsafe.Pointer(keyVal.Base)) = uint64(i)
		//keyVal = U64(uint64(i))
		if err = txn.Put(engine.rootDB, &keyVal, &dataVal, PutAppend); err != ErrSuccess {
			txn.Abort()
			b.Fatal(err)
		}
	}

	//var envInfo EnvInfo
	//if err = txn.EnvInfo(&envInfo); err != ErrSuccess {
	//	b.Fatal(err)
	//}
	//var info TxInfo
	//if err = txn.Info(&info); err != ErrSuccess {
	//	b.Fatal(err)
	//}
	if err = txn.Commit(); err != ErrSuccess {
		b.Fatal(err)
	}
	//engine.env.Sync(true, false)
	//engine.env.Sync(true, false)
}

func BenchmarkTxn_PutCursor(b *testing.B) {
	engine, err := initDB("./testdata/" + strconv.Itoa(b.N))
	if err != ErrSuccess {
		b.Fatal(err)
	}
	defer engine.env.Close(true)

	b.ReportAllocs()
	b.ResetTimer()
	key := uint64(0)
	data := []byte("hello")

	keyVal := U64(&key)
	dataVal := Bytes(&data)

	{
		insert := func(low, high uint64) {
			txn, err := engine.BeginWrite()
			if err != ErrSuccess {
				b.Fatal(err)
			}

			cursor, err := txn.OpenCursor(engine.rootDB)
			if err != ErrSuccess {
				b.Fatal(err)
			}

			for i := low; i < high; i++ {
				key = i
				if err = cursor.Put(&keyVal, &dataVal, PutAppend); err != ErrSuccess {
					cursor.Close()
					txn.Abort()
					b.Fatal(err)
				}
			}

			if err = cursor.Close(); err != ErrSuccess {
				b.Fatal(err)
			}
			if err = txn.Commit(); err != ErrSuccess {
				b.Fatal(err)
			}
		}

		const batchSize = 1000000
		for i := 0; i < b.N; i += batchSize {
			end := i + batchSize
			if end > b.N {
				end = b.N
			}
			insert(uint64(i), uint64(end))
		}
	}
}

func BenchmarkTxn_Get(b *testing.B) {
	engine, err := initDB("./testdata/" + strconv.Itoa(b.N))
	if err != ErrSuccess {
		b.Fatal(err)
	}
	defer engine.env.Close(true)

	key := uint64(0)
	data := []byte("hello")

	keyVal := U64(&key)
	dataVal := Bytes(&data)

	{
		insert := func(low, high uint64) {
			txn, err := engine.BeginWrite()
			if err != ErrSuccess {
				b.Fatal(err)
			}

			cursor, err := txn.OpenCursor(engine.rootDB)
			if err != ErrSuccess {
				b.Fatal(err)
			}

			for i := low; i < high; i++ {
				key = i
				if err = cursor.Put(&keyVal, &dataVal, PutAppend); err != ErrSuccess {
					cursor.Close()
					txn.Abort()
					b.Fatal(err)
				}
			}

			if err = cursor.Close(); err != ErrSuccess {
				b.Fatal(err)
			}
			if err = txn.Commit(); err != ErrSuccess {
				b.Fatal(err)
			}
		}

		const batchSize = 1000000
		for i := 0; i < b.N; i += batchSize {
			end := i + batchSize
			if end > b.N {
				end = b.N
			}
			insert(uint64(i), uint64(end))
		}
	}

	txn := &Tx{}

	//engine.env.Sync(true, false)
	//engine.env.Sync(true, false)

	if err = engine.env.Begin(txn, TxReadOnly); err != ErrSuccess {
		b.Fatal(err)
	}
	//txn, err = engine.BeginRead()
	//if err != ErrSuccess {
	//	b.Fatal(err)
	//}

	b.ResetTimer()
	b.ReportAllocs()

	cursor, err := txn.OpenCursor(engine.rootDB)
	if err != ErrSuccess {
		b.Fatal(err)
	}

	//binary.LittleEndian.PutUint64(key, uint64(b.N))

	//if err = txn.Get(engine.rootDB, &keyVal, &dataVal); err != ErrSuccess {
	//	b.Fatal(err)
	//}

	//keyInt := binary.LittleEndian.Uint64(key)

	//if err = cursor.Get(&keyVal, &dataVal, CursorSet); err != ErrSuccess {
	//	b.Fatal(err)
	//}

	dataVal = Val{}
	keyVal = Val{}

	//fmt.Println(dataVal.String())

	//binary.LittleEndian.PutUint64(key, 0)

	count := 0
	//
	for {
		if err = cursor.Get(&keyVal, &dataVal, CursorNextNoDup); err != ErrSuccess {
			break
		}
		//if keyVal.Base == nil {
		//	break
		//}
		count++
		//keyInt = binary.LittleEndian.Uint64(key)
		//_ = keyInt

		//keyVal = Val{}
		//dataVal = Val{}

		//if cursor.EOF() != 0 {
		//	break
		//}
	}

	//if count == 1000000 {
	//	println("1m")
	//}

	//for i := 0; i < b.N; i++ {
	//	*(*uint64)(unsafe.Pointer(&key[0])) = uint64(i)
	//	//binary.BigEndian.PutUint64(key, uint64(20))
	//	//binary.BigEndian.PutUint64(key[8:], uint64(i))
	//	//keyVal = U64(uint64(i))
	//	if err = txn.Get(engine.rootDB, &keyVal, &dataVal); err != ErrSuccess && err != ErrNotFound {
	//		txn.Reset()
	//		b.Fatal(err)
	//	}
	//}

	if err = cursor.Close(); err != ErrSuccess {
		b.Fatal(err)
	}
	if err = txn.Reset(); err != ErrSuccess {
		b.Fatal(err)
	}

	b.StopTimer()

	fmt.Println("count", count)

	//var envInfo EnvInfo
	//if err = txn.EnvInfo(&envInfo); err != ErrSuccess {
	//	b.Fatal(err)
	//}
	//var info TxInfo
	//if err = txn.Info(&info); err != ErrSuccess {
	//	b.Fatal(err)
	//}

	//engine.env.Sync(true, false)
	//engine.env.Sync(true, false)
}

func TestTxn_Cursor(b *testing.T) {
	iterations := 100
	engine, err := initDB("./testdata/" + strconv.Itoa(iterations))
	if err != ErrSuccess {
		b.Fatal(err)
	}

	key := make([]byte, 8)
	data := []byte("hello")

	keyVal := Bytes(&key)
	dataVal := Bytes(&data)

	txn, err := engine.BeginWrite()
	if err != ErrSuccess {
		b.Fatal(err)
	}

	for i := 0; i < iterations; i++ {
		//binary.BigEndian.PutUint64(key, uint64(20))
		//*(*uint64)(unsafe.Pointer(&key[0])) = uint64(i)
		binary.LittleEndian.PutUint64(key, uint64(i))
		//keyVal = U64(uint64(i))
		if err = txn.Put(engine.rootDB, &keyVal, &dataVal, 0); err != ErrSuccess {
			txn.Abort()
			b.Fatal(err)
		}
	}

	//*(*uint64)(unsafe.Pointer(&key[0])) = 0

	if err = txn.Commit(); err != ErrSuccess {
		b.Fatal(err)
	}

	txn = &Tx{}

	//engine.env.Sync(true, false)
	//engine.env.Sync(true, false)

	if err = engine.env.Begin(txn, TxReadOnly); err != ErrSuccess {
		b.Fatal(err)
	}
	//txn, err = engine.BeginRead()
	//if err != ErrSuccess {
	//	b.Fatal(err)
	//}

	cursor, err := txn.OpenCursor(engine.rootDB)
	if err != ErrSuccess {
		b.Fatal(err)
	}

	dataVal = Val{}
	keyVal = Val{}

	binary.LittleEndian.PutUint64(key, 0)

	count := 0
	//
	for {
		if err = cursor.Get(&keyVal, &dataVal, CursorNextNoDup); err != ErrSuccess {
			break
		}
		//if keyVal.Base == nil {
		//	break
		//}
		count++
		keyInt := keyVal.U64()
		println("key", keyInt)
		_ = keyInt

		//keyVal = Val{}
		//dataVal = Val{}

		//if cursor.EOF() != 0 {
		//	break
		//}
	}

	//for i := 0; i < b.N; i++ {
	//	*(*uint64)(unsafe.Pointer(&key[0])) = uint64(i)
	//	//binary.BigEndian.PutUint64(key, uint64(20))
	//	//binary.BigEndian.PutUint64(key[8:], uint64(i))
	//	//keyVal = U64(uint64(i))
	//	if err = txn.Get(engine.rootDB, &keyVal, &dataVal); err != ErrSuccess && err != ErrNotFound {
	//		txn.Reset()
	//		b.Fatal(err)
	//	}
	//}

	if err = cursor.Close(); err != ErrSuccess {
		b.Fatal(err)
	}
	if err = txn.Reset(); err != ErrSuccess {
		b.Fatal(err)
	}

	fmt.Println("count", count)

	//var envInfo EnvInfo
	//if err = txn.EnvInfo(&envInfo); err != ErrSuccess {
	//	b.Fatal(err)
	//}
	//var info TxInfo
	//if err = txn.Info(&info); err != ErrSuccess {
	//	b.Fatal(err)
	//}

	//engine.env.Sync(true, false)
	//engine.env.Sync(true, false)
}
