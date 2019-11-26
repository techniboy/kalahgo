package protocol

import (
	"bufio"
	"net"
)

type GameConnection struct {
	connection net.Conn
	connReader *bufio.Reader
	ln         net.Listener
}

func NewGameConnection(host string, port string) (*GameConnection, error) {
	g := new(GameConnection)
	var err error
	g.ln, err = net.Listen("tcp4", host+":"+port)
	if err != nil {
		return nil, err
	}
	g.connection, err = g.ln.Accept()
	if err != nil {
		return nil, err
	}

	g.connReader = bufio.NewReader(g.connection)
	return g, nil
}
