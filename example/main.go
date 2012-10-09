package main

import (
    "gorados"
    "io"
    "log"
)

func main() {
    rados := gorados.New()
    err := rados.ClusterCreateAsUser("pool.yunio2")
    if err != nil {
        log.Println("error is", err.Error())
        return
    }

    err = rados.ClusterConfig("./etc/ceph_yunio2.conf")
    if err != nil {
        log.Println("error is", err.Error())
        return
    }

    err = rados.ClusterConnect()
    if err != nil {
        log.Println("error is", err.Error())
        return
    }

    rados.PoolList()

    ctx, err := rados.IoCtxCreate("yunio2")
    if err != nil {
        log.Println("error is", err.Error())
        return
    }

    stat, err := ctx.Stat("30817")
    if err != nil {
        log.Println("error is", err.Error())
        return
    }
    log.Println(stat)

    poolname, err := ctx.GetPoolName()
    if err != nil {
        log.Println("error is", err.Error())
        return
    }
    log.Println(poolname)

    pool_stat, err := ctx.PoolStat()
    log.Println(pool_stat)

    list_ctx, err := ctx.ObjectsListOpen()
    if err != nil {
        log.Println("error is", err.Error())
        return
    }
    for i := 0; i < 20; i++ {
        err = list_ctx.Next()
        if err == io.EOF {
            list_ctx.Close()
            break
        }
    }

    log.Println("Done")
}
