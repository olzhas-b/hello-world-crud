package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olzhas-b/hello_world_crud/controllers"
	"github.com/olzhas-b/hello_world_crud/database"
	"gorm.io/gorm"
)

func handleRequest(r *gin.Engine) {
	r.GET("/user", controllers.GetUser)
	r.GET("/user/:id", controllers.GetUserByID)
	r.POST("/user", controllers.PostUser)
	r.PUT("/user", controllers.PutUser)
	r.DELETE("/user", controllers.DeleteUser)
}
var dataBase *gorm.DB
func main() {
	r := gin.Default()
	dataBase = database.InitDataBase()
	handleRequest(r)
	r.Run(":8000")
}
