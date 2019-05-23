package transprot

import (
	"bufio"
	"encoding/gob"
	"net"
	"time"
)

type tcpTransportSocket struct {
	conn    net.Conn
	buf     *bufio.Writer
	encode  *gob.Encoder
	decode  *gob.Decoder
	timeout time.Duration
}

func (tts *tcpTransportSocket) Recv(msg *Msg) error {
	if tts.timeout > 0 {
		tts.conn.SetDeadline(time.Now().Add(tts.timeout))
	}
	return tts.decode.Decode(msg)
}

func (tts *tcpTransportSocket) Send(msg *Msg) error {
	if tts.timeout > 0 {
		tts.conn.SetDeadline(time.Now().Add(tts.timeout))
	}
	if err := tts.encode.Encode(msg); err != nil {
		return err
	}
	return tts.buf.Flush()
}

func (tts *tcpTransportSocket) Close() error {
	return tts.conn.Close()
}

func (tts *tcpTransportSocket) Remote() string {
	return tts.conn.RemoteAddr().String()
}

func (tts *tcpTransportSocket) Local() string {
	return tts.conn.LocalAddr().String()
}
