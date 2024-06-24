# Имя исполняемого файла
BINARY_NAME=tasks

# Список всех файлов исходного кода
MAIN_SOURCES=$(wildcard *.go)
TASK_SOURCES=$(wildcard task/*.go)

# Цель по умолчанию - сборка проекта
all: build

# Инициализация модуля Go
init:
	@echo "Initializing Go module..."
	@go mod init main || true
	@go mod tidy

# Цель сборки проекта
build: init
	@echo "Building the project..."
	@go build -o $(BINARY_NAME) $(SOURCES)

# Цель для очистки
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

# Цель для запуска проекта
run: build
	@echo "Running the project..."
	@./$(BINARY_NAME) $(ARGS)

.PHONY: all build clean run test init
