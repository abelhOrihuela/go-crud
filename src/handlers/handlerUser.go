package handlersUser


import (
)

import (
	"net/http"	
	// "github.com/davecgh/go-spew/spew"	
	"github.com/labstack/echo"
		"github.com/jinzhu/gorm"

	"../models"
)

func Index(c echo.Context, db *gorm.DB) (err error) {
	users := []models.User{}
	list := db.Find(&users)
	return c.JSON(http.StatusCreated, list.Value)
}

func Show(c echo.Context, db *gorm.DB) (err error) {
	users := models.User{}
	id:= c.Param("id")
	list:= db.Where("Id = ?", id).Find(&users)
	return c.JSON(http.StatusCreated, list.Value)
}

func Create(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Hello, World!")
}
