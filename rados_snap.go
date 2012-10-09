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

const MAX_SNAP_LEN = 2000

func (r *RadosIoCtx) SnapCreate(snapname string) error {
    cerr := C.rados_ioctx_snap_create(*r.ctx, C.CString(snapname))
    if cerr < 0 {
        return errors.New("create snap failed")
    }

    return nil
}

func (r *RadosIoCtx) SnapRemove(snapname string) error {
    cerr := C.rados_ioctx_snap_remove(*r.ctx, C.CString(snapname))
    if cerr < 0 {
        return errors.New("remove snap failed")
    }

    return nil
}

func (r *RadosIoCtx) SnapRollBack(oid, snapname string) error {
    cerr := C.rados_rollback(*r.ctx, C.CString(oid), C.CString(snapname))
    if cerr < 0 {
        return errors.New("rollback snap failed")
    }

    return nil
}

type RadosSnapId uint64

func (r *RadosIoCtx) SnapList() ([]RadosSnapId, error) {
    var snap [MAX_SNAP_LEN]C.rados_snap_t
    cerr := C.rados_ioctx_snap_list(*r.ctx, &snap[0], MAX_SNAP_LEN)
    if cerr < 0 {
        return nil, errors.New("list snap failed")
    }

    if cerr > 0 {
        buf := make([]RadosSnapId, cerr)

        for k, snapid := range snap {
            buf[k] = RadosSnapId(snapid)
        }
        return buf, nil
    }
    return []RadosSnapId{}, nil
}
