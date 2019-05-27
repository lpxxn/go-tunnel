package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"

	"github.com/lpxxn/go-tunnel/transport"
)

func main() {
	tcpTransport := transport.NewTcpTransport()
	listener, err := tcpTransport.Listen(":7891")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fn := func(sock transport.Socket) {
		defer func() {
			fmt.Println("close socket")
			sock.Close()
		}()

		for {
			var m transport.Msg
			if err := sock.Recv(&m); err != nil {

				if err == io.EOF {
					fmt.Println("io.EOF")
				}
				if netOPErr, ok := err.(net.Error); ok {
					if netOPErr.Timeout() {
						fmt.Println("just timeout continue, err info: ", err)
						continue
					}
				}
				fmt.Println("socket Rev error", err)
				return
			}
			fmt.Println("received: ", m)
			if err := sock.Send(&m); err != nil {
				fmt.Println("socket Send error", err)
				return
			}
		}
	}
	go func() {
		if err := listener.Accept(fn); err != nil {
			fmt.Printf("Unexpected accept err: %#v\n", err)
		}
	}()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	fmt.Println("service is running")
	<-ch
	fmt.Println("stopped service")
}
