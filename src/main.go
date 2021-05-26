package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/mitchellh/hashstructure"
	"net/http"
	"strconv"

	//"github.com/gin-gonic/gin"
	//"net/http"
	//"strconv"
)

const(
	host ="db"
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


	// -------gin------------


	r:= gin.Default()

	r.LoadHTMLGlob("templates/form.html")

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


	r.POST("/marathon", func(c *gin.Context){

		name := c.PostForm("name")
		date := c.PostForm("date")

		newEvent := event{name,date,""}
		newEvent.Key = strconv.FormatUint(hashValue(newEvent), 20)
		_, err := conn.Exec("insert into events (name, date) values ( $1, $2)", newEvent.Name, newEvent.Date)
		fmt.Printf("name: %s; date: %s; key: %s",newEvent.Name,newEvent.Date, newEvent.Key)
	})

	r.Run()

}
