package main

import (
	"fmt"
	"net"
	"time"

	"github.com/lpxxn/go-tunnel/transport"
)

func main() {
	tcpTransport := transport.NewTcpTransport()
	client, err := tcpTransport.Dial("127.0.0.1:7891", func(options *transport.DialOptions) {
		options.Timeout = time.Second * 15
	})
	if err != nil {
		panic(err)
	}
	//defer client.Close()
	client.Close()
	msg := &transport.Msg{
		Metadata: map[string]string{"type": "tcp"},
		Body:     []byte("hello world"),
	}
	if err := client.Send(msg); err != nil {
		panic(err)
	}
	lab:
	time.Sleep(time.Second * 5)
	revMsg := transport.Msg{}
	if err := client.Recv(&revMsg); err != nil {
		if netOPErr, ok := err.(net.Error); ok {
			if netOPErr.Timeout() {
				goto lab
			}
		}
		panic(err)
	}
	fmt.Printf("rev data metada: %#v, body : %s\n", revMsg.Metadata, string(revMsg.Body))
}
