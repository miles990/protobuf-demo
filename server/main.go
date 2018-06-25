package main

import (
	"log"
	"net"
	"os"

	"fmt"
	//"encoding/binary"

	pb "protobuf-demo/server/protobuf"

	"github.com/golang/protobuf/proto"
)

const (
	host = "0.0.0.0"
	port = "50051"
)

type server struct{}

// Handles incoming requests.
func handleRequest(conn net.Conn) {

	// Make a buffer to hold incoming data.
	var err error
	//var message_length int32
	buf := make([]byte, 4096*3)

	// Read the incoming connection into the buffer.

	//   err = binary.Read(conn, binary.BigEndian, &message_length)
	//   if err != nil {
	//   fmt.Println("Error reading:", err.Error())
	// }else{
	//   //fmt.Println("Receive Length:", reqLen, "\nBuffer:\n", buf)
	//   fmt.Println("Receive Length:", message_length)
	// }

	//   println(message_length)
	//fmt.Println(buf)
	var reqLen int
	reqLen, err = conn.Read(buf)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
	} else {
		//fmt.Println("Receive Length:", reqLen, "\nBuffer:\n", buf)
		fmt.Println("Receive Length:", reqLen)
	}

	fmt.Println(conn)

	data := &pb.TheMsg{}
	_ = proto.Unmarshal(buf, data)
	/*if err != nil {
	      fmt.Println("Error Unmarshal:", err.Error())
	  }else{
	      fmt.Println("Receive Data:\n", data)
	  }*/
	//data.GetName()
	//fmt.Println([]byte(buf))
	fmt.Println(data.GetName(), data.GetNum())
	fmt.Println(data)

	// Send a response back to person contacting us.
	conn.Write([]byte("Server already get data."))
	// Close the connection when you're done with it.
	conn.Close()
}

func main() {
	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Listening on " + host + ":" + port)
	for {
		// Listen for an incoming connection.
		conn, err := lis.Accept()

		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}
