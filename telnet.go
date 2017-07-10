package main

import (
	"fmt"
	"net"
	"os"
	"bytes"
	"strings"
	"github.com/Vehsamrak/gomud/commands"
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

func executeCommand(connection net.Conn, commandName string) {
	fmt.Println("Command received: " + commandName)

	var command commands.Executable

	switch commandName {
	case "test":
		fallthrough
	case "тест":
		command = commands.Test{}

	case "who":
		fallthrough
	case "кто":
		command = commands.Who{}

	case "look" :
		fallthrough
	case "смотреть":
		command = commands.Look{}

	case "exit":
		fallthrough
	case "конец":
		command = commands.Exit{}
		defer connection.Close()
	default:
		respond(connection, "Command not found.")

		return
	}

	respond(connection, command.Execute())
}

func respond(connection net.Conn, message string)  {
	connection.Write([]byte(message + "\n\n"))
}
