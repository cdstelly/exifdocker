package rpcshared

import (
    "fmt"
    "os/exec"
    "log"
    "bytes"
	"time"
	"github.com/montanaflynn/stats"
)

type ExifTool struct {
	NumberRequests int
	RequestHistory []float64
}

type Args struct {
	Data   []byte
	FileID string
}

func (t *ExifTool) Execute(args *Args, reply *string) error {
    fmt.Println("I got data of length: ", len(args.Data))

    pathToExiftool := "exiftool"
    cmd := exec.Command(pathToExiftool, "-")
    cmd.Stdin = bytes.NewReader(args.Data)
    var out bytes.Buffer
    cmd.Stdout = &out

	startTime := time.Now()

    err := cmd.Run()
	fmt.Println(out.String())
	*reply = out.String()

	executionTime := time.Since(startTime).Seconds()  //use seconds as opposed to nanoseconds, returns float64 which is required with stats package

	//Update the counters
	t.NumberRequests += 1
	t.RequestHistory = append(t.RequestHistory, executionTime)

    if err != nil {
            log.Println(err)
    }
    return err
}

func (t *ExifTool) GetHistory(args *Args, reply *[]float64) error {
	*reply = t.RequestHistory
	return nil
}


func (t *ExifTool) MeanExecutionTime(args *Args, reply *float64) error {
	mean,merr := stats.Mean(t.RequestHistory)
	*reply = mean
	return merr
}

