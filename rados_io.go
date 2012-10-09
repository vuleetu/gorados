package gorados

/*
#cgo LDFLAGS: -lrados
#include <string.h>
#include <errno.h>
#include "rados/librados.h"
*/
import "C"

import (
    "errors"
    "unsafe"
    "io"
    "log"
)

type RadosIoCtx struct{
    ctx *C.rados_ioctx_t
}

func (r *Rados) IoCtxCreate(poolname string) (*RadosIoCtx, error) {
    var ctx C.rados_ioctx_t
    cerr := C.rados_ioctx_create(*r.cluster, C.CString(poolname), &ctx)
    if cerr < 0 {
        return nil, errors.New("create io contxt failed")
    }

    return &RadosIoCtx{&ctx}, nil
}

func (r *RadosIoCtx) Destroy(poolname string) {
    C.rados_ioctx_destroy(*r.ctx)
}

type RadosObjectStat struct{
    Size uint64
    Timestamp uint64
}

func (r *RadosIoCtx) Stat(oid string) (*RadosObjectStat, error){
    var size  C.uint64_t
    var time_t C.time_t
    cerr := C.rados_stat(*r.ctx, C.CString(oid), &size, &time_t)
    if cerr < 0 {
        log.Println("Get object stat failed, error is", C.GoString(C.strerror(-cerr)))
        return nil, errors.New("Get object stat failed")
    }

    return &RadosObjectStat{uint64(size), uint64(C.uint64_t(time_t))}, nil
}

func (r *RadosIoCtx) PoolSetAuid(uid uint64) error {
    cerr := C.rados_ioctx_pool_set_auid(*r.ctx, C.uint64_t(uid))
    if cerr < 0 {
        return errors.New("set auid failed")
    }

    return nil
}

func (r *RadosIoCtx) PoolGetAuid() (uint64, error) {
    var uid C.uint64_t
    cerr := C.rados_ioctx_pool_get_auid(*r.ctx, &uid)
    if cerr < 0 {
        return 0, errors.New("get auid failed")
    }

    return uint64(uid), nil
}

func (r *RadosIoCtx) GetId() uint64 {
    id := C.rados_ioctx_get_id(*r.ctx)
    return uint64(id)
}

func (r *RadosIoCtx) GetPoolName() (string, error) {
    var buf [MAX_NAME_LEN]C.char
    cerr := C.rados_ioctx_get_pool_name(*r.ctx, &buf[0], MAX_NAME_LEN-1)

    if cerr < 0 {
        return "", errors.New("get pool name failed")
    }

    return C.GoString(&buf[0]), nil
}

func (r *RadosIoCtx) Write(oid string, bin []byte, offset uint64) error {
    cerr := C.rados_write(*r.ctx, C.CString(oid), (*C.char)(unsafe.Pointer(&bin[0])), C.size_t(len(bin)), C.uint64_t(offset))
    if cerr < 0 {
        return errors.New("write data failed")
    }
    return nil
}

func (r *RadosIoCtx) WriteFull(oid string, bin []byte) error {
    cerr := C.rados_write_full(*r.ctx, C.CString(oid), (*C.char)(unsafe.Pointer(&bin[0])), C.size_t(len(bin)))
    if cerr < 0 {
        return errors.New("write full data failed")
    }
    return nil
}

func (r *RadosIoCtx) Append(oid string, bin []byte) error {
    cerr := C.rados_append(*r.ctx, C.CString(oid), (*C.char)(unsafe.Pointer(&bin[0])), C.size_t(len(bin)))
    if cerr < 0 {
        return errors.New("append data failed")
    }
    return nil
}

func (r *RadosIoCtx) Read(oid string, length, offset uint64) ([]byte, error) {
    var buf = make([]byte, length)
    cerr := C.rados_read(*r.ctx, C.CString(oid), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(length), C.uint64_t(offset))
    if cerr < 0 {
        return nil, errors.New("read data failed")
    }
    return buf, nil
}

func (r *RadosIoCtx) Remove(oid string) error {
    cerr := C.rados_remove(*r.ctx, C.CString(oid))
    if cerr < 0 {
        return errors.New("remove object failed")
    }
    return nil
}

func (r *RadosIoCtx) Trunc(oid string, length uint64) error {
    cerr := C.rados_trunc(*r.ctx, C.CString(oid), C.uint64_t(length))
    if cerr < 0 {
        return errors.New("resize object failed")
    }
    return nil
}

type RadosPoolStat struct {
    num_bytes, num_kb, num_objects, num_object_clones, num_object_copies uint64
}

func (r *RadosIoCtx) PoolStat() (*RadosPoolStat, error) {
    var pool_stat C.struct_rados_pool_stat_t
    cerr := C.rados_ioctx_pool_stat(*r.ctx, &pool_stat)
    if cerr < 0 {
        return nil, errors.New("Get pool status failed")
    }
    return &RadosPoolStat{
        num_bytes: uint64(pool_stat.num_bytes),
        num_kb: uint64(pool_stat.num_kb),
        num_objects: uint64(pool_stat.num_objects),
        num_object_clones: uint64(pool_stat.num_object_clones),
        num_object_copies: uint64(pool_stat.num_object_copies)}, nil
}

func (r *RadosIoCtx) ObjectsListOpen() (*RadosListCtx, error){
    var list_ctx C.rados_list_ctx_t
    cerr := C.rados_objects_list_open(*r.ctx, &list_ctx)
    if cerr < 0 {
        return nil, errors.New("list object failed")
    }
    return &RadosListCtx{&list_ctx}, nil
}

type RadosListCtx struct{
    list_ctx *C.rados_list_ctx_t
}

func (ctx *RadosListCtx) Next() error {
    var buf [1]*C.char
    cerr := C.rados_objects_list_next(*ctx.list_ctx, &buf[0], nil)
    if cerr == -C.ENOENT {
        log.Println("Next failed")
        return io.EOF
    }
    if cerr < 0 {
        log.Println("Next failed")
        return errors.New("next failed")
    }
    log.Println(C.GoString(buf[0]))
    return nil
}

func (ctx *RadosListCtx) Close() {
    C.rados_objects_list_close(*ctx.list_ctx)
}

