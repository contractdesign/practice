package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidStringMixed(t *testing.T) {
	assert.Equal(t, validUser("abcd124"), true, "mixed case")
	assert.Equal(t, validUser("abcd"), true, "mixed case")
	assert.Equal(t, validUser("abcd1234"), true, "mixed case")
	assert.Equal(t, validUser("abcd!"), false, "mixed case")
}

func (users Users) getUsers() (names []string) {
	for _, user := range users {
		names = append(names, user.name)
	}
	return names
}

func TestAddUsers(t *testing.T) {
	var users Users

	users.addUser(User{"a", nil})
	users.addUser(User{"b", nil})
	users.addUser(User{"c", nil})

	assert.Equal(t, users.getUsers(), []string{"a", "b", "c"}, "addUser")

	users.deleteUser("b")

	assert.Equal(t, users.getUsers(), []string{"a", "c"}, "addUser")
	log.Println(users.getUsers())

}
