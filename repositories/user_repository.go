package repositories

import "go-crud-api/models"

type UserRepository interface {
    Create(user *models.User) error
    GetAll() ([]models.User, error)
    GetByID(id uint) (*models.User, error)
    Update(user *models.User) error
    Delete(id uint) error
}

type userRepo struct {
    users []models.User // Simulando una base de datos en memoria
}

func NewUserRepository() UserRepository {
    return &userRepo{users: []models.User{}}
}

func (r *userRepo) Create(user *models.User) error {
    user.ID = uint(len(r.users) + 1)
    r.users = append(r.users, *user)
    return nil
}

func (r *userRepo) GetAll() ([]models.User, error) {
    return r.users, nil
}

func (r *userRepo) GetByID(id uint) (*models.User, error) {
    for _, user := range r.users {
        if user.ID == id {
            return &user, nil
        }
    }
    return nil, nil
}

func (r *userRepo) Update(user *models.User) error {
    for i, u := range r.users {
        if u.ID == user.ID {
            r.users[i] = *user
            return nil
        }
    }
    return nil
}

func (r *userRepo) Delete(id uint) error {
    for i, user := range r.users {
        if user.ID == id {
            r.users = append(r.users[:i], r.users[i+1:]...)
            return nil
        }
    }
    return nil
}
