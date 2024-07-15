package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/lib/pq"
)

type ConnectionSettings struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type DataBaseContext struct {
	Class             string
	Name              string
	Country           string
	Organization      string
	Displacement      string
	Length            string
	Width             string
	Draft             string
	Propulsion        string
	Engines           string
	Power             string
	Speed             string
	Ranges            string
	Autonomy          string
	Crew              string
	Artillery         string
	AntiAircraft      string
	TacticalStrike    string
	Navigation        string
	Radar             string
	ElectronicWarfare string
	Missile           string
	AntiSubmarine     string
	MineTorpedo       string
	AviationGroup     string
	ModelPath         string
}

func main() {

	//"localhost", 5432, "postgres", "123456", "test_db"

	//localhost 5432 postgres 123456 test_db

	connection := ConnectionSettings{}

	fmt.Println("Начинаем, введите данные для подключения к бд с клавиатуры в формате: '<host> <port> <user> <password> <dbname>'")
	fmt.Scanf("%s", &connection.Host)
	fmt.Scanf("%d", &connection.Port)
	fmt.Scanf("%s", &connection.User)
	fmt.Scanf("%s", &connection.Password)
	fmt.Scanf("%s", &connection.DBName)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", connection.Host, connection.Port, connection.User, connection.Password, connection.DBName)
	// psqlconn := fmt.Sprintf("host=localhost port=5432 user=postgres password=123456 dbname=test_db sslmode=disable")

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Подключение прошло успешно!")
	fmt.Println("")

	var choice string
	for true {
		fmt.Println("Введите желаемые операции с клавиатуры: create/read/update/delete (для выхода введите q)")
		fmt.Scanf("%s", &choice)
		if choice == "create" {
			db_create(db)
		} else if choice == "read" {
			db_read(db)
		} else if choice == "update" {
			db_update(db)
		} else if choice == "delete" {
			db_delete(db)
		} else if choice == "q" {
			break
		} else {
			fmt.Println("Неверное введена команда\n")
		}
		choice = ""
	}

}

func db_create(db *sql.DB) {
	_, err := db.Exec(`
    INSERT INTO ships (
        class, name, country, organization, displacement, length, width, draft,
        propulsion, engines, power, speed, range, autonomy, crew, artillery,
        anti_aircraft, tactical_strike, navigation, radar, electronic_warfare,
        missile, anti_submarine, mine_torpedo, aviation_group, model_path
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16,
               $17, $18, $19, $20, $21, $22, $23, $24, $25, $26)
	`, "test", "test", "test", "test", "test", "test", "test", "test", "test", "test",
		"test", "test", "test", "test", "test", "test", "test", "test", "test", "test",
		"test", "test", "test", "test", "test", "test")
	CheckError(err)
}

func db_read(db *sql.DB) {
	check_info := "SELECT * FROM ships;"
	rows, e := db.Query(check_info)
	CheckError(e)
	CheckError(rows.Err())

	var db_content []DataBaseContext

	for rows.Next() {
		row_content := new(DataBaseContext)
		s := reflect.ValueOf(row_content).Elem()
		/*
			в Go нельзя получить доступ к неэкспортируемым полям через рефлексию, поскольку это нарушает концепцию инкапсуляции языка.
			поэтому атрибуты с маленькой - неэкспортируемые, а с большой - экспортируемые
		*/
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		e = rows.Scan(columns...)
		CheckError(e)

		db_content = append(db_content, *row_content)
	}

	// err = rows.Scan(&name, &roll)    -   то есть нужно перечислять для какой переменной если не использовать классы

	fmt.Println(db_content)
}

func db_update(db *sql.DB) {
	_, err := db.Exec(`
    UPDATE "ships" SET "class" = $1, "length" = $2 WHERE "name" = $3`, "test333", "test222", "test")
	CheckError(err)
}

func db_delete(db *sql.DB) {
	_, err := db.Exec(`
	DELETE FROM ships where name = $1`, "test")
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
