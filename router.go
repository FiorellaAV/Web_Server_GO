package main

import (
	"net/http"
)

//Struct router para hacer request en el servidor
type Router struct {
	//Reglas para definir de que rutas pasan a que handler, mapa que pasa de strings a handler
	//mapa que tenga como llaves strings y que mapee a HandlerFunc
	rules map[string]map[string]http.HandlerFunc
}

//forma de instanciar el router, similar al NewServer() del archivo servidor.go
/*a diferencia del servidor, aqui el router debe empezar en un estado vacio, creamos un mapa vacio*/
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

//El primer paramentro seria la "API", la ruta (/home), el segundo parametro es par el metodo que
//se quiera realizar (POST, GET, )
func (r *Router) findHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	// en los corchetes va la ruta (es la llave) para verificar si existe
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, exist
}

//Metodo ServeHTTP de router para poder implementar en el handler el atributo s.router en server.go
//parametros: el primero es el escritor, el segundo es el request en donde viene la informacion
//no olvidar colocar ServeHTTP con letras mayusculas
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {

	//impresion de mensaje respuesta que el servidor da a la ruta
	//Fprintf es un escritor, que recibe w que es el escritor asignado, y el mensaje que queremos mostrar
	//fmt.Fprintf(w, "Hello world!")

	handler, methodExist, exist := r.findHandler(request.URL.Path, request.Method)
	//w es quien le escribe al usuario, quien le dice que paso
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)
}
