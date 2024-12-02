package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todoapp/internal/db"
)

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// text/plain"
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		// w.Header().Set("Content-Type", "applications/json")
		// w.Header().Set("Content-Disposition", "inline")
		todoRecords, err := db.ReadAllRec()
		//HT--001
		if err != nil {
			fmt.Printf("HT-001: Ошибка: %v\n", err)
		}

		// data, err := json.Marshal(todoRecords)
		// if err != nil {
		// 	return
		// }
		// // for
		// fmt.Printf("%v ", data)

		// err = json.NewEncoder(w).Encode(todoRecords)
		data, err := json.MarshalIndent(todoRecords, "", " ")
		if err != nil {
			http.Error(w,
				//HT--002
				fmt.Sprintf("HT--002: Ошибка при кодировании JSON: %v", err),
				http.StatusInternalServerError)

			return
		}

		w.Write(data)
		return
	}
	w.WriteHeader((http.StatusMethodNotAllowed))
}
