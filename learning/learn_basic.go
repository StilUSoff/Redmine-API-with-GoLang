package main

import (
	"errors"
	"fmt"
	"reflect"

	/// СВОЯ БИБЛИОТЕКА, КОТОРАЯ ПОДКЛЮЧАЕТСЯ БЛАГОДАРЯ go.mod (go mod init learning)
	"learning/package_learn"
)

/////////////////

var mes_init int

func init() { // init выполняется перед main
	fmt.Println("НАЧАЛО")
	mes_init = 1
}

func handler_panic() { // функция для примера работы ddefer
	if r := recover(); r != nil { // recover помогает справиться с паникой, чтобы не вылетать их кода
		fmt.Println("произошла паника:", r)
	} else {
		fmt.Println("теперь точно КОНЕЦ программы")
	}

}

func main() {

	///////////////// пример использование defer и panic

	//defer откладывает срабатывание строчки на самый конец, перед выходом из тела функции в случае ошибки или успеха, не важно
	//panic() просто вывод текста в терминал "в виде" ошибки и заканчивание программы
	defer handler_panic()

	// ЕСЛИ РАССКОМЕНТИРОВАТЬ СТРОКИ НИЖЕ БУДЕТ ПРИМЕР ВЫЗОВА ДЕФЕРА И ПАНИКИ
	// messageasdf := []string{
	// 	"mes1",
	// 	"mess2",
	// }
	// messageasdf[2] = "mess3"

	///////////////// использование глобальных переменных из инита

	fmt.Println(fmt.Sprintln("mes_init: ", mes_init))

	////////// использование своего пакета
	package_learn.Math_Print(package_learn.Math_Factorial(4))

	///////////////// способы комментирования

	// sdfsdf

	/* asd
	fghj
	sdfgasg
	kjfgh
	*/

	///////////////// про типы данных и виды переменных, а также способы их вывода в терминал; про пропуск переменной с помощью _

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

	aa, _, cc = 7, 8, 9
	fmt.Println(aa, bb, cc)

	///////////////// про использование глобальных переменных; про возможность передать в функцию не только переменную, но и сразу данные

	print()

	messages := "10"
	print_mes(messages)
	print_mes("123")

	///////////////// про return нескольких переменных

	var returning_val_1, ret_val_1 string
	returning_val_1, ret_val_1 = test_f_of_return(10, 10, 5, 3)
	fmt.Println(returning_val_1, ret_val_1)
	returning_val, ret_val := test_f_of_return(10, 20, 5, 3)
	fmt.Println(returning_val, ret_val)
	fmt.Println(test_f_of_return(15, 11, 7, 2))

	ret_val_one := test_f_of_return_one(17, 11, 7, 2)
	fmt.Println(ret_val_one)
	fmt.Println(test_f_of_return_one(18, 14, 7, 5))

	///////////////// && = и ; || = или ; ! = не ;

	///////////////// про использование ... для указывания длины масива, равную количеству элементов внесенных при объявлении переменных; про использование Sprintf и иной от Println способ вывода в консоль

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

	///////////////// про switch-case и про передачу элемента массива в функцию

	i = 0
	days := [8]string{"пн", "вт", "ср", "чт", "пт", "сб", "вс", "ывапрв"}
	for i <= len(days)-1 {
		messa := prediction(days[i])
		fmt.Println(messa)
		i += 1
	}

	///////////////// про функции с известным размером входного массива и неизвестным; про функции с неограниченным количеством входных переменных; про анонимные функции; про безымянные функции внутри основной функции, которые сразу же исполняются

	i = 0
	// ниже - слайс
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

	f := func(x, y int) int { return x + y }
	fmt.Println(f(3, 4))
	fmt.Println(f(6, 7))

	var ff func(int, int) int = func(x, y int) int { return x + y } //тоже самое что и определние выше
	fmt.Println(ff(3, 4))

	var fff func(int, int) int
	fff = func(x, y int) int { return x + y } //тоже самое что и определние выше
	fmt.Println(fff(3, 4))

	fmt.Println(find_min_numbs(10, 23, -12, 42, -12, -43, -64, 23, 64))

	mas_of_int := []int{10, 23, -12, 42, -12, -43, -64, 23, 64}
	fmt.Println(find_min_numbs(mas_of_int...)) // с помощью ... можно сразу передать срез без создания нового, если массив неопределенной длины

	func() {
		fmt.Println("sdf")
	}()

	///////////////// про возвращение функции от функции

	inc := increment()
	for ii := 0; ii < 5; ii++ {
		fmt.Println(inc())
	}

	inc = increment()
	for ii := 0; ii < 5; ii++ {
		fmt.Println(inc())
	}

	///////////////// про ссылки и указатели в функциях
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

	///////////////// про ссылки и указатели относительно памяти и как они друг с другом работают
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

	///////////////// про len и cap, а также про изменине массива в функции

	messagege := [3]string{"1", "2", "3"}
	fmt.Println(messagege)
	print_messeg(&messagege)
	fmt.Println(messagege)

	messagegeeg := make([]string, 5)
	fmt.Println(messagegeeg)
	fmt.Println(len(messagegeeg))
	fmt.Println(cap(messagegeeg)) // cap - это емкость слайса (slice, или же срез массива), максимальное количество элементов, которое может вместить слайс, прежде, чем потребуется изменить его размер
	messagegeeg = append(messagegeeg, "6")
	fmt.Println(messagegeeg)
	fmt.Println(len(messagegeeg))
	fmt.Println(cap(messagegeeg))

	///////////////// базовый математический цикл для матрицы и то, как матрицу создавать

	matrix := make([][]int, 6)

	for i := 0; i <= 5; i++ {
		matrix[i] = make([]int, 6)
		for j := 0; j <= 5; j++ {
			matrix[i][j] = i + j
		}
		matrix[i] = append(matrix[i], 12)
		fmt.Println(matrix[i])
	}

	///////////////// дополнение про лупы (циклы)

	for x := 0; true; x++ {
		fmt.Println(x)
		if x >= 12 {
			break
		}
	}

	x := 0
	for true {
		fmt.Println(x)
		x++
		if x >= 12 {
			break
		}
	}

	stringing := []string{"asd", "asf", "gasdf"}
	for ij := range stringing {
		fmt.Println(stringing[ij])
	}

	for index, value := range stringing {
		fmt.Println(index, value)
	}

	for _, value := range stringing {
		fmt.Println(value)
	}

	///////////////// мапы

	users := map[string]int{
		"Vasya":  15,
		"Petya":  23,
		"Kostya": 48,
	}
	delete(users, "Vasya")
	for key, value := range users {
		fmt.Println(key, value)
	}

	age, exist := users["Kostya"]
	if exist {
		fmt.Println("EXIST: Kostya", age)
	}

	users_1 := make(map[string]int)
	users_1["Petya"] = 17
	users_1["Pasha"] = 21
	for key, value := range users_1 {
		fmt.Println(key, value)
	}

	///////////////// классы

	user_1 := User{"Vasya", 23, "Male", 75, 185}
	user_2 := NewUser("Olga", 34, "Female", 63, 168)
	fmt.Printf("%+v\n", NewUser("Sasha", 42, "Male", 94, 195))

	user_1.print_info()
	user_2.print_info()
	user_2.set_name("Oksana")
	user_2.print_info()
	fmt.Println("User_1 is adult?", user_1.age.isAdult())

	///////////////// интерфейсы

	var r Shape = Rectangle{Width: 10, Height: 5}
	fmt.Println("r.Area()", r.Area())
	fmt.Println("r.Perimeter()", r.Perimeter())

	var s Shape = Sphere{Radius: 5}
	fmt.Println("s.Area()", s.Area())
	fmt.Println("s.Perimeter()", s.Perimeter())

	///////////////// func inside func
	func() {
		fmt.Println("asdfasdf")
	}()

}

