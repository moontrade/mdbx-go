package mdbx

/*
//#cgo !windows CFLAGS: -O2 -g -DMDBX_BUILD_FLAGS='' -DMDBX_DEBUG=0 -DNDEBUG=1 -DMDBX_FORCE_ASSERTIONS=1 -std=gnu11 -fvisibility=hidden -ffast-math  -fPIC -pthread -Wno-error=attributes -W -Wall -Werror -Wextra -Wpedantic -Wno-deprecated-declarations -Wno-format -Wno-implicit-fallthrough -Wno-unused-parameter -Wno-format-extra-args -Wno-missing-field-initializers
#cgo !windows CFLAGS: -O2 -g -DMDBX_BUILD_FLAGS='' -DMDBX_DEBUG=0 -DNDEBUG=1 -fPIC -ffast-math -std=gnu11 -fvisibility=hidden -pthread
#cgo linux LDFLAGS: -lrt

#include <stdlib.h>
#include <string.h>
#include <inttypes.h>
#include "mdbx.h"
#include "mdbx_utils.h"

#ifndef likely
#   if (defined(__GNUC__) || __has_builtin(__builtin_expect)) && !defined(__COVERITY__)
#       define likely(cond) __builtin_expect(!!(cond), 1)
#   else
#       define likely(x) (!!(x))
#   endif
#endif

#ifndef unlikely
#   if (defined(__GNUC__) || __has_builtin(__builtin_expect)) && !defined(__COVERITY__)
#       define unlikely(cond) __builtin_expect(!!(cond), 0)
#   else
#       define unlikely(x) (!!(x))
#   endif
#endif

static int cmp_lexical(const MDBX_val *a, const MDBX_val *b) {
  if (a->iov_len == b->iov_len)
    return a->iov_len ? memcmp(a->iov_base, b->iov_base, a->iov_len) : 0;

  const int diff_len = (a->iov_len < b->iov_len) ? -1 : 1;
  const size_t shortest = (a->iov_len < b->iov_len) ? a->iov_len : b->iov_len;
  int diff_data = shortest ? memcmp(a->iov_base, b->iov_base, shortest) : 0;
  return likely(diff_data) ? diff_data : diff_len;
}

int mdbx_cmp_u16(const MDBX_val *a, const MDBX_val *b) {
  //if (unlikely(a->iov_len < 2 || b->iov_len < 2)) {
  //  return cmp_lexical(a, b);
  //}
  uint16_t aa = *((uint16_t*)a->iov_base);
  uint16_t bb = *((uint16_t*)b->iov_base);
  return bb > aa ? -1 : aa > bb;
}

int mdbx_cmp_u32(const MDBX_val *a, const MDBX_val *b) {
  //if (unlikely(a->iov_len < 4 || b->iov_len < 4)) {
  //  return cmp_lexical(a, b);
  //}
  uint32_t aa = *((uint32_t*)a->iov_base);
  uint32_t bb = *((uint32_t*)b->iov_base);
  return bb > aa ? -1 : aa > bb;
}

int mdbx_cmp_u64(const MDBX_val *a, const MDBX_val *b) {
  //if (unlikely(a->iov_len < 8 || b->iov_len < 8)) {
  //  return cmp_lexical(a, b);
  //}
  uint64_t aa = *((uint64_t*)a->iov_base);
  uint64_t bb = *((uint64_t*)b->iov_base);
  return bb > aa ? -1 : aa > bb;
}

int mdbx_cmp_u16_prefix_lexical(const MDBX_val *a, const MDBX_val *b) {
  if (unlikely(a->iov_len < 2 || b->iov_len < 2)) {
    return cmp_lexical(a, b);
  }
  uint16_t aa = *((uint16_t*)a->iov_base);
  uint16_t bb = *((uint16_t*)a->iov_base);
  if (aa < bb) {
	return -1;
  }
  if (aa > bb) {
    return 1;
  }
  if (a->iov_len == b->iov_len)
    return a->iov_len ? memcmp(a->iov_base+2, b->iov_base+2, a->iov_len-2) : 0;

  const int diff_len = (a->iov_len < b->iov_len) ? -1 : 1;
  const size_t shortest = (a->iov_len < b->iov_len) ? a->iov_len : b->iov_len;
  int diff_data = shortest ? memcmp(a->iov_base+2, b->iov_base+2, shortest-2) : 0;
  return likely(diff_data) ? diff_data : diff_len;
}

int mdbx_cmp_u16_prefix_u64(const MDBX_val *a, const MDBX_val *b) {
  if (unlikely(a->iov_len < 10 || b->iov_len < 10)) {
    return cmp_lexical(a, b);
  }
  uint16_t aa = *((uint16_t*)a->iov_base);
  uint16_t bb = *((uint16_t*)a->iov_base);
  if (aa < bb) return -1;
  if (aa > bb) return 1;
  uint64_t aa2 = *((uint64_t*)a->iov_base+2);
  uint64_t bb2 = *((uint64_t*)b->iov_base+2);
  if (aa2 < bb2) return -1;
  if (aa2 > bb2) return 1;
  return 0;
}

int mdbx_cmp_u32_prefix_lexical(const MDBX_val *a, const MDBX_val *b) {
  if (unlikely(a->iov_len < 4 || b->iov_len < 4)) {
    return cmp_lexical(a, b);
  }
  uint32_t aa = *((uint32_t*)a->iov_base);
  uint32_t bb = *((uint32_t*)a->iov_base);
  if (aa < bb) {
	return -1;
  }
  if (aa > bb) {
    return 1;
  }
  if (a->iov_len == b->iov_len)
    return a->iov_len ? memcmp(a->iov_base+4, b->iov_base+4, a->iov_len-4) : 0;

  const int diff_len = (a->iov_len < b->iov_len) ? -1 : 1;
  const size_t shortest = (a->iov_len < b->iov_len) ? a->iov_len : b->iov_len;
  int diff_data = shortest ? memcmp(a->iov_base+4, b->iov_base+4, shortest-4) : 0;
  return likely(diff_data) ? diff_data : diff_len;
}

int mdbx_cmp_u32_prefix_u64(const MDBX_val *a, const MDBX_val *b) {
  if (unlikely(a->iov_len < 12 || b->iov_len < 12)) {
   return cmp_lexical(a, b);
  }
  uint32_t aa = *((uint32_t*)a->iov_base);
  uint32_t bb = *((uint32_t*)a->iov_base);
  if (aa < bb) return -1;
  if (aa > bb) return 1;
  uint64_t aa2 = *((uint64_t*)a->iov_base+4);
  uint64_t bb2 = *((uint64_t*)b->iov_base+4);
  if (aa2 < bb2) return -1;
  if (aa2 > bb2) return 1;
  return 0;
}

int mdbx_cmp_u64_prefix_lexical(const MDBX_val *a, const MDBX_val *b) {
  if (unlikely(a->iov_len < 8 || b->iov_len < 8)) {
    return cmp_lexical(a, b);
  }
  uint64_t aa = *((uint64_t*)a->iov_base);
  uint64_t bb = *((uint64_t*)a->iov_base);
  if (aa < bb) {
	return -1;
  }
  if (aa > bb) {
    return 1;
  }
  if (a->iov_len == b->iov_len)
    return a->iov_len ? memcmp(a->iov_base+8, b->iov_base+8, a->iov_len-8) : 0;

  const int diff_len = (a->iov_len < b->iov_len) ? -1 : 1;
  const size_t shortest = (a->iov_len < b->iov_len) ? a->iov_len : b->iov_len;
  int diff_data = shortest ? memcmp(a->iov_base+8, b->iov_base+8, shortest-8) : 0;
  return likely(diff_data) ? diff_data : diff_len;
}

int mdbx_cmp_u64_prefix_u64(const MDBX_val *a, const MDBX_val *b) {
  if (unlikely(a->iov_len < 16 || b->iov_len < 16)) {
   return cmp_lexical(a, b);
  }
  uint64_t aa = *((uint32_t*)a->iov_base);
  uint64_t bb = *((uint32_t*)a->iov_base);
  if (aa < bb) return -1;
  if (aa > bb) return 1;
  aa = *((uint64_t*)a->iov_base+8);
  bb = *((uint64_t*)b->iov_base+8);
  if (aa < bb) return -1;
  if (aa > bb) return 1;
  return 0;
}

typedef struct mdbx_strerror_t {
	size_t result;
	int32_t code;
} mdbx_strerror_t;

void do_mdbx_strerror(size_t arg0, size_t arg1) {
	mdbx_strerror_t* args = (mdbx_strerror_t*)(void*)arg0;
	args->result = (size_t)(void*)mdbx_strerror((int)args->code);
}

typedef struct mdbx_env_set_geometry_t {
	size_t env;
	size_t size_lower;
	size_t size_now;
	size_t size_upper;
	size_t growth_step;
	size_t shrink_threshold;
	size_t page_size;
	int32_t result;
} mdbx_env_set_geometry_t;

void do_mdbx_env_set_geometry(size_t arg0, size_t arg1) {
	mdbx_env_set_geometry_t* args = (mdbx_env_set_geometry_t*)(void*)arg0;
	args->result = (int32_t)mdbx_env_set_geometry(
		(MDBX_env*)(void*)args->env,
		args->size_lower,
		args->size_now,
		args->size_upper,
		args->growth_step,
		args->shrink_threshold,
		args->page_size
	);
}

typedef struct mdbx_env_info_t {
	size_t env;
	size_t txn;
	size_t info;
	size_t size;
	int32_t result;
} mdbx_env_info_t;

void do_mdbx_env_info_ex(size_t arg0, size_t arg1) {
	mdbx_env_info_t* args = (mdbx_env_info_t*)(void*)arg0;
	args->result = (int32_t)mdbx_env_info_ex(
		(MDBX_env*)(void*)args->env,
		(MDBX_txn*)(void*)args->txn,
		(MDBX_envinfo*)(void*)args->info,
		args->size
	);
}

typedef struct mdbx_txn_info_t {
	size_t txn;
	size_t info;
	int32_t scan_rlt;
	int32_t result;
} mdbx_txn_info_t;

void do_mdbx_txn_info(size_t arg0, size_t arg1) {
	mdbx_txn_info_t* args = (mdbx_txn_info_t*)(void*)arg0;
	args->result = (int32_t)mdbx_txn_info(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_txn_info*)(void*)args->info,
		args->scan_rlt >= 0
	);
}

typedef struct mdbx_txn_flags_t {
	size_t txn;
	int32_t flags;
} mdbx_txn_flags_t;

void do_mdbx_txn_flags(size_t arg0, size_t arg1) {
	mdbx_txn_flags_t* args = (mdbx_txn_flags_t*)(void*)arg0;
	args->flags = (int32_t)mdbx_txn_flags(
		(MDBX_txn*)(void*)args->txn
	);
}

typedef struct mdbx_txn_id_t {
	size_t txn;
	uint64_t id;
} mdbx_txn_id_t;

void do_mdbx_txn_id(size_t arg0, size_t arg1) {
	mdbx_txn_id_t* args = (mdbx_txn_id_t*)(void*)arg0;
	args->id = mdbx_txn_id(
		(MDBX_txn*)(void*)args->txn
	);
}

typedef struct mdbx_txn_commit_ex_t {
	size_t txn;
	size_t latency;
	int32_t result;
} mdbx_txn_commit_ex_t;

void do_mdbx_txn_commit_ex(size_t arg0, size_t arg1) {
	mdbx_txn_commit_ex_t* args = (mdbx_txn_commit_ex_t*)(void*)arg0;
	args->result = (int32_t)mdbx_txn_commit_ex(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_commit_latency*)(void*)args->latency
	);
}

typedef struct mdbx_txn_result_t {
	size_t txn;
	int32_t result;
} mdbx_txn_result_t;

void do_mdbx_txn_abort(size_t arg0, size_t arg1) {
	mdbx_txn_result_t* args = (mdbx_txn_result_t*)(void*)arg0;
	args->result = (int32_t)mdbx_txn_abort(
		(MDBX_txn*)(void*)args->txn
	);
}

void do_mdbx_txn_break(size_t arg0, size_t arg1) {
	mdbx_txn_result_t* args = (mdbx_txn_result_t*)(void*)arg0;
	args->result = (int32_t)mdbx_txn_break(
		(MDBX_txn*)(void*)args->txn
	);
}

void do_mdbx_txn_reset(size_t arg0, size_t arg1) {
	mdbx_txn_result_t* args = (mdbx_txn_result_t*)(void*)arg0;
	args->result = (int32_t)mdbx_txn_reset(
		(MDBX_txn*)(void*)args->txn
	);
}

void do_mdbx_txn_renew(size_t arg0, size_t arg1) {
	mdbx_txn_result_t* args = (mdbx_txn_result_t*)(void*)arg0;
	args->result = (int32_t)mdbx_txn_renew(
		(MDBX_txn*)(void*)args->txn
	);
}

typedef struct mdbx_txn_canary_t {
	size_t txn;
	size_t canary;
	int32_t result;
} mdbx_txn_canary_t;

void do_mdbx_canary_put(size_t arg0, size_t arg1) {
	mdbx_txn_canary_t* args = (mdbx_txn_canary_t*)(void*)arg0;
	args->result = (int32_t)mdbx_canary_put(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_canary*)(void*)args->canary
	);
}

void do_mdbx_canary_get(size_t arg0, size_t arg1) {
	mdbx_txn_canary_t* args = (mdbx_txn_canary_t*)(void*)arg0;
	args->result = (int32_t)mdbx_canary_get(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_canary*)(void*)args->canary
	);
}

typedef struct mdbx_dbi_stat_t {
	size_t txn;
	size_t stat;
	size_t size;
	uint32_t dbi;
	int32_t result;
} mdbx_dbi_stat_t;

void do_mdbx_dbi_stat(size_t arg0, size_t arg1) {
	mdbx_dbi_stat_t* args = (mdbx_dbi_stat_t*)(void*)arg0;
	args->result = (int32_t)mdbx_dbi_stat(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_dbi)args->dbi,
		(MDBX_stat*)(void*)args->stat,
		args->size
	);
}

typedef struct mdbx_dbi_flags_t {
	size_t txn;
	size_t flags;
	size_t state;
	uint32_t dbi;
	int32_t result;
} mdbx_dbi_flags_t;

void do_mdbx_dbi_flags_ex(size_t arg0, size_t arg1) {
	mdbx_dbi_flags_t* args = (mdbx_dbi_flags_t*)(void*)arg0;
	args->result = (int32_t)mdbx_dbi_flags_ex(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_dbi)args->dbi,
		(unsigned*)(void*)args->flags,
		(unsigned*)(void*)args->state
	);
}

typedef struct mdbx_drop_t {
	size_t txn;
	size_t del;
	uint32_t dbi;
	int32_t result;
} mdbx_drop_t;

void do_mdbx_drop(size_t arg0, size_t arg1) {
	mdbx_drop_t* args = (mdbx_drop_t*)(void*)arg0;
	args->result = (int32_t)mdbx_drop(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_dbi)args->dbi,
		args->del > 0
	);
}

typedef struct mdbx_get_t {
	size_t txn;
	size_t key;
	size_t data;
	uint32_t dbi;
	int32_t result;
} mdbx_get_t;

void do_mdbx_get(size_t arg0, size_t arg1) {
	mdbx_get_t* args = (mdbx_get_t*)(void*)arg0;
	args->result = (int32_t)mdbx_get(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_dbi)args->dbi,
		(MDBX_val*)(void*)args->key,
		(MDBX_val*)(void*)args->data
	);
}

void do_mdbx_get_equal_or_great(size_t arg0, size_t arg1) {
	mdbx_get_t* args = (mdbx_get_t*)(void*)arg0;
	args->result = (int32_t)mdbx_get_equal_or_great(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_dbi)args->dbi,
		(MDBX_val*)(void*)args->key,
		(MDBX_val*)(void*)args->data
	);
}

typedef struct mdbx_get_ex_t {
	size_t txn;
	size_t key;
	size_t data;
	size_t values_count;
	uint32_t dbi;
	int32_t result;
} mdbx_get_ex_t;

void do_mdbx_get_ex(size_t arg0, size_t arg1) {
	mdbx_get_ex_t* args = (mdbx_get_ex_t*)(void*)arg0;
	args->result = (int32_t)mdbx_get_ex(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_dbi)args->dbi,
		(MDBX_val*)(void*)args->key,
		(MDBX_val*)(void*)args->data,
		(size_t*)(void*)args->values_count
	);
}

typedef struct mdbx_put_t {
	size_t txn;
	size_t key;
	size_t data;
	uint32_t dbi;
	uint32_t flags;
	int32_t result;
} mdbx_put_t;

void do_mdbx_put(size_t arg0, size_t arg1) {
	mdbx_put_t* args = (mdbx_put_t*)(void*)arg0;
	args->result = (int32_t)mdbx_put(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_dbi)args->dbi,
		(MDBX_val*)(void*)args->key,
		(MDBX_val*)(void*)args->data,
		(MDBX_put_flags_t)args->flags
	);
}

typedef struct mdbx_replace_t {
	size_t txn;
	size_t key;
	size_t data;
	size_t old_data;
	uint32_t dbi;
	uint32_t flags;
	int32_t result;
} mdbx_replace_t;

void do_mdbx_replace(size_t arg0, size_t arg1) {
	mdbx_replace_t* args = (mdbx_replace_t*)(void*)arg0;
	args->result = (int32_t)mdbx_replace(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_dbi)args->dbi,
		(MDBX_val*)(void*)args->key,
		(MDBX_val*)(void*)args->data,
		(MDBX_val*)(void*)args->old_data,
		(MDBX_put_flags_t)args->flags
	);
}

typedef struct mdbx_del_t {
	size_t txn;
	size_t key;
	size_t data;
	uint32_t dbi;
	int32_t result;
} mdbx_del_t;

void do_mdbx_del(size_t arg0, size_t arg1) {
	mdbx_del_t* args = (mdbx_del_t*)(void*)arg0;
	args->result = (int32_t)mdbx_del(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_dbi)args->dbi,
		(MDBX_val*)(void*)args->key,
		(MDBX_val*)(void*)args->data
	);
}

typedef struct mdbx_txn_begin_t {
	size_t env;
	size_t parent;
	size_t txn;
	size_t context;
	uint32_t flags;
	int32_t result;
} mdbx_txn_begin_t;

void do_mdbx_txn_begin_ex(size_t arg0, size_t arg1) {
	mdbx_txn_begin_t* args = (mdbx_txn_begin_t*)(void*)arg0;
	args->result = (int32_t)mdbx_txn_begin_ex(
		(MDBX_env*)(void*)args->env,
		//(MDBX_txn*)(void*)args->parent,
		NULL,
		(MDBX_txn_flags_t)args->flags,
		(MDBX_txn**)(void*)args->txn,
		(void*)args->context
	);
}

typedef struct mdbx_cursor_create_t {
	size_t context;
	size_t cursor;
} mdbx_cursor_create_t;

void do_mdbx_cursor_create(size_t arg0, size_t arg1) {
	mdbx_cursor_create_t* args = (mdbx_cursor_create_t*)(void*)arg0;
	args->cursor = (size_t)mdbx_cursor_create(
		(void*)args->context
	);
}

typedef struct mdbx_cursor_bind_t {
	size_t txn;
	size_t cursor;
	uint32_t dbi;
	int32_t result;
} mdbx_cursor_bind_t;

void do_mdbx_cursor_bind(size_t arg0, size_t arg1) {
	mdbx_cursor_bind_t* args = (mdbx_cursor_bind_t*)(void*)arg0;
	args->result = (int32_t)mdbx_cursor_bind(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_cursor*)(void*)args->cursor,
		(MDBX_dbi)args->dbi
	);
}

typedef struct mdbx_cursor_open_t {
	size_t txn;
	size_t cursor;
	uint32_t dbi;
	int32_t result;
} mdbx_cursor_open_t;

void do_mdbx_cursor_open(size_t arg0, size_t arg1) {
	mdbx_cursor_open_t* args = (mdbx_cursor_open_t*)(void*)arg0;
	args->result = (int32_t)mdbx_cursor_open(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_dbi)args->dbi,
		(MDBX_cursor**)(void*)args->cursor
	);
}

void do_mdbx_cursor_close(size_t arg0, size_t arg1) {
	mdbx_cursor_close((MDBX_cursor*)(void*)arg0);
}

typedef struct mdbx_cursor_renew_t {
	size_t txn;
	size_t cursor;
	int32_t result;
} mdbx_cursor_renew_t;

void do_mdbx_cursor_renew(size_t arg0, size_t arg1) {
	mdbx_cursor_renew_t* args = (mdbx_cursor_renew_t*)(void*)arg0;
	args->result = (int32_t)mdbx_cursor_renew(
		(MDBX_txn*)(void*)args->txn,
		(MDBX_cursor*)(void*)args->cursor
	);
}

typedef struct mdbx_cursor_txn_t {
	size_t cursor;
	size_t txn;
} mdbx_cursor_txn_t;

void do_mdbx_cursor_txn(size_t arg0, size_t arg1) {
	mdbx_cursor_txn_t* args = (mdbx_cursor_txn_t*)(void*)arg0;
	args->txn = (size_t)mdbx_cursor_txn(
		(MDBX_cursor*)(void*)args->cursor
	);
}

typedef struct mdbx_cursor_dbi_t {
	size_t cursor;
	uint32_t dbi;
} mdbx_cursor_dbi_t;

void do_mdbx_cursor_dbi(size_t arg0, size_t arg1) {
	mdbx_cursor_dbi_t* args = (mdbx_cursor_dbi_t*)(void*)arg0;
	args->dbi = (uint32_t)mdbx_cursor_dbi(
		(MDBX_cursor*)(void*)args->cursor
	);
}

typedef struct mdbx_cursor_copy_t {
	size_t src;
	size_t dest;
	int32_t result;
} mdbx_cursor_copy_t;

void do_mdbx_cursor_copy(size_t arg0, size_t arg1) {
	mdbx_cursor_copy_t* args = (mdbx_cursor_copy_t*)(void*)arg0;
	args->result = (int32_t)mdbx_cursor_copy(
		(MDBX_cursor*)(void*)args->src,
		(MDBX_cursor*)(void*)args->dest
	);
}

typedef struct mdbx_cursor_get_t {
	size_t cursor;
	size_t key;
	size_t data;
	uint32_t op;
	int32_t result;
} mdbx_cursor_get_t;

void do_mdbx_cursor_get(size_t arg0, size_t arg1) {
	mdbx_cursor_get_t* args = (mdbx_cursor_get_t*)(void*)arg0;
	args->result = (int32_t)mdbx_cursor_get(
		(MDBX_cursor*)(void*)args->cursor,
		(MDBX_val*)(void*)args->key,
		(MDBX_val*)(void*)args->data,
		(MDBX_cursor_op)args->op
	);
}

typedef struct mdbx_cursor_put_t {
	size_t cursor;
	size_t key;
	size_t data;
	uint32_t flags;
	int32_t result;
} mdbx_cursor_put_t;

void do_mdbx_cursor_put(size_t arg0, size_t arg1) {
	mdbx_cursor_put_t* args = (mdbx_cursor_put_t*)(void*)arg0;
	args->result = (int32_t)mdbx_cursor_put(
		(MDBX_cursor*)(void*)args->cursor,
		(MDBX_val*)(void*)args->key,
		(MDBX_val*)(void*)args->data,
		(MDBX_put_flags_t)args->flags
	);
}

typedef struct mdbx_cursor_del_t {
	size_t cursor;
	uint32_t flags;
	int32_t result;
} mdbx_cursor_del_t;

void do_mdbx_cursor_del(size_t arg0, size_t arg1) {
	mdbx_cursor_del_t* args = (mdbx_cursor_del_t*)(void*)arg0;
	args->result = (int32_t)mdbx_cursor_del(
		(MDBX_cursor*)(void*)args->cursor,
		(MDBX_put_flags_t)args->flags
	);
}

typedef struct mdbx_cursor_count_t {
	size_t cursor;
	size_t count;
	int32_t result;
} mdbx_cursor_count_t;

void do_mdbx_cursor_count(size_t arg0, size_t arg1) {
	mdbx_cursor_count_t* args = (mdbx_cursor_count_t*)(void*)arg0;
	args->result = (int32_t)mdbx_cursor_count(
		(MDBX_cursor*)(void*)args->cursor,
		(size_t*)(void*)args->count
	);
}

typedef struct mdbx_cursor_eof_t {
	size_t cursor;
	int32_t result;
} mdbx_cursor_eof_t;

void do_mdbx_cursor_eof(size_t arg0, size_t arg1) {
	mdbx_cursor_eof_t* args = (mdbx_cursor_eof_t*)(void*)arg0;
	args->result = (int32_t)mdbx_cursor_eof(
		(MDBX_cursor*)(void*)args->cursor
	);
}

typedef struct mdbx_cursor_on_first_t {
	size_t cursor;
	int32_t result;
} mdbx_cursor_on_first_t;

void do_mdbx_cursor_on_first(size_t arg0, size_t arg1) {
	mdbx_cursor_on_first_t* args = (mdbx_cursor_on_first_t*)(void*)arg0;
	args->result = (int32_t)mdbx_cursor_on_first(
		(MDBX_cursor*)(void*)args->cursor
	);
}

typedef struct mdbx_cursor_on_last_t {
	size_t cursor;
	int32_t result;
} mdbx_cursor_on_last_t;

void do_mdbx_cursor_on_last(size_t arg0, size_t arg1) {
	mdbx_cursor_on_last_t* args = (mdbx_cursor_on_last_t*)(void*)arg0;
	args->result = (int32_t)mdbx_cursor_on_last(
		(MDBX_cursor*)(void*)args->cursor
	);
}

typedef struct mdbx_estimate_distance_t {
	size_t first;
	size_t last;
	int64_t distance_items;
	int32_t result;
} mdbx_estimate_distance_t;

void do_mdbx_estimate_distance(size_t arg0, size_t arg1) {
	mdbx_estimate_distance_t* args = (mdbx_estimate_distance_t*)(void*)arg0;
	args->result = (int32_t)mdbx_estimate_distance(
		(MDBX_cursor*)(void*)args->first,
		(MDBX_cursor*)(void*)args->last,
		(ptrdiff_t*)(void*)args->distance_items
	);
}

//typedef struct mdbx_estimate_move_t {
//	size_t txn;
//	size_t last;
//	int64_t distance_items;
//	int32_t result;
//} mdbx_estimate_move_t;

*/
import "C"
import (
	"github.com/moontrade/mdbx-go/internal/capture"
	"github.com/moontrade/mdbx-go/internal/unsafecgo"
	"os"
	"reflect"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

const (
	MaxDBI      = uint32(C.MDBX_MAX_DBI)
	MaxDataSize = uint32(C.MDBX_MAXDATASIZE)
	MinPageSize = int(C.MDBX_MIN_PAGESIZE)
	MaxPageSize = int(C.MDBX_MAX_PAGESIZE)
)

func init() {
	sz0 := unsafe.Sizeof(C.MDBX_envinfo{})
	sz1 := unsafe.Sizeof(EnvInfo{})
	if sz0 != sz1 {
		panic("sizeof(C.MDBX_envinfo) != sizeof(EnvInfo{})")
	}
}

type Cmp C.MDBX_cmp_func

var (
	CmpU16              = (*Cmp)(C.mdbx_cmp_u16)
	CmpU32              = (*Cmp)(C.mdbx_cmp_u32)
	CmpU64              = (*Cmp)(C.mdbx_cmp_u64)
	CmpU16PrefixLexical = (*Cmp)(C.mdbx_cmp_u16_prefix_lexical)
	CmpU16PrefixU64     = (*Cmp)(C.mdbx_cmp_u16_prefix_u64)
	CmpU32PrefixLexical = (*Cmp)(C.mdbx_cmp_u32_prefix_lexical)
	CmpU32PrefixU64     = (*Cmp)(C.mdbx_cmp_u32_prefix_u64)
	CmpU64PrefixLexical = (*Cmp)(C.mdbx_cmp_u64_prefix_lexical)
	CmpU64PrefixU64     = (*Cmp)(C.mdbx_cmp_u64_prefix_u64)
)

// Chk invokes the embedded mdbx_chk utility
// usage: mdbx_chk [-V] [-v] [-q] [-c] [-0|1|2] [-w] [-d] [-i] [-s subdb] dbpath
//  -V            print version and exit
//  -v            more verbose, could be used multiple times
//  -q            be quiet
//  -c            force cooperative mode (don't try exclusive)
//  -w            write-mode checking
//  -d            disable page-by-page traversal of B-tree
//  -i            ignore wrong order errors (for custom comparators case)
//  -s subdb      process a specific subdatabase only
//  -0|1|2        force using specific meta-page 0, or 2 for checking
//  -t            turn to a specified meta-page on successful check
//  -T            turn to a specified meta-page EVEN ON UNSUCCESSFUL CHECK!
func Chk(args ...string) (result int32, output []byte, err error) {
	argv := make([]*C.char, len(args)+1)
	argv[0] = C.CString("mdbx_chk")
	for i, arg := range args {
		argv[i+1] = C.CString(arg)
	}
	defer func() {
		for _, arg := range argv {
			C.free(unsafe.Pointer(arg))
		}
	}()

	//	ch := make(chan string)
	//	go func() {
	//		defer close(ch)
	//		err = capture.CaptureWithCGoChan(ch, func() {
	//			result = int32(C.mdbx_chk((C.int)(len(argv)), (**C.char)(unsafe.Pointer(&argv[0]))))
	//		})
	//	}()
	//
	//	var lines []string
	//LOOP:
	//	for {
	//		select {
	//		case line, ok := <-ch:
	//			if !ok {
	//				break LOOP
	//			}
	//			//fmt.Fprintf(out, "%s\n", line)
	//			lines = append(lines, line)
	//		}
	//	}
	//	for _, line := range lines {
	//		println(line)
	//	}

	output, err = capture.CaptureWithCGo(func() {
		result = int32(C.mdbx_chk((C.int)(len(argv)), (**C.char)(unsafe.Pointer(&argv[0]))))
	})
	return
}

// ChkMain invokes the embedded mdbx_chk utility and exits the program.
// usage: mdbx_chk [-V] [-v] [-q] [-c] [-0|1|2] [-w] [-d] [-i] [-s subdb] dbpath
//  -V            print version and exit
//  -v            more verbose, could be used multiple times
//  -q            be quiet
//  -c            force cooperative mode (don't try exclusive)
//  -w            write-mode checking
//  -d            disable page-by-page traversal of B-tree
//  -i            ignore wrong order errors (for custom comparators case)
//  -s subdb      process a specific subdatabase only
//  -0|1|2        force using specific meta-page 0, or 2 for checking
//  -t            turn to a specified meta-page on successful check
//  -T            turn to a specified meta-page EVEN ON UNSUCCESSFUL CHECK!
func ChkMain(args ...string) {
	argv := make([]*C.char, len(args)+1)
	argv[0] = C.CString("mdbx_chk")
	for i, arg := range args {
		argv[i+1] = C.CString(arg)
	}
	defer func() {
		for _, arg := range argv {
			C.free(unsafe.Pointer(arg))
		}
	}()

	os.Exit(int(C.mdbx_chk((C.int)(len(argv)), (**C.char)(unsafe.Pointer(&argv[0])))))
}

// Stat invokes the embedded mdbx_stat utility.
// usage: mdbx_stat [-V] [-q] [-e] [-f[f[f]]] [-r[r]] [-a|-s name] dbpath
//  -V            print version and exit
//  -q            be quiet
//  -p            show statistics of page operations for current session
//  -e            show whole DB info
//  -f            show GC info
//  -r            show readers
//  -a            print stat of main DB and all subDBs
//  -s name       print stat of only the specified named subDB
//                by default print stat of only the main DB
func Stat(args ...string) (result int32, output []byte, err error) {
	argv := make([]*C.char, len(args)+1)
	argv[0] = C.CString("mdbx_stat")
	for i, arg := range args {
		argv[i+1] = C.CString(arg)
	}
	defer func() {
		for _, arg := range argv {
			C.free(unsafe.Pointer(arg))
		}
	}()
	output, err = capture.CaptureWithCGo(func() {
		result = int32(C.mdbx_stat((C.int)(len(argv)), (**C.char)(unsafe.Pointer(&argv[0]))))
	})
	return
}

// Stat invokes the embedded mdbx_stat utility and exits the program.
// usage: mdbx_stat [-V] [-q] [-e] [-f[f[f]]] [-r[r]] [-a|-s name] dbpath
//  -V            print version and exit
//  -q            be quiet
//  -p            show statistics of page operations for current session
//  -e            show whole DB info
//  -f            show GC info
//  -r            show readers
//  -a            print stat of main DB and all subDBs
//  -s name       print stat of only the specified named subDB
//                by default print stat of only the main DB
func StatMain(args ...string) {
	argv := make([]*C.char, len(args)+1)
	argv[0] = C.CString("mdbx_chk")
	for i, arg := range args {
		argv[i+1] = C.CString(arg)
	}
	defer func() {
		for _, arg := range argv {
			C.free(unsafe.Pointer(arg))
		}
	}()

	os.Exit(int(C.mdbx_stat((C.int)(len(argv)), (**C.char)(unsafe.Pointer(&argv[0])))))
}

type LogLevel int32

const (
	// LogFatal Critical conditions, i.e. assertion failures
	LogFatal = LogLevel(C.MDBX_LOG_FATAL)

	// LogError Enables logging for error conditions and \ref MDBX_LOG_FATAL
	LogError = LogLevel(C.MDBX_LOG_ERROR)

	// LogWarn Enables logging for warning conditions and \ref MDBX_LOG_ERROR ...
	// \ref MDBX_LOG_FATAL
	LogWarn = LogLevel(C.MDBX_LOG_WARN)

	// LogNotice Enables logging for normal but significant condition and
	// \ref MDBX_LOG_WARN ... \ref MDBX_LOG_FATAL
	LogNotice = LogLevel(C.MDBX_LOG_NOTICE)

	// LogVerbose Enables logging for verbose informational and \ref MDBX_LOG_NOTICE ...
	// \ref MDBX_LOG_FATAL
	LogVerbose = LogLevel(C.MDBX_LOG_VERBOSE)

	// LogDebug Enables logging for debug-level messages and \ref MDBX_LOG_VERBOSE ...
	// \ref MDBX_LOG_FATAL
	LogDebug = LogLevel(C.MDBX_LOG_DEBUG)

	// LogTrace Enables logging for trace debug-level messages and \ref MDBX_LOG_DEBUG ...
	// \ref MDBX_LOG_FATAL
	LogTrace = LogLevel(C.MDBX_LOG_TRACE)

	// LogExtra Enables extra debug-level messages (dump pgno lists) and all other log-messages
	LogExtra = LogLevel(C.MDBX_LOG_EXTRA)
	LogMax   = LogLevel(7)

	// LogDontChange for \ref mdbx_setup_debug() only: Don't change current settings
	LogDontChange = LogLevel(C.MDBX_LOG_DONTCHANGE)
)

type Error int32

func (e Error) Error() string {
	args := struct {
		result uintptr
		code   int32
	}{
		code: int32(e),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_strerror), ptr, 0)
	str := C.GoString((*C.char)(unsafe.Pointer(args.result)))
	return str
}

