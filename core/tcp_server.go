package core

import (
	"fmt"
	"net"
)

type Handler func(string) (string, error)

type Server struct {
	port     string
	handlers map[byte]Handler
}

func (server *Server) Init() {
	server.handlers = make(map[byte]Handler)
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
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n == 0 {
		return
	}

	methodByte := buf[0]
	handler, exists := server.handlers[methodByte]
	if !exists {
		if _, writeErr := conn.Write([]byte("ERR unknown method\n")); writeErr != nil {
			fmt.Println(writeErr)
		}
		return
	}

	content := string(buf[1:n])

	if len(content) > 0 && content[len(content)-1] == '\n' {
		content = content[:len(content)-1]
	}
	if len(content) > 0 && content[len(content)-1] == '\r' {
		content = content[:len(content)-1]
	}

	result, handlerErr := handler(content)
	if handlerErr != nil {
		fmt.Println(handlerErr)
		if _, returnErr := conn.Write([]byte(fmt.Sprintf("ERR %s\n", handlerErr.Error()))); returnErr != nil {
			fmt.Println(returnErr)
		}
		return
	}

	if _, writeErr := conn.Write([]byte(result + "\n")); writeErr != nil {
		fmt.Println(writeErr)
	}
}

func (server *Server) Register(methodCode byte, handler Handler) {
	server.handlers[methodCode] = handler
}
