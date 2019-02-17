package mock_maxnumber_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	mn "maxnumber/maxnumber"
	mnmock "maxnumber/mock_maxnumber"
	"testing"
	"time"
)


var msg = &mn.NumberMesg{Number: 6}

/**
* Test GPRC messages request and reply
*/


func TestSendNumber(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	stream := mnmock.NewMockNumber_SendNumberClient(ctrl)
	stream.EXPECT().Send(gomock.Any()).Return(nil)
	stream.EXPECT().Recv().Return(msg , nil)

	// Create mock for the client interface.
	rgclient := mnmock.NewMockNumberClient(ctrl)
	// Set expectation on RouteChat
	rgclient.EXPECT().SendNumber(
		gomock.Any(),
	).Return(stream, nil)
	if err := testSendNumber(rgclient); err != nil {
		t.Fatalf("Test failed: %v", err)
	}

}



func testSendNumber(client mn.NumberClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.SendNumber(ctx)
	req := mn.NumberMesg{Number: 6}

	if err != nil {
		log.Println(err)
		return err
	}
	if err := stream.Send(&req); err != nil {
		log.Println(err)
		return err
	}

	got, err := stream.Recv()
	if err != nil {
		log.Println(err)
		return err
	}
	if !proto.Equal(got, &req) {
		return fmt.Errorf("stream.Recv() = %v, want %v", got, msg)
	}

	return nil
}

/**
* Check server if it is able to connect and get expected result
*/
func TestServer_Run(t *testing.T) {
	// Simply check that the server is up and can accept connections.
	// Create the client TLS credentials
	address     := "localhost:7777"

	creds, err := credentials.NewClientTLSFromFile("../cert/server.crt", "")
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		t.Fatalf("Test failed: %v", err)
	}
	defer conn.Close()

	c := mn.NewNumberClient(conn)

	stream, err := c.SendNumber(context.Background())
	if err != nil {
		t.Fatalf("openn stream error %v", err)
	}

	if err := testStreamSendNumberRequests(stream, t); err != nil {
		t.Fatalf("Test failed: %v", err)
	}
}


func testStreamSendNumberRequests(stream mn.Number_SendNumberClient, t *testing.T) error {
	tc := struct {
		test    string
		payload []int32
		want    []int32
	}{
			"Sending test series",
			[]int32{1,5,3,6,2,20},
			[]int32{1,5,6,20},
	}


		t.Run(tc.test, func(t *testing.T) {

			for _, value := range tc.payload{
				req := mn.NumberMesg{Number: value}
				if err := stream.Send(&req); err != nil {
					t.Fatalf("can not send %v", err)
				}
			}


			for _, want := range tc.want{
				resp, err := stream.Recv()
				if err == io.EOF {
					return
				}
				if err != nil {
					t.Fatalf("can not receive %v", err)
				}
				max := resp.Number
				//log.Printf("new max %d received", max)
				//log.Printf("want = %d ", want)

				if max != want{
					t.Fatalf("Wrong Server Output")
				}

			}



		});
	return nil
}
