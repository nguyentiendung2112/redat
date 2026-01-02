package main

import (
	"bufio"
	"fmt"
	"os"
	"redat/core"
	"strings"
)

func main() {
	fmt.Println("Welcome !!")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("REDAT> ")

		line, readErr := reader.ReadString('\n')
		if readErr != nil {
			fmt.Println("input error:", readErr)
			return
		}

		command := strings.TrimSpace(line)
		if command == "" {
			continue
		}
		client := core.TCPClient{}
		if err := client.Connect(":6380"); err != nil {
			fmt.Println("connect error:", err)
			continue
		}

		if err := ProcessCommand(&client, command); err != nil {
			fmt.Println(err)
		}
		if err := client.Disconnect(); err != nil {
			fmt.Println("disconnect error:", err)
		}
	}
}

func ProcessCommand(client core.Client, command string) error {
	parts := strings.SplitN(command, " ", 2)
	methodName := parts[0]
	methodByte, ok := core.METHOD_NAME_BYTE_MAP[methodName]
	if !ok {
		return fmt.Errorf("unknown method: %s", methodName)
	}

	content := ""
	if len(parts) > 1 {
		content = parts[1]
	}

	resp, err := client.SendRequest(methodByte, content)
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}
