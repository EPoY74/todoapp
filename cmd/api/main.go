package main

import (
	"fmt"
	"log"
	"net/http"
	"todoapp/internal/controllers"
)

func main() {

	http.HandleFunc("/get_tasks", controllers.HandleTasks)

	//Заглушка, первый код
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
