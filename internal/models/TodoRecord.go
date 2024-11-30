package models

import "database/sql"

// Структура для записи задачи todo
type TodoRecord struct {
	ID               int
	Data_of_creation sql.NullString
	Date_max         sql.NullString
	Todo_text        sql.NullString
	Is_gone          int
	Date_of_gone     sql.NullString
}
