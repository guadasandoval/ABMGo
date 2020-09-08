package controller

import (
	"encoding/json"

	"/ABMGo/pkg/model"
	"net/http"
)

var NuevoUsuario model.Usuario

func CrearUsuario(w http.ResponseWriter, r *http.Request) {
	CrearUsuario := &model.Usuario{}
	utils.ParseBody(r, CrearUsuario)
	b := CrearUsuario.CrearUsuario()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	listaUsuarios := model.ListarTodosUsuarios()
	res, _ := json.Marshal(listaUsuarios)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
