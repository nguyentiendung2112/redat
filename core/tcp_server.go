package core

import (
	"fmt"
	"net"
)

type Server struct {
	port string
}

func (server *Server) Start() {
	listener, err := net.Listen("tcp", server.port)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, acceptErr := listener.Accept()
		if acceptErr != nil {
			fmt.Println(acceptErr)
		}
		go server.handleConnection(conn)
	}
}

func (server *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read incoming data
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the incoming data
	fmt.Printf("Received: %s", buf)
}
