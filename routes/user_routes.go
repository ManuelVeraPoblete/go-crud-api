package routes

import (
    "go-crud-api/controllers"
    "go-crud-api/repositories"
    "go-crud-api/services"
    "github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
    repo := repositories.NewUserRepository()
    service := services.NewUserService(repo)
    controller := controllers.NewUserController(service)

    router.HandleFunc("/users", controller.CreateUser).Methods("POST")
    router.HandleFunc("/users", controller.GetAllUsers).Methods("GET")
    router.HandleFunc("/users/{id}", controller.GetUserByID).Methods("GET")
    router.HandleFunc("/users/{id}", controller.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")
}
