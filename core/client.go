package core

import (
	"bufio"
	"fmt"
	"net"
)

type Client interface {
	SendRequest(method string, content string) (string, error)
	Connect(addr string) error
	Disconnect() error
}

type TCPClient struct {
	connection net.Conn
}

func (T *TCPClient) Connect(addr string) error {
	var err error
	T.connection, err = net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	return nil
}

func (T *TCPClient) SendRequest(method string, content string) (string, error) {

	if T.connection == nil {
		return "", fmt.Errorf("not connect")
	}
	msg := fmt.Sprintf("%s %s\n", method, content)
	if _, err := T.connection.Write([]byte(msg)); err != nil {
		return "", err
	}

	reader := bufio.NewReader(T.connection)
	resp, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return resp, nil
}
