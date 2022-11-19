package main

import (
	"fmt"
	"net"
	"net-cat/pkg"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	port := ""
	if len(args) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	if len(args) == 0 {
		port = "4000"
	}

	if len(args) == 1 {
		_, err := strconv.Atoi(args[0])
		pkg.Error(err)
		port = args[0]
	}

	fmt.Printf("Listening on the port: " + port)
	ln, err := net.Listen("tcp", ":"+port)
	pkg.Error(err)

	count := 0

	for {

		connection, err := ln.Accept()
		pkg.Error(err)
		count++
		if count > 10 {
			connection.Write([]byte("Chat is full. Try again later!"))
			connection.Close()
		}
		// go pkg.Chat(connection)
	}
}
