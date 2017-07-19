package main

import (
	"fmt"
	"net"
	"os"
	"bytes"
	"strings"
	"github.com/Vehsamrak/gomud/commands"
	"github.com/Vehsamrak/gomud/console"
)

const MUD_NAME = "Экспериментальный Полигон"
const MUD_PORT = "7000"

func main() {
	listener, error := net.Listen("tcp", ":" + MUD_PORT)
	if error != nil {
		console.Server("Error listening:", error.Error())
		os.Exit(1)
	}

	defer listener.Close()

	console.Server(fmt.Sprintf("Mud is listening connections on port %s. Press Ctrl+C to exit.\n", MUD_PORT))

	connectionPool := map[string]*net.Conn{}

	for {
		connection, error := listener.Accept()

		connectionId := fmt.Sprint(&connection)
		connectionPool[connectionId] = &connection

		if error != nil {
			console.Server("Error accepting: ", error.Error())
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
		console.Server(fmt.Sprintf("New user connected! Players online: %v", numberOfPlayersOnline))
		console.Client(
			connectionPointer,
			fmt.Sprintf(
				"\nДобро пожаловать в %v!\nИгроков онлайн: %v\n\nВведите имя вашего персонажа:",
				MUD_NAME,
				numberOfPlayersOnline,
			),
		)

		for {
			userInput := make([]byte, 512)
			_, error := connection.Read(userInput)

			if error != nil {
				delete(connectionPool, fmt.Sprint(connectionPointer))
				connection.Close()

				console.Server(fmt.Sprintf("Connection was closed. Players online: %v", len(connectionPool)))

				return
			}

			ch<- userInput
		}
	}(channel)

	var executer commands.Executer

	executer = &commands.LoginCommander{
		ConnectionPointer: connectionPointer,
		ConnectionPool: connectionPool,
		LoginStage: "initial",
	}

	for {
		select {
		case userInput := <-channel:
			rawCommand := string(bytes.Trim(userInput, "\r\n\x00"))
			rawCommand = strings.TrimSpace(rawCommand)

			commandResult := executer.ExecuteCommand(rawCommand)

			if commandResult.Executer != nil {
				executer = commandResult.Executer
			}
		}
	}
}
