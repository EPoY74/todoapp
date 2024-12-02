package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"todoapp/internal/db"
)

// Обрабатывает запросы для командной строки
// err: HCC-001
func HandleConsoleCommand(command string, todo_id int) {

	if command == "show" {
		id_rec, err := db.ReadAllRec()
		if err != nil {
			fmt.Printf("HCC-001: Ошибка: %v\n", err)
		} else {
			// fmt.Printf("Прочитано: %v", id_rec)
			for _, todo := range id_rec {
				fmt.Println(
					todo.ID,
					todo.Data_of_creation,
					todo.Date_max,
					todo.Todo_text,
					todo.Is_gone,
					todo.Date_of_gone,
				)
			}
		}
	} else {
		fmt.Println("Пожалуйста, введите поддерживаемую команду.")
		fmt.Printf("Для получения помощи:\n%v -help", filepath.Base(os.Args[0]))

		os.Exit(0)
	}

}