const (
	ErrSuccess     = Error(C.MDBX_SUCCESS)
	ErrResultFalse = ErrSuccess

	// ErrResultTrue Successful result with special meaning or a flag
	ErrResultTrue = Error(C.MDBX_RESULT_TRUE)

	// ErrKeyExist key/data pair already exist
	ErrKeyExist = Error(C.MDBX_KEYEXIST)

	// ErrFirstLMDBErrCode The first LMDB-compatible defined error code
	ErrFirstLMDBErrCode = ErrKeyExist

	// ErrNotFound key/data pair not found (EOF)
	ErrNotFound = Error(C.MDBX_NOTFOUND)

	// ErrPageNotFound Requested page not found -this usually indicates corruption
	ErrPageNotFound = Error(C.MDBX_PAGE_NOTFOUND)

	// ErrCorrupted Database is corrupted (page was wrong type and so on)
	ErrCorrupted = Error(C.MDBX_CORRUPTED)

	// ErrPanic Environment had fatal error, i.e. update of meta page failed and so on.
	ErrPanic = Error(C.MDBX_PANIC)

	// ErrVersionMismatch DB file version mismatch with libmdbx
	ErrVersionMismatch = Error(C.MDBX_VERSION_MISMATCH)

	// ErrInvalid File is not a valid MDBX file
	ErrInvalid = Error(C.MDBX_INVALID)

	// ErrMapFull Environment mapsize reached
	ErrMapFull = Error(C.MDBX_MAP_FULL)

	// ErrDBSFull Environment maxdbs reached
	ErrDBSFull = Error(C.MDBX_DBS_FULL)

	// ErrReadersFull Environment maxreaders reached
	ErrReadersFull = Error(C.MDBX_READERS_FULL)

	// ErrTXNFull Transaction has too many dirty pages, i.e transaction too big
	ErrTXNFull = Error(C.MDBX_TXN_FULL)

	// ErrCursorFull Cursor stack too deep -this usually indicates corruption, i.e branchC.pages loop
	ErrCursorFull = Error(C.MDBX_CURSOR_FULL)

	// ErrPageFull Page has not enough space -internal error
	ErrPageFull = Error(C.MDBX_PAGE_FULL)

	// ErrUnableExtendMapSize Database engine was unable to extend mapping, e.g. since address space
	// is unavailable or busy. This can mean:
	//  - Database size extended by other process beyond to environment mapsize
	//    and engine was unable to extend mapping while starting read
	//    transaction. Environment should be reopened to continue.
	//  - Engine was unable to extend mapping during write transaction
	//    or explicit call of \ref mdbx_env_set_geometry().
	ErrUnableExtendMapSize = Error(C.MDBX_UNABLE_EXTEND_MAPSIZE)

	// ErrIncompatible Environment or database is not compatible with the requested operation
	// or the specified flags. This can mean:
	//  - The operation expects an \ref MDBX_DUPSORT / \ref MDBX_DUPFIXED
	//    database.
	//  - Opening a named DB when the unnamed DB has \ref MDBX_DUPSORT /
	//    \ref MDBX_INTEGERKEY.
	//  - Accessing a data record as a database, or vice versa.
	//  - The database was dropped and recreated with different flags.
	ErrIncompatible = Error(C.MDBX_INCOMPATIBLE)

	// ErrBadRSlot Invalid reuse of reader locktable slot
	// e.g. readC.transaction already run for current thread
	ErrBadRSlot = Error(C.MDBX_BAD_RSLOT)

	// ErrBadTXN Transaction is not valid for requested operation,
	// e.g. had errored and be must aborted, has a child, or is invalid
	ErrBadTXN = Error(C.MDBX_BAD_TXN)

	// ErrBadValSize Invalid size or alignment of key or data for target database,
	// either invalid subDB name
	ErrBadValSize = Error(C.MDBX_BAD_VALSIZE)

	// ErrBadDBI The specified DBIC.handle is invalid
	// or changed by another thread/transaction.
	ErrBadDBI = Error(C.MDBX_BAD_DBI)

	// ErrProblem Unexpected internal error, transaction should be aborted
	ErrProblem = Error(C.MDBX_PROBLEM)

	// ErrLastLMDBErrCode The last LMDBC.compatible defined error code
	ErrLastLMDBErrCode = ErrProblem

	// ErrBusy Another write transaction is running or environment is already used while
	// opening with \ref MDBX_EXCLUSIVE flag
	ErrBusy              = Error(C.MDBX_BUSY)
	ErrFirstAddedErrCode = ErrBusy                 // The first of MDBXC.added error codes
	ErrEMultiVal         = Error(C.MDBX_EMULTIVAL) // The specified key has more than one associated value

	// ErrEBadSign Bad signature of a runtime object(s), this can mean:
	//  - memory corruption or doubleC.free;
	//  - ABI version mismatch (rare case);
	ErrEBadSign = Error(C.MDBX_EBADSIGN)

	// ErrWannaRecovery Database should be recovered, but this could NOT be done for now
	// since it opened in readC.only mode.
	ErrWannaRecovery = Error(C.MDBX_WANNA_RECOVERY)

	// ErrEKeyMismatch The given key value is mismatched to the current cursor position
	ErrEKeyMismatch = Error(C.MDBX_EKEYMISMATCH)

	// ErrTooLarge Database is too large for current system,
	// e.g. could NOT be mapped into RAM.
	ErrTooLarge = Error(C.MDBX_TOO_LARGE)

	// ErrThreadMismatch A thread has attempted to use a not owned object,
	// e.g. a transaction that started by another thread.
	ErrThreadMismatch = Error(C.MDBX_THREAD_MISMATCH)

	// ErrTXNOverlapping Overlapping read and write transactions for the current thread
	ErrTXNOverlapping = Error(C.MDBX_TXN_OVERLAPPING)

	ErrLastAddedErrcode = ErrTXNOverlapping

	ErrENODAT  = Error(C.MDBX_ENODATA)
	ErrEINVAL  = Error(C.MDBX_EINVAL)
	ErrEACCESS = Error(C.MDBX_EACCESS)
	ErrENOMEM  = Error(C.MDBX_ENOMEM)
	ErrEROFS   = Error(C.MDBX_EROFS)
	ErrENOSYS  = Error(C.MDBX_ENOSYS)
	ErrEIO     = Error(C.MDBX_EIO)
	ErrEPERM   = Error(C.MDBX_EPERM)
	ErrEINTR   = Error(C.MDBX_EINTR)
	ErrENOENT  = Error(C.MDBX_ENOFILE)
	ErrENOTBLK = Error(C.MDBX_EREMOTE)
)

type EnvFlags uint32

