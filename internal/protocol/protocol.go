package protocol

import "net"

type Protocol interface {
	ConnLoop(conn net.Conn) error
}
