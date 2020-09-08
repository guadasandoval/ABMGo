package rutas

import (
	"ABMGo/pkg/controller"

	"github.com/gorilla/mux"
)

var ABMUsuarios = func(router *mux.Router) {
	router.HandleFunc("/usuario/", controller.CrearUsuario).Methods("POST")
	router.HandleFunc("/usuario/", controller.ListarUsuarios).Methods("GET")

}
