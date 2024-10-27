package main

import (
    "context"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v4"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

type User struct {
	Username string `json:"username"`
	Userid int `json:"userid"`
}

var db *pgx.Conn

func initDB() {
    var err error
    // connStr := os.Getenv("DATABASE_URL")
	connStr := "postgres://admin:admin@postgres:5432/taskdb"
    db, err = pgx.Connect(context.Background(), connStr)
    if err != nil {
        log.Fatal(err)
    }

    createTable := `
    CREATE TABLE IF NOT EXISTS users (
		userid SERIAL PRIMARY KEY,
        username TEXT
    );
    `
    _, err = db.Exec(context.Background(), createTable)
    if err != nil {
        log.Fatal(err)
    }
}

func createUser(c *gin.Context) {
    var newUser User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var lastInsertId int
    err := db.QueryRow(context.Background(), "INSERT INTO users (username) VALUES ($1) RETURNING userid", newUser.Username).Scan(&lastInsertId)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    newUser.Userid =  lastInsertId
    c.JSON(http.StatusCreated, newUser)
}

func findUserByName(c *gin.Context) {
	userName := c.Param("name")
    row := db.QueryRow(context.Background(), "SELECT username, userid FROM users WHERE username = $1", userName)

    var user User
    err := row.Scan(&user.Username, &user.Userid)
    if err == pgx.ErrNoRows {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    } else if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

func getAllUsers(c *gin.Context) {
    rows, err := db.Query(context.Background(), "SELECT username, userid FROM users")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.Username, &user.Userid); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        users = append(users, user)
    }

    c.JSON(http.StatusOK, users)
}

func deleteUserByName(c *gin.Context) {
    userName := c.Param("name")
    _, err := db.Exec(context.Background(), "DELETE FROM users WHERE username = $1", userName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User has been deleted"})
}

func updateUserByName(c *gin.Context) {
    userName := c.Param("name")

    var updatedUser User
    if err := c.ShouldBindJSON(&updatedUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := db.Exec(context.Background(), "UPDATE users SET userid = $1 WHERE username = $2", updatedUser.Userid, userName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedUser)
}

func main() {
	initDB()
    defer db.Close(context.Background())

	r := gin.Default()

	// Root 
	r.GET("/",func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message":"Welcome to Root"})
	})

	// create User
	r.POST("/user",createUser)

	// Get all Users
	r.GET("/user",getAllUsers)

	// Get User by name
	r.GET("/user/:name",findUserByName)

	// Delete a user
	r.DELETE("/user/:name",deleteUserByName)

	// Update a user
	r.PUT("/user/:name",updateUserByName)

	// Start the server on port 8081 for the REST API
	go func() {
		if err := r.Run(":8083"); err != nil {
			panic(err)
		}
	}()

	// Start a separate server for metrics on port 8080
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}