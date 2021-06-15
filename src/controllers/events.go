package controllers

import (
	"Med/src/db"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckEvent godoc
// @Summary Find event by key
// @Description search for an event by key in the database
// @ID get-string-by-int
// @Accept json
// @Produce json
// @Success 200 {object} event
// @Router /marathon [get]

func (c *Controller) CheckEvent (ctx *gin.Context){
	key := ctx.Query("key")
	f := db.Connect().QueryRow("SELECT * FROM events WHERE key = $1",key)
	ev := new(Event)
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

// AddEvents godoc
// @Summary Add new event
// @Description add new event
// @Accept json
// @Produce json
// @Success 200 {string} event.Key
// @Router /marathon [post]

func (c *Controller) AddEvents( ctx *gin.Context)  {
	name := ctx.PostForm("name")
	date := ctx.PostForm("date")
	newEvent := AddEvent(name, date)
	db.Connect().Exec("insert into events (name, date) values ( $1, $2, $3)", newEvent.Name, newEvent.Date, newEvent.Key)
	ctx.JSON(200,gin.H{"key": newEvent.Key})
}

