package main

import (
	"fmt"
	"log"
	"net"
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
		if err != nil {
			fmt.Println("Your port isn't number")
			return
		}
		port = args[0]
	}

	fmt.Printf("Listening on the port: " + port)
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	count := 0

	for {

		connection, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		count++
		if count > 10 {
			connection.Write([]byte("Chat is full. Try again later!"))
			connection.Close()
		}
		// go pkg.Chat(connection)
	}
}
