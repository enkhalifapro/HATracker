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
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

//______________Functions_________________________
//---GetUserData--
func ScanUserData(user *User, r *echo.Echo) func(c echo.Context) error {

	return func(c echo.Context) error {

		defer c.Request().Body.Close()

		err := json.NewDecoder(c.Request().Body).Decode(user)

		if err != nil {
			log.Printf("Faild Proccessing User Data: %s", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.String(http.StatusOK, "WE got your Data")
	}
}
func main() {

	r := echo.New()

	user := User{}

	r.POST("/sign_up", ScanUserData(&user, r))

	r.GET("/v", func(c echo.Context) error {

		fmt.Printf("hi")
		fmt.Printf("%v", user.FirstName)
		return c.String(http.StatusOK, "")
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
