package gorados

/*
#cgo LDFLAGS: -lrados
#include <stdio.h>
#include "rados/librados.h"
*/
import "C"

import (
    "errors"
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

func (r *RadosIoCtx) IoCtxDestroy(poolname string) {
    C.rados_ioctx_destroy(*r.ctx)
}

func (r *RadosIoCtx) IoCtxPoolSetAuid(uid uint64) error {
    cerr := C.rados_ioctx_pool_set_auid(*r.ctx, C.uint64_t(uid))
    if cerr < 0 {
        return errors.New("set auid failed")
    }

    return nil
}

func (r *RadosIoCtx) IoCtxPoolGetAuid() (uint64, error) {
    var uid C.uint64_t
    cerr := C.rados_ioctx_pool_get_auid(*r.ctx, &uid)
    if cerr < 0 {
        return 0, errors.New("get auid failed")
    }

    return uint64(uid), nil
}

func (r *RadosIoCtx) IoCtxGetId() uint64 {
    id := C.rados_ioctx_get_id(*r.ctx)
    return uint64(id)
}
