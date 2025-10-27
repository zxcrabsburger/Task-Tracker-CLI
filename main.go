package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
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
	msg := fmt.Sprintf("Задача обновлена успешно (ID: %d)", TaskID)
	fmt.Println(msg)
	saveTasks()
	return Tasks[TaskID-1]
}

func findUnusedID() int {
	used := make(map[int]bool, len(Tasks))
	max := 0
	for _, t := range Tasks {
		used[t.ID] = true
		if t.ID > max {
			max = t.ID
		}
	}
	for i := 1; i <= max; i++ {
		if !used[i] {
			return i
		}
	}
	return max + 1
}

func addTask(description string) Task {
	id := findUnusedID()

	var task = Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format(time.RFC1123),
		UpdatedAt:   "none",
	}

	Tasks = append(Tasks, task)
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
	fmt.Println("  exit                      - Выйти из программы")
}

func main() {
	loadTasks()
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.SplitN(input, " ", 3)

		switch args[0] {
		case "add":
			if len(args) < 2 {
				fmt.Println("Ошибка: описание задачи отсутствует.")
				continue
			}
			addTask(args[1])
		case "update":
			if len(args) < 3 {
				fmt.Println("Ошибка: недостаточно аргументов для обновления задачи.")
				continue
			}
			id := 0
			fmt.Sscanf(args[1], "%d", &id)
			updateTask(id, args[2])
		case "delete":
			if len(args) < 2 {
				fmt.Println("Ошибка: ID задачи отсутствует.")
				continue
			}
			id := 0
			fmt.Sscanf(args[1], "%d", &id)
			deleteTask(id)
		case "mark-todo":
			if len(args) < 2 {
				fmt.Println("Ошибка: ID задачи отсутствует.")
				continue
			}
			id := 0
			fmt.Sscanf(args[1], "%d", &id)
			markToDo(id)
		case "mark-in-progress":
			if len(args) < 2 {
				fmt.Println("Ошибка: ID задачи отсутствует.")
				continue
			}
			id := 0
			fmt.Sscanf(args[1], "%d", &id)
			markInProgress(id)
		case "mark-done":
			if len(args) < 2 {
				fmt.Println("Ошибка: ID задачи отсутствует.")
				continue
			}
			id := 0
			fmt.Sscanf(args[1], "%d", &id)
			markDone(id)
		case "list":
			status := ""
			if len(args) == 2 {
				status = args[1]
			}
			tasks := getTasks(status)
			for _, task := range tasks {
				fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
					task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
			}
		case "help":
			printUsage()
			continue
		default:
			fmt.Println("Неизвестная команда. Введите 'help' для списка команд.")
		}

		if args[0] == "exit" {
			break
		}
	}
}