const (
	EnvEnvDefaults EnvFlags = 0

	// EnvNoSubDir No environment directory.
	//
	// By default, MDBX creates its environment in a directory whose pathname is
	// given in path, and creates its data and lock files under that directory.
	// With this option, path is used as-is for the database rootDB data file.
	// The database lock file is the path with "-lck" appended.
	//
	// - with `MDBX_NOSUBDIR` = in a filesystem we have the pair of MDBX-files
	//   which names derived from given pathname by appending predefined suffixes.
	//
	// - without `MDBX_NOSUBDIR` = in a filesystem we have the MDBX-directory with
	//   given pathname, within that a pair of MDBX-files with predefined names.
	//
	// This flag affects only at new environment creating by \ref mdbx_env_open(),
	// otherwise at opening an existing environment libmdbx will choice this
	// automatically.
	EnvNoSubDir = EnvFlags(C.MDBX_NOSUBDIR)

	// EnvReadOnly Read only mode.
	//
	// Open the environment in read-only mode. No write operations will be
	// allowed. MDBX will still modify the lock file - except on read-only
	// filesystems, where MDBX does not use locks.
	//
	// - with `MDBX_RDONLY` = open environment in read-only mode.
	//   MDBX supports pure read-only mode (i.e. without opening LCK-file) only
	//   when environment directory and/or both files are not writable (and the
	//   LCK-file may be missing). In such case allowing file(s) to be placed
	//   on a network read-only share.
	//
	// - without `MDBX_RDONLY` = open environment in read-write mode.
	//
	// This flag affects only at environment opening but can't be changed after.
	EnvReadOnly = EnvFlags(C.MDBX_RDONLY)

	// EnvExclusive Open environment in exclusive/monopolistic mode.
	//
	// `MDBX_EXCLUSIVE` flag can be used as a replacement for `MDB_NOLOCK`,
	// which don't supported by MDBX.
	// In this way, you can get the minimal overhead, but with the correct
	// multi-process and multi-thread locking.
	//
	// - with `MDBX_EXCLUSIVE` = open environment in exclusive/monopolistic mode
	//   or return \ref MDBX_BUSY if environment already used by other process.
	//   The rootDB feature of the exclusive mode is the ability to open the
	//   environment placed on a network share.
	//
	// - without `MDBX_EXCLUSIVE` = open environment in cooperative mode,
	//   i.e. for multi-process access/interaction/cooperation.
	//   The rootDB requirements of the cooperative mode are:
	//
	//   1. data files MUST be placed in the LOCAL file system,
	//      but NOT on a network share.
	//   2. environment MUST be opened only by LOCAL processes,
	//      but NOT over a network.
	//   3. OS kernel (i.e. file system and memory mapping implementation) and
	//      all processes that open the given environment MUST be running
	//      in the physically single RAM with cache-coherency. The only
	//      exception for cache-consistency requirement is Linux on MIPS
	//      architecture, but this case has not been tested for a long time).
	//
	// This flag affects only at environment opening but can't be changed after.
	EnvExclusive = EnvFlags(C.MDBX_EXCLUSIVE)

	// EnvAccede Using database/environment which already opened by another process(es).
	//
	// The `MDBX_ACCEDE` flag is useful to avoid \ref MDBX_INCOMPATIBLE error
	// while opening the database/environment which is already used by another
	// process(es) with unknown mode/flags. In such cases, if there is a
	// difference in the specified flags (\ref MDBX_NOMETASYNC,
	// \ref MDBX_SAFE_NOSYNC, \ref MDBX_UTTERLY_NOSYNC, \ref MDBX_LIFORECLAIM,
	// \ref MDBX_COALESCE and \ref MDBX_NORDAHEAD), instead of returning an error,
	// the database will be opened in a compatibility with the already used mode.
	//
	// `MDBX_ACCEDE` has no effect if the current process is the only one either
	// opening the DB in read-only mode or other process(es) uses the DB in
	// read-only mode.
	EnvAccede = EnvFlags(C.MDBX_ACCEDE)

	// EnvWriteMap Map data into memory with write permission.
	//
	// Use a writeable memory map unless \ref MDBX_RDONLY is set. This uses fewer
	// mallocs and requires much less work for tracking database pages, but
	// loses protection from application bugs like wild pointer writes and other
	// bad updates into the database. This may be slightly faster for DBs that
	// fit entirely in RAM, but is slower for DBs larger than RAM. Also adds the
	// possibility for stray application writes thru pointers to silently
	// corrupt the database.
	//
	// - with `MDBX_WRITEMAP` = all data will be mapped into memory in the
	//   read-write mode. This offers a significant performance benefit, since the
	//   data will be modified directly in mapped memory and then flushed to disk
	//   by single system call, without any memory management nor copying.
	//
	// - without `MDBX_WRITEMAP` = data will be mapped into memory in the
	//   read-only mode. This requires stocking all modified database pages in
	//   memory and then writing them to disk through file operations.
	//
	// \warning On the other hand, `MDBX_WRITEMAP` adds the possibility for stray
	// application writes thru pointers to silently corrupt the database.
	//
	// \note The `MDBX_WRITEMAP` mode is incompatible with nested transactions,
	// since this is unreasonable. I.e. nested transactions requires mallocation
	// of database pages and more work for tracking ones, which neuters a
	// performance boost caused by the `MDBX_WRITEMAP` mode.
	//
	// This flag affects only at environment opening but can't be changed after.
	EnvWriteMap = EnvFlags(C.MDBX_WRITEMAP)

	// EnvNoTLS Tie reader locktable slots to read-only transactions
	// instead of to threads.
	//
	// Don't use Thread-Local Storage, instead tie reader locktable slots to
	// \ref MDBX_txn objects instead of to threads. So, \ref mdbx_txn_reset()
	// keeps the slot reserved for the \ref MDBX_txn object. A thread may use
	// parallel read-only transactions. And a read-only transaction may span
	// threads if you synchronizes its use.
	//
	// Applications that multiplex many user threads over individual OS threads
	// need this option. Such an application must also serialize the write
	// transactions in an OS thread, since MDBX's write locking is unaware of
	// the user threads.
	//
	// \note Regardless to `MDBX_NOTLS` flag a write transaction entirely should
	// always be used in one thread from start to finish. MDBX checks this in a
	// reasonable manner and return the \ref MDBX_THREAD_MISMATCH error in rules
	// violation.
	//
	// This flag affects only at environment opening but can't be changed after.
	EnvNoTLS = EnvFlags(C.MDBX_NOTLS)
	//MDBX_NOTLS = UINT32_C(0x200000)

	// EnvNoReadAhead Don't do readahead.
	//
	// Turn off readahead. Most operating systems perform readahead on read
	// requests by default. This option turns it off if the OS supports it.
	// Turning it off may help random read performance when the DB is larger
	// than RAM and system RAM is full.
	//
	// By default libmdbx dynamically enables/disables readahead depending on
	// the actual database size and currently available memory. On the other
	// hand, such automation has some limitation, i.e. could be performed only
	// when DB size changing but can't tracks and reacts changing a free RAM
	// availability, since it changes independently and asynchronously.
	//
	// \note The mdbx_is_readahead_reasonable() function allows to quickly find
	// out whether to use readahead or not based on the size of the data and the
	// amount of available memory.
	//
	// This flag affects only at environment opening and can't be changed after.
	EnvNoReadAhead = EnvFlags(C.MDBX_NORDAHEAD)

	// EnvNoMemInit Don't initialize malloc'ed memory before writing to datafile.
	//
	// Don't initialize malloc'ed memory before writing to unused spaces in the
	// data file. By default, memory for pages written to the data file is
	// obtained using malloc. While these pages may be reused in subsequent
	// transactions, freshly malloc'ed pages will be initialized to zeroes before
	// use. This avoids persisting leftover data from other code (that used the
	// heap and subsequently freed the memory) into the data file.
	//
	// Note that many other system libraries may allocate and free memory from
	// the heap for arbitrary uses. E.g., stdio may use the heap for file I/O
	// buffers. This initialization step has a modest performance cost so some
	// applications may want to disable it using this flag. This option can be a
	// problem for applications which handle sensitive data like passwords, and
	// it makes memory checkers like Valgrind noisy. This flag is not needed
	// with \ref MDBX_WRITEMAP, which writes directly to the mmap instead of using
	// malloc for pages. The initialization is also skipped if \ref MDBX_RESERVE
	// is used; the caller is expected to overwrite all of the memory that was
	// reserved in that case.
	//
	// This flag may be changed at any time using `mdbx_env_set_flags()`.
	EnvNoMemInit = EnvFlags(C.MDBX_NOMEMINIT)

	// EnvCoalesce Aims to coalesce a Garbage Collection items.
	//
	// With `MDBX_COALESCE` flag MDBX will aims to coalesce items while recycling
	// a Garbage Collection. Technically, when possible short lists of pages
	// will be combined into longer ones, but to fit on one database page. As a
	// result, there will be fewer items in Garbage Collection and a page lists
	// are longer, which slightly increases the likelihood of returning pages to
	// Unallocated space and reducing the database file.
	//
	// This flag may be changed at any time using mdbx_env_set_flags().
	EnvCoalesce = EnvFlags(C.MDBX_COALESCE)

	// EnvLIFOReclaim LIFO policy for recycling a Garbage Collection items.
	//
	// `MDBX_LIFORECLAIM` flag turns on LIFO policy for recycling a Garbage
	// Collection items, instead of FIFO by default. On systems with a disk
	// write-back cache, this can significantly increase write performance, up
	// to several times in a best case scenario.
	//
	// LIFO recycling policy means that for reuse pages will be taken which became
	// unused the lastest (i.e. just now or most recently). Therefore the loop of
	// database pages circulation becomes as short as possible. In other words,
	// the number of pages, that are overwritten in memory and on disk during a
	// series of write transactions, will be as small as possible. Thus creates
	// ideal conditions for the efficient operation of the disk write-back cache.
	//
	// \ref MDBX_LIFORECLAIM is compatible with all no-sync flags, but gives NO
	// noticeable impact in combination with \ref MDBX_SAFE_NOSYNC or
	// \ref MDBX_UTTERLY_NOSYN-Because MDBX will reused pages only before the
	// last "steady" MVCC-snapshot, i.e. the loop length of database pages
	// circulation will be mostly defined by frequency of calling
	// \ref mdbx_env_sync() rather than LIFO and FIFO difference.
	//
	// This flag may be changed at any time using mdbx_env_set_flags().
	EnvLIFOReclaim = EnvFlags(C.MDBX_LIFORECLAIM)

	// EnvPagPerTurb Debugging option, fill/perturb released pages.
	EnvPagePerTurb = EnvFlags(C.MDBX_PAGEPERTURB)

	// SYNC MODES

	// \defgroup sync_modes SYNC MODES
	//
	// \attention Using any combination of \ref MDBX_SAFE_NOSYNC, \ref
	// MDBX_NOMETASYNC and especially \ref MDBX_UTTERLY_NOSYNC is always a deal to
	// reduce durability for gain write performance. You must know exactly what
	// you are doing and what risks you are taking!
	//
	// \note for LMDB users: \ref MDBX_SAFE_NOSYNC is NOT similar to LMDB_NOSYNC,
	// but \ref MDBX_UTTERLY_NOSYNC is exactly match LMDB_NOSYN-See details
	// below.
	//
	// THE SCENE:
	// - The DAT-file contains several MVCC-snapshots of B-tree at same time,
	//   each of those B-tree has its own root page.
	// - Each of meta pages at the beginning of the DAT file contains a
	//   pointer to the root page of B-tree which is the result of the particular
	//   transaction, and a number of this transaction.
	// - For data durability, MDBX must first write all MVCC-snapshot data
	//   pages and ensure that are written to the disk, then update a meta page
	//   with the new transaction number and a pointer to the corresponding new
	//   root page, and flush any buffers yet again.
	// - Thus during commit a I/O buffers should be flushed to the disk twice;
	//   i.e. fdatasync(), FlushFileBuffers() or similar syscall should be
	//   called twice for each commit. This is very expensive for performance,
	//   but guaranteed durability even on unexpected system failure or power
	//   outage. Of course, provided that the operating system and the
	//   underlying hardware (e.g. disk) work correctly.
	//
	// TRADE-OFF:
	// By skipping some stages described above, you can significantly benefit in
	// speed, while partially or completely losing in the guarantee of data
	// durability and/or consistency in the event of system or power failure.
	// Moreover, if for any reason disk write order is not preserved, then at
	// moment of a system crash, a meta-page with a pointer to the new B-tree may
	// be written to disk, while the itself B-tree not yet. In that case, the
	// database will be corrupted!
	//
	// \see MDBX_SYNC_DURABLE \see MDBX_NOMETASYNC \see MDBX_SAFE_NOSYNC
	// \see MDBX_UTTERLY_NOSYNC
	//
	// @{

	// EnvSyncDurable Default robust and durable sync mode.
	//
	// Metadata is written and flushed to disk after a data is written and
	// flushed, which guarantees the integrity of the database in the event
	// of a crash at any time.
	//
	// \attention Please do not use other modes until you have studied all the
	// details and are sure. Otherwise, you may lose your users' data, as happens
	// in [Miranda NG](https://www.miranda-ng.org/) messenger.
	EnvSyncDurable = EnvFlags(C.MDBX_SYNC_DURABLE)

	// EnvNoMetaSync Don't sync the meta-page after commit.
	//
	// Flush system buffers to disk only once per transaction commit, omit the
	// metadata flush. Defer that until the system flushes files to disk,
	// or next non-\ref MDBX_RDONLY commit or \ref mdbx_env_sync(). Depending on
	// the platform and hardware, with \ref MDBX_NOMETASYNC you may get a doubling
	// of write performance.
	//
	// This trade-off maintains database integrity, but a system crash may
	// undo the last committed transaction. I.e. it preserves the ACI
	// (atomicity, consistency, isolation) but not D (durability) database
	// property.
	//
	// `MDBX_NOMETASYNC` flag may be changed at any time using
	// \ref mdbx_env_set_flags() or by passing to \ref mdbx_txn_begin() for
	// particular write transaction. \see sync_modes
	EnvNoMetaSync = EnvFlags(C.MDBX_NOMETASYNC)

	// EnvSafeNoSync Don't sync anything but keep previous steady commits.
	//
	// Like \ref MDBX_UTTERLY_NOSYNC the `MDBX_SAFE_NOSYNC` flag disable similarly
	// flush system buffers to disk when committing a transaction. But there is a
	// huge difference in how are recycled the MVCC snapshots corresponding to
	// previous "steady" transactions (see below).
	//
	// With \ref MDBX_WRITEMAP the `MDBX_SAFE_NOSYNC` instructs MDBX to use
	// asynchronous mmap-flushes to disk. Asynchronous mmap-flushes means that
	// actually all writes will scheduled and performed by operation system on it
	// own manner, i.e. unordered. MDBX itself just notify operating system that
	// it would be nice to write data to disk, but no more.
	//
	// Depending on the platform and hardware, with `MDBX_SAFE_NOSYNC` you may get
	// a multiple increase of write performance, even 10 times or more.
	//
	// In contrast to \ref MDBX_UTTERLY_NOSYNC mode, with `MDBX_SAFE_NOSYNC` flag
	// MDBX will keeps untouched pages within B-tree of the last transaction
	// "steady" which was synced to disk completely. This has big implications for
	// both data durability and (unfortunately) performance:
	//  - a system crash can't corrupt the database, but you will lose the last
	//    transactions; because MDBX will rollback to last steady commit since it
	//    kept explicitly.
	//  - the last steady transaction makes an effect similar to "long-lived" read
	//    transaction (see above in the \ref restrictions section) since prevents
	//    reuse of pages freed by newer write transactions, thus the any data
	//    changes will be placed in newly allocated pages.
	//  - to avoid rapid database growth, the system will sync data and issue
	//    a steady commit-point to resume reuse pages, each time there is
	//    insufficient space and before increasing the size of the file on disk.
	//
	// In other words, with `MDBX_SAFE_NOSYNC` flag MDBX insures you from the
	// whole database corruption, at the cost increasing database size and/or
	// number of disk IOPs. So, `MDBX_SAFE_NOSYNC` flag could be used with
	// \ref mdbx_env_sync() as alternatively for batch committing or nested
	// transaction (in some cases). As well, auto-sync feature exposed by
	// \ref mdbx_env_set_syncbytes() and \ref mdbx_env_set_syncperiod() functions
	// could be very useful with `MDBX_SAFE_NOSYNC` flag.
	//
	// The number and volume of of disk IOPs with MDBX_SAFE_NOSYNC flag will
	// exactly the as without any no-sync flags. However, you should expect a
	// larger process's [work set](https://bit.ly/2kA2tFX) and significantly worse
	// a [locality of reference](https://bit.ly/2mbYq2J), due to the more
	// intensive allocation of previously unused pages and increase the size of
	// the database.
	//
	// `MDBX_SAFE_NOSYNC` flag may be changed at any time using
	// \ref mdbx_env_set_flags() or by passing to \ref mdbx_txn_begin() for
	// particular write transaction.
	EnvSafeNoSync = EnvFlags(C.MDBX_SAFE_NOSYNC)

	// EnvUtterlyNoSync Don't sync anything and wipe previous steady commits.
	//
	// Don't flush system buffers to disk when committing a transaction. This
	// optimization means a system crash can corrupt the database, if buffers are
	// not yet flushed to disk. Depending on the platform and hardware, with
	// `MDBX_UTTERLY_NOSYNC` you may get a multiple increase of write performance,
	// even 100 times or more.
	//
	// If the filesystem preserves write order (which is rare and never provided
	// unless explicitly noted) and the \ref MDBX_WRITEMAP and \ref
	// MDBX_LIFORECLAIM flags are not used, then a system crash can't corrupt the
	// database, but you can lose the last transactions, if at least one buffer is
	// not yet flushed to disk. The risk is governed by how often the system
	// flushes dirty buffers to disk and how often \ref mdbx_env_sync() is called.
	// So, transactions exhibit ACI (atomicity, consistency, isolation) properties
	// and only lose `D` (durability). I.e. database integrity is maintained, but
	// a system crash may undo the final transactions.
	//
	// Otherwise, if the filesystem not preserves write order (which is
	// typically) or \ref MDBX_WRITEMAP or \ref MDBX_LIFORECLAIM flags are used,
	// you should expect the corrupted database after a system crash.
	//
	// So, most important thing about `MDBX_UTTERLY_NOSYNC`:
	//  - a system crash immediately after commit the write transaction
	//    high likely lead to database corruption.
	//  - successful completion of mdbx_env_sync(force = true) after one or
	//    more committed transactions guarantees consistency and durability.
	//  - BUT by committing two or more transactions you back database into
	//    a weak state, in which a system crash may lead to database corruption!
	//    In case single transaction after mdbx_env_sync, you may lose transaction
	//    itself, but not a whole database.
	//
	// Nevertheless, `MDBX_UTTERLY_NOSYNC` provides "weak" durability in case
	// of an application crash (but no durability on system failure), and
	// therefore may be very useful in scenarios where data durability is
	// not required over a system failure (e.g for short-lived data), or if you
	// can take such risk.
	//
	// `MDBX_UTTERLY_NOSYNC` flag may be changed at any time using
	// \ref mdbx_env_set_flags(), but don't has effect if passed to
	// \ref mdbx_txn_begin() for particular write transaction. \see sync_modes
	EnvUtterlyNoSync = EnvFlags(C.MDBX_UTTERLY_NOSYNC)
)

type TxFlags uint32

const (
	// TxReadWrite Start read-write transaction.
	//
	// Only one write transaction may be active at a time. Writes are fully
	// serialized, which guarantees that writers can never deadlock.
	TxReadWrite = TxFlags(C.MDBX_TXN_READWRITE)

	// TxReadOnly Start read-only transaction.
	//
	// There can be multiple read-only transactions simultaneously that do not
	// block each other and a write transactions.
	TxReadOnly = TxFlags(C.MDBX_TXN_RDONLY)

	// TxReadOnlyPrepare Prepare but not start read-only transaction.
	//
	// Transaction will not be started immediately, but created transaction handle
	// will be ready for use with \ref mdbx_txn_renew(). This flag allows to
	// preallocate memory and assign a reader slot, thus avoiding these operations
	// at the next start of the transaction.
	TxReadOnlyPrepare = TxFlags(C.MDBX_TXN_RDONLY_PREPARE)

	// TxTry Do not block when starting a write transaction.
	TxTry = TxFlags(C.MDBX_TXN_TRY)

	// TxNoMetaSync Exactly the same as \ref MDBX_NOMETASYNC,
	// but for this transaction only
	TxNoMetaSync = TxFlags(C.MDBX_TXN_NOMETASYNC)

	// TxNoSync Exactly the same as \ref MDBX_SAFE_NOSYNC,
	// but for this transaction only
	TxNoSync = TxFlags(C.MDBX_TXN_NOSYNC)
)

type DBFlags uint32

const (
	DBDefaults = DBFlags(C.MDBX_DB_DEFAULTS)

	// DBReverseKey Use reverse string keys
	DBReverseKey = DBFlags(C.MDBX_REVERSEKEY)

	// DBDupSort Use sorted duplicates, i.e. allow multi-values
	DBDupSort = DBFlags(C.MDBX_DUPSORT)

	// DBIntegerKey Numeric keys in native byte order either uint32_t or uint64_t. The keys
	// must all be of the same size and must be aligned while passing as
	// arguments.
	DBIntegerKey = DBFlags(C.MDBX_INTEGERKEY)

	// DBDupFixed With \ref MDBX_DUPSORT; sorted dup items have fixed size
	DBDupFixed = DBFlags(C.MDBX_DUPFIXED)

	// DBIntegerGroup With \ref MDBX_DUPSORT and with \ref MDBX_DUPFIXED; dups are fixed size
	// \ref MDBX_INTEGERKEY -style integers. The data values must all be of the
	// same size and must be aligned while passing as arguments.
	DBIntegerGroup = DBFlags(C.MDBX_INTEGERDUP)

	// DBReverseDup With \ref MDBX_DUPSORT; use reverse string comparison
	DBReverseDup = DBFlags(C.MDBX_REVERSEDUP)

	// DBCreate Create DB if not already existing
	DBCreate = DBFlags(C.MDBX_CREATE)

	// DBAccede Opens an existing sub-database created with unknown flags.
	//
	// The `MDBX_DB_ACCEDE` flag is intend to open a existing sub-database which
	// was created with unknown flags (\ref MDBX_REVERSEKEY, \ref MDBX_DUPSORT,
	// \ref MDBX_INTEGERKEY, \ref MDBX_DUPFIXED, \ref MDBX_INTEGERDUP and
	// \ref MDBX_REVERSEDUP).
	//
	// In such cases, instead of returning the \ref MDBX_INCOMPATIBLE error, the
	// sub-database will be opened with flags which it was created, and then an
	// application could determine the actual flags by \ref mdbx_dbi_flags().
	DBAccede = DBFlags(C.MDBX_DB_ACCEDE)
)

type PutFlags uint32

const (
	// PutUpsert Upsertion by default (without any other flags)
	PutUpsert = PutFlags(C.MDBX_UPSERT)

	// PutNoOverwrite For insertion: Don't write if the key already exists.
	PutNoOverwrite = PutFlags(C.MDBX_NOOVERWRITE)

	// PutNoDupData Has effect only for \ref MDBX_DUPSORT databases.
	// For upsertion: don't write if the key-value pair already exist.
	// For deletion: remove all values for key.
	PutNoDupData = PutFlags(C.MDBX_NODUPDATA)

	// PutCurrent For upsertion: overwrite the current key/data pair.
	// MDBX allows this flag for \ref mdbx_put() for explicit overwrite/update
	// without insertion.
	// For deletion: remove only single entry at the current cursor position.
	PutCurrent = PutFlags(C.MDBX_CURRENT)

	// PutAllDups Has effect only for \ref MDBX_DUPSORT databases.
	// For deletion: remove all multi-values (aka duplicates) for given key.
	// For upsertion: replace all multi-values for given key with a new one.
	PutAllDups = PutFlags(C.MDBX_ALLDUPS)

	// PutReserve For upsertion: Just reserve space for data, don't copy it.
	// Return a pointer to the reserved space.
	PutReserve = PutFlags(C.MDBX_RESERVE)

	// PutAppend Data is being appended.
	// Don't split full pages, continue on a new instead.
	PutAppend = PutFlags(C.MDBX_APPEND)

	// PutAppendDup Has effect only for \ref MDBX_DUPSORT databases.
	// Duplicate data is being appended.
	// Don't split full pages, continue on a new instead.
	PutAppendDup = PutFlags(C.MDBX_APPENDDUP)

	// PutMultiple Only for \ref MDBX_DUPFIXED.
	// Store multiple data items in one call.
	PutMultiple = PutFlags(C.MDBX_MULTIPLE)
)

type CopyFlags uint32

const (
	CopyDefaults = CopyFlags(C.MDBX_CP_DEFAULTS)

	// CopyCompact Copy and compact: Omit free space from copy and renumber all
	// pages sequentially
	CopyCompact = CopyFlags(C.MDBX_CP_COMPACT)

	// CopyForceDynamicSize Force to make resizeable copy, i.e. dynamic size instead of fixed
	CopyForceDynamicSize = CopyFlags(C.MDBX_CP_FORCE_DYNAMIC_SIZE)
)

type CursorOp int32

const (
	// CursorFirst Position at first key/data item
	CursorFirst = CursorOp(C.MDBX_FIRST)

	// CursorFirstDup \ref MDBX_DUPSORT -only: Position at first data item of current key.
	CursorFirstDup = CursorOp(C.MDBX_FIRST_DUP)

	// CursorGetBoth \ref MDBX_DUPSORT -only: Position at key/data pair.
	CursorGetBoth = CursorOp(C.MDBX_GET_BOTH)

	// CursorGetBothRange \ref MDBX_DUPSORT -only: Position at given key and at first data greater
	// than or equal to specified data.
	CursorGetBothRange = CursorOp(C.MDBX_GET_BOTH_RANGE)

	// CursorGetCurrent Return key/data at current cursor position
	CursorGetCurrent = CursorOp(C.MDBX_GET_CURRENT)

	// CursorGetMultiple \ref MDBX_DUPFIXED -only: Return up to a page of duplicate data items
	// from current cursor position. Move cursor to prepare
	// for \ref MDBX_NEXT_MULTIPLE.
	CursorGetMultiple = CursorOp(C.MDBX_GET_MULTIPLE)

	// CursorLast Position at last key/data item
	CursorLast = CursorOp(C.MDBX_LAST)

	// CursorLastDup \ref MDBX_DUPSORT -only: Position at last data item of current key.
	CursorLastDup = CursorOp(C.MDBX_LAST_DUP)

	// CursorNext Position at next data item
	CursorNext = CursorOp(C.MDBX_NEXT)

	// CursorNextDup \ref MDBX_DUPSORT -only: Position at next data item of current key.
	CursorNextDup = CursorOp(C.MDBX_NEXT_DUP)

	// CursorNextMultiple \ref MDBX_DUPFIXED -only: Return up to a page of duplicate data items
	// from next cursor position. Move cursor to prepare
	// for `MDBX_NEXT_MULTIPLE`.
	CursorNextMultiple = CursorOp(C.MDBX_NEXT_MULTIPLE)

	// CursorNextNoDup Position at first data item of next key
	CursorNextNoDup = CursorOp(C.MDBX_NEXT_NODUP)

	// CursorPrev Position at previous data item
	CursorPrev = CursorOp(C.MDBX_PREV)

	// CursorPrevDup \ref MDBX_DUPSORT -only: Position at previous data item of current key.
	CursorPrevDup = CursorOp(C.MDBX_PREV_DUP)

	// CursorPrevNoDup Position at last data item of previous key
	CursorPrevNoDup = CursorOp(C.MDBX_PREV_NODUP)

	// CursorSet Position at specified key
	CursorSet = CursorOp(C.MDBX_SET)

	// CursorSetKey Position at specified key, return both key and data
	CursorSetKey = CursorOp(C.MDBX_SET_KEY)

	// CursorSetRange Position at first key greater than or equal to specified key.
	CursorSetRange = CursorOp(C.MDBX_SET_RANGE)

	// CursorPrevMultiple \ref MDBX_DUPFIXED -only: Position at previous page and return up to
	// a page of duplicate data items.
	CursorPrevMultiple = CursorOp(C.MDBX_PREV_MULTIPLE)

	// CursorSetLowerBound Position at first key-value pair greater than or equal to specified,
	// return both key and data, and the return code depends on a exact match.
	//
	// For non DUPSORT-ed collections this work the same to \ref MDBX_SET_RANGE,
	// but returns \ref MDBX_SUCCESS if key found exactly and
	// \ref MDBX_RESULT_TRUE if greater key was found.
	//
	// For DUPSORT-ed a data value is taken into account for duplicates,
	// i.e. for a pairs/tuples of a key and an each data value of duplicates.
	// Returns \ref MDBX_SUCCESS if key-value pair found exactly and
	// \ref MDBX_RESULT_TRUE if the next pair was returned.
	CursorSetLowerBound = CursorOp(C.MDBX_SET_LOWERBOUND)
)

type Opt int32

