package socket

import "net"

type Conn interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
}

type Socket struct {
	onListen func(*Socket, error)
	onConnection func(Conn, error)
}

func Create() *Socket {
	return &Socket {}
}

func (s *Socket) SetOnListen (f func(*Socket, error)) *Socket {
	s.onListen = f
	return s
}

func (s *Socket) SetOnConnection (f func(Conn, error)) *Socket {
	s.onConnection = f
	return s
}

func (s *Socket) BeginListen (hostname string) {
	go func() {
		ln, err := net.Listen("tcp", hostname)
		s.onListen(s, err)
		for {
			conn, err := ln.Accept()
			s.onConnection(conn, err)
		}
	}()
}