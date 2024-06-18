package main

import (
	"errors"
	"fmt"
	"reflect"
)

/////////////////

var mes_init int

func init() {
	fmt.Println("НАЧАЛО")
	mes_init = 1
}

func main() {

	fmt.Println(fmt.Sprintln("mes_init: ", mes_init))

	/////////////////

	// sdfsdf

	/* asd
	fghj
	sdfgasg
	kjfgh
	*/

	/////////////////

	a := 1
	fmt.Println("Test", a)
	a = 2
	fmt.Println("Test", a)

	var check float32
	check = 1.24
	fmt.Println("Test", check)
	check = 42.21
	fmt.Println("Test", check)

	const check1 int = -12
	fmt.Println("Test", check1)

	const check2 uint = 10
	fmt.Println("Test", check1)

	var message = "sad"
	fmt.Println("Test", message)

	fmt.Println(reflect.TypeOf(message))

	var b_test bool = true
	fmt.Println("Test", b_test)

	mess := []byte("abc123")
	fmt.Println("Test", mess)

	var ab byte = 62
	fmt.Println(ab)
	fmt.Printf("%c\n", ab)

	var abc rune = '<'
	fmt.Println(abc)
	fmt.Printf("%c\n", abc)

	aa, bb, cc := 1, 2, 3
	fmt.Println(aa, bb, cc)

	aa, bb = bb, aa
	fmt.Println(aa, bb, cc)

	aa, _, cc = 1, 2, 3
	fmt.Println(aa, bb, cc)

	/////////////////

	print()

	messages := "10"
	print_mes(messages)
	print_mes("123")

	/////////////////

	var returning_val_1, ret_val_1 string
	returning_val_1, ret_val_1 = test_f_of_return(10, 10, 5, 3)
	fmt.Println(returning_val_1, ret_val_1)
	returning_val, ret_val := test_f_of_return(10, 20, 5, 3)
	fmt.Println(returning_val, ret_val)
	fmt.Println(test_f_of_return(15, 11, 7, 2))

	fmt.Println(test_f_of_return_one(15, 11, 7, 2))

	/////////////////

	ages := [...]int{4, 29, 47, 15, 75, 5}
	i := 0
	for i <= len(ages)-1 {
		if mess, err := ager_check(ages[i]); err != nil {
			fmt.Println(fmt.Sprintf("Сообщение: %s; Ошибка: %s", mess, err))
		} else {
			fmt.Println(mess)
		}
		i += 1
	}

	/////////////////

	i = 0
	days := [8]string{"пн", "вт", "ср", "чт", "пт", "сб", "вс", "ывапрв"}
	for i <= len(days)-1 {
		messa := prediction(days[i])
		fmt.Println(messa)
		i += 1
	}

	/////////////////

	i = 0
	numbs := [][]int{{1, 2, 3, 5}, {-1, 2, -4, 3}, {5, 3, 2, 8}}
	for i <= len(numbs[:])-1 {
		fmt.Println(find_min(numbs[i]))
		i += 1
	}

	i = 0
	numbss := [3][4]int{{1, 2, 3, 5}, {-1, 2, -4, 3}, {5, 3, 2, 8}}
	for i <= len(numbss[:])-1 {
		fmt.Println(find_min_known(numbss[i]))
		i += 1
	}

	fmt.Println(find_min_numbs(10, 23, -12, 42, -12, -43, -64, 23, 64))

	func() {
		fmt.Println("sdf")
	}()

	/////////////////

	inc := increment()
	for ii := 0; ii < 5; ii++ {
		fmt.Println(inc())
	}

	inc = increment()
	for ii := 0; ii < 5; ii++ {
		fmt.Println(inc())
	}

	/////////////////
	fmt.Println("переменная message_uk:")
	message_uk := "asd"
	fmt.Println("переменная до функции:", message_uk, "; адрес памяти:", &message_uk)
	print_mess_pointers(message_uk)
	fmt.Println("переменная вне функции:", message_uk, "; адрес памяти:", &message_uk)

	fmt.Println("переменная message_uk_1:")
	message_uk_1 := "asd"
	fmt.Println("переменная до функции:", message_uk_1, "; адрес памяти:", &message_uk_1)
	print_mess_pointers_2(&message_uk_1)
	fmt.Println(fmt.Sprintln("переменная message_uk_1 вне функции:", message_uk_1, "; ссылка на адрес памяти с переменной (&):", &message_uk_1, "; нахождение значения, которое хранится в участке памяти, указанном по ссылке (*):", *&message_uk_1)) // переменная

	/////////////////
	var p *int
	fmt.Println(p, &p)

	ne := 20
	p = &ne
	fmt.Println(p, &p, *p, ne, &ne)

	fmt.Println(p, &p, *p, *&p, **&p) // & указывает на адресс памяти, * указывает на данные по адресу памяти. комбинация *& невелируют друг друга, так как получается "найди данные по адресу этой переменной"

	var bc *int
	nem := 10
	bc = p
	fmt.Println(p, &p, *p, bc, &bc, *bc, ne, &ne, nem, &nem)
	bc = &nem
	fmt.Println(p, &p, *p, bc, &bc, *bc, ne, &ne, nem, &nem)

	fmt.Println(ne, &ne, p, *p, &p)
	*p = 5 // если n := 20 и p = &ne, то (*p = 5) == (ne = 5)
	fmt.Println(ne, &ne, p, *p, &p)

	ne = 50
	fmt.Println(ne, &ne, p, *p, &p)

	/////////////////

}

