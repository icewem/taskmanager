package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	// Читаем из файла данные
	// todo Добавить обработку ошибок и сделать вывод читаемым для пользователя
	byteValue, err := os.ReadFile("internal/data.json")

	var tasks []Task

	// Переводим в json
	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		// todo Добавить обработку ошибок и сделать вывод читаемым для пользователя
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
}
