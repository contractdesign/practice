/*

https://protohackers.com/problem/0

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"unicode"
)

// structure with user data: name and their socket
type User struct {
	name string
	conn *net.TCPConn
}

// TODO: handle duplicate user names. maybe a map is better than a slice
type Users []User

// global variable: list of users connected to server
var users Users

// add the user to the list of  users
func (users *Users) addUser(user User) {
	*users = append(*users, user)
}

// return users with a different name
func (users Users) getOtherUsers(name string) (otherUsers Users) {
	for _, user := range users {
		if user.name != name {
			otherUsers = append(otherUsers, user)
		}
	}
	return otherUsers
}

// given the name, delete the indicated user from active users
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

// valid usernames are comprised of letters and/or digits
func validName(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func writeClientString(conn *net.TCPConn, s string) {
	conn.Write([]byte(s)) // consider io.WriteString
	log.Printf(s)
}

func handleConnection(conn *net.TCPConn) {
	defer conn.CloseWrite()
	log.Printf("connection %s open\n", conn.RemoteAddr())

	writeClientString(conn, "Welcome to budgetchat! What shall I call you?\n")

	// create scanner to read newlines
	scanner := bufio.NewScanner(conn)
	scanner.Scan()

	// remove trailing spaces
	name := strings.TrimRight(scanner.Text(), "\r ")

	if len(name) == 0 || !validName(name) {
		writeClientString(conn, "invalid user name. closing connection\n")
		return
	}

	log.Printf("user %s logged in\n", name)
	users.addUser(User{name, conn})

	// TODO fix trailing commas
	// broadcast joining to other users
	var other_names = ""
	for _, user := range users.getOtherUsers(name) {
		other_names += user.name + ","
	}
	writeClientString(conn, "* The room contains: "+other_names+"\n")

	// wait for messages
	for scanner.Scan() {
		text := scanner.Text()
		for _, user := range users.getOtherUsers(name) {
			writeClientString(user.conn, fmt.Sprintf("[%s] %s\n", name, text))
		}
	}

	// exiting the for loop means that the user disconnected
	log.Printf("connection %s closed\n", conn.RemoteAddr())
	for _, user := range users.getOtherUsers(name) {
		writeClientString(user.conn, fmt.Sprintf("* %s has left the room\n", name))
	}
	users.deleteUser(name)
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
