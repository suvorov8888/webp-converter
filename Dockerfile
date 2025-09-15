# Этап 1: Сборка приложения
# Используем официальный образ Go для сборки
FROM golang:1.22-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для управления зависимостями
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем исполняемый файл для Linux
# CGO_ENABLED=0 делает бинарник статическим, не зависящим от библиотек системы
# -ldflags="-s -w" уменьшает размер файла, убирая символы отладки
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /webp-converter .

# Этап 2: Создание финального образа
# Используем минимальный образ Alpine, который весит всего несколько мегабайт
FROM alpine:3.19

# Устанавливаем рабочую директорию
WORKDIR /

# Копируем только исполняемый файл и папку с шаблонами из первого этапа
COPY --from=builder /webp-converter /webp-converter
COPY --from=builder /app/templates /templates

# Открываем порт 8080 для нашего приложения
EXPOSE 8080

# Команда, которая запускается при старте контейнера
ENTRYPOINT ["/webp-converter"]