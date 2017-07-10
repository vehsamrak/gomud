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
	listener, err := net.Listen("tcp", ":" + SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Printf("\nMud is listening connections on port %s\nPress Ctrl+C to exit.\n", SERVER_PORT)

	for {
		connection, error := listener.Accept()

		if error != nil {
			fmt.Println("Error accepting: ", error.Error())
			os.Exit(1)
		}

		go handleRequest(connection)
	}
}

func handleRequest(connection net.Conn) {
	channel := make(chan []byte)

	go func(ch chan []byte) {
		fmt.Println("New user connected!")

		for {
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
			commandName := string(bytes.Trim(data, "\r\n\x00"))
			commandName = strings.TrimSpace(commandName)
			executeCommand(connection, commandName)
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

	case "quit":
		fallthrough
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
