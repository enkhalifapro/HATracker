package controllers

import (
  "encoding/json"
  "fmt"
  "github.com/enkhalifapro/HATracker/models"
  "github.com/labstack/echo"
  "log"
  "net/http"
)

func Signup(c echo.Context) error {
  defer c.Request().Body.Close()
  user := models.User{}
  err := json.NewDecoder(c.Request().Body).Decode(user)
  if err != nil {
    log.Printf("Faild Proccessing User Data: %s", err)
    return echo.NewHTTPError(http.StatusInternalServerError)
  }

  fmt.Println("Our user data is")
  fmt.Println(user)

  return c.String(http.StatusOK, "WE got your Data")
}
