package main

import "sync"

type User struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Books    []string `json:"books"`
	Weight   float64  `json:"weight"`
}

type Readed struct {
	M      sync.Mutex
	Number int
}
