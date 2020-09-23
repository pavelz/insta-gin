package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("insta-go")

	users := sq.Select("id").From("users")

	sql_q, _, _ := users.ToSql()
	fmt.Println(sql_q)

	conn, err := sql.Open("postgres", "postgres:///insta_development?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}

	var id int
	err = conn.QueryRow(sql_q).Scan(&id)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("ID: %d\n", id)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
