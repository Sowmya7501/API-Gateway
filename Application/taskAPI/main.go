package main

import (
    "context"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v4"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

type Task struct {
	ID   int    `json:"id"`
	Assignee string `json:"assignee"`
	Assignor string `json:"assignor"`
    Name string `json:"name"`
}

var db *pgx.Conn

func initDB() {
    var err error
    // connStr := os.Getenv("DATABASE_URL")
	connStr := "postgresql://admin:admin@postgres:5432/taskdb"
    db, err = pgx.Connect(context.Background(), connStr)
    if err != nil {
        log.Fatal(err)
    }

    createTable := `
    CREATE TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        assignee TEXT,
        assignor TEXT,
        name TEXT
    );
    `
    _, err = db.Exec(context.Background(), createTable)
    if err != nil {
        log.Fatal(err)
    }
}

func createTask(c *gin.Context) {
    var newTask Task
    if err := c.ShouldBindJSON(&newTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var lastInsertId int
    err := db.QueryRow(context.Background(), "INSERT INTO tasks (assignee, assignor, name) VALUES ($1, $2, $3) RETURNING id", newTask.Assignee, newTask.Assignor, newTask.Name).Scan(&lastInsertId)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    newTask.ID = lastInsertId
    c.JSON(http.StatusCreated, newTask)
}

func getTaskByName(c *gin.Context) {
	taskName := c.Param("name")
    row := db.QueryRow(context.Background(), "SELECT id, assignee, assignor, name FROM tasks WHERE name = $1", taskName)

    var task Task
    err := row.Scan(&task.ID, &task.Assignee, &task.Assignor, &task.Name)
    if err == pgx.ErrNoRows {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    } else if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, task)
}

func getAllTasks(c *gin.Context) {
    rows, err := db.Query(context.Background(), "SELECT id, assignee, assignor, name FROM tasks")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var tasks []Task
    for rows.Next() {
        var task Task
        if err := rows.Scan(&task.ID, &task.Assignee, &task.Assignor, &task.Name); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        tasks = append(tasks, task)
    }

    c.JSON(http.StatusOK, tasks)
}

func deleteTaskByName(c *gin.Context) {
    name := c.Param("name")
    _, err := db.Exec(context.Background(), "DELETE FROM tasks WHERE name = $1", name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "The task has been deleted"})
}

func updateTaskByName(c *gin.Context) {
    name := c.Param("name")

    var updatedTask Task
    if err := c.ShouldBindJSON(&updatedTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := db.Exec(context.Background(), "UPDATE tasks SET name = $1, assignee = $2, assignor = $3 WHERE name = $4", updatedTask.Name, updatedTask.Assignee, updatedTask.Assignor, name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedTask)
}

func main() {

	initDB()
    defer db.Close(context.Background())

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

	// Start the server on port 8081 for the REST API
	go func() {
		if err := r.Run(":8082"); err != nil {
			panic(err)
		}
	}()

	// Start a separate server for metrics on port 8080
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}