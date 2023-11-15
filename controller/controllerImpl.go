package controller

import (
	"net/http"
	"todo-api/helper"
	"todo-api/middleware"
	"todo-api/model"
	"todo-api/service"

	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TodoControllerImpl struct {
	TodoService service.TodoService
}

func NewTodoController(todoService service.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService,
	}
}

func (controller *TodoControllerImpl) Registrasi(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var registrasiRequest model.RegistrasiRequest
	helper.ReadFromRequestBody(request, &registrasiRequest)

	hassPassword, err := model.HashPassword(registrasiRequest.Password)
	helper.IfError(err)

	arg := model.RegistrasiRequest{
		Email:    registrasiRequest.Email,
		Username: registrasiRequest.Username,
		Password: hassPassword,
	}

	registrasiResponse := controller.TodoService.Registrasi(request.Context(), arg)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   registrasiResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *TodoControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var loginRequest model.LoginRequest
	helper.ReadFromRequestBody(request, &loginRequest)

	loginResponse := controller.TodoService.Login(request.Context(), loginRequest)
	err := model.CheckPassword(loginRequest.Password, loginResponse.Password)
	if err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   err,
		}
		helper.WriteToResponse(writer, webResponse)
		return
	}

	config, err := model.LoadConfig("../")
	helper.IfError(err)

	maker, err := middleware.NewPasetoMaker(config.TokenSymmetricKey)
	token, err := maker.CreateToken(loginRequest.Username, int(loginResponse.Userid), config.AccessTokenDuration)

	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data: model.LoginResponse{
			Email:    loginResponse.Email,
			Username: loginResponse.Username,
			Userid:   loginResponse.Userid,
			Token:    token,
		},
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *TodoControllerImpl) GetAllTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userIdStr := params.ByName("userId")
	userId, err := strconv.Atoi(userIdStr)
	helper.IfError(err)

	todoRequest := model.GetAllTodoRequest{
		Userid: int32(userId),
	}

	todoResponse := controller.TodoService.GetAllTodo(request.Context(), todoRequest)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   todoResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *TodoControllerImpl) AddTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoRequest := model.AddTodoBuffer
	todoResponse := controller.TodoService.AddTodo(request.Context(), todoRequest)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   todoResponse,
	}

	helper.WriteToResponse(writer, webResponse)
}

func (controller *TodoControllerImpl) GetTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userIdStr := request.URL.Query().Get("userId")
	userid, err := strconv.Atoi(userIdStr)
	helper.IfError(err)

	idStr := request.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	helper.IfError(err)

	todoRequest := model.GetorDeleteTodoRequest{
		Userid: int32(userid),
		ID:     int32(id),
	}

	todoResponse := controller.TodoService.GetTodo(request.Context(), todoRequest)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   todoResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *TodoControllerImpl) UpdateStatusTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userIdStr := request.URL.Query().Get("userId")
	userid, err := strconv.Atoi(userIdStr)
	helper.IfError(err)

	idStr := request.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	helper.IfError(err)

	var todoRequest model.UpdateStatusTodoRequest
	helper.ReadFromRequestBody(request, &todoRequest)

	todoRequest.ID = int32(id)
	todoRequest.Userid = int32(userid)

	todoResponse := controller.TodoService.UpdateStatusTodo(request.Context(), todoRequest)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   todoResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *TodoControllerImpl) DeleteTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userIdStr := request.URL.Query().Get("userId")
	userid, err := strconv.Atoi(userIdStr)
	helper.IfError(err)

	idStr := request.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	helper.IfError(err)

	todoRequest := model.GetorDeleteTodoRequest{
		Userid: int32(userid),
		ID:     int32(id),
	}

	todoResponse := controller.TodoService.DeleteTodo(request.Context(), todoRequest)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   todoResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *TodoControllerImpl) GetRandomTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userIdStr := params.ByName("userId")
	userid, err := strconv.Atoi(userIdStr)
	helper.IfError(err)

	arg := model.GetAllTodoRequest{
		Userid: int32(userid),
	}

	todoResponse := controller.TodoService.GetRandomTodo(request.Context(), arg)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   todoResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *TodoControllerImpl) GetTodoFilter(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userIdStr := request.URL.Query().Get("userId")
	userid, err := strconv.Atoi(userIdStr)
	helper.IfError(err)

	limitStr := request.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	helper.IfError(err)

	offsetStr := request.URL.Query().Get("offset")
	offset, err := strconv.Atoi(offsetStr)
	helper.IfError(err)

	todoRequest := model.GetTodoFilterRequest{
		Userid: int32(userid),
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	todoResponse := controller.TodoService.GetTodoFilter(request.Context(), todoRequest)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   todoResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}
