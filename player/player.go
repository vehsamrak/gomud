package player

import "net"

type Player struct {
	Codepage string
	Name string
	ConnectionPointer *net.Conn
}
