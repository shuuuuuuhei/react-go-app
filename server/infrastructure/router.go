package infrastructure

import (
	"github.com/react-go-app/backend/interfaces/controller"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	userController := controller.CreateTodoController(NewSqlHandler())
	router.GET("/todoes", func(c *gin.Context) { userController.Index(c) })
}