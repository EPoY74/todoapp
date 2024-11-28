package controllers

import (
	"fmt"
	"todoapp/internal/db"
)

// Обрабатывает запросы для сомандной строки
// err: HCC-001
func HandleConsoleCommand(command string, todo_id int) {

	if command == "show" {
		id_rec, err := db.ReadAllRec()
		if err != nil {
			fmt.Printf("HCC-001: Ошибка: %v\n", err)
		} else {
			fmt.Printf("Прочитано: %v", id_rec)
		}
	}

}
