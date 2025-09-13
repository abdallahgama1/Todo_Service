package services

import (
	"Todo_Service/models"
	"Todo_Service/repositories"
	"Todo_Service/utils"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(input models.RegisterInput) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Username: input.Username,
		Password: hashedPassword,
	}

	err = s.repo.CreateUser(newUser)
	return newUser, err
}

func (s *UserService) Login(input models.LoginInput) (string, error) {
	user, err := s.repo.GetUserByUsername(input.Username)
	if err != nil {
		return "", err 
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return "", gorm.ErrInvalidValue 

	return utils.GenerateToken(user.ID)
}