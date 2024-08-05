package main

import (
	"kvstore/server"
)

func main() {
	dummyServer := server.NewServer("127.0.0.1:8080")
	dummyServer.Run()
}
