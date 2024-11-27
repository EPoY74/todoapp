package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	db "todoapp/internal/db"
)

func main() {
	// argCount := len(os.Args)
	// fmt.Printf("Количество элементов %d\n", argCount)
	// fmt.Println("Аргументы: ", os.Args)

	// for _, value := range os.Args {
	// 	fmt.Printf(" %s\n", value)
	// }

	var (
		command string
		todo_id int
		// height  float64
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: \n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	flag.StringVar(&command, "command", "show", "The command for ToDo")
	flag.IntVar(&todo_id, "id", 0, "Id for sowing todo")
	// flag.Float64Var(&height, "height", 180, "Ваш рост (см)")

	flag.Parse()
	if len(os.Args) < 2 {
		flag.Usage()
		fmt.Scanln()
		os.Exit(1)
	}

	fmt.Printf("Command is %s\n", command)
	// fmt.Printf("Age: %d\n", age)
	// fmt.Printf("Height: %.2f\n", height)

	if command == "show" {
		id_rec, err := db.ReadAllRec()
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Printf("Прочитано: %v", id_rec)
		}
	}

}
