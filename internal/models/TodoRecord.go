package models

import "database/sql"

// Структура для записи задачи todo
type TodoRecord struct {
	ID               int            `json:"id"`
	Data_of_creation sql.NullString `json:"data_of_creation"`
	Date_max         sql.NullString `json:"date_max"`
	Todo_text        sql.NullString `json:"todo_text"`
	Is_gone          int            `json:"is_gone"`
	Date_of_gone     sql.NullString `json:"date_of_gone"`
}
