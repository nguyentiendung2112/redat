package main

import (
	"fmt"
	"redat/core"
	"strings"
)

var storage = core.Store{}

func ValidateGetParams(params string) bool {
	if len(strings.Split(params, " ")) > 1 {
		return false
	}
	return true
}

func ValidateSetParams(params string) bool {
	if len(strings.Split(params, " ")) < 2 {
		return false
	}
	return true
}

func GetHandler(cmdParams string) (string, error) {
	if !ValidateGetParams(cmdParams) {
		return "", fmt.Errorf("syntax error, invalid parameters %s", cmdParams)
	}
	val, exist := storage.Get(cmdParams)
	if !exist {
		return "<nil>", nil
	}
	return val, nil
}

func SetHandler(cmdParams string) (string, error) {
	if !ValidateSetParams(cmdParams) {
		return "", fmt.Errorf("syntax error, invalid parameters %s", cmdParams)
	}
	parts := strings.SplitN(cmdParams, " ", 2)
	storage.Set(parts[0], parts[1])
	return fmt.Sprintf("SET \"%s\" \"%s\"", parts[0], parts[1]), nil
}

func main() {
	storage.Init()

	var server = core.Server{}
	server.Init()
	server.Register(core.METHOD_NAME_BYTE_MAP[core.GET], GetHandler)
	server.Register(core.METHOD_NAME_BYTE_MAP[core.SET], SetHandler)

	fmt.Println("server started")
	server.Start(":6380")
}
