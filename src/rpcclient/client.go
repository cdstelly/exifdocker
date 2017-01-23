package main

import (
	"io/ioutil"
	"fmt"
	"net/rpc"
//	"net/http"
	"log"
	"rpcshared"
	"github.com/montanaflynn/stats"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "0.0.0.0:5556")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	filepath := "test.file"
	fileData, err:= ioutil.ReadFile(filepath)
	if err != nil { 
		log.Fatal("error reading file: ", err)
	}

	// Execute the task
	args := &rpcshared.Args{FileID: "test", Data: fileData}
	var reply string
	err = client.Call("ExifTool.Execute", args, &reply)
	if err != nil {
		log.Fatal("exif error:", err)
	}
	fmt.Printf("Result: %s\n", reply)

	// Get the stats of the RPC client
	getArgs := &rpcshared.Args{}
	var ExecutionHistory []float64
	getErr := client.Call("ExifTool.GetHistory", getArgs, &ExecutionHistory)
	if getErr != nil {
		fmt.Println("Error getting history: ", getErr)
	}
	theMean, mathErr := stats.Mean(ExecutionHistory)
	theSum, mathErr := stats.Sum(ExecutionHistory)
	theMax, mathErr := stats.Max(ExecutionHistory)
	if mathErr != nil {
		fmt.Println("Arithmitic error: ", mathErr)
	}

	fmt.Printf("The average time it takes to run exiftools: %f\tSum: %f\tMax: %f\n", theMean, theSum, theMax)
}
