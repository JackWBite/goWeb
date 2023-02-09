package db

import (
	"fmt"

	"goWeb/entity"

	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	userName  string = "root"
	password  string = "123456"
	ipAddrees string = "127.0.0.1"
	port      int    = 3306
	dbName    string = "testdb"
	charset   string = "utf8"
)

func DbConnection() *sqlx.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	Db, err := sqlx.Open("mysql", dns)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}

	return Db
}

func QueryStudent(student entity.Student) entity.Student {
	dbConn := DbConnection()
	defer dbConn.Close()

	rows := dbConn.QueryRow("select sid, name, age, sex, school,class from student where name=?", student.Name)
	var result entity.Student
	rows.Scan(&result.Sid, &result.Name, &result.Age, &result.Sex, &result.School, &result.Class)
	res, _ := json.Marshal(student)
	fmt.Println(string(res))

	return result
}

func QueryAllStudent() ([]entity.Student, error) {
	dbConn := DbConnection()
	defer dbConn.Close()

	rows, err := dbConn.Query("select sid, name, age, sex, school,class from student")
	if err != nil {
		fmt.Println(err.Error())
		panic("db connection error")
	}
	var students []entity.Student
	fmt.Println("Sid, Name, Age, Sex, School,Class")
	for rows.Next() {
		var student entity.Student
		rows.Scan(&student.Sid, &student.Name, &student.Age, &student.Sex, &student.School, &student.Class)

		students = append(students, student)

		res, _ := json.Marshal(student)
		fmt.Println(string(res))
	}

	defer rows.Close()

	return students, err
}

func InsertNewStudent(student entity.Student) (int64, error) {
	dbConn := DbConnection()
	defer dbConn.Close()

	stm, err := dbConn.Prepare("INSERT INTO student set name=?,age=?,class=?,school=?,sex=?")
	if err != nil {
		fmt.Println(err.Error())
		panic("db connection error")
	}

	result, err := stm.Exec(student.Name, student.Age, student.Class, student.School, student.Sex)
	if err != nil {
		fmt.Printf("insert student error, %s", err.Error())
		return -1, err
	}
	return result.LastInsertId()
}

func DelStudent(sName string) (int64, error) {
	dbConn := DbConnection()
	defer dbConn.Close()

	stm, err := dbConn.Prepare("delete from student where name =?")
	if err != nil {
		fmt.Println(err.Error())
		panic("db connection error")
	}
	result, execErr := stm.Exec(sName)
	if execErr != nil {
		fmt.Printf("del student error, %s", err.Error())
		return -1, err
	}
	return result.RowsAffected()
}
