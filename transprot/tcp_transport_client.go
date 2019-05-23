package transprot

import (
	"bufio"
	"encoding/gob"
	"net"
	"time"
)

type tcpTransportClient struct {
	dialOpts DialOptions
	conn     net.Conn
	buf      *bufio.Writer
	encode   *gob.Encoder
	decode   *gob.Decoder
	timeout  time.Duration
}

func (c *tcpTransportClient) Recv(msg *Msg) error {
	if c.timeout > 0 {
		c.conn.SetDeadline(time.Now().Add(c.timeout))
	}
	return c.decode.Decode(msg)
}

func (c *tcpTransportClient) Send(msg *Msg) error {
	if c.timeout > 0 {
		c.conn.SetDeadline(time.Now().Add(c.timeout))
	}
	if err := c.encode.Encode(msg); err != nil {
		return err
	}
	return c.buf.Flush()
}

func (c *tcpTransportClient) Close() error {
	return c.conn.Close()
}

func (c *tcpTransportClient) Local() string {
	return c.conn.LocalAddr().String()
}

func (c *tcpTransportClient) Remote() string {
	return c.conn.RemoteAddr().String()
}