var aaa, bbb, ccc = 5, 6, 7

func print() {
	fmt.Println(aaa, bbb, ccc)
}

func print_mes(message string) {
	fmt.Println(message)
}

func test_f_of_return(val, val_2, val_3 int, val_4 int) (string, string) {
	// answer := val + val_2 + val_3 - val_4
	// end := "!!!"
	// return fmt.Sprintf("Ответ: %d, удачи%s", answer, end)
	return fmt.Sprintf("Ответ: %d, удачи%s", val+val_2+val_3-val_4, "!!!"), "КОНЕЦ"
}

func test_f_of_return_one(val, val_2, val_3 int, val_4 int) string {
	return fmt.Sprintf("Ответ: %d, удачи%s", val+val_2+val_3-val_4, "!!!")
}

func ager_check(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "Норм", nil
	} else if age <= 10 || age >= 50 {
		return "Точно не норм", nil
	} else {
		return "Не норм", errors.New("отмена")
	}

}

func prediction(day string) string {
	switch day {
	case "пн":
		return "Ужас"
	case "вт":
		return "Страх"
	case "ср":
		return "Сойдет"
	case "чт":
		return "Предвкушение"
	case "пт":
		return "Страсть"
	case "сб":
		return "Любовь"
	case "вс":
		return "Отчаяние"
	default:
		return "Неверный день"
	}
}

func find_min(numbers []int) int {
	lenght := len(numbers)
	if lenght == 0 {
		return 0
	} else if lenght == 1 {
		return numbers[0]
	}
	i, min := 0, numbers[0]
	for i < lenght-1 {
		if numbers[i] < min {
			min = numbers[i]
		}
		i += 1
	}
	return min
}

func find_min_known(numbers [4]int) int {
	lenght := len(numbers)
	if lenght == 0 {
		return 0
	} else if lenght == 1 {
		return numbers[0]
	}
	i, min := 0, numbers[0]
	for i < lenght-1 {
		if numbers[i] < min {
			min = numbers[i]
		}
		i += 1
	}
	return min
}

func find_min_numbs(numbers ...int) int {
	lenght := len(numbers)
	if lenght == 0 {
		return 0
	} else if lenght == 1 {
		return numbers[0]
	}
	min := numbers[0]
	for _, i := range numbers {
		if i < min {
			min = i
		}
		i += 1
	}
	return min
}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func print_mess_pointers(message string) {
	message += " aaaddd"
	fmt.Println("переменная в функции:", message, "; адрес памяти:", &message)
}

func print_mess_pointers_2(message *string) {
	*message += " aaaddd"
	fmt.Println("переменная в функции:", message, "значение переменной в функции:", *message, "; адрес памяти:", &message)
}
