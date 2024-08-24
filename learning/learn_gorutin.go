package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Hello concurrent world")

	// если не подождать, то программа закончится, не успев вывести сообщение
	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 5; i++ {
		go func(i int) { fmt.Println(i) }(i)
	}

	time.Sleep(100 * time.Millisecond)

	go fmt.Println("Hello concurrent world")

	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 5; i++ {
		go func() { fmt.Println(i) }()
	}

	time.Sleep(100 * time.Millisecond)

	i := 10
	go fmt.Printf("1. Значение переменной i равно %d\n", i)
	i++
	go fmt.Printf("2. Значение переменной i равно %d\n", i)
	go func() {
		i++
		go fmt.Printf("3. Значение переменной i равно %d\n", i)
	}()
	i++
	go fmt.Printf("4. Значение переменной i равно %d\n", i)
	time.Sleep(1000000)

}
