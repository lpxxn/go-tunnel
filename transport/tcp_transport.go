package transport

import (
	"bufio"
	"context"
	"encoding/gob"
	"net"
	"time"
)

type tcpTransport struct {
}

var defaultDial = &net.Dialer{KeepAlive: time.Second * 10}

func (t *tcpTransport) Dial(addr string, opts ...DialOption) (Client, error) {
	dialOpts := DialOptions{
		Timeout: defaultTimeout,
	}
	for _, v := range opts {
		v(&dialOpts)
	}
	ctx := context.Background()
	var cancel context.CancelFunc
	if dialOpts.Timeout >= 0 {
		ctx, cancel = context.WithTimeout(ctx, dialOpts.Timeout)
	}
	conn, err := defaultDial.DialContext(ctx, "tcp", addr)
	if cancel != nil {
		cancel()
	}
	if err != nil {
		return nil, err
	}
	buf := bufio.NewWriter(conn)
	return &tcpTransportClient{
		conn:     conn,
		buf:      buf,
		encode:   gob.NewEncoder(buf),
		decode:   gob.NewDecoder(conn),
		timeout:  dialOpts.Timeout,
		dialOpts: dialOpts,
	}, nil
}

func (t *tcpTransport) Listen(addr string, opts ...ListenOption) (Listener, error) {
	var listenOpts ListenOptions
	for _, v := range opts {
		v(&listenOpts)
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &tcpTransportListener{
		listener: listener,
		timeout:  listenOpts.Timeout,
	}, nil

}

func (t *tcpTransport) String() string {
	return "tcp"
}

func NewTcpTransport() *tcpTransport {
	return &tcpTransport{}
}
