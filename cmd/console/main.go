package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"todoapp/internal/controllers"
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
		fmt.Fprintf(
			os.Stderr,
			"For examlpe: %v -command=\"your_command\"",
			filepath.Base(os.Args[0]))
	}

	flag.StringVar(&command, "command", "none", "The command for ToDo")
	flag.IntVar(&todo_id, "id", 0, "Id for showing todo")
	// flag.Float64Var(&height, "height", 180, "Ваш рост (см)")

	flag.Parse()
	if len(os.Args) < 2 {
		flag.Usage()
		fmt.Scanln()
		os.Exit(1)
	}

	fmt.Printf("Command is %s\n", command)

	controllers.HandleConsoleCommand(command, todo_id)

}
