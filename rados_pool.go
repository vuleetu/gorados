package gorados

/*
#cgo LDFLAGS: -lrados
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "rados/librados.h"
*/
import "C"

import (
    "errors"
    "unsafe"
    "bytes"
)

func (r *Rados) PoolList() [][]byte {
    buflen := C.rados_pool_list(*r.cluster, nil, C.size_t(0))
    buf := make([]byte, buflen)
    C.rados_pool_list(*r.cluster, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(buflen))
    buf1 := buf[0:buflen-2]
    return bytes.Split(buf1, []byte{0})
}

func (r *Rados) PoolLookUp(poolname string) (int64, error) {
    cpoolname := C.CString(poolname)
    defer func(){
        C.free(unsafe.Pointer(cpoolname))
    }()
    cerr := C.rados_pool_lookup(*r.cluster, cpoolname)
    if cerr < 0 {
        return 0, errors.New("Pool not found")
    }

    return int64(cerr), nil
}

/*func (r *Rados) PoolReverseLookUp(poolid int64) (string, error) {
    var buf [MAX_NAME_LEN]C.char
    cerr := C.rados_pool_reverse_lookup(*r.cluster, C.int64_t(poolid), &buf[0], MAX_NAME_LEN)
    if cerr < 0 {
        return "", errors.New("Pool not found")
    }

    return C.GoString(&buf[0]), nil
}*/

func (r *Rados) PoolCreate(poolname string) error {
    cpoolname := C.CString(poolname)
    defer func(){
        C.free(unsafe.Pointer(cpoolname))
    }()
    cerr := C.rados_pool_create(*r.cluster, cpoolname)
    if cerr < 0 {
        return errors.New("create pool failed")
    }

    return nil
}

func (r *Rados) PoolCreateWithAuid(poolname string, auid uint64) error {
    cpoolname := C.CString(poolname)
    defer func(){
        C.free(unsafe.Pointer(cpoolname))
    }()
    cerr := C.rados_pool_create_with_auid(*r.cluster, cpoolname, C.uint64_t(auid))
    if cerr < 0 {
        return errors.New("create pool failed")
    }

    return nil
}

func (r *Rados) PoolDelete(poolname string) error {
    cpoolname := C.CString(poolname)
    defer func(){
        C.free(unsafe.Pointer(cpoolname))
    }()
    cerr := C.rados_pool_delete(*r.cluster, cpoolname)
    if cerr < 0 {
        return errors.New("delete pool failed")
    }

    return nil
}

