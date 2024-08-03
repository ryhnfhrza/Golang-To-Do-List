package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/helper"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/web"
)

func ErrorHandler(writer http.ResponseWriter , request *http.Request, err interface{}){

	if notFoundError(writer,request,err){
		return
	}
	if validationErrors(writer,request,err){
		return
	}
	if unauthorizedError(writer, request, err) {
		return
}


	internalServerError(writer,request,err)

}

func validationErrors(writer http.ResponseWriter , request *http.Request, err interface{}) bool{
	exception,ok := err.(validator.ValidationErrors)
	if ok{
		writer.Header().Set("Content-Type","application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code: http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: exception.Error(),
		}
		
		helper.WriteToResponseBody(writer,webResponse)
		return true
	}else{
		return false
	}
}

func notFoundError(writer http.ResponseWriter , request *http.Request, err interface{}) bool {
	exception,ok := err.(NotFoundError)
	if ok{
		writer.Header().Set("Content-Type","application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code: http.StatusNotFound,
			Status: "NOT FOUND",
			Data: exception.Error,
		}
		helper.WriteToResponseBody(writer,webResponse)
		return true
	}else{
		return false
	}
}

func internalServerError(writer http.ResponseWriter , request *http.Request, err interface{}){
	writer.Header().Set("Content-Type","application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		Code: http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data: err,
	}

	helper.WriteToResponseBody(writer,webResponse)
}

func unauthorizedError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	unauthorizedErr, ok := err.(*UnauthorizedError)
	if ok {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "UNAUTHORIZED",
					Data:   map[string]string{"error": unauthorizedErr.Error()},
			}

			helper.WriteToResponseBody(writer, webResponse)
			return true
	}
	return false

}

func WriteUnauthorizedError(writer http.ResponseWriter, errorMessage string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   map[string]string{"error": errorMessage},
	}

	helper.WriteToResponseBody(writer, webResponse)
}