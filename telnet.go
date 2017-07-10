package main

import (
	"fmt"
	"net"
	"os"
	"bytes"
	"strings"
)

const SERVER_PORT = "7000"

func main() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", ":" + SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes
	defer listener.Close()

	fmt.Println()
	fmt.Println("Mud is listening connections on port " + SERVER_PORT)
	fmt.Println("Press Ctrl+C to exit.")
	fmt.Println()

	for {
		// Listen for an incoming connection
		connection, error := listener.Accept()

		if error != nil {
			fmt.Println("Error accepting: ", error.Error())
			os.Exit(1)
		}

		// Handle connections in a new goroutine
		go handleRequest(connection)
	}
}

// Handles incoming requests
func handleRequest(connection net.Conn) {
	channel := make(chan []byte)

	// Start a goroutine to read from our net connection
	go func(ch chan []byte) {
		fmt.Println("New user connected!")

		for {
			// try to read the data
			data := make([]byte, 512)
			_, error := connection.Read(data)

			if error != nil {
				fmt.Println("Connection was closed.")
				connection.Close()

				return
			}

			ch<- data
		}
	}(channel)

	for {
		select {
		case data := <-channel:
			command := string(bytes.Trim(data, "\r\n\x00"))
			command = strings.TrimSpace(command)
			executeCommand(connection, command)
		}
	}
}

func executeCommand(connection net.Conn, command string) {
	fmt.Println("Command received: " + command)

	switch command {
	case "test":
		respond(connection, "Passed!")
	case "look" :
		respond(connection, "You see contours of new mud.")
	case "exit":
		respond(connection, "See you next time!")
		connection.Close()
	case "кто":
		respond(connection, "В этом мире нет никого лучше тебя.")
	case "смотреть":
		respond(connection, "Ты видишь контуры этого мира.")
	default:
		respond(connection, "Command not found.")
	}
}

func respond(connection net.Conn, message string)  {
	connection.Write([]byte(message + "\n\n"))
}
