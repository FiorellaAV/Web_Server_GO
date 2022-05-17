package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//se encarga de manejar la ruta principal
func HandlerRoot(w http.ResponseWriter, r *http.Request) {

	fmt.Println(fmt.Fprintf(w, "Hello World!"))

}

func HandlerHome(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "This is the API endpoint")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata MetaData
	//necesitamos pasarle un a variable de tipo interfaz para que vuelque la informacion
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "erorr: %v", err)
		return
	}

	fmt.Fprintf(w, "Payload %v\n", metadata)
}

//esta funcion agarra el json que se le envia y lo trasforma en un objeto utilizable
func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	//necesitamos pasarle un a variable de tipo interfaz para que vuelque la informacion
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "erorr: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
