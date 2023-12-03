/*

https://protohackers.com/problem/0

*/

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func handleError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

func sendData() {
	addr := &net.TCPAddr{
		IP:   []byte{},
		Port: 8080,
		Zone: "",
	}

	// send data
	data := []byte{1, 2, 3, 4, 5}
	conn, err := net.DialTCP("tcp", nil, addr)
	handleError(err, "Dial")

	num_iterations := rand.Intn(5) + 1
	for i := 0; i < num_iterations; i++ {
		conn.Write(data)
		num_seconds := rand.Intn(5) + 1
		time.Sleep(time.Duration(num_seconds) * time.Second)
		log.Println("data sent")
	}
	conn.Write(data)
	conn.CloseWrite()

	// receive echoed data
	log.Printf("data received from %s\n", conn.LocalAddr())
	buf := make([]byte, 0, 1024)
	tmp := make([]byte, 256)
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			// TODO change to EOF
			break
		}
		buf = append(buf, tmp[:n]...)
	}
	fmt.Println(buf)
	conn.CloseRead()
}

func main() {
	num_iterations := rand.Intn(5) + 5
	for i := 0; i < num_iterations; i++ {
		go sendData()
	}
	time.Sleep(20 * time.Second)
}
