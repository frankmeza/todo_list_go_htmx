package main

import (
	"log"
	"net/http"

	"github.com/frankmeza/todo_list_go_htmx/pkg/actions"
	"github.com/frankmeza/todo_list_go_htmx/pkg/utils"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", utils.RenderHome)

	// mux.HandleFunc("/add_task", actions.AddTask)
	mux.HandleFunc("/increment_counter", actions.Increment)
	mux.HandleFunc("/decrement_counter", actions.Decrement)
	mux.HandleFunc("/reset_counter", actions.ResetCounter)

	println("server running on localhost:3333")
	log.Fatal(http.ListenAndServe(":3333", mux))
}
