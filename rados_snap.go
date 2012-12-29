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

const MAX_SNAP_LEN = 2000

func (r *RadosIoCtx) SnapCreate(snapname string) error {
    csnapname := C.CString(snapname)
    defer func(){
        C.free(unsafe.Pointer(csnapname))
    }()
    cerr := C.rados_ioctx_snap_create(*r.ctx, csnapname)
    if cerr < 0 {
        return errors.New("create snap failed")
    }

    return nil
}

func (r *RadosIoCtx) SnapRemove(snapname string) error {
    csnapname := C.CString(snapname)
    defer func(){
        C.free(unsafe.Pointer(csnapname))
    }()
    cerr := C.rados_ioctx_snap_remove(*r.ctx, csnapname)
    if cerr < 0 {
        return errors.New("remove snap failed")
    }

    return nil
}

func (r *RadosIoCtx) SnapRollBack(oid, snapname string) error {
    csnapname := C.CString(snapname)
    coid := C.CString(oid)
    defer func(){
        C.free(unsafe.Pointer(csnapname))
        C.free(unsafe.Pointer(coid))
    }()
    cerr := C.rados_rollback(*r.ctx, coid, csnapname)
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

        for k, _ := range snap {
            buf[k] = RadosSnapId(snap[k])
        }
        return buf, nil
    }
    return []RadosSnapId{}, nil
}

func (r *RadosIoCtx) SnapLookup(snapname string) (RadosSnapId, error) {
    var snapid C.rados_snap_t
    csnapname := C.CString(snapname)
    defer func(){
        C.free(unsafe.Pointer(csnapname))
    }()
    cerr := C.rados_ioctx_snap_lookup(*r.ctx, csnapname, &snapid)
    if cerr < 0 {
        return 0, errors.New("lookup snap failed")
    }

    return RadosSnapId(snapid), nil
}

func (r *RadosIoCtx) SnapGetName(snapid RadosSnapId) (string, error) {
    var snapname [MAX_NAME_LEN]C.char
    cerr := C.rados_ioctx_snap_get_name(*r.ctx, C.rados_snap_t(snapid), &snapname[0], MAX_NAME_LEN)
    if cerr < 0 {
        return "", errors.New("get snap name failed")
    }

    return C.GoString(&snapname[0]), nil
}

func (r *RadosIoCtx) SnapGetStamp(snapid RadosSnapId) (uint64, error) {
    var stamp C.time_t
    cerr := C.rados_ioctx_snap_get_stamp(*r.ctx, C.rados_snap_t(snapid), &stamp)
    if cerr < 0 {
        return 0, errors.New("get snap stamp failed")
    }

    return uint64(C.uint64_t(stamp)), nil
}
