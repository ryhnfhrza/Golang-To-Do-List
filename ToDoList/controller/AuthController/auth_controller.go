package controller

import (
	"net/http"
)

type AuthController interface {
	Registration(writer http.ResponseWriter, request *http.Request)
	Login(writer http.ResponseWriter, request *http.Request)
	Logout(writer http.ResponseWriter, request *http.Request)
}