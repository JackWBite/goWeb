package main

import (
	"fmt"
	"goWeb/entity"
	"goWeb/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	c := router.Group("/")
	{
		c.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
		c.GET("/query", func(ctx *gin.Context) {
			sName := ctx.Query("name")
			var student entity.Student
			student.Name = sName

			stuInfo := service.QueryStudent(student)
			ctx.XML(http.StatusOK, stuInfo)
		})

		c.GET("/queryAll", func(ctx *gin.Context) {
			students := service.QueryAllStudent()
			ctx.XML(http.StatusOK, students)
		})

		c.POST("/add", func(ctx *gin.Context) {
			student := entity.Student{}
			if err := ctx.BindJSON(&student); err != nil {
				fmt.Println("bindJson fail")
				ctx.JSON(200, "bindJsonFail data is invalid")
				return
			}
			service.AddStudent(student)
			ctx.String(http.StatusOK, "student:%s added success", student.Name)
		})

		c.POST("/update", func(ctx *gin.Context) {
			student := entity.Student{}
			if err := ctx.BindJSON(&student); err != nil {
				fmt.Println("bindJson fail")
				ctx.JSON(200, "bindJsonFail data is invalid")
				return
			}
			name := service.UpdateStudent(student)
			ctx.String(http.StatusOK, "student:%s update success", name)
		})

		c.GET("/del", func(ctx *gin.Context) {
			sName := ctx.Query("name")
			name := service.DelStudent(sName)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
				"name":    name,
			})
		})
	}

	router.Run(":8080")
}
