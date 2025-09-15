package main

import (
	"image/jpeg"
	"log"
	"net/http"

	"golang.org/x/image/webp"
)

// ConvertHandler обрабатывает загрузку и конвертацию файла
func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// 1. Получаем файл из запроса
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error getting file from form: %v", err)
		http.Error(w, "Error getting file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 2. Декодируем WebP
	img, err := webp.Decode(file)
	if err != nil {
		log.Printf("Error decoding webp: %v", err)
		http.Error(w, "Invalid WebP file", http.StatusBadRequest)
		return
	}

	// 3. Отправляем JPG файл в ответе
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Disposition", `attachment; filename="converted.jpg"`)

	if err := jpeg.Encode(w, img, &jpeg.Options{Quality: 100}); err != nil {
		log.Printf("Error encoding to jpeg: %v", err)
		http.Error(w, "Error converting image", http.StatusInternalServerError)
		return
	}
}
