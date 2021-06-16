package main

import (
	"Med/src/controllers"
	_ "Med/src/controllers"
	"Med/src/db"
	_ "Med/src/db"
	_ "Med/src/docs"
	"github.com/docopt/docopt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/swag/example/celler/model"
	"strconv"
)


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

	usage := `
Usage:
	./main.go <value>
	
Options:
	server   
	populate  
	`

	arguments, _ := docopt.ParseDoc(usage)
	//fmt.Println(arguments)

	switch arguments["<value>"] {
		case "server": {
			r:= gin.Default()
			c := controllers.NewController()

			r.GET("/marathon", c.CheckEvent)
			r.POST("/marathon", c.AddEvents)

			url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
			r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

			r.Run()
		}
		case "populate": {
			var events []controllers.Event
			firstEvent := controllers.Event{ "firstevent", "20.05.2021", "" }
			events= append(events, firstEvent)
			secondEvent := controllers.Event{ "secondevent", "21.05.2021", ""}
			events= append(events,secondEvent)
			thirdEvent := controllers.Event{" thirdevent", "22.05.2021", ""}
			events= append(events, thirdEvent)

			for _, v := range events{
				v.Key = strconv.FormatUint(controllers.HashValue(v), 20)
				db.Connect().Exec("insert into events (name, date, key) values ($1, $2, $3)",v.Name, v.Date, v.Key)
			}
			//fmt.Println("command to fill db")
		}
	}

}
