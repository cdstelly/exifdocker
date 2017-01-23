package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/http"
	"log"
	"rpcshared"
)

func startServer() {
	Exif := new(rpcshared.ExifTool)
	rpc.Register(Exif)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":5556")
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	go http.Serve(l, nil)
}

func main() {
	startServer()
	meta := make(chan int)
	x := <- meta    // wait for a while, and listen
	fmt.Println(x)
}
