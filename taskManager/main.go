package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Task struct {
	Assignee string `json:"assignee"`
	Assignor string `json:"assignor"`
	ID   string    `json:"id"`
    Name string `json:"name"`
}

var totalTasks = []Task{}

func createTask(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask.ID = uuid.NewString()
	totalTasks = append(totalTasks,newTask)

	c.JSON(http.StatusCreated, newTask)
}

func getTaskByName(c *gin.Context) {
	taskName := c.Param("name")
	for _, task := range totalTasks {
		if taskName == task.Name {
			c.JSON(http.StatusOK,task)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error":"Task not found"})
}

func getAllTasks(c *gin.Context) {
	c.JSON(http.StatusOK,totalTasks)
}

func deleteTaskByName(c *gin.Context) {
	name := c.Param("name")
	for i, task := range totalTasks {
		if task.Name == name {
			totalTasks = append(totalTasks[:i], totalTasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message":"The task has been deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message":"The task does not exist"})
}

func updateTaskByName(c *gin.Context) {
	name := c.Param("name")

	var updatedTask Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	}
	for i,task := range totalTasks {
		if task.Name == name {
			totalTasks[i].Name = updatedTask.Name
			c.JSON(http.StatusOK, totalTasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error":"Task not found"})
}

func main() {
	r := gin.Default()

	// Root 
	r.GET("/",func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message":"Welcome to Root"})
	})

	// create Task
	r.POST("/task",createTask)

	// Get all tasks
	r.GET("/task",getAllTasks)

	// Get task by name
	r.GET("/task/:name",getTaskByName)

	// Delete a task
	r.DELETE("/task/:name",deleteTaskByName)

	// Update a task
	r.PUT("/task/:name",updateTaskByName)

	log.Println("The API started Running")
	r.Run(":8081")
}