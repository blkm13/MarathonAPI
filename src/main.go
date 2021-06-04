package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/mitchellh/hashstructure"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strconv"
)

const(
	host ="127.0.0.1"
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

// @title Marathon API
// @version 1.0
// @description API for the marathon service. Provides basic methods for managing a marathon.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email bilenkomaria02@gmail.com

// @license.name MIT
// @license.url https://git.tjump.ru/mariya.bilenko/med

// @host localhost:8080
// @BasePath /api/v1



func main(){

	var events []event

	firstEvent := event{ "firstevent", "20.05.2021", "key" }
	events= append(events, firstEvent)
	secondEvent := event{ "secondevent", "21.05.2021", "key"}
	events= append(events,secondEvent)
	thirdEvent := event{" thirdevent", "22.05.2021", "key"}
	events= append(events, thirdEvent)

	for _, v := range events{
		v.Key = strconv.FormatUint(hashValue(v), 20)
		fmt.Println(hashValue(v))
	}

	//------db connect ------
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil{
		panic(err)
	}

	defer conn.Close()

	err = conn.Ping()
	if err!= nil {
		panic(err)
	}

	_, err = conn.Exec("insert into events (name, date) values ( 'seventhEvent', '21.05.2021')")

	if err != nil {
		panic(err)
	}


	//fmt.Println(firstEvent.Key)

	// -------gin------------

	// -------gin------------


	r:= gin.Default()

	//r.LoadHTMLGlob("templates/form.html")

	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	// checkEvent godoc
	// @Summary Find event by key
	// @Produce json
	// @Param event.key string
	// @Success 200 {object}
	// @Router /marathon [get]
	r.GET("/marathon", func(c *gin.Context) {
		key := c.Query("key")
		flag := true
		for _, v := range events {
			if key == v.Key {
				c.JSON(http.StatusOK, v)
				flag = false
				break
			}
		}
		if flag{
			msg := "marathon "+key+" not found"
			c.JSON(404, gin.H{
				"message": msg,
			})
		}

	})


	// addEvent godoc
	// @Summary Add new event
	// @Produce json
	// @Success 200 {object}
	// @Router /marathon [get]
	r.POST("/marathon", func(c *gin.Context){

		name := c.PostForm("name" )
		date := c.PostForm("date")

		newEvent := event{name,date,""}
		newEvent.Key = strconv.FormatUint(hashValue(newEvent), 20)
		//_, err = conn.Exec("insert into events (name, date) values ( $1, $2)", newEvent.Name, newEvent.Date)
		_, err = conn.Exec("insert into events (name, date) values ( $1, $2)", newEvent.Name, newEvent.Date)
		fmt.Printf("name: %s; date: %s; key: %s",newEvent.Name,newEvent.Date, newEvent.Key)
		c.JSON(200,gin.H{"key": newEvent.Key})

	})

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run()

}