const (
	// OptMaxDB \brief Controls the maximum number of named databases for the environment.
	//
	// \details By default only unnamed key-value database could used and
	// appropriate value should set by `MDBX_opt_max_db` to using any more named
	// subDB(s). To reduce overhead, use the minimum sufficient value. This option
	// may only set after \ref mdbx_env_create() and before \ref mdbx_env_open().
	//
	// \see mdbx_env_set_maxdbs() \see mdbx_env_get_maxdbs()
	OptMaxDB = Opt(C.MDBX_opt_max_db)

	// OptMaxReaders \brief Defines the maximum number of threads/reader slots
	// for all processes interacting with the database.
	//
	// \details This defines the number of slots in the lock table that is used to
	// track readers in the the environment. The default is about 100 for 4K
	// system page size. Starting a read-only transaction normally ties a lock
	// table slot to the current thread until the environment closes or the thread
	// exits. If \ref MDBX_NOTLS is in use, \ref mdbx_txn_begin() instead ties the
	// slot to the \ref MDBX_txn object until it or the \ref MDBX_env object is
	// destroyed. This option may only set after \ref mdbx_env_create() and before
	// \ref mdbx_env_open(), and has an effect only when the database is opened by
	// the first process interacts with the database.
	//
	// \see mdbx_env_set_maxreaders() \see mdbx_env_get_maxreaders()
	OptMaxReaders = Opt(C.MDBX_opt_max_readers)

	// OptSyncBytes \brief Controls interprocess/shared threshold to force flush the data
	// buffers to disk, if \ref MDBX_SAFE_NOSYNC is used.
	//
	// \see mdbx_env_set_syncbytes() \see mdbx_env_get_syncbytes()
	OptSyncBytes = Opt(C.MDBX_opt_sync_bytes)

	// OptSyncPeriod \brief Controls interprocess/shared relative period since the last
	// unsteady commit to force flush the data buffers to disk,
	// if \ref MDBX_SAFE_NOSYNC is used.
	// \see mdbx_env_set_syncperiod() \see mdbx_env_get_syncperiod()
	OptSyncPeriod = Opt(C.MDBX_opt_sync_period)

	// OptRpAugmentLimit \brief Controls the in-process limit to grow a list of reclaimed/recycled
	// page's numbers for finding a sequence of contiguous pages for large data
	// items.
	//
	// \details A long values requires allocation of contiguous database pages.
	// To find such sequences, it may be necessary to accumulate very large lists,
	// especially when placing very long values (more than a megabyte) in a large
	// databases (several tens of gigabytes), which is much expensive in extreme
	// cases. This threshold allows you to avoid such costs by allocating new
	// pages at the end of the database (with its possible growth on disk),
	// instead of further accumulating/reclaiming Garbage Collection records.
	//
	// On the other hand, too small threshold will lead to unreasonable database
	// growth, or/and to the inability of put long values.
	//
	// The `MDBX_opt_rp_augment_limit` controls described limit for the current
	// process. Default is 262144, it is usually enough for most cases.
	OptRpAugmentLimit = Opt(C.MDBX_opt_rp_augment_limit)

	// OptLooseLimit \brief Controls the in-process limit to grow a cache of dirty
	// pages for reuse in the current transaction.
	//
	// \details A 'dirty page' refers to a page that has been updated in memory
	// only, the changes to a dirty page are not yet stored on disk.
	// To reduce overhead, it is reasonable to release not all such pages
	// immediately, but to leave some ones in cache for reuse in the current
	// transaction.
	//
	// The `MDBX_opt_loose_limit` allows you to set a limit for such cache inside
	// the current process. Should be in the range 0..255, default is 64.
	OptLooseLimit = Opt(C.MDBX_opt_loose_limit)

	// OptDpReserveLimit \brief Controls the in-process limit of a pre-allocated memory items
	// for dirty pages.
	//
	// \details A 'dirty page' refers to a page that has been updated in memory
	// only, the changes to a dirty page are not yet stored on disk.
	// Without \ref MDBX_WRITEMAP dirty pages are allocated from memory and
	// released when a transaction is committed. To reduce overhead, it is
	// reasonable to release not all ones, but to leave some allocations in
	// reserve for reuse in the next transaction(s).
	//
	// The `MDBX_opt_dp_reserve_limit` allows you to set a limit for such reserve
	// inside the current process. Default is 1024.
	OptDpReserveLimit = Opt(C.MDBX_opt_dp_reserve_limit)

	// OptTxnDpLimit \brief Controls the in-process limit of dirty pages
	// for a write transaction.
	//
	// \details A 'dirty page' refers to a page that has been updated in memory
	// only, the changes to a dirty page are not yet stored on disk.
	// Without \ref MDBX_WRITEMAP dirty pages are allocated from memory and will
	// be busy until are written to disk. Therefore for a large transactions is
	// reasonable to limit dirty pages collecting above an some threshold but
	// spill to disk instead.
	//
	// The `MDBX_opt_txn_dp_limit` controls described threshold for the current
	// process. Default is 65536, it is usually enough for most cases.
	OptTxnDpLimit = Opt(C.MDBX_opt_txn_dp_limit)

	// OptTxnDpInitial \brief Controls the in-process initial allocation size for dirty pages
	// list of a write transaction. Default is 1024.
	OptTxnDpInitial = Opt(C.MDBX_opt_txn_dp_initial)

	// OptSpillMaxDenomiator \brief Controls the in-process how maximal part of the dirty pages may be
	// spilled when necessary.
	//
	// \details The `MDBX_opt_spill_max_denominator` defines the denominator for
	// limiting from the top for part of the current dirty pages may be spilled
	// when the free room for a new dirty pages (i.e. distance to the
	// `MDBX_opt_txn_dp_limit` threshold) is not enough to perform requested
	// operation.
	// Exactly `max_pages_to_spill = dirty_pages - dirty_pages / N`,
	// where `N` is the value set by `MDBX_opt_spill_max_denominator`.
	//
	// Should be in the range 0..255, where zero means no limit, i.e. all dirty
	// pages could be spilled. Default is 8, i.e. no more than 7/8 of the current
	// dirty pages may be spilled when reached the condition described above.
	OptSpillMaxDenomiator = Opt(C.MDBX_opt_spill_max_denominator)

	// OptSpillMinDenomiator \brief Controls the in-process how minimal part of the dirty pages should
	// be spilled when necessary.
	//
	// \details The `MDBX_opt_spill_min_denominator` defines the denominator for
	// limiting from the bottom for part of the current dirty pages should be
	// spilled when the free room for a new dirty pages (i.e. distance to the
	// `MDBX_opt_txn_dp_limit` threshold) is not enough to perform requested
	// operation.
	// Exactly `min_pages_to_spill = dirty_pages / N`,
	// where `N` is the value set by `MDBX_opt_spill_min_denominator`.
	//
	// Should be in the range 0..255, where zero means no restriction at the
	// bottom. Default is 8, i.e. at least the 1/8 of the current dirty pages
	// should be spilled when reached the condition described above.
	OptSpillMinDenomiator = Opt(C.MDBX_opt_spill_min_denominator)

	// OptSpillParent4ChildDenominator \brief Controls the in-process how much of the parent transaction dirty
	// pages will be spilled while start each child transaction.
	//
	// \details The `MDBX_opt_spill_parent4child_denominator` defines the
	// denominator to determine how much of parent transaction dirty pages will be
	// spilled explicitly while start each child transaction.
	// Exactly `pages_to_spill = dirty_pages / N`,
	// where `N` is the value set by `MDBX_opt_spill_parent4child_denominator`.
	//
	// For a stack of nested transactions each dirty page could be spilled only
	// once, and parent's dirty pages couldn't be spilled while child
	// transaction(s) are running. Therefore a child transaction could reach
	// \ref MDBX_TXN_FULL when parent(s) transaction has  spilled too less (and
	// child reach the limit of dirty pages), either when parent(s) has spilled
	// too more (since child can't spill already spilled pages). So there is no
	// universal golden ratio.
	//
	// Should be in the range 0..255, where zero means no explicit spilling will
	// be performed during starting nested transactions.
	// Default is 0, i.e. by default no spilling performed during starting nested
	// transactions, that correspond historically behaviour.
	OptSpillParent4ChildDenominator = Opt(C.MDBX_opt_spill_parent4child_denominator)

	// OptMergeThreshold16Dot16Percent \brief Controls the in-process threshold of semi-empty pages merge.
	// \warning This is experimental option and subject for change or removal.
	// \details This option controls the in-process threshold of minimum page
	// fill, as used space of percentage of a page. Neighbour pages emptier than
	// this value are candidates for merging. The threshold value is specified
	// in 1/65536 of percent, which is equivalent to the 16-dot-16 fixed point
	// format. The specified value must be in the range from 12.5% (almost empty)
	// to 50% (half empty) which corresponds to the range from 8192 and to 32768
	// in units respectively.
	OptMergeThreshold16Dot16Percent = Opt(C.MDBX_opt_merge_threshold_16dot16_percent)
)

type DeleteMode int32

const (
	// DeleteModeJustDelete \brief Just delete the environment's files and directory if any.
	// \note On POSIX systems, processes already working with the database will
	// continue to work without interference until it close the environment.
	// \note On Windows, the behavior of `MDB_ENV_JUST_DELETE` is different
	// because the system does not support deleting files that are currently
	// memory mapped.
	DeleteModeJustDelete = DeleteMode(C.MDBX_ENV_JUST_DELETE)

	// DeleteModeEnsureUnused \brief Make sure that the environment is not being used by other
	// processes, or return an error otherwise.
	DeleteModeEnsureUnused = DeleteMode(C.MDBX_ENV_ENSURE_UNUSED)

	// DeleteModeWaitForUnused \brief Wait until other processes closes the environment before deletion.
	DeleteModeWaitForUnused = DeleteMode(C.MDBX_ENV_WAIT_FOR_UNUSED)
)

type DBIState uint32

const (
	DBIStateDirty = DBIState(C.MDBX_DBI_DIRTY) // DB was written in this txn
	DBIStateState = DBIState(C.MDBX_DBI_STALE) // Named-DB record is older than txnID
	DBIStateFresh = DBIState(C.MDBX_DBI_FRESH) // Named-DB handle opened in this txn
	DBIStateCreat = DBIState(C.MDBX_DBI_CREAT) // Named-DB handle created in this txn
)

// Delete \brief Delete the environment's files in a proper and multiprocess-safe way.
// \ingroup c_extra
//
// \param [in] pathname  The pathname for the database or the directory in which
//                       the database files reside.
//
// \param [in] mode      Special deletion mode for the environment. This
//                       parameter must be set to one of the values described
//                       above in the \ref MDBX_env_delete_mode_t section.
//
// \note The \ref MDBX_ENV_JUST_DELETE don't supported on Windows since system
// unable to delete a memory-mapped files.
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_RESULT_TRUE   No corresponding files or directories were found,
//                            so no deletion was performed.
func Delete(path string, mode DeleteMode) Error {
	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))
	return Error(C.mdbx_env_delete(p, (C.MDBX_env_delete_mode_t)(mode)))
}

//////////////////////////////////////////////////////////////////////////////////////////
// Env
//////////////////////////////////////////////////////////////////////////////////////////

type Env struct {
	env    *C.MDBX_env
	opened int64
	info   EnvInfo
	closed int64
	mu     sync.Mutex
}

// NewEnv \brief Create an MDBX environment instance.
// \ingroup c_opening
//
// This function allocates memory for a \ref MDBX_env structure. To release
// the allocated memory and discard the handle, call \ref mdbx_env_close().
// Before the handle may be used, it must be opened using \ref mdbx_env_open().
//
// Various other options may also need to be set before opening the handle,
// e.g. \ref mdbx_env_set_geometry(), \ref mdbx_env_set_maxreaders(),
// \ref mdbx_env_set_maxdbs(), depending on usage requirements.
//
// \param [out] penv  The address where the new handle will be stored.
//
// \returns a non-zero error value on failure and 0 on success.
func NewEnv() (*Env, Error) {
	env := &Env{}
	err := Error(C.mdbx_env_create((**C.MDBX_env)(unsafe.Pointer(&env.env))))
	if err != ErrSuccess {
		return nil, err
	}
	return env, err
}

// FD returns the open file descriptor (or Windows file handle) for the given
// environment.  An error is returned if the environment has not been
// successfully Opened (where C API just retruns an invalid handle).
//
// See mdbx_env_get_fd.
func (env *Env) FD() (uintptr, error) {
	// fdInvalid is the value -1 as a uintptr, which is used by MDBX in the
	// case that env has not been opened yet.  the strange construction is done
	// to avoid constant value overflow errors at compile time.
	const fdInvalid = ^uintptr(0)

	var mf C.mdbx_filehandle_t
	err := Error(C.mdbx_env_get_fd(env.env, &mf))
	//err := operrno("mdbx_env_get_fd", ret)
	if err != ErrSuccess {
		return 0, err
	}
	fd := uintptr(mf)

	if fd == fdInvalid {
		return 0, os.ErrClosed
	}
	return fd, nil
}

// ReaderList dumps the contents of the reader lock table as text.  Readers
// start on the second line as space-delimited fields described by the first
// line.
//
// See mdbx_reader_list.
//func (env *Env) ReaderList(fn func(string) error) error {
//	ctx, done := newMsgFunc(fn)
//	defer done()
//	if fn == nil {
//		ctx = 0
//	}
//
//	ret := C.mdbxgo_reader_list(env._env, C.size_t(ctx))
//	if ret >= 0 {
//		return nil
//	}
//	if ret < 0 && ctx != 0 {
//		err := ctx.get().err
//		if err != nil {
//			return err
//		}
//	}
//	return operrno("mdbx_reader_list", ret)
//}

// ReaderCheck clears stale entries from the reader lock table and returns the
// number of entries cleared.
//
// See mdbx_reader_check()
func (env *Env) ReaderCheck() (int, error) {
	var dead C.int
	err := Error(C.mdbx_reader_check(env.env, &dead))
	if err != ErrSuccess {
		return int(dead), err
	}
	return int(dead), nil
}

// Path returns the path argument passed to Open.  Path returns a non-nil error
// if env.Open() was not previously called.
//
// See mdbx_env_get_path.
func (env *Env) Path() (string, error) {
	var cpath *C.char
	err := Error(C.mdbx_env_get_path(env.env, &cpath))
	if err != ErrSuccess {
		return "", err
	}
	if cpath == nil {
		return "", os.ErrNotExist
	}
	return C.GoString(cpath), nil
}

// MaxKeySize returns the maximum allowed length for a key.
//
// See mdbx_env_get_maxkeysize.
func (env *Env) MaxKeySize() int {
	if env == nil {
		return int(C.mdbx_env_get_maxkeysize_ex(nil, 0))
	}
	return int(C.mdbx_env_get_maxkeysize_ex(env.env, 0))
}

// Close the environment and release the memory map.
// \ingroup c_opening
//
// Only a single thread may call this function. All transactions, databases,
// and cursors must already be closed before calling this function. Attempts
// to use any such handles after calling this function will cause a `SIGSEGV`.
// The environment handle will be freed and must not be used again after this
// call.
//
// \param [in] env        An environment handle returned by
//                        \ref mdbx_env_create().
//
// \param [in] dont_sync  A dont'sync flag, if non-zero the last checkpoint
//                        will be kept "as is" and may be still "weak" in the
//                        \ref MDBX_SAFE_NOSYNC or \ref MDBX_UTTERLY_NOSYNC
//                        modes. Such "weak" checkpoint will be ignored on
//                        opening next time, and transactions since the last
//                        non-weak checkpoint (meta-page update) will rolledback
//                        for consistency guarantee.
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_BUSY   The write transaction is running by other thread,
//                     in such case \ref MDBX_env instance has NOT be destroyed
//                     not released!
//                     \note If any OTHER error code was returned then
//                     given MDBX_env instance has been destroyed and released.
//
// \retval MDBX_EBADSIGN  Environment handle already closed or not valid,
//                        i.e. \ref mdbx_env_close() was already called for the
//                        `env` or was not created by \ref mdbx_env_create().
//
// \retval MDBX_PANIC  If \ref mdbx_env_close_ex() was called in the child
//                     process after `fork()`. In this case \ref MDBX_PANIC
//                     is expected, i.e. \ref MDBX_env instance was freed in
//                     proper manner.
//
// \retval MDBX_EIO    An error occurred during synchronization.
func (env *Env) Close(dontSync bool) Error {
	env.mu.Lock()
	defer env.mu.Unlock()
	if env.closed > 0 {
		return ErrSuccess
	}
	err := Error(C.mdbx_env_close_ex(env.env, (C.bool)(dontSync)))
	if err != ErrSuccess {
		return err
	}
	env.closed = time.Now().UnixNano()
	return err
}

// SetFlags Set environment flags.
// \ingroup c_settings
//
// This may be used to set some flags in addition to those from
// mdbx_env_open(), or to unset these flags.
// \see mdbx_env_get_flags()
//
// \note In contrast to LMDB, the MDBX serialize threads via mutex while
// changing the flags. Therefore this function will be blocked while a write
// transaction running by other thread, or \ref MDBX_BUSY will be returned if
// function called within a write transaction.
//
// \param [in] env      An environment handle returned
//                      by \ref mdbx_env_create().
// \param [in] flags    The \ref env_flags to change, bitwise OR'ed together.
// \param [in] onoff    A non-zero value sets the flags, zero clears them.
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_EINVAL  An invalid parameter was specified.
func (env *Env) SetFlags(flags EnvFlags, onoff bool) Error {
	return Error(C.mdbx_env_set_flags(env.env, (C.MDBX_env_flags_t)(flags), (C.bool)(onoff)))
}

// GetFlags Get environment flags.
// \ingroup c_statinfo
// \see mdbx_env_set_flags()
//
// \param [in] env     An environment handle returned by \ref mdbx_env_create().
// \param [out] flags  The address of an integer to store the flags.
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_EINVAL An invalid parameter was specified.
func (env *Env) GetFlags() (EnvFlags, Error) {
	flags := C.unsigned(0)
	err := Error(C.mdbx_env_get_flags(env.env, &flags))
	return EnvFlags(flags), err
}

// Copy an MDBX environment to the specified path, with options.
// \ingroup c_extra
//
// This function may be used to make a backup of an existing environment.
// No lockfile is created, since it gets recreated at need.
// \note This call can trigger significant file size growth if run in
// parallel with write transactions, because it employs a read-only
// transaction. See long-lived transactions under \ref restrictions section.
//
// \param [in] env    An environment handle returned by mdbx_env_create().
//                    It must have already been opened successfully.
// \param [in] dest   The pathname of a file in which the copy will reside.
//                    This file must not be already exist, but parent directory
//                    must be writable.
// \param [in] flags  Special options for this operation. This parameter must
//                    be set to 0 or by bitwise OR'ing together one or more
//                    of the values described here:
//
//  - \ref MDBX_CP_COMPACT
//      Perform compaction while copying: omit free pages and sequentially
//      renumber all pages in output. This option consumes little bit more
//      CPU for processing, but may running quickly than the default, on
//      account skipping free pages.
//
//  - \ref MDBX_CP_FORCE_DYNAMIC_SIZE
//      Force to make resizeable copy, i.e. dynamic size instead of fixed.
//
// \returns A non-zero error value on failure and 0 on success.
func (env *Env) Copy(dest string, flags CopyFlags) Error {
	if env.env == nil {
		return ErrSuccess
	}
	d := C.CString(dest)
	defer C.free(unsafe.Pointer(d))
	return Error(C.mdbx_env_copy(env.env, d, (C.MDBX_copy_flags_t)(flags)))
}

// Open \brief Open an environment instance.
// \ingroup c_opening
//
// Indifferently this function will fails or not, the \ref mdbx_env_close() must
// be called later to discard the \ref MDBX_env handle and release associated
// resources.
//
// \param [in] env       An environment handle returned
//                       by \ref mdbx_env_create()
//
// \param [in] pathname  The pathname for the database or the directory in which
//                       the database files reside. In the case of directory it
//                       must already exist and be writable.
//
// \param [in] flags     Special options for this environment. This parameter
//                       must be set to 0 or by bitwise OR'ing together one
//                       or more of the values described above in the
//                       \ref env_flags and \ref sync_modes sections.
//
// Flags set by mdbx_env_set_flags() are also used:
//  - \ref MDBX_NOSUBDIR, \ref MDBX_RDONLY, \ref MDBX_EXCLUSIVE,
//    \ref MDBX_WRITEMAP, \ref MDBX_NOTLS, \ref MDBX_NORDAHEAD,
//    \ref MDBX_NOMEMINIT, \ref MDBX_COALESCE, \ref MDBX_LIFORECLAIM.
//    See \ref env_flags section.
//
//  - \ref MDBX_NOMETASYNC, \ref MDBX_SAFE_NOSYNC, \ref MDBX_UTTERLY_NOSYNC.
//    See \ref sync_modes section.
//
// \note `MDB_NOLOCK` flag don't supported by MDBX,
//       try use \ref MDBX_EXCLUSIVE as a replacement.
//
// \note MDBX don't allow to mix processes with different \ref MDBX_SAFE_NOSYNC
//       flags on the same environment.
//       In such case \ref MDBX_INCOMPATIBLE will be returned.
//
// If the database is already exist and parameters specified early by
// \ref mdbx_env_set_geometry() are incompatible (i.e. for instance, different
// page size) then \ref mdbx_env_open() will return \ref MDBX_INCOMPATIBLE
// error.
//
// \param [in] mode   The UNIX permissions to set on created files.
//                    Zero value means to open existing, but do not create.
//
// \return A non-zero error value on failure and 0 on success,
//         some possible errors are:
// \retval MDBX_VERSION_MISMATCH The version of the MDBX library doesn't match
//                            the version that created the database environment.
// \retval MDBX_INVALID       The environment file headers are corrupted.
// \retval MDBX_ENOENT        The directory specified by the path parameter
//                            doesn't exist.
// \retval MDBX_EACCES        The user didn't have permission to access
//                            the environment files.
// \retval MDBX_EAGAIN        The environment was locked by another process.
// \retval MDBX_BUSY          The \ref MDBX_EXCLUSIVE flag was specified and the
//                            environment is in use by another process,
//                            or the current process tries to open environment
//                            more than once.
// \retval MDBX_INCOMPATIBLE  Environment is already opened by another process,
//                            but with different set of \ref MDBX_SAFE_NOSYNC,
//                            \ref MDBX_UTTERLY_NOSYNC flags.
//                            Or if the database is already exist and parameters
//                            specified early by \ref mdbx_env_set_geometry()
//                            are incompatible (i.e. different pagesize, etc).
//
// \retval MDBX_WANNA_RECOVERY The \ref MDBX_RDONLY flag was specified but
//                             read-write access is required to rollback
//                             inconsistent state after a system crash.
//
// \retval MDBX_TOO_LARGE      Database is too large for this process,
//                             i.e. 32-bit process tries to open >4Gb database.
func (env *Env) Open(path string, flags EnvFlags, mode os.FileMode) Error {
	if env.opened > 0 {
		return ErrSuccess
	}

	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))

	err := Error(C.mdbx_env_open(
		(*C.MDBX_env)(unsafe.Pointer(env.env)),
		p,
		(C.MDBX_env_flags_t)(flags),
		(C.mdbx_mode_t)(mode),
	))
	if err != ErrSuccess {
		return err
	}

	env.opened = time.Now().UnixNano()
	return err
}

type Geometry struct {
	env             uintptr
	SizeLower       uintptr
	SizeNow         uintptr
	SizeUpper       uintptr
	GrowthStep      uintptr
	ShrinkThreshold uintptr
	PageSize        uintptr
	err             Error
}

