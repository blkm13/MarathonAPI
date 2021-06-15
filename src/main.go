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

type Controller struct {
}

// NewController example
func NewController() *Controller {
	return &Controller{}
}

// Message example


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

func connect() *sql.DB {
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname))
	if err != nil{
		panic(err)
	}
	return conn
}


// checkEvent godoc
// @Summary Find event by key
// @Description search for an event by key in the database
// @ID get-string-by-int
// @Accept json
// @Produce json
// @Success 200 {object} event
// @Router /marathon [get]
func (c *Controller) checkEvent (ctx *gin.Context){
	key := ctx.Query("key")

	f := connect().QueryRow("SELECT * FROM events WHERE key = $1",key)

	ev := new(event)
	err := f.Scan(&ev.Name, &ev.Date, &ev.Key)
	if err != sql.ErrNoRows{
		ctx.JSON(http.StatusOK,ev)
	}


	if err != nil{
		msg := "marathon " + key + " not found"
		ctx.JSON(404, gin.H{
			"message": msg,
		})
	}
}

// addEvent godoc
// @Summary Add new event
// @Description add new event
// @Accept json
// @Produce json
// @Success 200 {string} event.Key
// @Router /marathon [post]
func (c *Controller) addEvents( ctx *gin.Context)  {
	name := ctx.PostForm("name")
	date := ctx.PostForm("date")
	newEvent := addEvent(name, date)
	connect().Exec("insert into events (name, date) values ( $1, $2, $3)", newEvent.Name, newEvent.Date, newEvent.Key)
	ctx.JSON(200,gin.H{"key": newEvent.Key})
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
	c := NewController()


	r.GET("/marathon", c.checkEvent )
	r.POST("/marathon", c.addEvents)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run()

}
