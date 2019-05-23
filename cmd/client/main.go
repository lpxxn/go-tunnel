package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
)

func main() {
	listener, err := net.Listen("tcp", ":7890")
	if err != nil {
		panic(err)
	}
	go func() {
		for  {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}
			
		}
	}()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	fmt.Println("running")
	<-ch
}
