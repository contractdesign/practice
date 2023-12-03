/*

https://protohackers.com/problem/0

*/

package main

import (
	"io"
	"log"
	"net"
)

func handleError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

func handleConnection(conn *net.TCPConn) {
	log.Printf("connection %s open\n", conn.RemoteAddr())

	defer conn.CloseWrite()

	io.Copy(conn, conn)
	log.Printf("connection %s closed\n", conn.RemoteAddr())
}

func main() {
    const port = 8080
	addr := &net.TCPAddr{
		IP:   []byte{},
		Port: port,
		Zone: "",
	}
	log.Printf("listening on %d\n", port)
	ln, err := net.ListenTCP("tcp", addr)
	handleError(err, "Listen")
	for {
		conn, err := ln.AcceptTCP()
		handleError(err, "Accept")
		go handleConnection(conn)
	}
}
