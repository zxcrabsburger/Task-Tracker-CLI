package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

var Tasks []Task
var idCounter int = 1

func markToDo(TaskID int) Task {
	for i, task := range Tasks {
		if task.ID == TaskID {
			Tasks[i].Status = "todo"
			Tasks[i].UpdatedAt = time.Now().Format(time.RFC1123)
		}
	}
	saveTasks()
	return Tasks[TaskID-1]
}

func markInProgress(TaskID int) Task {
	for i, task := range Tasks {
		if task.ID == TaskID {
			Tasks[i].Status = "in-progress"
			Tasks[i].UpdatedAt = time.Now().Format(time.RFC1123)
		}
	}
	saveTasks()
	return Tasks[TaskID-1]
}

func markDone(TaskID int) Task {
	for i, task := range Tasks {
		if task.ID == TaskID {
			Tasks[i].Status = "done"
			Tasks[i].UpdatedAt = time.Now().Format(time.RFC1123)
		}
	}
	saveTasks()
	return Tasks[TaskID-1]
}

func loadTasks() error {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &Tasks)
	if err != nil {
		return err
	}
	if len(Tasks) > 0 {
		idCounter = Tasks[len(Tasks)-1].ID + 1
	}
	return nil
}

func saveTasks() error {
	data, err := json.MarshalIndent(Tasks, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func getTasks(status string) []Task {
	switch status {
	case "todo":
		var todoTasks []Task
		for _, task := range Tasks {
			if task.Status == "todo" {
				todoTasks = append(todoTasks, task)
			}
		}
		return todoTasks
	case "done":
		var doneTasks []Task
		for _, task := range Tasks {
			if task.Status == "done" {
				doneTasks = append(doneTasks, task)
			}
		}
		return doneTasks
	case "in-progress":
		var inProgressTasks []Task
		for _, task := range Tasks {
			if task.Status == "in-progress" {
				inProgressTasks = append(inProgressTasks, task)
			}
		}
		return inProgressTasks
	default:
		return Tasks
	}
}

func updateTask(TaskID int, newDescription string) Task {
	for i, task := range Tasks {
		if task.ID == TaskID {
			Tasks[i].Description = newDescription
			Tasks[i].UpdatedAt = time.Now().Format(time.RFC1123)
		}
	}
	saveTasks()
	return Tasks[TaskID-1]
}

// TODO: добавить проверку на неиспользуемые ID при добавлении задачи
func addTask(description string) Task {
	var task = Task{
		ID:          idCounter,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format(time.RFC1123),
		UpdatedAt:   "none",
	}
	Tasks = append(Tasks, task)
	idCounter++
	msg := fmt.Sprintf("Задача добавлена успешно (ID: %d)", task.ID)
	fmt.Println(msg)
	saveTasks()
	return task
}

func deleteTask(TaskID int) {
	for i, task := range Tasks {
		if task.ID == TaskID {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			msg := fmt.Sprintf("Задача удалена успешно (ID: %d)", TaskID)
			fmt.Println(msg)
		}
	}
	saveTasks()
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add <описание>          - Добавить задачу")
	fmt.Println("  update <id> <описание>  - Обновить задачу")
	fmt.Println("  delete <id>                - Удалить задачу")
	fmt.Println("  mark-todo <id>             - Пометить задачу как 'to-do'")
	fmt.Println("  mark-in-progress <id>      - Пометить задачу как 'in-progress'")
	fmt.Println("  mark-done <id>             - Пометить задачу как 'done'")
	fmt.Println("  list [статус]              - Показать задачи (опционально: todo, in-progress, done)")
}

func main() {
	loadTasks()
	addTask("сегодня дрочим")
}
