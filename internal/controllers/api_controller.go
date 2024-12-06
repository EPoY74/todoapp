package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todoapp/internal/db"
	"todoapp/internal/models"
)

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	var (
		todoRecords []models.TodoRecord
		err         error
	)

	if r.Method == http.MethodGet {
		// text/plain"
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		// w.Header().Set("Content-Type", "applications/json")
		// w.Header().Set("Content-Disposition", "inline")
		url_recs_id := r.URL.Query().Get("recs_id")
		if url_recs_id != "" {
			recs_id, err := strconv.Atoi(url_recs_id)
			if err != nil {
				//HT--003
				fmt.Printf("HT-003: Ошибка: %v\n", err)
			}
			if recs_id > 0 {
				fmt.Printf("Получаю запись %v \n", recs_id)
				fmt.Fprintf(w, "Получаю запись %v \n", recs_id)
				// w.Write(data)
				return
			}
		} else {
			todoRecords, err = db.ReadAllRec()
			if err != nil {
				fmt.Printf("HT-001: Ошибка: %v\n", err)
			}

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

		//HT--001

		// data, err := json.Marshal(todoRecords)
		// if err != nil {
		// 	return
		// }
		// // for
		// fmt.Printf("%v ", data)

		// err = json.NewEncoder(w).Encode(todoRecords)

	}
	w.WriteHeader((http.StatusMethodNotAllowed))
}
