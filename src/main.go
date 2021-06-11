package main

import (
	_ "Med/src/docs"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/mitchellh/hashstructure"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/swag/example/celler/controller"
	_ "github.com/swaggo/swag/example/celler/model"
	"net/http"
	"strconv"
)

const (
	host = "db"
	port = 5432
	user = "postgres"
	password = "12340"
	dbname = "marathon"
)

type event struct {
	Name string
	Date string
	Key string
}

func hashValue(c event) uint64 {
	hash, err := hashstructure.Hash(c, nil)
	if err != nil {
		panic(err)
	}
	return hash
}

func addEvent(name string, date string) event {
	newEvent := event{name, date, ""}
	newEvent.Key = strconv.FormatUint(hashValue(newEvent), 20)
	return newEvent
}

// @title Marathon API
// @version 1.0
// @description API for the marathon service. Provides basic methods for managing a marathon
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email bilenkomaria02@gmail.com
// @license.name MIT
// @license.url https://git.tjump.ru/mariya.bilenko/med
// @host localhost:8080
// @BasePath /api/v1
func main(){

	//------db connect -----
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname))
	if err != nil{
		panic(err)
	}

	defer conn.Close()

	err = conn.Ping()
	if err!= nil {
		panic(err)
	}

	// -------gin------------
	r:= gin.Default()

	// checkEvent godoc
	// @Summary Find event by key
	// @Produce json
	// @Param event.key string
	// @Success 200 {object}
	// @Router /marathon [get]
	r.GET("/marathon", func(c *gin.Context) {

		key := c.Query("key")

		f := conn.QueryRow("SELECT * FROM events WHERE key = $1",key)
		if err != nil {
			panic(err)
		}


		ev := new(event)
		err := f.Scan(&ev.Name, &ev.Date, &ev.Key)
		if err != sql.ErrNoRows{
			c.JSON(http.StatusOK,ev)
		}


		if err != nil{
			msg := "marathon " + key + " not found"
			c.JSON(404, gin.H{
				"message": msg,
			})
		}
	})


	// addEvent godoc
	// @Summary Add new event
	// @Produce json
	// @Success 200 {object}
	// @Router /marathon [post]
	r.POST("/marathon", func(c *gin.Context) {

		name := c.PostForm("name")
		date := c.PostForm("date")
		newEvent := addEvent(name, date)
		_, err = conn.Exec("insert into events (name, date) values ( $1, $2, $3)", newEvent.Name, newEvent.Date, newEvent.Key)
		c.JSON(200,gin.H{"key": newEvent.Key})

	})

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run()

}
