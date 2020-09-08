package model

import (
	"ABMGo/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Usuario struct {
	gorm.Model

	Nombre   string `gorm:""json:"nombre"`
	Apellido string `json:"apellido"`
	Email    string `json:"email"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Usuario{})
}

func (b *Usuario) CrearUsuario() *Usuario {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func ListarTodosUsuarios() []Usuario {
	var Usuarios []Usuario
	db.Find(&Usuario)
	return Usuarios
}
