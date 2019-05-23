package transprot

import (
	"bufio"
	"encoding/gob"
	"net"
)

type tcpTransport struct {
}

func (t *tcpTransport) Dial(addr string, opts ...DialOption) (Client, error) {
	var dialOpts DialOptions
	for _, v := range opts {
		v(&dialOpts)
	}
	conn, err := net.DialTimeout("tcp", addr, dialOpts.Timeout)
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
