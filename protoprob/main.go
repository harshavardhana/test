package main

import (
	"fmt"

	"github.com/colinmarc/hdfs/v2"
	_ "go.etcd.io/etcd/v3/etcdserver/api/membership"
)

func main() {
	_, _ = hdfs.New("")
	fmt.Println("Initialized ok")
}
