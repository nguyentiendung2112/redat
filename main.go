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

func ValidateDelParams(params string) bool {
	if len(strings.Split(params, " ")) > 1 {
		return false
	}
	return true
}

func ValidateKeysParams(params string) bool {
	if len(params) != 0 {
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

func DelHandler(cmdParams string) (string, error) {
	if !ValidateDelParams(cmdParams) {
		return "", fmt.Errorf("syntax error, invalid parameters %s", cmdParams)
	}

	storage.Delete(cmdParams)
	return fmt.Sprintf("deleted \"%s\"", cmdParams), nil
}

func ListKeysHandler(cmdParams string) (string, error) {
	if !ValidateKeysParams(cmdParams) {
		return "", fmt.Errorf("syntax error, invalid parameters %s", cmdParams)
	}
	return strings.Join(storage.Keys(), ", "), nil
}

func main() {
	storage.Init()

	var server = core.Server{}
	server.Init()
	server.Register(core.MethodNameByteMap[core.GET], GetHandler)
	server.Register(core.MethodNameByteMap[core.SET], SetHandler)
	server.Register(core.MethodNameByteMap[core.DEL], DelHandler)
	server.Register(core.MethodNameByteMap[core.KEYS], ListKeysHandler)

	fmt.Println("server started")
	server.Start(":6380")
}
