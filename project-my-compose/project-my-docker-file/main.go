// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"os"
// 	"time"

// 	_ "github.com/lib/pq"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()

// 	dbUser := os.Getenv("DB_USER")
// 	dbPass := os.Getenv("DB_PASS")
// 	dbName := os.Getenv("DB_NAME")
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")

// 	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
// 		dbUser, dbPass, dbHost, dbPort, dbName)

// 	fmt.Println("Connect to: ", connStr)
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		fmt.Println("Error connecting to database:", err)
// 	}

// 	router.GET("/", func(c *gin.Context) {
// 		c.String(200, "Ouch!")
// 		fmt.Println("To Infinity and Beyond!")
// 	})

// 	router.GET("/ping", func(c *gin.Context) {
// 		_, err = db.Exec("INSERT INTO access_log (timestamp) VALUES ($1)", time.Now())
// 		if err := db.Ping(); err != nil {
// 			fmt.Println("Error pinging database:", err)
// 			c.String(500, "Database connection failed")
// 		} else {
// 			fmt.Println("Ping to database successful")
// 			c.String(200, "Ping to Database connection successful")
// 		}
// 	})

// 	// Start the server on port 78
// 	err = router.Run(":77")
// 	if err != nil {
// 		fmt.Println("Error starting server:", err)
// 	}

//		if err := db.Ping(); err != nil {
//			fmt.Printf("Error connecting to database: %v\n", err)
//		} else {
//			fmt.Println("Successfully connected to database!")
//			fmt.Println("Database: ", db)
//		}
//	}
package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func main() {
	// Get env
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Connect to Database"
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println("Connect to: ", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Connection to database failed: ", err)
	}

	// web services
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		fmt.Println("Ouch!")
		return c.HTML(http.StatusOK, "Aku abis operasi loh :) \n")
	})

	e.GET("/ping", func(c echo.Context) error {
		_, err = db.Exec("INSERT INTO access_log (timestamp) VALUES ($1)", time.Now())
		if err := db.Ping(); err != nil {
			fmt.Println("Ping Failed")
			return c.HTML(http.StatusBadRequest, "Ouch!")
		}
		fmt.Println("Ping Successfull")
		return c.HTML(http.StatusOK, "PONG! \n")
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "77"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
