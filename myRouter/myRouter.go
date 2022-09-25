package myRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/asherplotnik/golang-rest/controller"
)
	
func Init() {
	router := gin.Default()
	router.GET("/todos", controller.GetTodos)
	router.GET("/todo", controller.GetTodo)
	router.POST("/addTodo", controller.PostTodo)
	router.DELETE("/deleteTodo", controller.DeleteTodo)
	router.DELETE("/deleteAllTodos", controller.DeleteAllTodo)
	router.PATCH("/updateTodo", controller.UpdateTodo)
	router.Run("localhost:8080")
}