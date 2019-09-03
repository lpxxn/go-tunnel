package server

import (
	"net"
	"sync"
)

type ClientInfo struct {
	LastUpdate    int64
	Id            int64
	RemoteAddress string `json:"remote_address"`
	HostName      string `json:"host_name"`
	TcpPort       int    `json:"tcp_port"`
	Version       string `json:"version"`
}

type clientInfoList []*ClientInfo

type TunnelServer struct {
	sync.RWMutex
	opts             Options
	tcpListener      net.Listener
	clientIDSequence int64
	clients          map[int64]*ClientInfo
}

func NewTunnelServer(opts *Options) (*TunnelServer, error) {
	defaultOpts := NewOptions()
	tunnel := &TunnelServer{
		opts:        *defaultOpts,
		tcpListener: nil,
		clients:     nil,
	}
	return tunnel, nil
}

func (t *TunnelServer) AddClient(clientInfo *ClientInfo) {
	t.Lock()
	t.clients[clientInfo.Id] = clientInfo
	t.Unlock()
}

func (t *TunnelServer) RemoveClient(id int64) {
	t.Lock()
	delete(t.clients, id)
	t.Unlock()
}
