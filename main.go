package main

import (
	"fmt"
	"github.com/enkhalifapro/HATracker/controllers"
	"net/http"
	"github.com/labstack/echo"
)

func main() {

	// router config
	r := echo.New()

	r.POST("/signup", controllers.Signup)

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
