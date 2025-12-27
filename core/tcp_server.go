package core

import (
	"fmt"
	"net"
)

type Handler func([]byte) error

type Server struct {
	port     string
	handlers map[string]Handler
}

func (server *Server) Start(port string) {
	server.port = port
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
		go server.HandleConnection(conn)
	}
}

func (server *Server) HandleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	var firsThreeByteStr = string(buf[0:4])
	handler, exists := server.handlers[firsThreeByteStr]
	if exists {
		handlerErr := handler(buf[4:])
		if handlerErr != nil {
			fmt.Println(handlerErr)
		}
	}
}

func (server *Server) Register(name string, handler Handler) {
	server.handlers[name] = handler
}
