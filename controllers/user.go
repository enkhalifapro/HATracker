package controllers

import (
	"HATracker/models"
	"HATracker/services"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// UserCtrl for Database connection------
type UserCtrl struct {
	Service services.IUser
}

//_______________________________________________________________________________________________Sign_UP

// Signup Get User Data & Send it to Database------
func (s *UserCtrl) Signup(c echo.Context) error {

	defer c.Request().Body.Close()
	user := &models.User{}
	err := json.NewDecoder(c.Request().Body).Decode(user)

	if err != nil {
		log.Printf("Faild Proccessing User Data: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	err = s.Service.Add(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}
