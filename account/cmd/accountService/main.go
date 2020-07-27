package main

import (
	"flag"

	"github.com/golang/glog"

	"github.com/vmustillo/microservices/accountservice/internal"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := internal.Run(); err != nil {
		glog.Fatal(err)
	}
}
