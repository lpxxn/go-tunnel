package transprot

type Msg struct {
	Metadata map[string]string
	Body     []byte
}

type Socket interface {
	Recv(msg *Msg) error
	Send(msg *Msg) error
	Close() error
	Local() string
	Remote() string
}

type Listener interface {
	Addr() string
	Close() error
	Accept(func(socket Socket)) error
}

//Transport
type Transport interface {
	Dial(addr string, opts ...DialOption) (Client, error)
	Listen(addr string, opts ...ListenOption) (Listener, error)
	String() string
}

type DialOption func(*DialOptions)
type ListenOption func(*ListenOptions)

type Client interface {
	Socket
}
