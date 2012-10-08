package main

import (
    "gorados"
    "log"
)

func main() {
    rados := gorados.New()
    err := rados.ClusterCreate()
    if err != nil {
        log.Println("error is", err.Error())
        return
    }

    err = rados.ClusterAutoConfig()
    if err != nil {
        log.Println("error is", err.Error())
        return
    }

    err = rados.ClusterConnect()
    if err != nil {
        log.Println("error is", err.Error())
        return
    }

    log.Println("Done")
}
