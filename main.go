package main

import (
	"socket"
	"log"
)

const (
	hostname = "0.0.0.0:8080"
)

func main() {
	ch := make(chan socket.Conn)

	socket.
		Create().
		SetOnListen(func(s *(socket.Socket), err error) {
			if err != nil {
				log.Fatalln(err)
			}
			log.Printf("Listening on %v\n", hostname)
		}).
		SetOnConnection(func(c socket.Conn, err error) {
			if err != nil {
				log.Panicln(err)
			}
			ch <- c
		}).
		BeginListen(hostname)

	conns := make(map[socket.Conn]byte)

	for {
		conn := <- ch	
		go func(conn socket.Conn) {
			conns[conn] = 0
			for {
				msg := make([]byte, 1024)
				_, err := conn.Read(msg)
				if err != nil {
					break
				}
				for c := range conns {
					c.Write(msg)
				}
			}
			delete(conns, conn)
		}(conn)
	}
}