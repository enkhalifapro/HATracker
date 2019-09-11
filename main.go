package main

import (
  "github.com/labstack/echo"
  "net/http"
)

type Issue struct {
  ID    int
  Title string
}

func main() {
  r := echo.New()

  r.GET("/hager/ahmed", func(c echo.Context) error {
    myIssue := &Issue{ID: 1, Title: "My first issue"}
    return c.JSON(http.StatusOK, myIssue)
  })

  r.Start("localhost:7071")
}
