package gorados

/*
#cgo LDFLAGS: -lrados
#include "rados/librados.h"
*/
import "C"

type Rados struct {
    cluster *C.rados_t
}

func New() *Rados {
    return &Rados{}
}

