package main

import (
	"database/sql"
	"log"
	"os"

	// "time"
	"togo-cli/internal/cli"
	"togo-cli/internal/service"

	// "togo-cli/internal/domain"
	"togo-cli/internal/repository"
	"togo-cli/pkg/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Please Execute a command: add, list, delete, complete")
	}

	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	err = database.CreateTables(db)
	if err != nil {
		log.Fatal(err)
	}

	taskService := service.NewTaskService(
		repository.NewSQLiteTaskRepository(db),
	)

	taskHandler := cli.NewCLIHandler(taskService)

	switch os.Args[1] {
	case "add":
		{
			taskHandler.AddHandler()
		}
	case "delete":
		{
			taskHandler.DeleteHandler()
		}
	case "list":
		{
			taskHandler.ListHandler()
		}
	case "complete":
		{
			taskHandler.CompleteHandler()
		}
	default:
		cli.PrintHelp()
	}
}
