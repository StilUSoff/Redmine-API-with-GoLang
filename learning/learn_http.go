package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Обработка корневого пути - отдаём index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Обработка статики (CSS, JS, картинки и т.д.)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Дополнительный путь
	http.HandleFunc("/hello_world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	fmt.Println("Сервер запущен на порту 80")
	http.ListenAndServe(":80", nil)
}
