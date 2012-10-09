package gorados

/*
#cgo LDFLAGS: -lrados
#include <stdio.h>
#include <string.h>
#include "rados/librados.h"
*/
import "C"

import (
    //"errors"
    "unsafe"
    "bytes"
    "log"
)

func (r *Rados) PoolList() {
    buflen := C.rados_pool_list(*r.cluster, nil, C.size_t(0))
    buf := make([]byte, buflen)
    C.rados_pool_list(*r.cluster, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(buflen))
    list := bytes.Split(buf, []byte{0})
    for _, poolname := range list {
        if len(poolname) != 0 {
            log.Println(string(poolname), len(poolname), poolname)
        }
    }
}
