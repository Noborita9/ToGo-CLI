package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"togo-cli/internal/domain"
	"togo-cli/internal/service"
)

type CLIHandler struct {
	service *service.TaskService
}

func NewCLIHandler(service *service.TaskService) *CLIHandler {
	return &CLIHandler{service: service}
}

func (h *CLIHandler) AddHandler() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	priority := addCmd.String("priority", "normal", "Task Priority")

	addCmd.Parse(os.Args[2:])
	args := addCmd.Args()

	if len(args) == 0 {
		log.Fatal("Please add a non-empty description")
	}

	newTask, err := h.service.CreateNewTask(args[0], *priority)
	if err != nil {
		log.Fatal(err)
	}

	PrintTask(newTask)
}

func (h *CLIHandler) DeleteHandler() {
	args := os.Args[2:]
	if len(args) < 1 {
		log.Fatal("Please add an ID to be deleted")
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("Please pass an integer ID")
	}
	task, err := h.service.GetTaskByID(id)
	if err != nil {
		log.Fatal(err)
	}
	err = h.service.Delete(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Task %d: %s deleted Successfully\n", id, task.Description)
}

func (h *CLIHandler) ListHandler() {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	status := listCmd.String("status", "", "Task Status Completed|Incompleted")

	listCmd.Parse(os.Args[2:])
	
	var tasks []*domain.Task
	var err error
	

	if *status != "" {
		tasks, err = h.service.GetTasksByStatus(*status)		
	} else {
		tasks, err = h.service.GetTasks()
	}

	if err != nil {
		log.Fatal(err)
	}

	PrintTitles()
	for _, task := range tasks {
		PrintTask(task)
	}
}

func (h *CLIHandler) CompleteHandler() {
	args := os.Args[2:]
	if len(args) < 1 {
		log.Fatal("Please add an ID to be completed")
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("Please pass an integer ID")
	}
	task, err := h.service.GetTaskByID(id)
	if err != nil {
		log.Fatal("Could not find task with ID ", id)
	}
	err = h.service.Complete(task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Task %d: %s Completed Successfully\n", id, task.Description)
}
