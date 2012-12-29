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
)

const MAX_XATTR_LEN = 255

func (r *RadosIoCtx) GetXattr(oid, name string) (string, error) {
    var buf [MAX_XATTR_LEN]C.char
    cname := C.CString(name)
    coid := C.CString(oid)
    defer func(){
        C.free(unsafe.Pointer(cname))
        C.free(unsafe.Pointer(coid))
    }()
    cerr := C.rados_getxattr(*r.ctx, coid, cname, &buf[0], MAX_XATTR_LEN)
    if cerr < 0 {
        return "", errors.New("get xattr failed")
    }

    return C.GoString(&buf[0]), nil
}

func (r *RadosIoCtx) SetXattr(oid, name, value string) error {
    cname := C.CString(name)
    coid := C.CString(oid)
    cvalue := C.CString(value)
    defer func(){
        C.free(unsafe.Pointer(cname))
        C.free(unsafe.Pointer(coid))
        C.free(unsafe.Pointer(cvalue))
    }()
    cerr := C.rados_setxattr(*r.ctx, coid, cname, cvalue, C.size_t(len(value)))
    if cerr < 0 {
        return errors.New("set xattr failed")
    }

    return nil
}

func (r *RadosIoCtx) RmXattr(oid, name string) error {
    cname := C.CString(name)
    coid := C.CString(oid)
    defer func(){
        C.free(unsafe.Pointer(cname))
        C.free(unsafe.Pointer(coid))
    }()
    cerr := C.rados_rmxattr(*r.ctx, coid, cname)
    if cerr < 0 {
        return errors.New("delete xattr failed")
    }

    return nil
}
