package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sandbox/internal/service"
	"strconv"
	"strings"
)

type ToDoHandler struct {
	Service *service.TodoService
}

type TodoInput struct {
	S string `json:"title"`
}

func NewTodoHandler(s *service.TodoService) *ToDoHandler {
	return &ToDoHandler{
		Service: s,
	}
}

func (t *ToDoHandler) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== health ===")
	fmt.Println("Метод:", r.Method)
	fmt.Println("Путь:", r.URL.Path)

	if r.Method != http.MethodGet {
		fmt.Println("Ошибка: неправильный метод для /ping")
		http.Error(w, "Не правильный метод", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Go!"))
	fmt.Println("Ответ отправлен: Hello Go!")
}

func getIdFromPath(path string, prefix string) (int, error) {
	idStr := strings.TrimPrefix(path, prefix)
	return strconv.Atoi(idStr)
}

// func getToDoById(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("=== getToDoById ===")
// 	fmt.Println("Метод:", r.Method)
// 	fmt.Println("Путь:", r.URL.Path)

// 	if r.Method != http.MethodGet {
// 		fmt.Println("Ошибка: неправильный метод")
// 		http.Error(w, "Не правильный метод", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	id, err := getIdFromPath(r.URL.Path, "/todo/")
// 	if err != nil {
// 		fmt.Println("Ошибка: не удалось преобразовать id из пути:", r.URL.Path)
// 		http.Error(w, "Не правильно указан id", http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Println("Ищу ToDo по id:", id)

// 	for _, v := range todos {
// 		if v.Id == id {
// 			fmt.Println("ToDo найден:", v)

// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusOK)

// 			err := json.NewEncoder(w).Encode(v)
// 			if err != nil {
// 				fmt.Println("Ошибка при отправке JSON:", err)
// 			}

// 			return
// 		}
// 	}

// 	fmt.Println("ToDo не найден, id:", id)
// 	http.Error(w, "Запись не нашлась", http.StatusNotFound)
// }

func (h *ToDoHandler) CreateToDo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== createToDo ===")
	fmt.Println("Метод:", r.Method)
	fmt.Println("Путь:", r.URL.Path)

	if r.Method != http.MethodPost {
		fmt.Println("Ошибка: неправильный метод")
		http.Error(w, "Не правильный метод", http.StatusMethodNotAllowed)
		return
	}

	var t TodoInput
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		fmt.Println("Ошибка декодирования body:", err)
		http.Error(w, "Плохое тело", http.StatusBadRequest)
		return
	}
	todo, err :=	h.Service.Create(r.Context(), t.S)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	fmt.Println("Создана новая ToDo:", todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		fmt.Println("Ошибка при отправке JSON:", err)
	}
}

func (h *ToDoHandler) GetToDoById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("Неправильный метод")
		http.Error(w, "Не правильный метод", http.StatusMethodNotAllowed)
		return
	}

	id, err := getIdFromPath(r.URL.Path, "/todo/")
	if err != nil {
		fmt.Println("Неправильный id")
		http.Error(w, "Неправильный id", http.StatusBadRequest)
		return
	}

	todo, err := h.Service.GetById(r.Context(), id)
	if err != nil {
		fmt.Println("Задача не найдена")
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		fmt.Println("Ошибка при отправке JSON:", err)
	}
}


// func deleteToDo(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("=== deleteToDo ===")
// 	fmt.Println("Метод:", r.Method)
// 	fmt.Println("Путь:", r.URL.Path)

// 	if r.Method != http.MethodDelete {
// 		fmt.Println("Ошибка: неправильный метод")
// 		http.Error(w, "Не правильный метод", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	id, err := getIdFromPath(r.URL.Path, "/delete/")
// 	if err != nil {
// 		fmt.Println("Ошибка: не удалось преобразовать id из пути:", r.URL.Path)
// 		http.Error(w, "Не правильно указан id", http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Println("Удаляю ToDo с id:", id)

// 	index := -1
// 	for i, t := range todos {
// 		if t.Id == id {
// 			index = i
// 			break
// 		}
// 	}

// 	if index == -1 {
// 		fmt.Println("ToDo не найден, id:", id)
// 		http.Error(w, "ToDo не найден", http.StatusNotFound)
// 		return
// 	}

// 	deleted := todos[index]
// 	todos = append(todos[:index], todos[index+1:]...)

// 	fmt.Println("Удалена ToDo:", deleted)
// 	fmt.Println("Всего задач осталось:", len(todos))

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Удалено"))
// }

// func updateToDo(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("=== updateToDo ===")
// 	fmt.Println("Метод:", r.Method)
// 	fmt.Println("Путь:", r.URL.Path)

// 	if r.Method != http.MethodPut {
// 		fmt.Println("Ошибка: неправильный метод")
// 		http.Error(w, "Не правильный метод", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	id, err := getIdFromPath(r.URL.Path, "/update/")
// 	if err != nil {
// 		fmt.Println("Ошибка: не удалось преобразовать id из пути:", r.URL.Path)
// 		http.Error(w, "Неверный id", http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Println("Обновляю ToDo с id:", id)

// 	var updated ToDo
// 	err = json.NewDecoder(r.Body).Decode(&updated)
// 	if err != nil {
// 		fmt.Println("Ошибка чтения body:", err)
// 		http.Error(w, "Ошибка чтения body", http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Println("Новые данные:", updated)

// 	found := false
// 	for i, t := range todos {
// 		if t.Id == id {
// 			todos[i].Title = updated.Title
// 			todos[i].Completed = updated.Completed
// 			found = true

// 			fmt.Println("ToDo после обновления:", todos[i])
// 			break
// 		}
// 	}

// 	if !found {
// 		fmt.Println("ToDo не найден, id:", id)
// 		http.Error(w, "ToDo не найден", http.StatusNotFound)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Обновлено"))
// }