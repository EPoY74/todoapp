package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	//  Импортирую для инициализации драйвера
	// _ "github.com/mattn/go-sqlite3"

	_ "modernc.org/sqlite"

	"todoapp/internal/models"
)

// // Структура для записи задачи todo
// type todo_record struct {
// 	id               int
// 	data_of_creation sql.NullString
// 	date_max         sql.NullString
// 	todo_text        sql.NullString
// 	is_gone          int
// 	date_of_gone     sql.NullString
// }

// var s todo_record

func init() {

}

// Получаю путь до базовой папки приложения
func GetBaseDirPath() (string, error) {
	cutting := "\\cmd\\console"
	prorgammStartedDir, err := os.Getwd()
	//error GBDP--001
	if err != nil {
		fmt.Printf("GBDP--001: Ошибка при формировании пути к каталогу %v", err)
		fmt.Scanln()
		return "", err
	}
	baseDir := strings.Replace(prorgammStartedDir, cutting, "\\", 1)
	return baseDir, nil
}

// Устанавлваю путь до БД
func SetPathToDB(baseDir string, dbName string) string {
	pathToDB := filepath.Join(baseDir, "internal", "db", dbName)
	fmt.Printf("path: %v", pathToDB)
	return pathToDB
}

// Читаю все записи из таблицы
func ReadAllRec() ([]models.TodoRecord, error) {
	// последний ReadRec--4
	// Функция читает данные из БД
	// id ReadRec--1
	// if id_rec < 0 {
	// 	return 0, errors.New(
	// 		"id ReadRec--1: недопустимый идентификатор записи")
	// }
	fmt.Println("Привет! func ReadAllRec()")
	// basePath, err := GetBaseDirPath()
	// if err != nil {
	// 	//rr--001
	// 	fmt.Printf("rr--001: %v", err)
	// 	fmt.Scanln()
	// 	os.Exit(1)
	// }
	// pathToBD := SetPathToDB(basePath, "ep20231204_todo_cli.db")
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
	todoRecords := []models.TodoRecord{}

	// fmt.Printf("rows type is %T/n", rows)

	for rows.Next() {
		todo := models.TodoRecord{}

		err := rows.Scan(
			&todo.ID,
			&todo.Data_of_creation,
			&todo.Date_max,
			&todo.Todo_text,
			&todo.Is_gone,
			&todo.Date_of_gone,
		)
		if err != nil {
			//id ReadRec--4
			fmt.Printf("ReadRec--4: Не могу прочитать строку из БД %v", err)
			fmt.Scanln()
			continue
		}
		todoRecords = append(todoRecords, todo)
	}

	// for _, todo := range todoRecords {
	// 	fmt.Println(
	// 		todo.ID,
	// 		todo.Data_of_creation,
	// 		todo.Date_max,
	// 		todo.Todo_text,
	// 		todo.Is_gone,
	// 		todo.Date_of_gone,
	// 	)
	// }
	return todoRecords, nil
}
