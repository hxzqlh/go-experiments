package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/hxzqlh/go-experiments/rpc/service"
)

func main() {
	arith := new(service.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
