package controllers

import (
	"HATracker/DB"
	"HATracker/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// UserCtrl for Database connection------
type UserCtrl struct {
	Database DB.IPersistence
}

//_______________________________________________________________________________________________Sign_UP

// Signup Get User Data & Send it to Database------
func (user *UserCtrl) Signup(c echo.Context) error {

	defer c.Request().Body.Close()
	userdata := &models.Users{}
	err := json.NewDecoder(c.Request().Body).Decode(userdata)

	if err != nil {
		log.Printf("Faild Proccessing User Data: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// ---data access logic---
	err = user.Database.Connect()
	if err != nil {
		log.Fatalf("Faild to access Database %s", err)
	}

	// ---check fo Email
	query := fmt.Sprintf("SELECT * FROM users WHERE email ='%s'", userdata.Email)

	rusalt, err := user.Database.Select(query)
	if err != nil {
		fmt.Println(err)
	}
	if len(rusalt) == 0 {

		err = user.Database.Insert(*userdata)
		if err != nil {
			return err
		}

		/////////////////////////////GARBAGE///////////////
		fmt.Println("Our user data is")
		fmt.Println(user)
		//////////////////////////////////////////////////
		return c.String(http.StatusOK, "WE got your Data")

	} else {
		return c.String(http.StatusAlreadyReported, "this Email is already has an accont")
	}

}
