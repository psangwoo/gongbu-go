package main

import (
	_ "backend/docs"
	"backend/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var db map[string]model.User

// @title Wookiist Sample Swagger API
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.GET("/*", echoSwagger.WrapHandler)

	db = make(map[string]model.User)
	e.GET("/user/:id", getUser)
	e.POST("/user", createUser)
	e.DELETE("/user/:id", deleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}

// @Summary Get user
// @Description Get user's info
// @Accept json
// @Produce json
// @Param id path string true "ID of the user"
// @Success 200 {object} model.User
// @Router /user/{id} [get]
func getUser(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("GETUSER : ", id)
	val, have := db[id]
	if !have {
		return c.String(http.StatusBadRequest, "없는 ID")
	}
	return c.JSONPretty(http.StatusOK, val, "  ")
}

// @Summary Create user
// @Description Create new user
// @Accept json
// @Produce json
// @Param userBody body model.User true "User Info Body"
// @Success 200 {object} model.User
// @Router /user [post]
func createUser(c echo.Context) error {
	u := new(model.User)
	err := c.Bind(u)

	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request. Check POST body")
	}
	val, have := db[u.ID]
	fmt.Println(u.ID)
	if have {
		return c.String(http.StatusConflict, fmt.Sprintf("Given ID %v is in Use", val.ID))
	}
	db[u.ID] = *u
	fmt.Println("asd")
	return c.String(http.StatusCreated, "Successful Request")
}

// @Summary Delete user
// @Description Delete user's info
// @Accept json
// @Produce json
// @Param id path string true "ID of the user"
// @Success 200 {object} model.User
// @Router /user/{id} [delete]
func deleteUser(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("GETUSER : ", id)
	val, have := db[id]
	if !have {
		return c.String(http.StatusBadRequest, "없는 ID")
	}
	delete(db, id)
	return c.JSONPretty(http.StatusOK, val, "  ")

}
