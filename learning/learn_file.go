package main

// import "fmt"

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	///////РАБОТА С ФАЙЛАМИ

	var pause string

	fmt.Println("Начинаем, введите что-нибудь с клавиатуры")
	fmt.Scanf("%s\n", &pause)

	data := []byte("New text")
	err := os.WriteFile("test.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully.")

	fmt.Println("Текст успешно изменен, введите что-нибудь с клавиатуры")
	fmt.Scanf("%s\n", &pause)

	data, err = os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Содержимое файла:", string(data), ", введите что-нибудь с клавиатуры")
	fmt.Scanf("%s\n", &pause)

	err = os.Remove("test.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
	}

	fmt.Println("Файл удален, введите что-нибудь с клавиатуры")
	fmt.Scanf("%s\n", &pause)

	data = []byte("Hello, Golang!")
	err = os.WriteFile("new_test.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully.")

	fmt.Println("Файл создан, введите что-нибудь с клавиатуры")
	fmt.Scanf("%s\n", &pause)

	err = os.Rename("new_test.txt", "test.txt")
	if err != nil {
		fmt.Println("Error renaming file:", err)
	}

	fmt.Println("Файл переименован, введите что-нибудь с клавиатуры")
	fmt.Scanf("%s\n", &pause)

	///////РАБОТА С СТРИМАМИ

	file, err := os.Open("test_stream.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
