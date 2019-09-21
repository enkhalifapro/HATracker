package services

import (
	"HATracker/DB"
	"HATracker/models"
	"github.com/labstack/gommon/log"
	"fmt"
)

type IUser interface {
	Add(user *models.User) error
}

type UserService struct {
	Database DB.IPersistence
}

func (s *UserService) Add(user *models.User) error {
	// ---data access logic---
	err := s.Database.Connect()
	if err != nil {
		log.Error("Faild to access Database %s", err)
		return err
	}

	// ---check fo Email
	query := fmt.Sprintf("SELECT * FROM users WHERE email ='%s'", user.Email)

	result, err := s.Database.Select(query)
	if err != nil {
		log.Error(err)
	}

	if len(result) != 0 {
		return fmt.Errorf("E-mail is already exist")
	}

	err = s.Database.Insert(user)
	return err
}
