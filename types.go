package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(handler http.HandlerFunc) http.HandlerFunc

type User struct {
	//usando el `json:""` transforma en formato json las variables cuando se usan
	//en caso de que no se usen en formato json, se quedan de la forma definida.
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
type MetaData interface {
}

//funcion para encodear a json
func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
