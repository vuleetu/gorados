package gorados

/*
#cgo LDFLAGS: -lrados
#include <stdio.h>
#include <string.h>
#include "rados/librados.h"
*/
import "C"

import (
    "errors"
)

const MAX_XATTR_LEN = 255

func (r *RadosIoCtx) GetXattr(oid, name string) (string, error) {
    var buf [MAX_XATTR_LEN]C.char
    cerr := C.rados_getxattr(*r.ctx, C.CString(oid), C.CString(name), &buf[0], MAX_XATTR_LEN)
    if cerr < 0 {
        return "", errors.New("get xattr failed")
    }

    return C.GoString(&buf[0]), nil
}

func (r *RadosIoCtx) SetXattr(oid, name, value string) error {
    cerr := C.rados_setxattr(*r.ctx, C.CString(oid), C.CString(name), C.CString(value), C.size_t(len(value)))
    if cerr < 0 {
        return errors.New("set xattr failed")
    }

    return nil
}

func (r *RadosIoCtx) RmXattr(oid, name string) error {
    cerr := C.rados_rmxattr(*r.ctx, C.CString(oid), C.CString(name))
    if cerr < 0 {
        return errors.New("delete xattr failed")
    }

    return nil
}
