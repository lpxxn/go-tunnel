package server

import (
	"context"
	"log"
	"net"
	"sync"
)

type clientInfo struct {
	lastUpdate    int64
	id            string
	RemoteAddress string `json:"remote_address"`
	HostName      string `json:"host_name"`
	TcpPort       int    `json:"tcp_port"`
	Version       string `json:"version"`
}

type Server struct {
	sync.RWMutex
	opts        Options
	tcpListener net.Listener
	clients     map[string]*clientInfo
}

type Context struct {
	context.Context
	manager *Server
}

type tcpServer struct {
	ctx *Context
}

func (t *tcpServer) Handle(clientConn net.Conn) {
	log.Printf("TCP: new client address: %s", clientConn.RemoteAddr())

}
