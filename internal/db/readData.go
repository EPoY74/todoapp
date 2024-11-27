package db

import (
	"database/sql"
	"fmt"
	"os"

	//  Импортирую для инициализации драйвера
	// _ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

// Структура для записи задачи todo
type todo_record struct {
	id               int
	data_of_creation sql.NullString
	date_max         sql.NullString
	todo_text        sql.NullString
	is_gone          int
	date_of_gone     sql.NullString
}

var s todo_record

func init() {

}

// ЗАГЛУШКА: Читаю данные из БД
func ReadAllRec() (int, error) {
	// Функция читает данные из БД
	// id ReadRec--1
	// if id_rec < 0 {
	// 	return 0, errors.New(
	// 		"id ReadRec--1: недопустимый идентификатор записи")
	// }
	db_connect, err := sql.Open("sqlite", "ep20231204_todo_cli.db")
	// id ReadRec--2:
	if err != nil {
		fmt.Printf("ReadRec--2: Не могу открыть БД %v", err)
		fmt.Scanln()
		os.Exit(1)
	}
	defer db_connect.Close()

	query := `
		SELECT 
			id,
			data_of_creation,
			date_max,
			todo_text,
			is_gone,
			date_of_gone
		FROM my_todo_list;
	`
	rows, err := db_connect.Query(query)
	// id ReadRec--3
	if err != nil {
		fmt.Printf("ReadRec--3: Не могу прочитать из БД %v", err)
		fmt.Scanln()
		os.Exit(1)
	}
	defer rows.Close()
	todo_records := []todo_record{}

	fmt.Printf("rows type is %T/n", rows)

	for rows.Next() {
		todo := todo_record{}

		err := rows.Scan(
			&todo.id,
			&todo.data_of_creation,
			&todo.date_max,
			&todo.todo_text,
			&todo.is_gone,
			&todo.date_of_gone,
		)
		if err != nil {
			//id ReadRec--4
			fmt.Printf("ReadRec--4: Не могу прочитать строку из БД %v", err)
			fmt.Scanln()
			continue
		}
		todo_records = append(todo_records, todo)
	}

	for _, todo := range todo_records {
		fmt.Println(
			todo.id,
			todo.data_of_creation,
			todo.date_max,
			todo.todo_text,
			todo.is_gone,
			todo.date_of_gone,
		)
	}
	return 1, nil
}
