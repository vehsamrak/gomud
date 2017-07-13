package console

import (
	"net"
	"time"
	"fmt"
)

// Send message to external connection
func Client(connectionPointer *net.Conn, message string) {
	connection := *connectionPointer
	connection.Write([]byte(message + "\n\n"))
}

// Output to server console with current time stamp
func Server(message ...interface{})  {
	currentTime := time.Now()

	fmt.Printf("[%d-%02d-%02d %02d:%02d:%02d] ",
		currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute(), currentTime.Second())

	fmt.Println(message...)
}
