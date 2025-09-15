package main

import (
	"log"
	"net/http"
)

func main() {
	// Создаем мультиплексор, чтобы обрабатывать разные URL
	mux := http.NewServeMux()

	// Обработчик для главной страницы
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	// Обработчик для конвертации
	mux.HandleFunc("/convert", ConvertHandler)

	// Запускаем сервер на порту 8080.
	// Мы будем использовать ваш сервер, но для тестирования
	// на локальной машине лучше использовать стандартный порт.
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
