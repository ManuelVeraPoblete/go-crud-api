package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "go-crud-api/models"
    "go-crud-api/services"
    "github.com/gorilla/mux"
)

type UserController struct {
    UserService services.UserService
}

func NewUserController(service services.UserService) *UserController {
    return &UserController{UserService: service}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    
    err := c.UserService.CreateUser(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := c.UserService.GetAllUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    
    user, err := c.UserService.GetUserByID(uint(id))
    if err != nil || user == nil {
        http.Error(w, "Usuario no encontrado", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    user.ID = uint(id)
    
    err := c.UserService.UpdateUser(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    
    err := c.UserService.DeleteUser(uint(id))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
