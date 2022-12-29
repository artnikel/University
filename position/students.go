package position

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Uni interface {
	Insert_db()
	Select_db()
	Delete_db()
	Search_db()
	Update_db()
}

type Student struct {
	Id      int
	Name    string
	Surname string
	Course  int
	AvgMark float64
}

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var address = goDotEnvVariable("DB_ADDRESS")

var Lists = []Student{}

func (s Student) Insert_db(w http.ResponseWriter, r *http.Request) {
	s.Name = r.FormValue("name")
	s.Surname = r.FormValue("surname")
	s.Course, _ = strconv.Atoi(r.FormValue("course"))
	s.AvgMark, _ = strconv.ParseFloat(r.FormValue("avgmark"), 64)

	db, err := sql.Open("mysql", address)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `students`(`Name`, `Surname`, `Course`, `AvgMark`) VALUES ('%s','%s','%d','%g')", s.Name, s.Surname, s.Course, s.AvgMark))
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func (s Student) Select_db() {
	db, err := sql.Open("mysql", address)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM `students`")
	if err != nil {
		panic(err.Error())
	}
	Lists = []Student{}
	for results.Next() {
		var List Student
		err = results.Scan(&List.Id, &List.Surname, &List.Name, &List.Course, &List.AvgMark)
		if err != nil {
			panic(err.Error())
		}
		Lists = append(Lists, List)
	}
	defer results.Close()

}

func (s Student) Delete_db(w http.ResponseWriter, r *http.Request) {
	s.Id, _ = strconv.Atoi(r.FormValue("id"))

	db, err := sql.Open("mysql", address)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	delete, err := db.Query(fmt.Sprintf("DELETE FROM `students` WHERE ID = '%d'", s.Id))

	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
}

func (s Student) Search_db(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")
	db, err := sql.Open("mysql", address)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query(fmt.Sprintf("SELECT * FROM `students` WHERE ID = '%[1]s' || Surname = '%[1]s' || Name = '%[1]s' || Course = '%[1]s' || AvgMark = '%[1]s'", search))
	if err != nil {
		panic(err.Error())
	}
	Lists = []Student{}
	for results.Next() {
		var List Student
		err = results.Scan(&List.Id, &List.Surname, &List.Name, &List.Course, &List.AvgMark)
		if err != nil {
			panic(err.Error())
		}
		Lists = append(Lists, List)
	}
	defer results.Close()
}

func (s Student) Update_db(w http.ResponseWriter, r *http.Request) {
	s.Id, _ = strconv.Atoi(r.FormValue("id"))
	thisOption := r.FormValue("this")
	newOption := r.FormValue("new")

	db, err := sql.Open("mysql", address)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	delete, err := db.Query(fmt.Sprintf("UPDATE `students` SET `%s`='%s' WHERE ID = '%d' ", thisOption, newOption, s.Id))

	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()

}
