package server

import (
	"context"
	"log"
	"net"
)

type Context struct {
	context.Context
	manager *TunnelServer
}

type tcpServer struct {
	ctx *Context
}

func (t *tcpServer) Handle(clientConn net.Conn) {
	log.Printf("TCP: new client address: %s", clientConn.RemoteAddr())

}
