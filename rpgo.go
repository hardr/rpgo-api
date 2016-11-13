package main

import (
	"fmt"
	// "net/http"
	// "html"
	// "html/template"
	"log"
	// "encoding/json"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
	// "strconv"
)

type Char struct {
	Id    int
	Name  string
  Class string
  Url string
  Hp int
	Def int
}

var dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("postgres", "postgres://RyH:password@localhost/rpgo_db?sslmode=disable")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("/chars", GetChars)
		// v1.GET("/chars/:id", GetChar)
		// v1.POST("/users", PostUser)
		// v1.PUT("/chars/:id", UpdateChar)
		// v1.DELETE("/chars/:id", DeleteUser)
	}

	r.Run(":8080")
}

func GetChars(c *gin.Context) {
	var chars []Char
	_, err := dbmap.Select(&chars, "SELECT * FROM characters")

	if err == nil {
		fmt.Println(err)
		c.JSON(200, chars)
	} else {
		c.JSON(404, gin.H{"error": "characters can't be accessed"})
	}
	// curl -i http://localhost:8080/api/v1/chars
}
