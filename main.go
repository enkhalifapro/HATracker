package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

//______________STructData________________________
//---User---
type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

//______________Functions_________________________
//---GetUserData--

func signup(c echo.Context) error {
	defer c.Request().Body.Close()
	user := &User{}
	err := json.NewDecoder(c.Request().Body).Decode(user)
	if err != nil {
		log.Printf("Faild Proccessing User Data: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	fmt.Println("Our user data is")
	fmt.Println(user)

	return c.String(http.StatusOK, "WE got your Data")
}

func main() {

	r := echo.New()

	r.POST("/signup", signup)

	r.GET("/v", func(c echo.Context) error {

		fmt.Printf("hi")
		return c.String(http.StatusOK, "Service is working!!!")
	})

	r.Logger.Fatal(r.Start("localhost:1323"))
}

/*
{
    "first_name":"Ahmed",
    "last_name":"Ali",
    "password":"dv444",
    "email":"asdf"

}


*/
