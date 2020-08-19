package main

import (
	"testing"

	"github.com/colinmarc/hdfs/v2"
)

func TestInit(t *testing.T) {
	_, _ = hdfs.New("")
	t.Log("ok")
}
