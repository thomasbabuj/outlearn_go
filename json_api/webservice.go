package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/*
Todo Model
*/
type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

/*
Todos is a slice of Todo
*/
type Todos []Todo

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)
	log.Fatal(http.ListenAndServe(":8080", router))
}

/*
Index handler
*/
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Weclome!")
}

/*
TodoIndex handler
*/
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	//Static slice of todos
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

/*
TodoShow handler
*/
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show :", todoId)
}
