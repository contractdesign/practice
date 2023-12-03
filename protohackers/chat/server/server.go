/*

https://protohackers.com/problem/0

*/

package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"unicode"
)

type User struct {
	name string
	conn *net.TCPConn
}

type Users []User

var users Users

func (users *Users) addUser(user User) {
	*users = append(*users, user)
}

/*
func getOtherUsers(name string) (names []string) {
	names = nil
	for _, user := range users {
		if user.name != name {
			names = append(names, user.name)
		}
	}
	return names
}
*/

func (users *Users) deleteUser(name string) {
	temp := Users{}
	for _, user := range *users {
		if user.name != name {
			temp = append(temp, user)
		}
	}
	*users = temp
}

func handleError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

func validUser(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func writeClientString(s string, conn *net.TCPConn) {
	conn.Write([]byte(s)) // consider io.WriteString
}

func handleConnection(conn *net.TCPConn) {
	defer conn.CloseWrite()
	log.Printf("connection %s open\n", conn.RemoteAddr())

	writeClientString("Welcome to budgetchat! What shall I call you?\n", conn)

	// create scanner to read newlines
	scanner := bufio.NewScanner(conn)
	scanner.Scan()

	name := strings.TrimRight(scanner.Text(), "\r ")

	if len(name) == 0 || !validUser(name) {
		writeClientString("invalid user name. closing connection", conn)
		return
	}

	log.Printf("user %s logged in\n", name)
	users.addUser(User{name, conn})

	for scanner.Scan() {
		log.Println(scanner.Text())
	}

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
