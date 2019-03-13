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
	db.Create(&models.User{Id: 3000, Name: "L1212jj", Email: "abejl@commonsense.io"})


	e := echo.New()

	e.GET("/users",  func(c echo.Context) error {
		return handlersUser.Index(c, db)
	})
	e.GET("/users/:id",  func(c echo.Context) error {
		return handlersUser.Show(c, db)
	})

	e.POST("/users", handlersUser.Create)

	e.Logger.Fatal(e.Start(":1323"))
}
