//go:generate protoc -I maxnumber/ maxnumber/maxnumber.proto --go_out=plugins=grpc:maxnumber

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	mn "maxnumber/maxnumber"
)



// server is used to implement Number.SendNumber.
type server struct{}

// SendNumber implements  Number Server
func (s server) SendNumber(srv mn.Number_SendNumberServer) error {
	var max int32
	ctx := srv.Context()


	for {

		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}

		// continue if number reveived from stream
		// less than max
		if req.Number <= max {
			continue
		}

		// update max and send it to stream
		max = req.Number
		resp := mn.NumberMesg{Number: max}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("sent new max=%d", max)
	}


}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "localhost", 7777))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create the TLS credentials
	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}
	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	s := grpc.NewServer(opts...)
	mn.RegisterNumberServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
