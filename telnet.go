package main

import (
	"fmt"
	"net"
	"os"
	"bytes"
	"strings"
	"github.com/Vehsamrak/gomud/commands"
	"time"
)

const MUD_NAME = "Экспериментальный Полигон"
const MUD_PORT = "7000"

func main() {
	listener, err := net.Listen("tcp", ":" + MUD_PORT)
	if err != nil {
		consoleOutput("Error listening:", err.Error())
		os.Exit(1)
	}

	defer listener.Close()

	consoleOutput(fmt.Sprintf("Mud is listening connections on port %s. Press Ctrl+C to exit.\n", MUD_PORT))

	connectionPool := map[string]*net.Conn{}

	for {
		connection, error := listener.Accept()

		connectionId := fmt.Sprint(&connection)
		connectionPool[connectionId] = &connection

		if error != nil {
			consoleOutput("Error accepting: ", error.Error())
			os.Exit(1)
		}

		go handleConnection(&connection, connectionPool)
	}
}

func handleConnection(connectionPointer *net.Conn, connectionPool map[string]*net.Conn) {
	connection := *connectionPointer
	channel := make(chan []byte)

	go func(ch chan []byte) {
		numberOfPlayersOnline := len(connectionPool)
		consoleOutput(fmt.Sprintf("New user connected! Players online: %v", numberOfPlayersOnline))
		respond(connectionPointer, fmt.Sprintf("\nДобро пожаловать в %v!\nИгроков онлайн: %v", MUD_NAME, numberOfPlayersOnline))

		for {
			data := make([]byte, 512)
			_, error := connection.Read(data)

			if error != nil {
				delete(connectionPool, fmt.Sprint(connectionPointer))
				connection.Close()

				consoleOutput(fmt.Sprintf("Connection was closed. Players online: %v", len(connectionPool)))

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
			executeCommand(commandName, connectionPointer, connectionPool)
		}
	}
}

func executeCommand(fullCommand string, connectionPointer *net.Conn, connectionPool map[string]*net.Conn) {
	consoleOutput("Command received: " + fullCommand)

	commandWithParameters := strings.Fields(fullCommand)
	commandName := commandWithParameters[0]
	commandParameters := commandWithParameters[1:]

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

	case "chat" :
		fallthrough
	case "чат":
		command = commands.Chat{commandParameters, connectionPool}

	case "quit":
		fallthrough
	case "exit":
		fallthrough
	case "конец":
		command = commands.Exit{}
		connection := *connectionPointer
		defer connection.Close()

	default:
		respond(connectionPointer, "Command not found.")

		return
	}

	respond(connectionPointer, command.Execute())
}

// Send message to external connection
func respond(connectionPointer *net.Conn, message string) {
	connection := *connectionPointer
	connection.Write([]byte(message + "\n\n"))
}

// Output to server console with current time stamp
func consoleOutput(message ...interface{})  {
	currentTime := time.Now()

	fmt.Printf("[%d-%02d-%02d %02d:%02d:%02d] ",
		currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute(), currentTime.Second())

	fmt.Println(message...)
}
