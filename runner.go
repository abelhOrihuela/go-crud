package main

import (
	"net/http"	
	"github.com/labstack/echo"
	"./src/user"
)

func hello(c echo.Context) (err error) {
	u := new(user.User)
	u.Email = "Abel"
  if err = c.Bind(u); err != nil {
    return
  }
  return c.JSON(http.StatusOK, u)
}

func handleRoot(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Hello, World!")
}


func main() {
	e := echo.New()
	e.GET("/", handleRoot)
	e.GET("/users", hello)

	e.Logger.Fatal(e.Start(":1323"))
}
