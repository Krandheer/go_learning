package main

import (
	"fmt"
)

type User struct {
	UserName string
	UserID   string
}

type UserNotFoundError struct {
	UserID string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user with id '%s' not found", e.UserID)
}

var users = map[string]User{
	"123": {"randheer", "123"},
	"321": {"ravi", "321"},
}

func FindUser(id string) (*User, error) {
	user, ok := users[id]
	if !ok {
		return nil, &UserNotFoundError{UserID: id}
	}
	return &user, nil
}

func getUser(id string) (*User, error) {
	u, err := FindUser(id)
	if err != nil {
		return nil, fmt.Errorf("could not get user: %w", err)
	}
	return u, nil
}

func main() {
	u, err := getUser("123")
	if err != nil {
		fmt.Printf("error occurred: %v\n", err)
		return
	}
	fmt.Println(*u)
}
