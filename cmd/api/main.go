package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//Заглушка, первый код
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
