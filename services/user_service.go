package services

import (
    "go-crud-api/models"
    "go-crud-api/repositories"
)

type UserService interface {
    CreateUser(user *models.User) error
    GetAllUsers() ([]models.User, error)
    GetUserByID(id uint) (*models.User, error)
    UpdateUser(user *models.User) error
    DeleteUser(id uint) error
}

type userService struct {
    userRepository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
    return &userService{userRepository: repo}
}

func (s *userService) CreateUser(user *models.User) error {
    return s.userRepository.Create(user)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
    return s.userRepository.GetAll()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
    return s.userRepository.GetByID(id)
}

func (s *userService) UpdateUser(user *models.User) error {
    return s.userRepository.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
    return s.userRepository.Delete(id)
}
