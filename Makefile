# Имя исполняемого файла
BINARY_NAME=tasks

# Список всех файлов исходного кода
SOURCES=$(wildcard *.go)

# Цель по умолчанию - сборка проекта
all: build

# Цель сборки проекта
build:
	@echo "Building the project..."
	@go build -o $(BINARY_NAME) $(SOURCES)

# Цель для очистки
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

# Цель для запуска проекта
run: build
	@echo "Running the project..."
	@./$(BINARY_NAME)

.PHONY: all build clean run test
