package storage

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
)

const (
	StatusNew        = "new"
	StatusInProgress = "in_progress"
	StatusDone       = "done"
)

var validStatuses = []string{
	StatusNew,
	StatusInProgress,
	StatusDone,
}

// Task базовая структура задачи
type Task struct {
	ID          string `json:"ID"`
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

func CreateTask(description, status string, taskList []Task) error {

	if !isValidStatus(status) {
		log.Printf("Статус %v не доступен, доступные варианты: %s\n", status, validStatuses)
		return nil
	}

	id := uuid.New().String()

	newTask := Task{
		ID:          id,
		Description: description,
		Status:      status,
	}

	tasksSlice = append(taskList, newTask)

	outData, err := json.Marshal(tasksSlice)
	if err != nil {
		log.Fatalf("cannot marshal JSON: %v", err)
	}

	if err := os.WriteFile("storage/tasks.json", outData, 0644); err != nil {
		log.Fatalf("write file: %v", err)
	}

	log.Println("Задача добавлена успешно")

	return nil
}

func isValidStatus(s string) bool {
	for _, vs := range validStatuses {
		if vs == s {
			return true
		}
	}
	return false
}
