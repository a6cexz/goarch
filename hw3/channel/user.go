package channel

import "fmt"

// User struct
type User struct {
	SubscriberDefault
	Username string
}

// NewUser creates new user
func NewUser(username string) *User {
	return &User{
		Username: username,
	}
}

// OnReceive hanles new received message
func (u *User) OnReceive(msg string) {
	fmt.Printf("MESSAGE GOT: %s: %s\n", u.Username, msg)
}

// GetID returns user ID
func (u *User) GetID() string {
	return u.Username
}
