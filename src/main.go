package main

import (
	"Med/src/controllers"
	_ "Med/src/controllers"
	_ "Med/src/db"
	_ "Med/src/docs"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/swag/example/celler/model"
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
	r:= gin.Default()
	c := controllers.NewController()

	r.GET("/marathon", c.CheckEvent)
	r.POST("/marathon", c.AddEvents)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run()

}
