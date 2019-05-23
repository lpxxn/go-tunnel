package main

import (
	"fmt"
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
		defer sock.Close()

		for {
			var m transport.Msg
			if err := sock.Recv(&m); err != nil {
				fmt.Println("socket Rev error", err)
				return
			}
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