// SetGeometry Set all size-related parameters of environment, including page size
// and the min/max size of the memory map. \ingroup c_settings
//
// In contrast to LMDB, the MDBX provide automatic size management of an
// database according the given parameters, including shrinking and resizing
// on the fly. From user point of view all of these just working. Nevertheless,
// it is reasonable to know some details in order to make optimal decisions
// when choosing parameters.
//
// Both \ref mdbx_env_info_ex() and legacy \ref mdbx_env_info() are inapplicable
// to read-only opened environment.
//
// Both \ref mdbx_env_info_ex() and legacy \ref mdbx_env_info() could be called
// either before or after \ref mdbx_env_open(), either within the write
// transaction running by current thread or not:
//
//  - In case \ref mdbx_env_info_ex() or legacy \ref mdbx_env_info() was called
//    BEFORE \ref mdbx_env_open(), i.e. for closed environment, then the
//    specified parameters will be used for new database creation, or will be
//    applied during opening if database exists and no other process using it.
//
//    If the database is already exist, opened with \ref MDBX_EXCLUSIVE or not
//    used by any other process, and parameters specified by
//    \ref mdbx_env_set_geometry() are incompatible (i.e. for instance,
//    different page size) then \ref mdbx_env_open() will return
//    \ref MDBX_INCOMPATIBLE error.
//
//    In another way, if database will opened read-only or will used by other
//    process during calling \ref mdbx_env_open() that specified parameters will
//    silently discarded (open the database with \ref MDBX_EXCLUSIVE flag
//    to avoid this).
//
//  - In case \ref mdbx_env_info_ex() or legacy \ref mdbx_env_info() was called
//    after \ref mdbx_env_open() WITHIN the write transaction running by current
//    thread, then specified parameters will be applied as a part of write
//    transaction, i.e. will not be visible to any others processes until the
//    current write transaction has been committed by the current process.
//    However, if transaction will be aborted, then the database file will be
//    reverted to the previous size not immediately, but when a next transaction
//    will be committed or when the database will be opened next time.
//
//  - In case \ref mdbx_env_info_ex() or legacy \ref mdbx_env_info() was called
//    after \ref mdbx_env_open() but OUTSIDE a write transaction, then MDBX will
//    execute internal pseudo-transaction to apply new parameters (but only if
//    anything has been changed), and changes be visible to any others processes
//    immediately after succesful completion of function.
//
// Essentially a concept of "automatic size management" is simple and useful:
//  - There are the lower and upper bound of the database file size;
//  - There is the growth step by which the database file will be increased,
//    in case of lack of space.
//  - There is the threshold for unused space, beyond which the database file
//    will be shrunk.
//  - The size of the memory map is also the maximum size of the database.
//  - MDBX will automatically manage both the size of the database and the size
//    of memory map, according to the given parameters.
//
// So, there some considerations about choosing these parameters:
//  - The lower bound allows you to prevent database shrinking below some
//    rational size to avoid unnecessary resizing costs.
//  - The upper bound allows you to prevent database growth above some rational
//    size. Besides, the upper bound defines the linear address space
//    reservation in each process that opens the database. Therefore changing
//    the upper bound is costly and may be required reopening environment in
//    case of \ref MDBX_UNABLE_EXTEND_MAPSIZE errors, and so on. Therefore, this
//    value should be chosen reasonable as large as possible, to accommodate
//    future growth of the database.
//  - The growth step must be greater than zero to allow the database to grow,
//    but also reasonable not too small, since increasing the size by little
//    steps will result a large overhead.
//  - The shrink threshold must be greater than zero to allow the database
//    to shrink but also reasonable not too small (to avoid extra overhead) and
//    not less than growth step to avoid up-and-down flouncing.
//  - The current size (i.e. size_now argument) is an auxiliary parameter for
//    simulation legacy \ref mdbx_env_set_mapsize() and as workaround Windows
//    issues (see below).
//
// Unfortunately, Windows has is a several issues
// with resizing of memory-mapped file:
//  - Windows unable shrinking a memory-mapped file (i.e memory-mapped section)
//    in any way except unmapping file entirely and then map again. Moreover,
//    it is impossible in any way if a memory-mapped file is used more than
//    one process.
//  - Windows does not provide the usual API to augment a memory-mapped file
//    (that is, a memory-mapped partition), but only by using "Native API"
//    in an undocumented way.
//
// MDBX bypasses all Windows issues, but at a cost:
//  - Ability to resize database on the fly requires an additional lock
//    and release `SlimReadWriteLock during` each read-only transaction.
//  - During resize all in-process threads should be paused and then resumed.
//  - Shrinking of database file is performed only when it used by single
//    process, i.e. when a database closes by the last process or opened
//    by the first.
//  = Therefore, the size_now argument may be useful to set database size
//    by the first process which open a database, and thus avoid expensive
//    remapping further.
//
// For create a new database with particular parameters, including the page
// size, \ref mdbx_env_set_geometry() should be called after
// \ref mdbx_env_create() and before mdbx_env_open(). Once the database is
// created, the page size cannot be changed. If you do not specify all or some
// of the parameters, the corresponding default values will be used. For
// instance, the default for database size is 10485760 bytes.
//
// If the mapsize is increased by another process, MDBX silently and
// transparently adopt these changes at next transaction start. However,
// \ref mdbx_txn_begin() will return \ref MDBX_UNABLE_EXTEND_MAPSIZE if new
// mapping size could not be applied for current process (for instance if
// address space is busy).  Therefore, in the case of
// \ref MDBX_UNABLE_EXTEND_MAPSIZE error you need close and reopen the
// environment to resolve error.
//
// \note Actual values may be different than your have specified because of
// rounding to specified database page size, the system page size and/or the
// size of the system virtual memory management unit. You can get actual values
// by \ref mdbx_env_sync_ex() or see by using the tool `mdbx_chk` with the `-v`
// option.
//
// Legacy \ref mdbx_env_set_mapsize() correspond to calling
// \ref mdbx_env_set_geometry() with the arguments `size_lower`, `size_now`,
// `size_upper` equal to the `size` and `-1` (i.e. default) for all other
// parameters.
//
// \param [in] env         An environment handle returned
//                         by \ref mdbx_env_create()
//
// \param [in] size_lower  The lower bound of database size in bytes.
//                         Zero value means "minimal acceptable",
//                         and negative means "keep current or use default".
//
// \param [in] size_now    The size in bytes to setup the database size for
//                         now. Zero value means "minimal acceptable", and
//                         negative means "keep current or use default". So,
//                         it is recommended always pass -1 in this argument
//                         except some special cases.
//
// \param [in] size_upper The upper bound of database size in bytes.
//                        Zero value means "minimal acceptable",
//                        and negative means "keep current or use default".
//                        It is recommended to avoid change upper bound while
//                        database is used by other processes or threaded
//                        (i.e. just pass -1 in this argument except absolutely
//                        necessary). Otherwise you must be ready for
//                        \ref MDBX_UNABLE_EXTEND_MAPSIZE error(s), unexpected
//                        pauses during remapping and/or system errors like
//                        "address busy", and so on. In other words, there
//                        is no way to handle a growth of the upper bound
//                        robustly because there may be a lack of appropriate
//                        system resources (which are extremely volatile in
//                        a multi-process multi-threaded environment).
//
// \param [in] growth_step  The growth step in bytes, must be greater than
//                          zero to allow the database to grow. Negative value
//                          means "keep current or use default".
//
// \param [in] shrink_threshold  The shrink threshold in bytes, must be greater
//                               than zero to allow the database to shrink and
//                               greater than growth_step to avoid shrinking
//                               right after grow.
//                               Negative value means "keep current
//                               or use default". Default is 2*growth_step.
//
// \param [in] pagesize          The database page size for new database
//                               creation or -1 otherwise. Must be power of 2
//                               in the range between \ref MDBX_MIN_PAGESIZE and
//                               \ref MDBX_MAX_PAGESIZE. Zero value means
//                               "minimal acceptable", and negative means
//                               "keep current or use default".
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_EINVAL    An invalid parameter was specified,
//                        or the environment has an active write transaction.
// \retval MDBX_EPERM     Specific for Windows: Shrinking was disabled before
//                        and now it wanna be enabled, but there are reading
//                        threads that don't use the additional `SRWL` (that
//                        is required to avoid Windows issues).
// \retval MDBX_EACCESS   The environment opened in read-only.
// \retval MDBX_MAP_FULL  Specified size smaller than the space already
//                        consumed by the environment.
// \retval MDBX_TOO_LARGE Specified size is too large, i.e. too many pages for
//                        given size, or a 32-bit process requests too much
//                        bytes for the 32-bit address space.
func (env *Env) SetGeometry(args Geometry) Error {
	args.env = uintptr(unsafe.Pointer(env.env))
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_env_set_geometry), ptr, 0)
	return args.err
}

// GetOption \brief Gets the value of runtime options from an environment.
// \ingroup c_settings
//
// \param [in] env     An environment handle returned by \ref mdbx_env_create().
// \param [in] option  The option from \ref MDBX_option_t to get value of it.
// \param [out] pvalue The address where the option's value will be stored.
//
// \see MDBX_option_t
// \see mdbx_env_get_option()
// \returns A non-zero error value on failure and 0 on success.
func (env *Env) GetOption(option Opt) (uint64, Error) {
	value := uint64(0)
	err := Error(C.mdbx_env_get_option(
		(*C.MDBX_env)(unsafe.Pointer(env.env)),
		(C.MDBX_option_t)(option),
		(*C.uint64_t)(unsafe.Pointer(&value))),
	)
	return value, err
}

// SetOption \brief Sets the value of a runtime options for an environment.
// \ingroup c_settings
//
// \param [in] env     An environment handle returned by \ref mdbx_env_create().
// \param [in] option  The option from \ref MDBX_option_t to set value of it.
// \param [in] value   The value of option to be set.
//
// \see MDBX_option_t
// \see mdbx_env_get_option()
// \returns A non-zero error value on failure and 0 on success.
func (env *Env) SetOption(option Opt, value uint64) Error {
	return Error(C.mdbx_env_set_option(
		(*C.MDBX_env)(unsafe.Pointer(env.env)),
		(C.MDBX_option_t)(option),
		C.uint64_t(value)),
	)
}

type EnvInfo struct {
	Geo struct {
		Lower   uint64 // Lower limit for datafile size
		Upper   uint64 // Upper limit for datafile size
		Current uint64 // Current datafile size
		Shrink  uint64 // Shrink threshold for datafile
		Grow    uint64 // Growth step for datafile
	}
	MapSize               uint64 // Size of the data memory map
	LastPageNumber        uint64 // Number of the last used page
	RecentTxnID           uint64 // ID of the last committed transaction
	LatterReaderTxnID     uint64 // ID of the last reader transaction
	SelfLatterReaderTxnID uint64 // ID of the last reader transaction of caller process

	Meta0TxnID, MIMeta0Sign uint64
	Meta1TxnID, MIMeta1Sign uint64
	Meta2TxnID, MIMeta2Sign uint64

	MaxReaders  uint32 // Total reader slots in the environment
	NumReaders  uint32 // Max reader slots used in the environment
	DXBPageSize uint32 // Database pagesize
	SysPageSize uint32 // System pagesize

	// BootID A mostly unique ID that is regenerated on each boot.
	// As such it can be used to identify the local machine's current boot. MDBX
	// uses such when open the database to determine whether rollback required to
	// the last steady sync point or not. I.e. if current bootid is differ from the
	// value within a database then the system was rebooted and all changes since
	// last steady sync must be reverted for data integrity. Zeros mean that no
	// relevant information is available from the system.
	BootID struct {
		Current, Meta0, Meta1, Meta2 struct{ X, Y uint64 }
	}

	UnSyncVolume                   uint64 // Bytes not explicitly synchronized to disk.
	AutoSyncThreshold              uint64 // Current auto-sync threshold, see \ref mdbx_env_set_syncbytes().
	SinceSyncSeconds16Dot16        uint32 // Time since the last steady sync in 1/65536 of second
	AutoSyncPeriodSeconds16Dot16   uint32 // Current auto-sync period in 1/65536 of second, see \ref mdbx_env_set_syncperiod().
	SinceReaderCheckSeconds16Dot16 uint32 // Time since the last readers check in 1/65536 of second, see \ref mdbx_reader_check()
	Mode                           uint32 // Current environment mode. The same as \ref mdbx_env_get_flags() returns.

	// Statistics of page operations.
	// details Overall statistics of page operations of all (running, completed
	// and aborted) transactions in the current multi-process session (since the
	// first process opened the database after everyone had previously closed it).
	PGOpStat struct {
		Newly   uint64 // Quantity of a new pages added
		Cow     uint64 // Quantity of pages copied for update
		Clone   uint64 // Quantity of parent's dirty pages clones for nested transactions
		Split   uint64 // Page splits
		Merge   uint64 // Page merges
		Spill   uint64 // Quantity of spilled dirty pages
		UnSpill uint64 // Quantity of unspilled/reloaded pages
		Wops    uint64 // Number of explicit write operations (not a pages) to a disk
	}
}

// Sync Flush the environment data buffers to disk.
// \ingroup c_extra
//
// Unless the environment was opened with no-sync flags (\ref MDBX_NOMETASYNC,
// \ref MDBX_SAFE_NOSYNC and \ref MDBX_UTTERLY_NOSYNC), then
// data is always written an flushed to disk when \ref mdbx_txn_commit() is
// called. Otherwise \ref mdbx_env_sync() may be called to manually write and
// flush unsynced data to disk.
//
// Besides, \ref mdbx_env_sync_ex() with argument `force=false` may be used to
// provide polling mode for lazy/asynchronous sync in conjunction with
// \ref mdbx_env_set_syncbytes() and/or \ref mdbx_env_set_syncperiod().
//
// \note This call is not valid if the environment was opened with MDBX_RDONLY.
//
// \param [in] env      An environment handle returned by \ref mdbx_env_create()
// \param [in] force    If non-zero, force a flush. Otherwise, If force is
//                      zero, then will run in polling mode,
//                      i.e. it will check the thresholds that were
//                      set \ref mdbx_env_set_syncbytes()
//                      and/or \ref mdbx_env_set_syncperiod() and perform flush
//                      if at least one of the thresholds is reached.
//
// \param [in] nonblock Don't wait if write transaction
//                      is running by other thread.
//
// \returns A non-zero error value on failure and \ref MDBX_RESULT_TRUE or 0 on
//     success. The \ref MDBX_RESULT_TRUE means no data pending for flush
//     to disk, and 0 otherwise. Some possible errors are:
//
// \retval MDBX_EACCES   the environment is read-only.
// \retval MDBX_BUSY     the environment is used by other thread
//                       and `nonblock=true`.
// \retval MDBX_EINVAL   an invalid parameter was specified.
// \retval MDBX_EIO      an error occurred during synchronization.
func (env *Env) Sync(force, nonblock bool) Error {
	return Error(C.mdbx_env_sync_ex(env.env, (C.bool)(force), (C.bool)(nonblock)))
}

// CloseDBI Close a database handle. Normally unnecessary.
// \ingroup c_dbi
//
// Closing a database handle is not necessary, but lets \ref mdbx_dbi_open()
// reuse the handle value. Usually it's better to set a bigger
// \ref mdbx_env_set_maxdbs(), unless that value would be large.
//
// \note Use with care.
// This call is synchronized via mutex with \ref mdbx_dbi_close(), but NOT with
// other transactions running by other threads. The "next" version of libmdbx
// (\ref MithrilDB) will solve this issue.
//
// Handles should only be closed if no other threads are going to reference
// the database handle or one of its cursors any further. Do not close a handle
// if an existing transaction has modified its database. Doing so can cause
// misbehavior from database corruption to errors like \ref MDBX_BAD_DBI
// (since the DB name is gone).
//
// \param [in] env  An environment handle returned by \ref mdbx_env_create().
// \param [in] dbi  A database handle returned by \ref mdbx_dbi_open().
//
// \returns A non-zero error value on failure and 0 on success.
func (env *Env) CloseDBI(dbi DBI) Error {
	return Error(C.mdbx_dbi_close(env.env, (C.MDBX_dbi)(dbi)))
}

// GetMaxDBS Controls the maximum number of named databases for the environment.
//
// \details By default only unnamed key-value database could used and
// appropriate value should set by `MDBX_opt_max_db` to using any more named
// subDB(s). To reduce overhead, use the minimum sufficient value. This option
// may only set after \ref mdbx_env_create() and before \ref mdbx_env_open().
//
// \see mdbx_env_set_maxdbs() \see mdbx_env_get_maxdbs()
func (env *Env) GetMaxDBS() (uint64, Error) {
	return env.GetOption(OptMaxDB)
}

// SetMaxDBS Controls the maximum number of named databases for the environment.
//
// \details By default only unnamed key-value database could used and
// appropriate value should set by `MDBX_opt_max_db` to using any more named
// subDB(s). To reduce overhead, use the minimum sufficient value. This option
// may only set after \ref mdbx_env_create() and before \ref mdbx_env_open().
//
// \see mdbx_env_set_maxdbs() \see mdbx_env_get_maxdbs()
func (env *Env) SetMaxDBS(max uint16) Error {
	return env.SetOption(OptMaxDB, uint64(max))
}

// GetMaxReaders Defines the maximum number of threads/reader slots
// for all processes interacting with the database.
//
// \details This defines the number of slots in the lock table that is used to
// track readers in the the environment. The default is about 100 for 4K
// system page size. Starting a read-only transaction normally ties a lock
// table slot to the current thread until the environment closes or the thread
// exits. If \ref MDBX_NOTLS is in use, \ref mdbx_txn_begin() instead ties the
// slot to the \ref MDBX_txn object until it or the \ref MDBX_env object is
// destroyed. This option may only set after \ref mdbx_env_create() and before
// \ref mdbx_env_open(), and has an effect only when the database is opened by
// the first process interacts with the database.
//
// \see mdbx_env_set_maxreaders() \see mdbx_env_get_maxreaders()
func (env *Env) GetMaxReaders() (uint64, Error) {
	return env.GetOption(OptMaxReaders)
}

// SetMaxReaders Defines the maximum number of threads/reader slots
// for all processes interacting with the database.
//
// \details This defines the number of slots in the lock table that is used to
// track readers in the the environment. The default is about 100 for 4K
// system page size. Starting a read-only transaction normally ties a lock
// table slot to the current thread until the environment closes or the thread
// exits. If \ref MDBX_NOTLS is in use, \ref mdbx_txn_begin() instead ties the
// slot to the \ref MDBX_txn object until it or the \ref MDBX_env object is
// destroyed. This option may only set after \ref mdbx_env_create() and before
// \ref mdbx_env_open(), and has an effect only when the database is opened by
// the first process interacts with the database.
//
// \see mdbx_env_set_maxreaders() \see mdbx_env_get_maxreaders()
func (env *Env) SetMaxReaders(max uint64) Error {
	return env.SetOption(OptMaxReaders, max)
}

// GetSyncBytes Controls interprocess/shared threshold to force flush the data
// buffers to disk, if \ref MDBX_SAFE_NOSYNC is used.
//
// \see mdbx_env_set_syncbytes() \see mdbx_env_get_syncbytes()
func (env *Env) GetSyncBytes() (uint64, Error) {
	return env.GetOption(OptSyncBytes)
}

// SetSyncBytes Controls interprocess/shared threshold to force flush the data
// buffers to disk, if \ref MDBX_SAFE_NOSYNC is used.
//
// \see mdbx_env_set_syncbytes() \see mdbx_env_get_syncbytes()
func (env *Env) SetSyncBytes(bytes uint64) Error {
	return env.SetOption(OptSyncBytes, bytes)
}

// GetSyncPeriod Controls interprocess/shared relative period since the last
// unsteady commit to force flush the data buffers to disk,
// if \ref MDBX_SAFE_NOSYNC is used.
// \see mdbx_env_set_syncperiod() \see mdbx_env_get_syncperiod()
func (env *Env) GetSyncPeriod() (uint64, Error) {
	return env.GetOption(OptSyncPeriod)
}

// SetSyncPeriod Controls interprocess/shared relative period since the last
// unsteady commit to force flush the data buffers to disk,
// if \ref MDBX_SAFE_NOSYNC is used.
// \see mdbx_env_set_syncperiod() \see mdbx_env_get_syncperiod()
func (env *Env) SetSyncPeriod(period uint64) Error {
	return env.SetOption(OptSyncPeriod, period)
}

// GetRPAugmentLimit Controls the in-process limit to grow a list of reclaimed/recycled
// page's numbers for finding a sequence of contiguous pages for large data
// items.
//
// \details A long values requires allocation of contiguous database pages.
// To find such sequences, it may be necessary to accumulate very large lists,
// especially when placing very long values (more than a megabyte) in a large
// databases (several tens of gigabytes), which is much expensive in extreme
// cases. This threshold allows you to avoid such costs by allocating new
// pages at the end of the database (with its possible growth on disk),
// instead of further accumulating/reclaiming Garbage Collection records.
//
// On the other hand, too small threshold will lead to unreasonable database
// growth, or/and to the inability of put long values.
//
// The `MDBX_opt_rp_augment_limit` controls described limit for the current
// process. Default is 262144, it is usually enough for most cases.
func (env *Env) GetRPAugmentLimit() (uint64, Error) {
	return env.GetOption(OptRpAugmentLimit)
}

// SetRPAugmentLimit Controls the in-process limit to grow a list of reclaimed/recycled
// page's numbers for finding a sequence of contiguous pages for large data
// items.
//
// \details A long values requires allocation of contiguous database pages.
// To find such sequences, it may be necessary to accumulate very large lists,
// especially when placing very long values (more than a megabyte) in a large
// databases (several tens of gigabytes), which is much expensive in extreme
// cases. This threshold allows you to avoid such costs by allocating new
// pages at the end of the database (with its possible growth on disk),
// instead of further accumulating/reclaiming Garbage Collection records.
//
// On the other hand, too small threshold will lead to unreasonable database
// growth, or/and to the inability of put long values.
//
// The `MDBX_opt_rp_augment_limit` controls described limit for the current
// process. Default is 262144, it is usually enough for most cases.
func (env *Env) SetRPAugmentLimit(limit uint64) Error {
	return env.SetOption(OptRpAugmentLimit, limit)
}

// GetLooseLimit Controls the in-process limit to grow a cache of dirty
// pages for reuse in the current transaction.
//
// \details A 'dirty page' refers to a page that has been updated in memory
// only, the changes to a dirty page are not yet stored on disk.
// To reduce overhead, it is reasonable to release not all such pages
// immediately, but to leave some ones in cache for reuse in the current
// transaction.
//
// The `MDBX_opt_loose_limit` allows you to set a limit for such cache inside
// the current process. Should be in the range 0..255, default is 64.
func (env *Env) GetLooseLimit() (uint64, Error) {
	return env.GetOption(OptLooseLimit)
}

// SetLooseLimit Controls the in-process limit to grow a cache of dirty
// pages for reuse in the current transaction.
//
// \details A 'dirty page' refers to a page that has been updated in memory
// only, the changes to a dirty page are not yet stored on disk.
// To reduce overhead, it is reasonable to release not all such pages
// immediately, but to leave some ones in cache for reuse in the current
// transaction.
//
// The `MDBX_opt_loose_limit` allows you to set a limit for such cache inside
// the current process. Should be in the range 0..255, default is 64.
func (env *Env) SetLooseLimit(limit uint64) Error {
	return env.SetOption(OptLooseLimit, limit)
}

// GetDPReserveLimit Controls the in-process limit of a pre-allocated memory items
// for dirty pages.
//
// \details A 'dirty page' refers to a page that has been updated in memory
// only, the changes to a dirty page are not yet stored on disk.
// Without \ref MDBX_WRITEMAP dirty pages are allocated from memory and
// released when a transaction is committed. To reduce overhead, it is
// reasonable to release not all ones, but to leave some allocations in
// reserve for reuse in the next transaction(s).
//
// The `MDBX_opt_dp_reserve_limit` allows you to set a limit for such reserve
// inside the current process. Default is 1024.
func (env *Env) GetDPReserveLimit() (uint64, Error) {
	return env.GetOption(OptDpReserveLimit)
}

// SetDPReserveLimit Controls the in-process limit of a pre-allocated memory items
// for dirty pages.
//
// \details A 'dirty page' refers to a page that has been updated in memory
// only, the changes to a dirty page are not yet stored on disk.
// Without \ref MDBX_WRITEMAP dirty pages are allocated from memory and
// released when a transaction is committed. To reduce overhead, it is
// reasonable to release not all ones, but to leave some allocations in
// reserve for reuse in the next transaction(s).
//
// The `MDBX_opt_dp_reserve_limit` allows you to set a limit for such reserve
// inside the current process. Default is 1024.
func (env *Env) SetDPReserveLimit(limit uint64) Error {
	return env.SetOption(OptDpReserveLimit, limit)
}

// GetTxDPLimit Controls the in-process limit of dirty pages
// for a write transaction.
//
// \details A 'dirty page' refers to a page that has been updated in memory
// only, the changes to a dirty page are not yet stored on disk.
// Without \ref MDBX_WRITEMAP dirty pages are allocated from memory and will
// be busy until are written to disk. Therefore for a large transactions is
// reasonable to limit dirty pages collecting above an some threshold but
// spill to disk instead.
//
// The `MDBX_opt_txn_dp_limit` controls described threshold for the current
// process. Default is 65536, it is usually enough for most cases.
func (env *Env) GetTxDPLimit() (uint64, Error) {
	return env.GetOption(OptTxnDpLimit)
}

