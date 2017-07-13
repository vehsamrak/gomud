package main

import (
	"fmt"
	"net"
	"os"
	"bytes"
	"strings"
	"github.com/Vehsamrak/gomud/commands"
	"github.com/golang-collections/collections/set"
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

	var connectionPool = set.New()

	for {
		connection, error := listener.Accept()
		connectionPool.Insert(connection)

		if error != nil {
			consoleOutput("Error accepting: ", error.Error())
			os.Exit(1)
		}

		go handleRequest(connection, connectionPool)
	}
}

func handleRequest(connection net.Conn, connectionPool *set.Set) {
	channel := make(chan []byte)

	go func(ch chan []byte) {
		numberOfPlayersOnline := connectionPool.Len()
		consoleOutput(fmt.Sprintf("New user connected! Players online: %v", numberOfPlayersOnline))
		respond(connection, fmt.Sprintf("\nДобро пожаловать в %v!\nИгроков онлайн: %v", MUD_NAME, numberOfPlayersOnline))

		for {
			data := make([]byte, 512)
			_, error := connection.Read(data)

			if error != nil {
				connectionPool.Remove(connection)
				connection.Close()

				consoleOutput(fmt.Sprintf("Connection was closed. Players online: %v", connectionPool.Len()))

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
			executeCommand(commandName, connection)
		}
	}
}

func executeCommand(fullCommand string, connection net.Conn) {
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
		command = commands.Chat{commandParameters}

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

// Send message to external connection
func respond(connection net.Conn, message string)  {
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
