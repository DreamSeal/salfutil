package main

import (
	_ "net/http/pprof"
	"salfutil/test"
	"time"
)

func init() {

}

func main() {
	i := 0
	for i < 100 {
		test.RedPackage(10, 100)
		time.Sleep(2111)
		i++
	}

}
