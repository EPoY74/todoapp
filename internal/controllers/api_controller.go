package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todoapp/internal/db"
)

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "applications/json")
		todoRecords, err := db.ReadAllRec()
		//HT--001
		if err != nil {
			fmt.Printf("HT-001: Ошибка: %v\n", err)
		}

		data, err := json.Marshal(todoRecords)
		if err != nil {
			return
		}
		// for
		json.NewEncoder(w).Encode(data)
		return
	}
	w.WriteHeader((http.StatusMethodNotAllowed))
}
