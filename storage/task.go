package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Task — модель одной задачи
type Task struct {
	ID          int    `json:"ID"`
	Description string `json:"Description"`
	Status      string `json:"Status"`
}

// tasksSlice слайс для хранения задач
var tasksSlice []Task

// LoadTasks читает и парсит JSON-массив задач из файла
func LoadTasks(path string) ([]Task, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("open tasks file: %w", err)
	}

	var tasks []Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("unmarshal tasks json: %w", err)
	}
	return tasks, nil
}

func createTask(description, status string) {
	newTask := Task{
		ID:          len(tasksSlice) + 1,
		Description: description,
		Status:      status,
	}

	tasksSlice = append(tasksSlice, newTask)

	outData, err := json.Marshal(tasksSlice)
	if err != nil {
		log.Fatalf("cannot marshal JSON: %v", err)
	}

	if err := os.WriteFile("storage/tasks.json", outData, 0644); err != nil {
		log.Fatalf("write file: %v", err)
	}

	log.Println("Задача добавлена успешно")
}
