package cli

import (
	"fmt"
	"togo-cli/internal/domain"
)

func getIdLenght(id int) int {
	length := 1
	id /= 10
	for id > 0 {
		length++
		id /= 10
	}
	return length
}

func PrintTask(t *domain.Task) {
	idLen := getIdLenght(t.Id)
	fmt.Printf("%v", t.Id)
	for i := 0; i < 9-idLen; i++ {
		fmt.Print(" ")
	}

	fmt.Print(t.Priority)
	for i := 0; i < 16-len(t.Priority); i++ {
		fmt.Print(" ")
	}

	completedText := "P"
	if t.Completed {
		completedText = "C"
	}

	fmt.Printf("[%v]              %v \n", completedText, t.Description)
}

func PrintTitles() {
	fmt.Println("ID  ||   Priority   ||   Completed   ||   Description ")
	fmt.Println("----  --------------  ---------------  ---------------")
}

func PrintHelp() {
	fmt.Println("Task Manager")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  task add [--priority=high|normal|low] <description>")
	fmt.Println("  task list [--status=pending|completed]")
	fmt.Println("  task complete <id>")
	fmt.Println("  task delete <id>")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  task add --priority=high \"Call dentist\"")
	fmt.Println("  task add \"Buy groceries\"")
	fmt.Println("  task list --status=pending")
}
