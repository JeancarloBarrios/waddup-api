package main

import (
	"time"
)

type User struct {
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	CreationDate time.Time `json:"creationdate"`
	Active       bool      `json:"active"`
	Location     string    `json:"location"`
}

type Users []User
