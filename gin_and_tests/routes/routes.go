package routes

import (
	"github.com/gin-gonic/gin"
	"gin_and_tests/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ListStudents)
	r.GET("/:name", controllers.Greet)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.GetStudentById)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/students/registration/:registration", controllers.SearchStudentByFederalRegistration)
	r.Run()
}
