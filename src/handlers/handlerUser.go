package handlersUser


import (
)

import (
	"net/http"	
	"github.com/davecgh/go-spew/spew"	
	"github.com/labstack/echo"
	// "../models"
)

// func Index(name string, email string) *models.User {
//     user := models.User{Name: name, Email: email}
//     return &user
// }

func Index(c echo.Context) (err error) {
	id:= c.QueryParam("id")
	spew.Println("id", id)
	return c.String(http.StatusOK, "Hello, World!" + id)
}
func Create(c echo.Context) (err error) {
	// u := &models.User{
	// 	ID: seq,
	// }
	// if err := c.Bind(u); err != nil {
	// 	return err
	// }
	// users[u.ID] = u
	// seq++
	// return c.JSON(http.StatusCreated, u)


	// id:= c.QueryParam("id")
	// spew.Println("id", id)
	return c.String(http.StatusOK, "Hello, World!")
}
