package protocol

import (
	"log"
	"net"
)

type TcpHandler interface {
	Handle(conn net.Conn)
}

func TcpServer(listener net.Listener, handler TcpHandler) error {
	for {
		clientConn, err := listener.Accept()
		if err != nil {
			if nErr, ok := err.(net.Error); ok && nErr.Temporary() {
				log.Printf("temporary Accept() failure %s\n", err)
				continue
			}
			log.Printf("listener.Accept() error %s \n", err)
			break
		}
		go handler.Handle(clientConn)
	}
	log.Printf("TCP: closing %s", listener.Addr())
	return nil
}
