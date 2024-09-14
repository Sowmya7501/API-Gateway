package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	Username string `json:"username"`
	Userid string `json:"userid"`
}

var totalUsers = []User{}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.Userid = uuid.NewString()
	totalUsers = append(totalUsers,newUser)

	c.JSON(http.StatusCreated, newUser)
}

func findUserByName(c *gin.Context) {
	userName := c.Param("name")
	for _, user := range totalUsers {
		if userName == user.Username {
			c.JSON(http.StatusOK,user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error":"User not found"})
}

func getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK,totalUsers)
}

// func deleteTaskByName(c *gin.Context) {
// 	name := c.Param("name")
// 	for i, task := range totalTasks {
// 		if task.Name == name {
// 			totalTasks = append(totalTasks[:i], totalTasks[i+1:]...)
// 			c.JSON(http.StatusOK, gin.H{"message":"The task has been deleted"})
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, gin.H{"message":"The task does not exist"})
// }

// func updateTaskByName(c *gin.Context) {
// 	name := c.Param("name")

// 	var updatedTask Task

// 	if err := c.ShouldBindJSON(&updatedTask); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
// 	}
// 	for i,task := range totalTasks {
// 		if task.Name == name {
// 			totalTasks[i].Name = updatedTask.Name
// 			c.JSON(http.StatusOK, totalTasks[i])
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, gin.H{"error":"Task not found"})
// }

func main() {
	r := gin.Default()

	// Root 
	r.GET("/",func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message":"Welcome to Root"})
	})

	// create Task
	r.POST("/user",createUser)

	// Get all tasks
	r.GET("/user",getAllUsers)

	// Get task by name
	r.GET("/user/:name",findUserByName)

	// // Delete a task
	// r.DELETE("/task/:name",deleteTaskByName)

	// // Update a task
	// r.PUT("/task/:name",updateTaskByName)

	log.Println("The API started Running")
	r.Run(":8082")
}