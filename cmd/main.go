package main

import (
    "log"
    "net/http"
    "go-crud-api/routes"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    
    // Configurar las rutas de usuario
    routes.UserRoutes(router)
    
    log.Println("Servidor iniciado en el puerto 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