type Age int

func (a Age) isAdult() bool {
	return a <= 18
}

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

func (u User) print_info() {
	fmt.Println(fmt.Sprintf("name: %s, age: %#v, sex: %s, weight: %#v, height: %#v", u.name, u.age, u.sex, u.weight, u.height))
}

func (u *User) set_name(new_name string) {
	u.name = new_name
	fmt.Println(fmt.Sprintf("New name: %s", u.name))
}

func NewUser(name string, age int, sex string, weight int, height int) User {
	return User{
		name:   name,
		age:    Age(age),
		sex:    sex,
		weight: weight,
		height: height,
	} // ИЛИ return User{name, age, sex, weight, height}
}

/* ИНТЕРФЕЙСЫ НАПРИМЕР НУЖНЫ В ТАКИХ ВОТ СЛУЧАЯХ, КОГДА НЕ ХОЧЕТСЯ
НАЗЫВАТЬ ПО-РАЗНОМУ ФУНКЦИИ КОТОРЫЕ ЗНАЧАТ ДЛЯ РАЗНЫХ ТИПОВ ОБЪЕКТОВ
ОДНО И ТО ЖЕ (пусть могут и работать по-разному, но главное что return и данные на вход
одинаковые из-за того, что написано в определении интерфейса)
*/

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Sphere struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (s Sphere) Area() float64 {
	return 2 * 3.14 * s.Radius
}

func (s Sphere) Perimeter() float64 {
	return 3.14 * s.Radius * s.Radius
}

var aaa, bbb, ccc = 123, 34, 2345

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

// /// func with 2 return
func ager_check(age int) (string, error) { // пример возврата переменной и nil/ошибки (nil ТОЖЕ САМОЕ ЧТО None)
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

func print_messeg(message *[3]string) error {
	lenght := len(message)
	if lenght == 0 {
		return errors.New("empty array")
	}
	message[1] = "5"
	fmt.Println(*message)
	return nil
}
