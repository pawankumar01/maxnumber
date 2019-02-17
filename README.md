## MaxNumber
MaxNumber is simple Go server and client over the gRPC protocol buffers.
It allows users to send stream of numbers to server.
Server returns a stream of responses that represent the current maximum between all these integers.

Client and Server also encrypt and authenticate request with the cryptographic key.

### Make Sure
The MaxNumber directory is placed under $GOPATH/src/ .Please copy directory to above path to work correctly

### Server Port
Server starts on the port 7777

### Directory Structure

```
maxnumber   |
            |-number_server/
            |        |- main.go (server file - Run this file to start server)
            |-number_client/
            |        |- main.go (client file)
            |-maxnumber/
            |        |- maxnumber.proto (Defining a service)   
            |        |- maxnumber.pb.go (Generated gRPC code in go)
            |-cert/ (public and private keys)
            |-mock_maxnumber/
            |        |- mn_mock.go (generated mock object file from gomock/mockgen)
            |        |- mn_mock_test.go (Test cases, to run go test -v maxnumber/mock_maxnumber)  
```

### Technology Used
```Go, gRPC, ProtocolBuffer, GoMock, gRPC Credentials```


### Functionality
The app implements following feature:-
* Server
    - Support Signing Key
    - Return if new max is found
    
* Client
    - Each request message is signed by key
    - Send Number Streams
    - Receive max number from the server
    
    
### Test Cases
* mn_mock_test.go
    - Test for the gRPC message
    - Test for correctness of server code.

### Assumption
Project assumes that $GOPATH, $GOROOT and protoc is set in the system path.
