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

const MUD_NAME = "Experimental Polygon"
const MUD_PORT = "7000"

// Starts mud server and listens for incoming telnet connections
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

// Handle single connection in separated goroutine
func handleConnection(connectionPointer *net.Conn, connectionPool map[string]*net.Conn) {
	connection := *connectionPointer
	channel := make(chan []byte)

	go func(ch chan []byte) {
		numberOfPlayersOnline := len(connectionPool)
		console.Server(fmt.Sprintf("New user connected! Players online: %v", numberOfPlayersOnline))
		console.Client(
			connectionPointer,
			fmt.Sprintf(
				"\nWelcome to \x1b[37;1m%v\x1b[0m!\nPlayers online: %v",
				MUD_NAME,
				numberOfPlayersOnline,
			),
		)
		console.Client(connectionPointer, "BBEDUTE HOMEP KODUPOBKU: 1) UTF-8, 2) KOI8-R, 3) Windows-1251")

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
		Sender: commands.Sender{ConnectionPointer: connectionPointer},
		ConnectionPool: connectionPool,
		LoginStage: "charset",
	}

	for {
		userInput := <-channel
		rawCommand := string(bytes.Trim(userInput, "\r\n\x00"))
		rawCommand = strings.TrimSpace(rawCommand)

		commandResult := executer.ExecuteCommand(rawCommand)

		if commandResult.Executer != nil {
			executer = commandResult.Executer
		}
	}
}