// SetTxDPLimit Controls the in-process limit of dirty pages
// for a write transaction.
//
// \details A 'dirty page' refers to a page that has been updated in memory
// only, the changes to a dirty page are not yet stored on disk.
// Without \ref MDBX_WRITEMAP dirty pages are allocated from memory and will
// be busy until are written to disk. Therefore for a large transactions is
// reasonable to limit dirty pages collecting above an some threshold but
// spill to disk instead.
//
// The `MDBX_opt_txn_dp_limit` controls described threshold for the current
// process. Default is 65536, it is usually enough for most cases.
func (env *Env) SetTxDPLimit(limit uint64) Error {
	return env.SetOption(OptTxnDpLimit, limit)
}

// GetTxDPInitial Controls the in-process initial allocation size for dirty pages
// list of a write transaction. Default is 1024.
func (env *Env) GetTxDPInitial() (uint64, Error) {
	return env.GetOption(OptTxnDpInitial)
}

// SetTxDPInitial Controls the in-process initial allocation size for dirty pages
// list of a write transaction. Default is 1024.
func (env *Env) SetTxDPInitial(initial uint64) Error {
	return env.SetOption(OptTxnDpInitial, initial)
}

// GetSpillMinDenominator Controls the in-process how minimal part of the dirty pages should
// be spilled when necessary.
//
// \details The `MDBX_opt_spill_min_denominator` defines the denominator for
// limiting from the bottom for part of the current dirty pages should be
// spilled when the free room for a new dirty pages (i.e. distance to the
// `MDBX_opt_txn_dp_limit` threshold) is not enough to perform requested
// operation.
// Exactly `min_pages_to_spill = dirty_pages / N`,
// where `N` is the value set by `MDBX_opt_spill_min_denominator`.
//
// Should be in the range 0..255, where zero means no restriction at the
// bottom. Default is 8, i.e. at least the 1/8 of the current dirty pages
// should be spilled when reached the condition described above.
func (env *Env) GetSpillMinDenominator() (uint64, Error) {
	return env.GetOption(OptSpillMinDenomiator)
}

// SetSpillMinDenominator Controls the in-process how minimal part of the dirty pages should
// be spilled when necessary.
//
// \details The `MDBX_opt_spill_min_denominator` defines the denominator for
// limiting from the bottom for part of the current dirty pages should be
// spilled when the free room for a new dirty pages (i.e. distance to the
// `MDBX_opt_txn_dp_limit` threshold) is not enough to perform requested
// operation.
// Exactly `min_pages_to_spill = dirty_pages / N`,
// where `N` is the value set by `MDBX_opt_spill_min_denominator`.
//
// Should be in the range 0..255, where zero means no restriction at the
// bottom. Default is 8, i.e. at least the 1/8 of the current dirty pages
// should be spilled when reached the condition described above.
func (env *Env) SetSpillMinDenominator(min uint64) Error {
	return env.SetOption(OptSpillMinDenomiator, min)
}

// GetSpillMaxDenominator Controls the in-process how maximal part of the dirty pages may be
// spilled when necessary.
//
// \details The `MDBX_opt_spill_max_denominator` defines the denominator for
// limiting from the top for part of the current dirty pages may be spilled
// when the free room for a new dirty pages (i.e. distance to the
// `MDBX_opt_txn_dp_limit` threshold) is not enough to perform requested
// operation.
// Exactly `max_pages_to_spill = dirty_pages - dirty_pages / N`,
// where `N` is the value set by `MDBX_opt_spill_max_denominator`.
//
// Should be in the range 0..255, where zero means no limit, i.e. all dirty
// pages could be spilled. Default is 8, i.e. no more than 7/8 of the current
// dirty pages may be spilled when reached the condition described above.
func (env *Env) GetSpillMaxDenominator() (uint64, Error) {
	return env.GetOption(OptSpillMaxDenomiator)
}

// SetSpillMaxDenominator Controls the in-process how maximal part of the dirty pages may be
// spilled when necessary.
//
// \details The `MDBX_opt_spill_max_denominator` defines the denominator for
// limiting from the top for part of the current dirty pages may be spilled
// when the free room for a new dirty pages (i.e. distance to the
// `MDBX_opt_txn_dp_limit` threshold) is not enough to perform requested
// operation.
// Exactly `max_pages_to_spill = dirty_pages - dirty_pages / N`,
// where `N` is the value set by `MDBX_opt_spill_max_denominator`.
//
// Should be in the range 0..255, where zero means no limit, i.e. all dirty
// pages could be spilled. Default is 8, i.e. no more than 7/8 of the current
// dirty pages may be spilled when reached the condition described above.
func (env *Env) SetSpillMaxDenominator(max uint64) Error {
	return env.SetOption(OptSpillMaxDenomiator, max)
}

// GetSpillParent4ChildDeominator Controls the in-process how much of the parent transaction dirty
// pages will be spilled while start each child transaction.
//
// \details The `MDBX_opt_spill_parent4child_denominator` defines the
// denominator to determine how much of parent transaction dirty pages will be
// spilled explicitly while start each child transaction.
// Exactly `pages_to_spill = dirty_pages / N`,
// where `N` is the value set by `MDBX_opt_spill_parent4child_denominator`.
//
// For a stack of nested transactions each dirty page could be spilled only
// once, and parent's dirty pages couldn't be spilled while child
// transaction(s) are running. Therefore a child transaction could reach
// \ref MDBX_TXN_FULL when parent(s) transaction has  spilled too less (and
// child reach the limit of dirty pages), either when parent(s) has spilled
// too more (since child can't spill already spilled pages). So there is no
// universal golden ratio.
//
// Should be in the range 0..255, where zero means no explicit spilling will
// be performed during starting nested transactions.
// Default is 0, i.e. by default no spilling performed during starting nested
// transactions, that correspond historically behaviour.
func (env *Env) GetSpillParent4ChildDeominator() (uint64, Error) {
	return env.GetOption(OptSpillParent4ChildDenominator)
}

// SetSpillParent4ChildDeominator Controls the in-process how much of the parent transaction dirty
// pages will be spilled while start each child transaction.
//
// \details The `MDBX_opt_spill_parent4child_denominator` defines the
// denominator to determine how much of parent transaction dirty pages will be
// spilled explicitly while start each child transaction.
// Exactly `pages_to_spill = dirty_pages / N`,
// where `N` is the value set by `MDBX_opt_spill_parent4child_denominator`.
//
// For a stack of nested transactions each dirty page could be spilled only
// once, and parent's dirty pages couldn't be spilled while child
// transaction(s) are running. Therefore a child transaction could reach
// \ref MDBX_TXN_FULL when parent(s) transaction has  spilled too less (and
// child reach the limit of dirty pages), either when parent(s) has spilled
// too more (since child can't spill already spilled pages). So there is no
// universal golden ratio.
//
// Should be in the range 0..255, where zero means no explicit spilling will
// be performed during starting nested transactions.
// Default is 0, i.e. by default no spilling performed during starting nested
// transactions, that correspond historically behaviour.
func (env *Env) SetSpillParent4ChildDeominator(value uint64) Error {
	return env.SetOption(OptSpillParent4ChildDenominator, value)
}

// GetMergeThreshold16Dot16Percent Controls the in-process threshold of semi-empty pages merge.
// \warning This is experimental option and subject for change or removal.
// \details This option controls the in-process threshold of minimum page
// fill, as used space of percentage of a page. Neighbour pages emptier than
// this value are candidates for merging. The threshold value is specified
// in 1/65536 of percent, which is equivalent to the 16-dot-16 fixed point
// format. The specified value must be in the range from 12.5% (almost empty)
// to 50% (half empty) which corresponds to the range from 8192 and to 32768
// in units respectively.
func (env *Env) GetMergeThreshold16Dot16Percent() (uint64, Error) {
	return env.GetOption(OptMergeThreshold16Dot16Percent)
}

// SetMergeThreshold16Dot16Percent Controls the in-process threshold of semi-empty pages merge.
// \warning This is experimental option and subject for change or removal.
// \details This option controls the in-process threshold of minimum page
// fill, as used space of percentage of a page. Neighbour pages emptier than
// this value are candidates for merging. The threshold value is specified
// in 1/65536 of percent, which is equivalent to the 16-dot-16 fixed point
// format. The specified value must be in the range from 12.5% (almost empty)
// to 50% (half empty) which corresponds to the range from 8192 and to 32768
// in units respectively.
func (env *Env) SetMergeThreshold16Dot16Percent(percent uint64) Error {
	return env.SetOption(OptMergeThreshold16Dot16Percent, percent)
}

//////////////////////////////////////////////////////////////////////////////////////////
// DBI
//////////////////////////////////////////////////////////////////////////////////////////

type DBI uint32

//////////////////////////////////////////////////////////////////////////////////////////
// Val
//////////////////////////////////////////////////////////////////////////////////////////

type Val syscall.Iovec

func (v *Val) String() string {
	b := make([]byte, v.Len)
	copy(b, *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(v.Base)),
		Len:  int(v.Len),
		Cap:  int(v.Len),
	})))
	return string(b)
}

func (v *Val) UnsafeString() string {
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: uintptr(unsafe.Pointer(v.Base)),
		Len:  int(v.Len),
	}))
}

func (v *Val) Bytes() []byte {
	b := make([]byte, v.Len)
	copy(b, *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(v.Base)),
		Len:  int(v.Len),
		Cap:  int(v.Len),
	})))
	return b
}

func (v *Val) UnsafeBytes() []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(v.Base)),
		Len:  int(v.Len),
		Cap:  int(v.Len),
	}))
}

func (v *Val) Copy(dst []byte) []byte {
	src := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(v.Base)),
		Len:  int(v.Len),
		Cap:  int(v.Len),
	}))
	if cap(dst) >= int(v.Len) {
		dst = dst[0:v.Len]
		copy(dst, src)
		return dst
	}
	dst = make([]byte, v.Len)
	copy(dst, src)
	return dst
}

func U8(v *uint8) Val {
	return Val{
		Base: v,
		Len:  1,
	}
}

func I8(v *int8) Val {
	return Val{
		Base: (*byte)(unsafe.Pointer(v)),
		Len:  1,
	}
}

func U16(v *uint16) Val {
	return Val{
		Base: (*byte)(unsafe.Pointer(v)),
		Len:  2,
	}
}

func I16(v *int16) Val {
	return Val{
		Base: (*byte)(unsafe.Pointer(v)),
		Len:  2,
	}
}

func U32(v *uint32) Val {
	return Val{
		Base: (*byte)(unsafe.Pointer(v)),
		Len:  4,
	}
}

func I32(v *int32) Val {
	return Val{
		Base: (*byte)(unsafe.Pointer(v)),
		Len:  4,
	}
}

func F32(v *float32) Val {
	return Val{
		Base: (*byte)(unsafe.Pointer(v)),
		Len:  4,
	}
}

func U64(v *uint64) Val {
	return Val{
		Base: (*byte)(unsafe.Pointer(v)),
		Len:  8,
	}
}

func I64(v *int64) Val {
	return Val{
		Base: (*byte)(unsafe.Pointer(v)),
		Len:  8,
	}
}

func F64(v *float64) Val {
	return Val{
		Base: (*byte)(unsafe.Pointer(v)),
		Len:  8,
	}
}

func Bytes(b *[]byte) Val {
	return Val{
		Base: &(*b)[0],
		Len:  uint64(len(*b)),
	}
}

func String(s *string) Val {
	h := *(*reflect.StringHeader)(unsafe.Pointer(s))
	return Val{
		Base: (*byte)(unsafe.Pointer(h.Data)),
		Len:  uint64(h.Len),
	}
}

func StringConst(s string) Val {
	h := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	return Val{
		Base: (*byte)(unsafe.Pointer(h.Data)),
		Len:  uint64(h.Len),
	}
}

func (v *Val) I8() int8 {
	if v.Len < 1 {
		return 0
	}
	return *(*int8)(unsafe.Pointer(v.Base))
}

func (v *Val) U8() uint8 {
	if v.Len < 1 {
		return 0
	}
	return *v.Base
}

func (v *Val) I16() int16 {
	if v.Len < 2 {
		return 0
	}
	return *(*int16)(unsafe.Pointer(v.Base))
}

func (v *Val) U16() uint16 {
	if v.Len < 2 {
		return 0
	}
	return *(*uint16)(unsafe.Pointer(v.Base))
}

func (v *Val) I32() int32 {
	if v.Len < 4 {
		return 0
	}
	return *(*int32)(unsafe.Pointer(v.Base))
}

func (v *Val) U32() uint32 {
	if v.Len < 4 {
		return 0
	}
	return *(*uint32)(unsafe.Pointer(v.Base))
}

func (v *Val) I64() int64 {
	if v.Len < 8 {
		return 0
	}
	return *(*int64)(unsafe.Pointer(v.Base))
}

func (v *Val) U64() uint64 {
	if v.Len < 8 {
		return 0
	}
	return *(*uint64)(unsafe.Pointer(v.Base))
}

func (v *Val) F32() float32 {
	if v.Len < 4 {
		return 0
	}
	return *(*float32)(unsafe.Pointer(v.Base))
}

func (v *Val) F64() float64 {
	if v.Len < 8 {
		return 0
	}
	return *(*float64)(unsafe.Pointer(v.Base))
}

//////////////////////////////////////////////////////////////////////////////////////////
// Tx
//////////////////////////////////////////////////////////////////////////////////////////

type Tx struct {
	env       *Env
	txn       *C.MDBX_txn
	shared    bool
	reset     bool
	aborted   bool
	committed bool
}

func NewTransaction(env *Env) *Tx {
	txn := &Tx{}
	txn.env = env
	txn.shared = true
	return txn
}

func (tx *Tx) IsReset() bool {
	return tx.reset
}

func (tx *Tx) IsAborted() bool {
	return tx.aborted
}

func (tx *Tx) IsCommitted() bool {
	return tx.committed
}

func (env *Env) Begin(txn *Tx, flags TxFlags) Error {
	txn.env = env
	txn.txn = nil
	txn.reset = false
	txn.aborted = false
	txn.committed = false
	args := struct {
		env     uintptr
		parent  uintptr
		txn     uintptr
		context uintptr
		flags   TxFlags
		result  Error
	}{
		env:    uintptr(unsafe.Pointer(env.env)),
		parent: 0,
		txn:    uintptr(unsafe.Pointer(&txn.txn)),
		flags:  flags,
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_txn_begin_ex), ptr, 0)
	return args.result
}

// TxInfo Information about the transaction
type TxInfo struct {
	// The ID of the transaction. For a READ-ONLY transaction, this corresponds to the snapshot being read.
	ID uint64

	// For READ-ONLY transaction: the lag from a recent MVCC-snapshot, i.e. the
	// number of committed transaction since read transaction started.
	// For WRITE transaction (provided if `scan_rlt=true`): the lag of the oldest
	// reader from current transaction (i.e. at least 1 if any reader running).
	ReaderLag uint64

	// Used space by this transaction, i.e. corresponding to the last used database page.
	SpaceUsed uint64

	// Current size of database file.
	SpaceLimitSoft uint64

	// Upper bound for size the database file, i.e. the value `size_upper`
	// argument of the appropriate call of \ref mdbx_env_set_geometry().
	SpaceLimitHard uint64

	// For READ-ONLY transaction: The total size of the database pages that were
	// retired by committed write transactions after the reader's MVCC-snapshot,
	// i.e. the space which would be freed after the Reader releases the
	// MVCC-snapshot for reuse by completion read transaction.
	//
	// For WRITE transaction: The summarized size of the database pages that were
	// retired for now due Copy-On-Write during this transaction.
	SpaceRetired uint64

	// For READ-ONLY transaction: the space available for writer(s) and that
	// must be exhausted for reason to call the Handle-Slow-Readers callback for
	// this read transaction.
	//
	// For WRITE transaction: the space inside transaction
	// that left to `MDBX_TXN_FULL` error.
	SpaceLeftover uint64

	// For READ-ONLY transaction (provided if `scan_rlt=true`): The space that
	// actually become available for reuse when only this transaction will be finished.
	//
	// For WRITE transaction: The summarized size of the dirty database
	// pages that generated during this transaction.
	SpaceDirty uint64
}

// Info Return information about the MDBX transaction.
// \ingroup c_statinfo
//
// \param [in] txn        A transaction handle returned by \ref mdbx_txn_begin()
// \param [out] info      The address of an \ref MDBX_txn_info structure
//                        where the information will be copied.
// \param [in] scan_rlt   The boolean flag controls the scan of the read lock
//                        table to provide complete information. Such scan
//                        is relatively expensive and you can avoid it
//                        if corresponding fields are not needed.
//                        See description of \ref MDBX_txn_info.
//
// \returns A non-zero error value on failure and 0 on success.
func (tx *Tx) Info(info *TxInfo) Error {
	args := struct {
		txn     uintptr
		info    uintptr
		scanRlt int32
		result  Error
	}{
		txn:  uintptr(unsafe.Pointer(tx.txn)),
		info: uintptr(unsafe.Pointer(info)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_txn_info), ptr, 0)
	return args.result
}

// Flags Return the transaction's flags.
// \ingroup c_transactions
//
// This returns the flags associated with this transaction.
//
// \param [in] txn  A transaction handle returned by \ref mdbx_txn_begin().
//
// \returns A transaction flags, valid if input is an valid transaction,
//          otherwise -1.
func (tx *Tx) Flags() int32 {
	args := struct {
		txn   uintptr
		flags int32
	}{
		txn: uintptr(unsafe.Pointer(tx.txn)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_txn_flags), ptr, 0)
	return args.flags
}

// ID Return the transaction's ID.
// \ingroup c_statinfo
//
// This returns the identifier associated with this transaction. For a
// read-only transaction, this corresponds to the snapshot being read;
// concurrent readers will frequently have the same transaction ID.
//
// \param [in] txn  A transaction handle returned by \ref mdbx_txn_begin().
//
// \returns A transaction ID, valid if input is an active transaction,
//          otherwise 0.
func (tx *Tx) ID() uint64 {
	args := struct {
		txn uintptr
		id  uint64
	}{
		txn: uintptr(unsafe.Pointer(tx.txn)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_txn_id), ptr, 0)
	return args.id
}

// CommitLatency of commit stages in 1/65536 of seconds units.
// \warning This structure may be changed in future releases.
// \see mdbx_txn_commit_ex()
type CommitLatency struct {
	// Duration of preparation (commit child transactions, update sub-databases records and cursors destroying).
	Preparation uint32
	// Duration of GC/freeDB handling & updation.
	GC uint32
	// Duration of internal audit if enabled.
	Audit uint32
	// Duration of writing dirty/modified data pages.
	Write uint32
	// Duration of syncing written data to the dist/storage.
	Sync uint32
	// Duration of transaction ending (releasing resources).
	Ending uint32
	// The total duration of a commit.
	Whole uint32
}

// CommitEx commit all the operations of a transaction into the database and
// collect latency information.
// \see mdbx_txn_commit()
// \ingroup c_statinfo
// \warning This function may be changed in future releases.
func (tx *Tx) CommitEx(latency *CommitLatency) Error {
	args := struct {
		txn     uintptr
		latency uintptr
		result  Error
	}{
		txn:     uintptr(unsafe.Pointer(tx.txn)),
		latency: uintptr(unsafe.Pointer(latency)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_txn_commit_ex), ptr, 0)
	return args.result
}

// Commit all the operations of a transaction into the database.
// \ingroup c_transactions
//
// If the current thread is not eligible to manage the transaction then
// the \ref MDBX_THREAD_MISMATCH error will returned. Otherwise the transaction
// will be committed and its handle is freed. If the transaction cannot
// be committed, it will be aborted with the corresponding error returned.
//
// Thus, a result other than \ref MDBX_THREAD_MISMATCH means that the
// transaction is terminated:
//  - Resources are released;
//  - Transaction handle is invalid;
//  - Cursor(s) associated with transaction must not be used, except with
//    mdbx_cursor_renew() and \ref mdbx_cursor_close().
//    Such cursor(s) must be closed explicitly by \ref mdbx_cursor_close()
//    before or after transaction commit, either can be reused with
//    \ref mdbx_cursor_renew() until it will be explicitly closed by
//    \ref mdbx_cursor_close().
//
// \param [in] txn  A transaction handle returned by \ref mdbx_txn_begin().
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_RESULT_TRUE      Transaction was aborted since it should
//                               be aborted due to previous errors.
// \retval MDBX_PANIC            A fatal error occurred earlier
//                               and the environment must be shut down.
// \retval MDBX_BAD_TXN          Transaction is already finished or never began.
// \retval MDBX_EBADSIGN         Transaction object has invalid signature,
//                               e.g. transaction was already terminated
//                               or memory was corrupted.
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_EINVAL           Transaction handle is NULL.
// \retval MDBX_ENOSPC           No more disk space.
// \retval MDBX_EIO              A system-level I/O error occurred.
// \retval MDBX_ENOMEM           Out of memory.
func (tx *Tx) Commit() Error {
	tx.committed = true
	return tx.CommitEx(nil)
}

// Abort Abandon all the operations of the transaction instead of saving them.
// \ingroup c_transactions
//
// The transaction handle is freed. It and its cursors must not be used again
// after this call, except with \ref mdbx_cursor_renew() and
// \ref mdbx_cursor_close().
//
// If the current thread is not eligible to manage the transaction then
// the \ref MDBX_THREAD_MISMATCH error will returned. Otherwise the transaction
// will be aborted and its handle is freed. Thus, a result other than
// \ref MDBX_THREAD_MISMATCH means that the transaction is terminated:
//  - Resources are released;
//  - Transaction handle is invalid;
//  - Cursor(s) associated with transaction must not be used, except with
//    \ref mdbx_cursor_renew() and \ref mdbx_cursor_close().
//    Such cursor(s) must be closed explicitly by \ref mdbx_cursor_close()
//    before or after transaction abort, either can be reused with
//    \ref mdbx_cursor_renew() until it will be explicitly closed by
//    \ref mdbx_cursor_close().
//
// \param [in] txn  A transaction handle returned by \ref mdbx_txn_begin().
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_PANIC            A fatal error occurred earlier and
//                               the environment must be shut down.
// \retval MDBX_BAD_TXN          Transaction is already finished or never began.
// \retval MDBX_EBADSIGN         Transaction object has invalid signature,
//                               e.g. transaction was already terminated
//                               or memory was corrupted.
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_EINVAL           Transaction handle is NULL.
func (tx *Tx) Abort() Error {
	args := struct {
		txn    uintptr
		result Error
	}{
		txn: uintptr(unsafe.Pointer(tx.txn)),
	}
	tx.aborted = true
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_txn_abort), ptr, 0)
	return args.result
}

// Break Marks transaction as broken.
// \ingroup c_transactions
//
// Function keeps the transaction handle and corresponding locks, but makes
// impossible to perform any operations within a broken transaction.
// Broken transaction must then be aborted explicitly later.
//
// \param [in] txn  A transaction handle returned by \ref mdbx_txn_begin().
//
// \see mdbx_txn_abort() \see mdbx_txn_reset() \see mdbx_txn_commit()
// \returns A non-zero error value on failure and 0 on success.
func (tx *Tx) Break() Error {
	args := struct {
		txn    uintptr
		result Error
	}{
		txn: uintptr(unsafe.Pointer(tx.txn)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_txn_break), ptr, 0)
	return args.result
}

// Reset a read-only transaction.
// \ingroup c_transactions
//
// Abort the read-only transaction like \ref mdbx_txn_abort(), but keep the
// transaction handle. Therefore \ref mdbx_txn_renew() may reuse the handle.
// This saves allocation overhead if the process will start a new read-only
// transaction soon, and also locking overhead if \ref MDBX_NOTLS is in use. The
// reader table lock is released, but the table slot stays tied to its thread
// or \ref MDBX_txn. Use \ref mdbx_txn_abort() to discard a reset handle, and to
// free its lock table slot if \ref MDBX_NOTLS is in use.
//
// Cursors opened within the transaction must not be used again after this
// call, except with \ref mdbx_cursor_renew() and \ref mdbx_cursor_close().
//
// Reader locks generally don't interfere with writers, but they keep old
// versions of database pages allocated. Thus they prevent the old pages from
// being reused when writers commit new data, and so under heavy load the
// database size may grow much more rapidly than otherwise.
//
// \param [in] txn  A transaction handle returned by \ref mdbx_txn_begin().
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_PANIC            A fatal error occurred earlier and
//                               the environment must be shut down.
// \retval MDBX_BAD_TXN          Transaction is already finished or never began.
// \retval MDBX_EBADSIGN         Transaction object has invalid signature,
//                               e.g. transaction was already terminated
//                               or memory was corrupted.
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_EINVAL           Transaction handle is NULL.
func (tx *Tx) Reset() Error {
	args := struct {
		txn    uintptr
		result Error
	}{
		txn: uintptr(unsafe.Pointer(tx.txn)),
	}
	tx.reset = true
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_txn_reset), ptr, 0)
	return args.result
}

