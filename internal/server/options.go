package server

import (
	"os"
	"time"
)

type Options struct {
	HostName              string
	TcpAddress            string        `flag:"tcp-address" json:"tcp_address"`
	HttpAddress           string        `flag:"http-address" json:"http_address"`
	InactiveClientTimeout time.Duration `flat:"inactive-client-timeout"`
}

func NewOptions() *Options {
	hostName, _ := os.Hostname()

	return &Options{
		HostName:              hostName,
		TcpAddress:            ":7890",
		HttpAddress:           ":7891",
		InactiveClientTimeout: 5 * time.Minute,
	}
}
