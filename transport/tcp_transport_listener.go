package transport

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

type tcpTransportListener struct {
	listener net.Listener
	timeout  time.Duration
}

func (t *tcpTransportListener) Addr() string {
	return t.listener.Addr().String()
}

func (t *tcpTransportListener) Close() error {
	return t.listener.Close()
}

func (t *tcpTransportListener) Accept(acceptFun func(socket Socket)) error {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			if netError, ok := err.(net.Error); ok && netError.Temporary() {
				fmt.Println("temp error: ", netError.Error())
				time.Sleep(time.Second / 2)
				continue
			}
			fmt.Println("tcpTransportListener Accept func have error and returned", err)
			return err
		}
		newBuffer := bufio.NewWriter(conn)
		newSocket := &tcpTransportSocket{
			conn:    conn,
			buf:     newBuffer,
			encode:  gob.NewEncoder(newBuffer),
			decode:  gob.NewDecoder(conn),
			timeout: t.timeout,
		}
		go func() {
			defer func() {
				if err := recover(); err != nil {
					newSocket.Close()
				}
			}()
			acceptFun(newSocket)
		}()
	}
}
