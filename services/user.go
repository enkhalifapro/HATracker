package services

import (
	"HATracker/DB"
	"HATracker/helpers"
	"HATracker/models"
	"github.com/labstack/gommon/log"
	"fmt"
)

type IUser interface {
	Add(user *models.User) error
}

type UserService struct {
	Database       DB.IPersistence
	PasswordHelper helpers.Password
}

func (s *UserService) Add(user *models.User) error {
	// db start connection
	err := s.Database.Connect()
	defer s.Database.Close()
	if err != nil {
		log.Error("Faild to access Database %s", err)
		return err
	}

	// ---check Email existance
	query := fmt.Sprintf("SELECT * FROM users WHERE email ='%s'", user.Email)
	result, err := s.Database.Select(query)
	if err != nil {
		log.Error(err)
	}

	if len(result) != 0 {
		return fmt.Errorf("E-mail is already exist")
	}

	// Hash password
	user.Password = s.PasswordHelper.Hash(user.Password)

	// insert into db
	err = s.Database.Insert(user)
	return err
}
