package main

import "time"

/*
Todo Model
*/
type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

/*
Todos is a slice of Todo
*/
type Todos []Todo
