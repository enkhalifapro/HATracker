package models

type User struct {
  FirstName string `json:"firstName"`
  LastName  string `json:"lastName"`
  Password  string `json:"password"`
  Email     string `json:"email"`
}
