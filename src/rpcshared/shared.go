package rpcshared

import (
    "fmt"
    "os/exec"
    "log"
    "bytes"
)

type BulkExtractor string

type Args struct {
	Data   []byte
	FileID string
}

func (t *BulkExtractor) Extract(args *Args, reply *string) error {
    fmt.Println("I got data of length: ", len(args.Data))
    pathToExiftool := "exiftool"
    fmt.Println("Path to Exiftools: " , pathToExiftool)
    cmd := exec.Command(pathToExiftool, "-")
    cmd.Stdin = bytes.NewReader(args.Data)
    var out bytes.Buffer
    cmd.Stdout = &out

    err := cmd.Run()
	fmt.Println(out.String())
	*reply = out.String()
    if err != nil {
            log.Println(err)
    }
    return err
}

