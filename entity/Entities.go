package entity

type Human struct {
	Name string `json:"name" form:"name"` //姓名
	Sex  int    `json:"sex" form:"sex"`   //、性别
	Age  int    `json:"age" form:"age"`   //年龄
}

type Student struct {
	Human
	Sid    int    `json:"sid" form:"sid"`       // 学生ID
	Class  string `json:"class" form:"class"`   //班级
	School string `json:"school" form:"school"` //学校
}

type Teacher struct {
	Human
	Tid    int    `json:"tid" form:"tid"`       //老师ID
	Course string `json:"course" form:"course"` //学科
	Office string `json:"office" form:"office"` //科室
}
