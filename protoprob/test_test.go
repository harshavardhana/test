package main

import (
	"testing"

	"github.com/colinmarc/hdfs/v2"
	_ "go.etcd.io/etcd/v3/etcdserver/api/membership"
)

func TestInit(t *testing.T) {
	_, _ = hdfs.New("")
	t.Log("ok")
}
