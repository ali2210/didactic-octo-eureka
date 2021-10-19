package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	connect, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Grpc dial connection fail .... ", err.Error())
		return
	}

	ctx, request_cancel := context.WithTimeout(connect, time.Second*2)
	defer request_cancel()

}
