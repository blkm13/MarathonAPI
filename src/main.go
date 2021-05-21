package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/mitchellh/hashstructure"

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

type events struct {
	name string
	date string
}

func hashValue(c events) uint64 {
	hash, err := hashstructure.Hash(c, nil)
	if err != nil {
		panic(err)
	}
	return hash
}


func main(){

	var event []events

	firstEvent := events{ "firstevent", "20.05.2021" }
	event= append(event, firstEvent)
	secondEvent := events{ "secondevent", "21.05.2021"}
	event= append(event,secondEvent)
	thirdEvent := events{" rthirdevent", "22.05.2021"}
	event= append(event, thirdEvent)

	var hashValues []uint64

	for _, v := range event{
		hashValues = append(hashValues, hashValue(v))
	}

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

	_, err = conn.Exec("insert into events (name, date) values ( 'fourthEvent', '21.05.2021')")

	if err != nil {
		panic(err)
	}

	/*fmt.Println(hashValues[0])

	rout := "/"+ strconv.Itoa(int(hashValues[0]))+"/"

	r:= gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET(rout, func(context *gin.Context) {
		context.HTML(http.StatusOK, "form.html", "title: Form page")
	})

	r.Run()*/

}
