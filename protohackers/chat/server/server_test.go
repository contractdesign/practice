package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidStringMixed(t *testing.T) {
	assert.Equal(t, validName("abcd124"), true, "mixed case")
	assert.Equal(t, validName("abcd"), true, "mixed case")
	assert.Equal(t, validName("1234"), true, "digits")
	assert.Equal(t, validName("1234 "), false, "digits with space")
	assert.Equal(t, validName("abcd!"), false, "invalid")
	assert.Equal(t, validName("abcd_adfd"), false, "invalid")
}

func (users Users) getUsers() (names []string) {
	for _, user := range users {
		names = append(names, user.name)
	}
	return names
}

func TestAddUsers(t *testing.T) {
	var users Users
	names := []string{"a", "b", "c"}

	for _, name := range names {
		users.addUser(User{name, nil})
	}

	assert.Equal(t, users.getUsers(), names, "addUser")

	assert.Equal(t, users.getOtherUsers("b"),
		Users{User{"a", nil}, User{"c", nil}}, "OtherUsers")

	users.deleteUser("b")
	assert.Equal(t, users.getUsers(), []string{"a", "c"}, "addUser")

	users.addUser(User{"e", nil})
	assert.Equal(t, users.getUsers(), []string{"a", "c", "e"}, "addUser")

}
