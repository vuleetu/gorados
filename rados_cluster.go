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
    "log"
)

//type MapCluster map[int64]*C.rados_t

//var map_cluster = MapCluster{}

func (r *Rados) ClusterCreate() error {
    var cluster C.rados_t
    //cerr := C.rados_create(&cluster, (*C.char)(unsafe.Pointer(uintptr(0))))
    cerr := C.rados_create(&cluster, nil)
    if cerr < 0 {
        return errors.New("create cluster handler failed")
    }
    //id := int64(C.random())
    //map_cluster[id] = &cluster
    r.cluster = &cluster
    return nil
}

func (r *Rados) ClusterCreateAsUser(id string) error {
    var cluster C.rados_t
    //cerr := C.rados_create(&cluster, (*C.char)(unsafe.Pointer(uintptr(0))))
    cerr := C.rados_create(&cluster, C.CString(id))
    if cerr < 0 {
        return errors.New("create cluster handler failed")
    }
    //id := int64(C.random())
    //map_cluster[id] = &cluster
    r.cluster = &cluster
    return nil
}

/*
Configure the cluster handle using a Ceph config file.

If path is NULL, the default locations are searched, and the first found is used. The locations are:

$CEPH_CONF (environment variable)
/etc/ceph/ceph.conf
~/.ceph/config
ceph.conf (in the current working directory)
*/
func (r *Rados) ClusterAutoConfig() error {
    cerr := C.rados_conf_read_file(*r.cluster, nil)
    if cerr < 0 {
        return errors.New("read config failed")
    }

    return nil
}

func (r *Rados) ClusterConfig(filename string) error {
    cerr := C.rados_conf_read_file(*(r.cluster), C.CString(filename))
    if cerr < 0 {
        return errors.New("read config failed")
    }

    return nil
}

func (r *Rados) ClusterSetConfig(option, value string) error {
    cerr := C.rados_conf_set(*r.cluster, C.CString(option), C.CString(value))
    if cerr < 0 {
        return errors.New("set config failed")
    }

    return nil
}

func (r *Rados) ClusterConnect() error {
    cerr := C.rados_connect(*r.cluster)
    if cerr < 0 {
        log.Println("error is", C.GoString(C.strerror(-cerr)))
        return errors.New("connect to ceph failed")
    }

    return nil
}


func (r *Rados) ClusterShutDown() {
    C.rados_shutdown(*r.cluster)
}

func (r *Rados) ClusterGetInstanceId() uint64 {
    instance_id := C.rados_get_instance_id(*r.cluster)
    return uint64(instance_id)
}

