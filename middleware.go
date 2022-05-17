package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//funcionan en cadena ya que va chequeando cada handler
func CheckAuth() Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {

			//logica de la funcion
			flag := true
			fmt.Println("Checking authentication")

			if flag {
				f(writer, request)
			} else {
				return
			}
		}
	}
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {

			start := time.Now()
			defer func() {
				log.Println(request.URL, time.Since(start))
				//estos parentesis al final son para llamar a la funcion anonima creada
			}()
			//con esto salta al siguiente middleware
			f(writer, request)
		}
	}
}
