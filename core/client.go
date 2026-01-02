package core

import (
	"bufio"
	"fmt"
	"net"
)

type Client interface {
	SendRequest(method byte, content string) (string, error)
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

func (T *TCPClient) SendRequest(method byte, content string) (string, error) {

	if T.connection == nil {
		return "", fmt.Errorf("not connect")
	}
	msg := fmt.Sprintf("%s%s\n", string(method), content)
	if _, err := T.connection.Write([]byte(msg)); err != nil {
		return "", err
	}

	reader := bufio.NewReader(T.connection)
	resp, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// trim the newline for cleaner output
	resp = resp[:len(resp)-1]
	if len(resp) > 0 && resp[len(resp)-1] == '\r' {
		resp = resp[:len(resp)-1]
	}

	// check if server returned an error
	if len(resp) >= 4 && resp[:4] == "ERR " {
		return "", fmt.Errorf("%s", resp[4:])
	}

	return resp, nil
}

func (T *TCPClient) Disconnect() error {
	err := T.connection.Close()
	return err
}
