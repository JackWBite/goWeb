package service

import (
	"goWeb/db"
	"goWeb/entity"
)

func AddStudent(student entity.Student) (m string) {
	db.InsertNewStudent(student)
	return "add success"
}

func QueryStudent(student entity.Student) entity.Student {
	return db.QueryStudent(student)
}

func UpdateStudent(student entity.Student) (m string) {
	db.InsertNewStudent(student)
	return student.Name
}

func DelStudent(sName string) (m string) {
	db.DelStudent(sName)
	return sName
}

func QueryAllStudent() []entity.Student {
	students, err := db.QueryAllStudent()
	if err == nil {
		return students
	}
	return nil
}
