package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/hxzqlh/go-experiments/rpc/service"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var args *service.Args
	var reply int

	// simultaneous way
	args = &service.Args{7, 8}
	err = client.Call("Arith.Add", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d+%d=%d \n", args.A, args.B, reply)

	// async way
	args = &service.Args{4, 3}
	c := client.Go("Arith.Add", args, &reply, nil)
	<-c.Done
	fmt.Printf("Arith: %d+%d=%d \n", args.A, args.B, reply)
}
