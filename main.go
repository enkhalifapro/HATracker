package main

import (
	"HATracker/DB"
	"HATracker/controllers"

	"github.com/labstack/echo"
)

func main() {

	// router config
	r := echo.New()

	userCtrl := &controllers.UserCtrl{
		Database: &DB.PostgresHelper{},
	}

	r.POST("/signup", userCtrl.Signup)

	r.Logger.Fatal(r.Start("localhost:1323"))
}
