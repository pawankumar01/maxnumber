package main

import (
	"context"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	"math/rand"
	"time"
	"google.golang.org/grpc"
	mn "maxnumber/maxnumber"
)

const (
	address     = "localhost:7777"
)

func main() {
	rand.Seed(time.Now().Unix())

	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := mn.NewNumberClient(conn)


	stream, err := c.SendNumber(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	var max int32
	ctx := stream.Context()
	done := make(chan bool)

	// first goroutine sends random increasing numbers to stream
	// and closes int after 20 iterations
	go func() {
		for i := 1; i <= 20; i++ {
			// generate random number and send it to stream
			rnd := int32(rand.Intn(i))
			req := mn.NumberMesg{Number: rnd}
			if err := stream.Send(&req); err != nil {
				log.Fatalf("can not send %v", err)
			}
			log.Printf("number = %d", req.Number)
			time.Sleep(time.Millisecond * 200)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	// second goroutine receives data from stream
	// and saves result in max variable
	//
	// if stream is finished it closes done channel
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			max = resp.Number
			log.Printf("new max = %d", max)
		}
	}()

	// third goroutine closes done channel
	// if context is done
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done
	log.Printf("finished with max=%d", max)
}
