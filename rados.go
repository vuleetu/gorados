package rados

/*
#cgo LDFLAGS: -lrados
#include <stdio.h>
#include "rados/librados.h"
*/
import "C"

import (
    "errors"
    //"unsafe"
)

type Rados struct {
    cluster *C.rados_t
}

func New() (*Rados, error) {
    var cluster C.rados_t
    //cerr := C.rados_create(&cluster, (*C.char)(unsafe.Pointer(uintptr(0))))
    cerr := C.rados_create(&cluster, nil)
    if cerr < 0 {
        return nil, errors.New("create cluster handler failed")
    }

    return &Rados{&cluster}, nil
}

func (r *Rados) Config(filename string) error {
    cerr := C.rados_conf_read_file(*r.cluster, C.CString(filename))
    if cerr < 0 {
        return errors.New("read config failed")
    }

    return nil
}

func (r *Rados) Connect() error {
    cerr := C.rados_connect(*r.cluster)
    if cerr < 0 {
        return errors.New("connect to ceph failed")
    }

    return nil
}

func (r *Rados) Create() error {
    var ctx C.rados_ioctx_t
    cerr := C.rados_ioctx_create(*r.cluster, &ctx)
    if cerr < 0 {
        return errors.New("connect to ceph failed")
    }

    return nil
}
