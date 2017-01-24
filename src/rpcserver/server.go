package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/http"
	"log"
	"rpcshared"
	"os"
	"time"
	"rpclogger"
)

var (
	MyName     string
	BrokerHost string
	MyType     string
)

func init() {
	MyName = Generate(2, "-")
	BrokerHost = os.Getenv("BROKERHOST")
	if len(BrokerHost) == 0 {
		BrokerHost = "trex1:5050"
	}
	MyType = "ExifTools"
}

func startServer() {
	Exif := new(rpcshared.ExifTool)
	rpc.Register(Exif)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":5556")
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	go http.Serve(l, nil)
	go PeriodicUpdate(Exif)
}

// Send the whole request history periodically
// TODO: Decay the RequestHistory buffer. This struct will eventually get huge..
func PeriodicUpdate(myRPCInstance *rpcshared.ExifTool) {
	for {
		time.Sleep(time.Millisecond * 5000)
		rpclogger.SubmitReport(BrokerHost, MyName, MyType, myRPCInstance.RequestHistory)
	}
}


func main() {
	startServer()
	meta := make(chan int)
	x := <- meta    // wait for a while, and listen
	fmt.Println(x)
}
