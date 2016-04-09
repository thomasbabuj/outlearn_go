package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

/*
TodoShow handler
*/
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show :", todoId)
}
