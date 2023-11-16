package exception

import (
	"net/http"
	"regexp"

	"todo-api/helper"
	"todo-api/model"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if validationError(writer, request, err) {
		return
	}
	if UnauthorizedError(writer, request, err) {
		return
	}
	if notFoundError(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}
		helper.WriteToResponse(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	patternDuplicate := "duplicate key value violates unique constraint"
	regex := regexp.MustCompile(patternDuplicate)
	errDuplicate := exception.Error

	if ok && regex.MatchString(errDuplicate) {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusConflict)

		webResponse := model.WebResponse{
			Code:   http.StatusConflict,
			Status: "CONFLICT",
			Data:   err,
		}
		helper.WriteToResponse(writer, webResponse)
		return true
	}
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := model.WebResponse{
			Code:   http.StatusNotFound,
			Status: "REQUEST NOT FOUND",
			Data:   exception.Error,
		}
		helper.WriteToResponse(writer, webResponse)
		return true
	} else {
		return false
	}
}

func UnauthorizedError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(UnauthorizedErr)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   exception.Error,
		}
		helper.WriteToResponse(writer, webResponse)
		return true
	} else {
		return false
	}

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := model.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR`",
		Data:   err,
	}
	helper.WriteToResponse(writer, webResponse)
}
