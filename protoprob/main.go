package main

import (
	"fmt"

	"github.com/colinmarc/hdfs/v2"
)

func main() {
	_, _ = hdfs.New("")
	fmt.Println("Initialized ok")
}