// Renew a read-only transaction.
// \ingroup c_transactions
//
// This acquires a new reader lock for a transaction handle that had been
// released by \ref mdbx_txn_reset(). It must be called before a reset
// transaction may be used again.
//
// \param [in] txn  A transaction handle returned by \ref mdbx_txn_begin().
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_PANIC            A fatal error occurred earlier and
//                               the environment must be shut down.
// \retval MDBX_BAD_TXN          Transaction is already finished or never began.
// \retval MDBX_EBADSIGN         Transaction object has invalid signature,
//                               e.g. transaction was already terminated
//                               or memory was corrupted.
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_EINVAL           Transaction handle is NULL.
func (tx *Tx) Renew() Error {
	args := struct {
		txn    uintptr
		result Error
	}{
		txn: uintptr(unsafe.Pointer(tx.txn)),
	}
	tx.reset = false
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_txn_renew), ptr, 0)
	return args.result
}

type Canary struct {
	X, Y, Z, V uint64
}

// PutCanary Set integers markers (aka "canary") associated with the environment.
// \ingroup c_crud
// \see mdbx_canary_get()
//
// \param [in] txn     A transaction handle returned by \ref mdbx_txn_begin()
// \param [in] canary  A optional pointer to \ref MDBX_canary structure for `x`,
//              `y` and `z` values from.
//            - If canary is NOT NULL then the `x`, `y` and `z` values will be
//              updated from given canary argument, but the 'v' be always set
//              to the current transaction number if at least one `x`, `y` or
//              `z` values have changed (i.e. if `x`, `y` and `z` have the same
//              values as currently present then nothing will be changes or
//              updated).
//            - if canary is NULL then the `v` value will be explicitly update
//              to the current transaction number without changes `x`, `y` nor
//              `z`.
//
// \returns A non-zero error value on failure and 0 on success.
func (tx *Tx) PutCanary(canary *Canary) Error {
	args := struct {
		txn    uintptr
		canary uintptr
		result Error
	}{
		txn:    uintptr(unsafe.Pointer(tx.txn)),
		canary: uintptr(unsafe.Pointer(canary)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_canary_put), ptr, 0)
	return args.result
}

// GetCanary Returns fours integers markers (aka "canary") associated with the
// environment.
// \ingroup c_crud
// \see mdbx_canary_set()
//
// \param [in] txn     A transaction handle returned by \ref mdbx_txn_begin().
// \param [in] canary  The address of an MDBX_canary structure where the
//                     information will be copied.
//
// \returns A non-zero error value on failure and 0 on success.
func (tx *Tx) GetCanary(canary *Canary) Error {
	args := struct {
		txn    uintptr
		canary uintptr
		result Error
	}{
		txn:    uintptr(unsafe.Pointer(tx.txn)),
		canary: uintptr(unsafe.Pointer(canary)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_canary_get), ptr, 0)
	return args.result
}

// EnvInfo Return information about the MDBX environment.
// \ingroup c_statinfo
//
// At least one of env or txn argument must be non-null. If txn is passed
// non-null then stat will be filled accordingly to the given transaction.
// Otherwise, if txn is null, then stat will be populated by a snapshot from
// the last committed write transaction, and at next time, other information
// can be returned.
//
// Legacy \ref mdbx_env_info() correspond to calling \ref mdbx_env_info_ex()
// with the null `txn` argument.
//
// \param [in] env     An environment handle returned by \ref mdbx_env_create()
// \param [in] txn     A transaction handle returned by \ref mdbx_txn_begin()
// \param [out] info   The address of an \ref MDBX_envinfo structure
//                     where the information will be copied
// \param [in] bytes   The size of \ref MDBX_envinfo.
//
// \returns A non-zero error value on failure and 0 on success.
func (tx *Tx) EnvInfo(info *EnvInfo) Error {
	if info == nil {
		return ErrInvalid
	}
	args := struct {
		env    uintptr
		txn    uintptr
		info   uintptr
		size   uintptr
		result int32
	}{
		env:  uintptr(unsafe.Pointer(tx.env.env)),
		txn:  uintptr(unsafe.Pointer(tx.txn)),
		info: uintptr(unsafe.Pointer(info)),
		size: unsafe.Sizeof(C.MDBX_envinfo{}),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_env_info_ex), ptr, 0)
	return Error(args.result)
}

// OpenDBI Open or Create a database in the environment.
// \ingroup c_dbi
//
// A database handle denotes the name and parameters of a database,
// independently of whether such a database exists. The database handle may be
// discarded by calling \ref mdbx_dbi_close(). The old database handle is
// returned if the database was already open. The handle may only be closed
// once.
//
// \note A notable difference between MDBX and LMDB is that MDBX make handles
// opened for existing databases immediately available for other transactions,
// regardless this transaction will be aborted or reset. The REASON for this is
// to avoiding the requirement for multiple opening a same handles in
// concurrent read transactions, and tracking of such open but hidden handles
// until the completion of read transactions which opened them.
//
// Nevertheless, the handle for the NEWLY CREATED database will be invisible
// for other transactions until the this write transaction is successfully
// committed. If the write transaction is aborted the handle will be closed
// automatically. After a successful commit the such handle will reside in the
// shared environment, and may be used by other transactions.
//
// In contrast to LMDB, the MDBX allow this function to be called from multiple
// concurrent transactions or threads in the same process.
//
// To use named database (with name != NULL), \ref mdbx_env_set_maxdbs()
// must be called before opening the environment. Table names are
// keys in the internal unnamed database, and may be read but not written.
//
// \param [in] txn    transaction handle returned by \ref mdbx_txn_begin().
// \param [in] name   The name of the database to open. If only a single
//                    database is needed in the environment,
//                    this value may be NULL.
// \param [in] flags  Special options for this database. This parameter must
//                    be set to 0 or by bitwise OR'ing together one or more
//                    of the values described here:
//  - \ref MDBX_REVERSEKEY
//      Keys are strings to be compared in reverse order, from the end
//      of the strings to the beginning. By default, Keys are treated as
//      strings and compared from beginning to end.
//  - \ref MDBX_INTEGERKEY
//      Keys are binary integers in native byte order, either uint32_t or
//      uint64_t, and will be sorted as such. The keys must all be of the
//      same size and must be aligned while passing as arguments.
//  - \ref MDBX_DUPSORT
//      Duplicate keys may be used in the database. Or, from another point of
//      view, keys may have multiple data items, stored in sorted order. By
//      default keys must be unique and may have only a single data item.
//  - \ref MDBX_DUPFIXED
//      This flag may only be used in combination with \ref MDBX_DUPSORT. This
//      option tells the library that the data items for this database are
//      all the same size, which allows further optimizations in storage and
//      retrieval. When all data items are the same size, the
//      \ref MDBX_GET_MULTIPLE, \ref MDBX_NEXT_MULTIPLE and
//      \ref MDBX_PREV_MULTIPLE cursor operations may be used to retrieve
//      multiple items at once.
//  - \ref MDBX_INTEGERDUP
//      This option specifies that duplicate data items are binary integers,
//      similar to \ref MDBX_INTEGERKEY keys. The data values must all be of the
//      same size and must be aligned while passing as arguments.
//  - \ref MDBX_REVERSEDUP
//      This option specifies that duplicate data items should be compared as
//      strings in reverse order (the comparison is performed in the direction
//      from the last byte to the first).
//  - \ref MDBX_CREATE
//      Create the named database if it doesn't exist. This option is not
//      allowed in a read-only transaction or a read-only environment.
//
// \param [out] dbi     Address where the new \ref MDBX_dbi handle
//                      will be stored.
//
// For \ref mdbx_dbi_open_ex() additional arguments allow you to set custom
// comparison functions for keys and values (for multimaps).
// \see avoid_custom_comparators
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_NOTFOUND   The specified database doesn't exist in the
//                         environment and \ref MDBX_CREATE was not specified.
// \retval MDBX_DBS_FULL   Too many databases have been opened.
//                         \see mdbx_env_set_maxdbs()
// \retval MDBX_INCOMPATIBLE  Database is incompatible with given flags,
//                         i.e. the passed flags is different with which the
//                         database was created, or the database was already
//                         opened with a different comparison function(s).
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
func (tx *Tx) OpenDBI(name string, flags DBFlags) (DBI, Error) {
	if len(name) == 0 {
		var dbi DBI
		err := Error(C.mdbx_dbi_open(tx.txn, nil, (C.MDBX_db_flags_t)(flags), (*C.MDBX_dbi)(unsafe.Pointer(&dbi))))
		return dbi, err
	} else {
		n := C.CString(name)
		defer C.free(unsafe.Pointer(n))
		var dbi DBI
		err := Error(C.mdbx_dbi_open(tx.txn, n, (C.MDBX_db_flags_t)(flags), (*C.MDBX_dbi)(unsafe.Pointer(&dbi))))
		return dbi, err
	}
}

// OpenDBIEx OpenDBI with custom comparators.
// \ref avoid_custom_comparators "avoid using custom comparators" and use
// \ref mdbx_dbi_open() instead.
//
// \ingroup c_dbi
//
// \param [in] txn    transaction handle returned by \ref mdbx_txn_begin().
// \param [in] name   The name of the database to open. If only a single
//                    database is needed in the environment,
//                    this value may be NULL.
// \param [in] flags  Special options for this database.
// \param [in] keycmp  Optional custom key comparison function for a database.
// \param [in] datacmp Optional custom data comparison function for a database.
// \param [out] dbi    Address where the new MDBX_dbi handle will be stored.
// \returns A non-zero error value on failure and 0 on success.
func (tx *Tx) OpenDBIEx(name string, flags DBFlags, keyCompare, dataCompare *Cmp) (DBI, Error) {
	if len(name) == 0 {
		var dbi DBI
		err := Error(C.mdbx_dbi_open_ex(tx.txn, nil, (C.MDBX_db_flags_t)(flags), (*C.MDBX_dbi)(unsafe.Pointer(&dbi)),
			(*C.MDBX_cmp_func)(unsafe.Pointer(keyCompare)), (*C.MDBX_cmp_func)(unsafe.Pointer(dataCompare))))
		return dbi, err
	} else {
		n := C.CString(name)
		defer C.free(unsafe.Pointer(n))
		var dbi DBI
		err := Error(C.mdbx_dbi_open_ex(tx.txn, n, (C.MDBX_db_flags_t)(flags), (*C.MDBX_dbi)(unsafe.Pointer(&dbi)),
			(*C.MDBX_cmp_func)(unsafe.Pointer(keyCompare)), (*C.MDBX_cmp_func)(unsafe.Pointer(dataCompare))))
		return dbi, err
	}
}

// Stats Statistics for a database in the environment
// \ingroup c_statinfo
// \see mdbx_env_stat_ex() \see mdbx_dbi_stat()
type Stats struct {
	PageSize      uint32 // Size of a database page. This is the same for all databases.
	Depth         uint32 // Depth (height) of the B-tree
	BranchPages   uint64 // Number of internal (non-leaf) pages
	LeafPages     uint64 // Number of leaf pages
	OverflowPages uint64 // Number of overflow pages
	Entries       uint64 // Number of data items
	ModTxnID      uint64 // Transaction ID of committed last modification
}

// DBIStat Retrieve statistics for a database.
// \ingroup c_statinfo
//
// \param [in] txn     A transaction handle returned by \ref mdbx_txn_begin().
// \param [in] dbi     A database handle returned by \ref mdbx_dbi_open().
// \param [out] stat   The address of an \ref MDBX_stat structure where
//                     the statistics will be copied.
// \param [in] bytes   The size of \ref MDBX_stat.
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_EINVAL   An invalid parameter was specified.
func (tx *Tx) DBIStat(dbi DBI, stat *Stats) Error {
	args := struct {
		txn    uintptr
		stat   uintptr
		size   uintptr
		dbi    uint32
		result Error
	}{
		txn:  uintptr(unsafe.Pointer(tx.txn)),
		stat: uintptr(unsafe.Pointer(stat)),
		size: unsafe.Sizeof(Stats{}),
		dbi:  uint32(dbi),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_dbi_stat), ptr, 0)
	return args.result
}

// DBIFlags Retrieve the DB flags and status for a database handle.
// \ingroup c_statinfo
//
// \param [in] txn     A transaction handle returned by \ref mdbx_txn_begin().
// \param [in] dbi     A database handle returned by \ref mdbx_dbi_open().
// \param [out] flags  Address where the flags will be returned.
// \param [out] state  Address where the state will be returned.
//
// \returns A non-zero error value on failure and 0 on success.
func (tx *Tx) DBIFlags(dbi DBI) (DBFlags, DBIState, Error) {
	var flags DBFlags
	var state DBIState

	args := struct {
		txn    uintptr
		flags  uintptr
		state  uintptr
		dbi    uint32
		result Error
	}{
		txn:   uintptr(unsafe.Pointer(tx.txn)),
		flags: uintptr(unsafe.Pointer(&flags)),
		state: uintptr(unsafe.Pointer(&state)),
		dbi:   uint32(dbi),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_dbi_flags_ex), ptr, 0)
	return flags, state, args.result
}

// Drop Empty or delete and close a database.
// \ingroup c_crud
//
// \see mdbx_dbi_close() \see mdbx_dbi_open()
//
// \param [in] txn  A transaction handle returned by \ref mdbx_txn_begin().
// \param [in] dbi  A database handle returned by \ref mdbx_dbi_open().
// \param [in] del  `false` to empty the DB, `true` to delete it
//                  from the environment and close the DB handle.
//
// \returns A non-zero error value on failure and 0 on success.
func (tx *Tx) Drop(dbi DBI, del bool) Error {
	args := struct {
		txn    uintptr
		del    uintptr
		dbi    uint32
		result Error
	}{
		txn: uintptr(unsafe.Pointer(tx.txn)),
		dbi: uint32(dbi),
	}
	if del {
		args.del = 1
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_drop), ptr, 0)
	return args.result
}

// Get items from a database.
// \ingroup c_crud
//
// This function retrieves key/data pairs from the database. The address
// and length of the data associated with the specified key are returned
// in the structure to which data refers.
// If the database supports duplicate keys (\ref MDBX_DUPSORT) then the
// first data item for the key will be returned. Retrieval of other
// items requires the use of \ref mdbx_cursor_get().
//
// \note The memory pointed to by the returned values is owned by the
// database. The caller need not dispose of the memory, and may not
// modify it in any way. For values returned in a read-only transaction
// any modification attempts will cause a `SIGSEGV`.
//
// \note Values returned from the database are valid only until a
// subsequent update operation, or the end of the transaction.
//
// \param [in] txn       A transaction handle returned by \ref mdbx_txn_begin().
// \param [in] dbi       A database handle returned by \ref mdbx_dbi_open().
// \param [in] key       The key to search for in the database.
// \param [in,out] data  The data corresponding to the key.
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_NOTFOUND  The key was not in the database.
// \retval MDBX_EINVAL    An invalid parameter was specified.
func (tx *Tx) Get(dbi DBI, key *Val, data *Val) Error {
	args := struct {
		txn    uintptr
		key    uintptr
		data   uintptr
		dbi    uint32
		result Error
	}{
		txn:  uintptr(unsafe.Pointer(tx.txn)),
		key:  uintptr(unsafe.Pointer(key)),
		data: uintptr(unsafe.Pointer(data)),
		dbi:  uint32(dbi),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_get), ptr, 0)
	return args.result
}

// GetEqualOrGreat Get equal or great item from a database.
// \ingroup c_crud
//
// Briefly this function does the same as \ref mdbx_get() with a few
// differences:
// 1. Return equal or great (due comparison function) key-value
//    pair, but not only exactly matching with the key.
// 2. On success return \ref MDBX_SUCCESS if key found exactly,
//    and \ref MDBX_RESULT_TRUE otherwise. Moreover, for databases with
//    \ref MDBX_DUPSORT flag the data argument also will be used to match over
//    multi-value/duplicates, and \ref MDBX_SUCCESS will be returned only when
//    BOTH the key and the data match exactly.
// 3. Updates BOTH the key and the data for pointing to the actual key-value
//    pair inside the database.
//
// \param [in] txn           A transaction handle returned
//                           by \ref mdbx_txn_begin().
// \param [in] dbi           A database handle returned by \ref mdbx_dbi_open().
// \param [in,out] key       The key to search for in the database.
// \param [in,out] data      The data corresponding to the key.
//
// \returns A non-zero error value on failure and \ref MDBX_RESULT_FALSE
//          or \ref MDBX_RESULT_TRUE on success (as described above).
//          Some possible errors are:
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_NOTFOUND      The key was not in the database.
// \retval MDBX_EINVAL        An invalid parameter was specified.
func (tx *Tx) GetEqualOrGreat(dbi DBI, key *Val, data *Val) Error {
	args := struct {
		txn    uintptr
		key    uintptr
		data   uintptr
		dbi    uint32
		result Error
	}{
		txn:  uintptr(unsafe.Pointer(tx.txn)),
		key:  uintptr(unsafe.Pointer(key)),
		data: uintptr(unsafe.Pointer(data)),
		dbi:  uint32(dbi),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_get_equal_or_great), ptr, 0)
	return args.result
}

// GetEx Get items from a database
// and optionally number of data items for a given key.
//
// \ingroup c_crud
//
// Briefly this function does the same as \ref mdbx_get() with a few
// differences:
//  1. If values_count is NOT NULL, then returns the count
//     of multi-values/duplicates for a given key.
//  2. Updates BOTH the key and the data for pointing to the actual key-value
//     pair inside the database.
//
// \param [in] txn           A transaction handle returned
//                           by \ref mdbx_txn_begin().
// \param [in] dbi           A database handle returned by \ref mdbx_dbi_open().
// \param [in,out] key       The key to search for in the database.
// \param [in,out] data      The data corresponding to the key.
// \param [out] values_count The optional address to return number of values
//                           associated with given key:
//                            = 0 - in case \ref MDBX_NOTFOUND error;
//                            = 1 - exactly for databases
//                                  WITHOUT \ref MDBX_DUPSORT;
//                            >= 1 for databases WITH \ref MDBX_DUPSORT.
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_NOTFOUND  The key was not in the database.
// \retval MDBX_EINVAL    An invalid parameter was specified.
func (tx *Tx) GetEx(dbi DBI, key *Val, data *Val) (int, Error) {
	var valuesCount uintptr
	args := struct {
		txn         uintptr
		key         uintptr
		data        uintptr
		valuesCount uintptr
		dbi         uint32
		result      Error
	}{
		txn:         uintptr(unsafe.Pointer(tx.txn)),
		key:         uintptr(unsafe.Pointer(key)),
		data:        uintptr(unsafe.Pointer(data)),
		valuesCount: uintptr(unsafe.Pointer(&valuesCount)),
		dbi:         uint32(dbi),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_get_ex), ptr, 0)
	return int(valuesCount), args.result
}

// Put Store items into a database.
// \ingroup c_crud
//
// This function stores key/data pairs in the database. The default behavior
// is to enter the new key/data pair, replacing any previously existing key
// if duplicates are disallowed, or adding a duplicate data item if
// duplicates are allowed (see \ref MDBX_DUPSORT).
//
// \param [in] txn        A transaction handle returned
//                        by \ref mdbx_txn_begin().
// \param [in] dbi        A database handle returned by \ref mdbx_dbi_open().
// \param [in] key        The key to store in the database.
// \param [in,out] data   The data to store.
// \param [in] flags      Special options for this operation.
//                        This parameter must be set to 0 or by bitwise OR'ing
//                        together one or more of the values described here:
//   - \ref MDBX_NODUPDATA
//      Enter the new key-value pair only if it does not already appear
//      in the database. This flag may only be specified if the database
//      was opened with \ref MDBX_DUPSORT. The function will return
//      \ref MDBX_KEYEXIST if the key/data pair already appears in the database.
//
//  - \ref MDBX_NOOVERWRITE
//      Enter the new key/data pair only if the key does not already appear
//      in the database. The function will return \ref MDBX_KEYEXIST if the key
//      already appears in the database, even if the database supports
//      duplicates (see \ref  MDBX_DUPSORT). The data parameter will be set
//      to point to the existing item.
//
//  - \ref MDBX_CURRENT
//      Update an single existing entry, but not add new ones. The function will
//      return \ref MDBX_NOTFOUND if the given key not exist in the database.
//      In case multi-values for the given key, with combination of
//      the \ref MDBX_ALLDUPS will replace all multi-values,
//      otherwise return the \ref MDBX_EMULTIVAL.
//
//  - \ref MDBX_RESERVE
//      Reserve space for data of the given size, but don't copy the given
//      data. Instead, return a pointer to the reserved space, which the
//      caller can fill in later - before the next update operation or the
//      transaction ends. This saves an extra memcpy if the data is being
//      generated later. MDBX does nothing else with this memory, the caller
//      is expected to modify all of the space requested. This flag must not
//      be specified if the database was opened with \ref MDBX_DUPSORT.
//
//  - \ref MDBX_APPEND
//      Append the given key/data pair to the end of the database. This option
//      allows fast bulk loading when keys are already known to be in the
//      correct order. Loading unsorted keys with this flag will cause
//      a \ref MDBX_EKEYMISMATCH error.
//
//  - \ref MDBX_APPENDDUP
//      As above, but for sorted dup data.
//
//  - \ref MDBX_MULTIPLE
//      Store multiple contiguous data elements in a single request. This flag
//      may only be specified if the database was opened with
//      \ref MDBX_DUPFIXED. With combination the \ref MDBX_ALLDUPS
//      will replace all multi-values.
//      The data argument must be an array of two \ref MDBX_val. The `iov_len`
//      of the first \ref MDBX_val must be the size of a single data element.
//      The `iov_base` of the first \ref MDBX_val must point to the beginning
//      of the array of contiguous data elements which must be properly aligned
//      in case of database with \ref MDBX_INTEGERDUP flag.
//      The `iov_len` of the second \ref MDBX_val must be the count of the
//      number of data elements to store. On return this field will be set to
//      the count of the number of elements actually written. The `iov_base` of
//      the second \ref MDBX_val is unused.
//
// \see \ref c_crud_hints "Quick reference for Insert/Update/Delete operations"
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_KEYEXIST  The key/value pair already exists in the database.
// \retval MDBX_MAP_FULL  The database is full, see \ref mdbx_env_set_mapsize().
// \retval MDBX_TXN_FULL  The transaction has too many dirty pages.
// \retval MDBX_EACCES    An attempt was made to write
//                        in a read-only transaction.
// \retval MDBX_EINVAL    An invalid parameter was specified.
func (tx *Tx) Put(dbi DBI, key *Val, data *Val, flags PutFlags) Error {
	args := struct {
		txn    uintptr
		key    uintptr
		data   uintptr
		dbi    uint32
		flags  uint32
		result Error
	}{
		txn:   uintptr(unsafe.Pointer(tx.txn)),
		key:   uintptr(unsafe.Pointer(key)),
		data:  uintptr(unsafe.Pointer(data)),
		dbi:   uint32(dbi),
		flags: uint32(flags),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_put), ptr, 0)
	return args.result
}

