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

const MAX_NAME_LEN = 1024
