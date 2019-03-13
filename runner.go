package main

import (
	// "net/http"	
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
	"./src/handlers"
	"./src/models"
)

func main() {

	db, err := gorm.Open("sqlite3", "test.db")
  if err != nil {
    panic("failed to connect database")
  }
	defer db.Close()

	db.AutoMigrate(&models.User{})
	db.Create(&models.User{Id: 200, Name: "L1212", Email: "abel@commonsense.io"})


	e := echo.New()
	e.GET("/users", handlersUser.Index)
	e.POST("/users", handlersUser.Create)

	e.Logger.Fatal(e.Start(":1323"))
}