// Replace items in a database.
// \ingroup c_crud
//
// This function allows to update or delete an existing value at the same time
// as the previous value is retrieved. If the argument new_data equal is NULL
// zero, the removal is performed, otherwise the update/insert.
//
// The current value may be in an already changed (aka dirty) page. In this
// case, the page will be overwritten during the update, and the old value will
// be lost. Therefore, an additional buffer must be passed via old_data
// argument initially to copy the old value. If the buffer passed in is too
// small, the function will return \ref MDBX_RESULT_TRUE by setting iov_len
// field pointed by old_data argument to the appropriate value, without
// performing any changes.
//
// For databases with non-unique keys (i.e. with \ref MDBX_DUPSORT flag),
// another use case is also possible, when by old_data argument selects a
// specific item from multi-value/duplicates with the same key for deletion or
// update. To select this scenario in flags should simultaneously specify
// \ref MDBX_CURRENT and \ref MDBX_NOOVERWRITE. This combination is chosen
// because it makes no sense, and thus allows you to identify the request of
// such a scenario.
//
// \param [in] txn           A transaction handle returned
//                           by \ref mdbx_txn_begin().
// \param [in] dbi           A database handle returned by \ref mdbx_dbi_open().
// \param [in] key           The key to store in the database.
// \param [in] new_data      The data to store, if NULL then deletion will
//                           be performed.
// \param [in,out] old_data  The buffer for retrieve previous value as describe
//                           above.
// \param [in] flags         Special options for this operation.
//                           This parameter must be set to 0 or by bitwise
//                           OR'ing together one or more of the values
//                           described in \ref mdbx_put() description above,
//                           and additionally
//                           (\ref MDBX_CURRENT | \ref MDBX_NOOVERWRITE)
//                           combination for selection particular item from
//                           multi-value/duplicates.
//
// \see \ref c_crud_hints "Quick reference for Insert/Update/Delete operations"
//
// \returns A non-zero error value on failure and 0 on success.
func (tx *Tx) Replace(dbi DBI, key *Val, data *Val, oldData *Val, flags PutFlags) Error {
	args := struct {
		txn     uintptr
		key     uintptr
		data    uintptr
		oldData uintptr
		dbi     uint32
		flags   uint32
		result  Error
	}{
		txn:     uintptr(unsafe.Pointer(tx.txn)),
		key:     uintptr(unsafe.Pointer(key)),
		data:    uintptr(unsafe.Pointer(data)),
		oldData: uintptr(unsafe.Pointer(oldData)),
		dbi:     uint32(dbi),
		flags:   uint32(flags),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_replace), ptr, 0)
	return args.result
}

// Delete items from a database.
// \ingroup c_crud
//
// This function removes key/data pairs from the database.
//
// \note The data parameter is NOT ignored regardless the database does
// support sorted duplicate data items or not. If the data parameter
// is non-NULL only the matching data item will be deleted. Otherwise, if data
// parameter is NULL, any/all value(s) for specified key will be deleted.
//
// This function will return \ref MDBX_NOTFOUND if the specified key/data
// pair is not in the database.
//
// \see \ref c_crud_hints "Quick reference for Insert/Update/Delete operations"
//
// \param [in] txn   A transaction handle returned by \ref mdbx_txn_begin().
// \param [in] dbi   A database handle returned by \ref mdbx_dbi_open().
// \param [in] key   The key to delete from the database.
// \param [in] data  The data to delete.
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_EACCES   An attempt was made to write
//                       in a read-only transaction.
// \retval MDBX_EINVAL   An invalid parameter was specified.
func (tx *Tx) Delete(dbi DBI, key *Val, data *Val) Error {
	args := struct {
		txn    uintptr
		key    uintptr
		data   uintptr
		dbi    uint32
		result Error
	}{
		txn:  uintptr(unsafe.Pointer(tx.txn)),
		key:  uintptr(unsafe.Pointer(key)),
		data: uintptr(unsafe.Pointer(data)),
		dbi:  uint32(dbi),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_del), ptr, 0)
	return args.result
}

//////////////////////////////////////////////////////////////////////////////////////////
// Cursor
//////////////////////////////////////////////////////////////////////////////////////////

type Cursor C.MDBX_cursor

// NewCursor Create a cursor handle but not bind it to transaction nor DBI handle.
// \ingroup c_cursors
//
// An capable of operation cursor is associated with a specific transaction and
// database. A cursor cannot be used when its database handle is closed. Nor
// when its transaction has ended, except with \ref mdbx_cursor_bind() and
// \ref mdbx_cursor_renew().
// Also it can be discarded with \ref mdbx_cursor_close().
//
// A cursor must be closed explicitly always, before or after its transaction
// ends. It can be reused with \ref mdbx_cursor_bind()
// or \ref mdbx_cursor_renew() before finally closing it.
//
// \note In contrast to LMDB, the MDBX required that any opened cursors can be
// reused and must be freed explicitly, regardless ones was opened in a
// read-only or write transaction. The REASON for this is eliminates ambiguity
// which helps to avoid errors such as: use-after-free, double-free, i.e.
// memory corruption and segfaults.
//
// \param [in] context A pointer to application context to be associated with
//                     created cursor and could be retrieved by
//                     \ref mdbx_cursor_get_userctx() until cursor closed.
//
// \returns Created cursor handle or NULL in case out of memory.
func NewCursor() *Cursor {
	args := struct {
		context uintptr
		cursor  uintptr
	}{}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_cursor_create), ptr, 0)
	return (*Cursor)(unsafe.Pointer(args.cursor))
}

// Bind cursor to specified transaction and DBI handle.
// \ingroup c_cursors
//
// Using of the `mdbx_cursor_bind()` is equivalent to calling
// \ref mdbx_cursor_renew() but with specifying an arbitrary dbi handle.
//
// An capable of operation cursor is associated with a specific transaction and
// database. The cursor may be associated with a new transaction,
// and referencing a new or the same database handle as it was created with.
// This may be done whether the previous transaction is live or dead.
//
// \note In contrast to LMDB, the MDBX required that any opened cursors can be
// reused and must be freed explicitly, regardless ones was opened in a
// read-only or write transaction. The REASON for this is eliminates ambiguity
// which helps to avoid errors such as: use-after-free, double-free, i.e.
// memory corruption and segfaults.
//
// \param [in] txn      A transaction handle returned by \ref mdbx_txn_begin().
// \param [in] dbi      A database handle returned by \ref mdbx_dbi_open().
// \param [out] cursor  A cursor handle returned by \ref mdbx_cursor_create().
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_EINVAL  An invalid parameter was specified.
func (tx *Tx) Bind(cursor *Cursor, dbi DBI) Error {
	args := struct {
		txn    uintptr
		cursor uintptr
		dbi    DBI
		result Error
	}{
		txn:    uintptr(unsafe.Pointer(tx.txn)),
		cursor: uintptr(unsafe.Pointer(cursor)),
		dbi:    dbi,
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_cursor_bind), ptr, 0)
	return args.result
}

// OpenCursor Create a cursor handle for the specified transaction and DBI handle.
// \ingroup c_cursors
//
// Using of the `mdbx_cursor_open()` is equivalent to calling
// \ref mdbx_cursor_create() and then \ref mdbx_cursor_bind() functions.
//
// An capable of operation cursor is associated with a specific transaction and
// database. A cursor cannot be used when its database handle is closed. Nor
// when its transaction has ended, except with \ref mdbx_cursor_bind() and
// \ref mdbx_cursor_renew().
// Also it can be discarded with \ref mdbx_cursor_close().
//
// A cursor must be closed explicitly always, before or after its transaction
// ends. It can be reused with \ref mdbx_cursor_bind()
// or \ref mdbx_cursor_renew() before finally closing it.
//
// \note In contrast to LMDB, the MDBX required that any opened cursors can be
// reused and must be freed explicitly, regardless ones was opened in a
// read-only or write transaction. The REASON for this is eliminates ambiguity
// which helps to avoid errors such as: use-after-free, double-free, i.e.
// memory corruption and segfaults.
//
// \param [in] txn      A transaction handle returned by \ref mdbx_txn_begin().
// \param [in] dbi      A database handle returned by \ref mdbx_dbi_open().
// \param [out] cursor  Address where the new \ref MDBX_cursor handle will be
//                      stored.
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_EINVAL  An invalid parameter was specified.
func (tx *Tx) OpenCursor(dbi DBI) (*Cursor, Error) {
	var cursor *C.MDBX_cursor
	args := struct {
		txn    uintptr
		cursor uintptr
		dbi    DBI
		result Error
	}{
		txn:    uintptr(unsafe.Pointer(tx.txn)),
		cursor: uintptr(unsafe.Pointer(&cursor)),
		dbi:    dbi,
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_cursor_open), ptr, 0)
	return (*Cursor)(unsafe.Pointer(cursor)), args.result
}

// Close a cursor handle.
// \ingroup c_cursors
//
// The cursor handle will be freed and must not be used again after this call,
// but its transaction may still be live.
//
// \note In contrast to LMDB, the MDBX required that any opened cursors can be
// reused and must be freed explicitly, regardless ones was opened in a
// read-only or write transaction. The REASON for this is eliminates ambiguity
// which helps to avoid errors such as: use-after-free, double-free, i.e.
// memory corruption and segfaults.
//
// \param [in] cursor  A cursor handle returned by \ref mdbx_cursor_open()
//                     or \ref mdbx_cursor_create().
func (cur *Cursor) Close() Error {
	ptr := uintptr(unsafe.Pointer(cur))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_cursor_close), ptr, 0)
	return ErrSuccess
}

// Renew a cursor handle.
// \ingroup c_cursors
//
// An capable of operation cursor is associated with a specific transaction and
// database. The cursor may be associated with a new transaction,
// and referencing a new or the same database handle as it was created with.
// This may be done whether the previous transaction is live or dead.
//
// Using of the `mdbx_cursor_renew()` is equivalent to calling
// \ref mdbx_cursor_bind() with the DBI handle that previously
// the cursor was used with.
//
// \note In contrast to LMDB, the MDBX allow any cursor to be re-used by using
// \ref mdbx_cursor_renew(), to avoid unnecessary malloc/free overhead until it
// freed by \ref mdbx_cursor_close().
//
// \param [in] txn      A transaction handle returned by \ref mdbx_txn_begin().
// \param [in] cursor   A cursor handle returned by \ref mdbx_cursor_open().
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_EINVAL  An invalid parameter was specified.
func (cur *Cursor) Renew(tx *Tx) Error {
	args := struct {
		txn    uintptr
		cursor uintptr
		result Error
	}{
		txn:    uintptr(unsafe.Pointer(tx.txn)),
		cursor: uintptr(unsafe.Pointer(cur)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_cursor_renew), ptr, 0)
	return args.result
}

// Tx Return the cursor's transaction handle.
// \ingroup c_cursors
//
// \param [in] cursor A cursor handle returned by \ref mdbx_cursor_open().
func (cur *Cursor) Tx() *C.MDBX_txn {
	args := struct {
		cursor uintptr
		txn    uintptr
	}{
		cursor: uintptr(unsafe.Pointer(cur)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_cursor_txn), ptr, 0)
	return (*C.MDBX_txn)(unsafe.Pointer(args.txn))
}

// DBI Return the cursor's database handle.
// \ingroup c_cursors
//
// \param [in] cursor  A cursor handle returned by \ref mdbx_cursor_open().
func (cur *Cursor) DBI() DBI {
	args := struct {
		cursor uintptr
		dbi    DBI
	}{
		cursor: uintptr(unsafe.Pointer(cur)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_cursor_dbi), ptr, 0)
	return args.dbi
}

// Copy cursor position and state.
// \ingroup c_cursors
//
// \param [in] src       A source cursor handle returned
// by \ref mdbx_cursor_create() or \ref mdbx_cursor_open().
//
// \param [in,out] dest  A destination cursor handle returned
// by \ref mdbx_cursor_create() or \ref mdbx_cursor_open().
//
// \returns A non-zero error value on failure and 0 on success.
func (cur *Cursor) Copy(dest *Cursor) Error {
	args := struct {
		src    uintptr
		dest   uintptr
		result Error
	}{
		src:  uintptr(unsafe.Pointer(cur)),
		dest: uintptr(unsafe.Pointer(dest)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_cursor_copy), ptr, 0)
	return args.result
}

// Get Retrieve by cursor.
// \ingroup c_crud
//
// This function retrieves key/data pairs from the database. The address and
// length of the key are returned in the object to which key refers (except
// for the case of the \ref MDBX_SET option, in which the key object is
// unchanged), and the address and length of the data are returned in the object
// to which data refers.
// \see mdbx_get()
//
// \param [in] cursor    A cursor handle returned by \ref mdbx_cursor_open().
// \param [in,out] key   The key for a retrieved item.
// \param [in,out] data  The data of a retrieved item.
// \param [in] op        A cursor operation \ref MDBX_cursor_op.
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_NOTFOUND  No matching key found.
// \retval MDBX_EINVAL    An invalid parameter was specified.
func (cur *Cursor) Get(key *Val, data *Val, op CursorOp) Error {
	args := struct {
		cursor uintptr
		key    uintptr
		data   uintptr
		op     CursorOp
		result Error
	}{
		cursor: uintptr(unsafe.Pointer(cur)),
		key:    uintptr(unsafe.Pointer(key)),
		data:   uintptr(unsafe.Pointer(data)),
		op:     op,
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_cursor_get), ptr, 0)
	return args.result
}

// Put Store by cursor.
// \ingroup c_crud
//
// This function stores key/data pairs into the database. The cursor is
// positioned at the new item, or on failure usually near it.
//
// \param [in] cursor    A cursor handle returned by \ref mdbx_cursor_open().
// \param [in] key       The key operated on.
// \param [in,out] data  The data operated on.
// \param [in] flags     Options for this operation. This parameter
//                       must be set to 0 or by bitwise OR'ing together
//                       one or more of the values described here:
//  - \ref MDBX_CURRENT
//      Replace the item at the current cursor position. The key parameter
//      must still be provided, and must match it, otherwise the function
//      return \ref MDBX_EKEYMISMATCH. With combination the
//      \ref MDBX_ALLDUPS will replace all multi-values.
//
//      \note MDBX allows (unlike LMDB) you to change the size of the data and
//      automatically handles reordering for sorted duplicates
//      (see \ref MDBX_DUPSORT).
//
//  - \ref MDBX_NODUPDATA
//      Enter the new key-value pair only if it does not already appear in the
//      database. This flag may only be specified if the database was opened
//      with \ref MDBX_DUPSORT. The function will return \ref MDBX_KEYEXIST
//      if the key/data pair already appears in the database.
//
//  - \ref MDBX_NOOVERWRITE
//      Enter the new key/data pair only if the key does not already appear
//      in the database. The function will return \ref MDBX_KEYEXIST if the key
//      already appears in the database, even if the database supports
//      duplicates (\ref MDBX_DUPSORT).
//
//  - \ref MDBX_RESERVE
//      Reserve space for data of the given size, but don't copy the given
//      data. Instead, return a pointer to the reserved space, which the
//      caller can fill in later - before the next update operation or the
//      transaction ends. This saves an extra memcpy if the data is being
//      generated later. This flag must not be specified if the database
//      was opened with \ref MDBX_DUPSORT.
//
//  - \ref MDBX_APPEND
//      Append the given key/data pair to the end of the database. No key
//      comparisons are performed. This option allows fast bulk loading when
//      keys are already known to be in the correct order. Loading unsorted
//      keys with this flag will cause a \ref MDBX_KEYEXIST error.
//
//  - \ref MDBX_APPENDDUP
//      As above, but for sorted dup data.
//
//  - \ref MDBX_MULTIPLE
//      Store multiple contiguous data elements in a single request. This flag
//      may only be specified if the database was opened with
//      \ref MDBX_DUPFIXED. With combination the \ref MDBX_ALLDUPS
//      will replace all multi-values.
//      The data argument must be an array of two \ref MDBX_val. The `iov_len`
//      of the first \ref MDBX_val must be the size of a single data element.
//      The `iov_base` of the first \ref MDBX_val must point to the beginning
//      of the array of contiguous data elements which must be properly aligned
//      in case of database with \ref MDBX_INTEGERDUP flag.
//      The `iov_len` of the second \ref MDBX_val must be the count of the
//      number of data elements to store. On return this field will be set to
//      the count of the number of elements actually written. The `iov_base` of
//      the second \ref MDBX_val is unused.
//
// \see \ref c_crud_hints "Quick reference for Insert/Update/Delete operations"
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_EKEYMISMATCH  The given key value is mismatched to the current
//                            cursor position
// \retval MDBX_MAP_FULL      The database is full,
//                             see \ref mdbx_env_set_mapsize().
// \retval MDBX_TXN_FULL      The transaction has too many dirty pages.
// \retval MDBX_EACCES        An attempt was made to write in a read-only
//                            transaction.
// \retval MDBX_EINVAL        An invalid parameter was specified.
func (cur *Cursor) Put(key *Val, data *Val, flags PutFlags) Error {
	args := struct {
		cursor uintptr
		key    uintptr
		data   uintptr
		flags  PutFlags
		result Error
	}{
		cursor: uintptr(unsafe.Pointer(cur)),
		key:    uintptr(unsafe.Pointer(key)),
		data:   uintptr(unsafe.Pointer(data)),
		flags:  flags,
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_cursor_put), ptr, 0)
	return args.result
}

// Delete current key/data pair.
// \ingroup c_crud
//
// This function deletes the key/data pair to which the cursor refers. This
// does not invalidate the cursor, so operations such as \ref MDBX_NEXT can
// still be used on it. Both \ref MDBX_NEXT and \ref MDBX_GET_CURRENT will
// return the same record after this operation.
//
// \param [in] cursor  A cursor handle returned by mdbx_cursor_open().
// \param [in] flags   Options for this operation. This parameter must be set
// to one of the values described here.
//
//  - \ref MDBX_CURRENT Delete only single entry at current cursor position.
//  - \ref MDBX_ALLDUPS
//    or \ref MDBX_NODUPDATA (supported for compatibility)
//      Delete all of the data items for the current key. This flag has effect
//      only for database(s) was created with \ref MDBX_DUPSORT.
//
// \see \ref c_crud_hints "Quick reference for Insert/Update/Delete operations"
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_MAP_FULL      The database is full,
//                            see \ref mdbx_env_set_mapsize().
// \retval MDBX_TXN_FULL      The transaction has too many dirty pages.
// \retval MDBX_EACCES        An attempt was made to write in a read-only
//                            transaction.
// \retval MDBX_EINVAL        An invalid parameter was specified.
func (cur *Cursor) Delete(flags PutFlags) Error {
	args := struct {
		cursor uintptr
		flags  PutFlags
		result Error
	}{
		cursor: uintptr(unsafe.Pointer(cur)),
		flags:  flags,
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_cursor_del), ptr, 0)
	return args.result
}

// Count Return count of duplicates for current key.
// \ingroup c_crud
//
// This call is valid for all databases, but reasonable only for that support
// sorted duplicate data items \ref MDBX_DUPSORT.
//
// \param [in] cursor    A cursor handle returned by \ref mdbx_cursor_open().
// \param [out] pcount   Address where the count will be stored.
//
// \returns A non-zero error value on failure and 0 on success,
//          some possible errors are:
// \retval MDBX_THREAD_MISMATCH  Given transaction is not owned
//                               by current thread.
// \retval MDBX_EINVAL   Cursor is not initialized, or an invalid parameter
//                       was specified.
func (cur *Cursor) Count() (int, Error) {
	var count uintptr
	args := struct {
		cursor uintptr
		count  uintptr
		result Error
	}{
		cursor: uintptr(unsafe.Pointer(cur)),
		count:  uintptr(unsafe.Pointer(&count)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_cursor_count), ptr, 0)
	return int(count), args.result
}

// EOF Determines whether the cursor is pointed to a key-value pair or not,
// i.e. was not positioned or points to the end of data.
// \ingroup c_cursors
//
// \param [in] cursor    A cursor handle returned by \ref mdbx_cursor_open().
//
// \returns A \ref MDBX_RESULT_TRUE or \ref MDBX_RESULT_FALSE value,
//          otherwise the error code:
// \retval MDBX_RESULT_TRUE    No more data available or cursor not
//                             positioned
// \retval MDBX_RESULT_FALSE   A data is available
// \retval Otherwise the error code
func (cur *Cursor) EOF() Error {
	args := struct {
		cursor uintptr
		result Error
	}{
		cursor: uintptr(unsafe.Pointer(cur)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_cursor_eof), ptr, 0)
	return args.result
}

// OnFirst Determines whether the cursor is pointed to the first key-value pair
// or not. \ingroup c_cursors
//
// \param [in] cursor    A cursor handle returned by \ref mdbx_cursor_open().
//
// \returns A MDBX_RESULT_TRUE or MDBX_RESULT_FALSE value,
//          otherwise the error code:
// \retval MDBX_RESULT_TRUE   Cursor positioned to the first key-value pair
// \retval MDBX_RESULT_FALSE  Cursor NOT positioned to the first key-value
// pair \retval Otherwise the error code
func (cur *Cursor) OnFirst() Error {
	args := struct {
		cursor uintptr
		result Error
	}{
		cursor: uintptr(unsafe.Pointer(cur)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_cursor_on_first), ptr, 0)
	return args.result
}

// OnLast Determines whether the cursor is pointed to the last key-value pair
// or not. \ingroup c_cursors
//
// \param [in] cursor    A cursor handle returned by \ref mdbx_cursor_open().
//
// \returns A \ref MDBX_RESULT_TRUE or \ref MDBX_RESULT_FALSE value,
//          otherwise the error code:
// \retval MDBX_RESULT_TRUE   Cursor positioned to the last key-value pair
// \retval MDBX_RESULT_FALSE  Cursor NOT positioned to the last key-value pair
// \retval Otherwise the error code
func (cur *Cursor) OnLast() Error {
	args := struct {
		cursor uintptr
		result Error
	}{
		cursor: uintptr(unsafe.Pointer(cur)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_cursor_on_last), ptr, 0)
	return args.result
}

// EstimateDistance
// \details \note The estimation result varies greatly depending on the filling
// of specific pages and the overall balance of the b-tree:
//
// 1. The number of items is estimated by analyzing the height and fullness of
// the b-tree. The accuracy of the result directly depends on the balance of
// the b-tree, which in turn is determined by the history of previous
// insert/delete operations and the nature of the data (i.e. variability of
// keys length and so on). Therefore, the accuracy of the estimation can vary
// greatly in a particular situation.
//
// 2. To understand the potential spread of results, you should consider a
// possible situations basing on the general criteria for splitting and merging
// b-tree pages:
//  - the page is split into two when there is no space for added data;
//  - two pages merge if the result fits in half a page;
//  - thus, the b-tree can consist of an arbitrary combination of pages filled
//    both completely and only 1/4. Therefore, in the worst case, the result
//    can diverge 4 times for each level of the b-tree excepting the first and
//    the last.
//
// 3. In practice, the probability of extreme cases of the above situation is
// close to zero and in most cases the error does not exceed a few percent. On
// the other hand, it's just a chance you shouldn't overestimate.///

// EstimateDistance the distance between cursors as a number of elements.
// \ingroup c_rqest
//
// This function performs a rough estimate based only on b-tree pages that are
// common for the both cursor's stacks. The results of such estimation can be
// used to build and/or optimize query execution plans.
//
// Please see notes on accuracy of the result in the details
// of \ref c_rqest section.
//
// Both cursors must be initialized for the same database and the same
// transaction.
//
// \param [in] first            The first cursor for estimation.
// \param [in] last             The second cursor for estimation.
// \param [out] distance_items  The pointer to store estimated distance value,
//                              i.e. `*distance_items = distance(first, last)`.
//
// \returns A non-zero error value on failure and 0 on success.
func EstimateDistance(first, last *Cursor) (int64, Error) {
	var distance int64
	args := struct {
		first    uintptr
		last     uintptr
		distance uintptr
		result   Error
	}{
		first:    uintptr(unsafe.Pointer(first)),
		last:     uintptr(unsafe.Pointer(last)),
		distance: uintptr(unsafe.Pointer(&distance)),
	}
	ptr := uintptr(unsafe.Pointer(&args))
	unsafecgo.NonBlocking((*byte)(C.do_mdbx_estimate_distance), ptr, 0)
	return distance, args.result
}
