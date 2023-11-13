package service

import (
	"github.com/user-assignment/app/db"
	"github.com/user-assignment/app/models"
)

type UserService struct {
	userDB *db.UserDB
}

func NewService(userDB *db.UserDB) *UserService {
	return &UserService{userDB: userDB}
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	return s.userDB.CreateUser(user)
}

func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {
	return s.userDB.UpdateUser(user)
}

func (s *UserService) DeleteUser(user *models.User) error {
	return s.userDB.DeleteUser(user)
}

func (s *UserService) GetAllUsers() *[]models.User {
	return s.userDB.GetAllUsers()
}
