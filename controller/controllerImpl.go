package controller

import (
	"net/http"
	"todo-api/helper"
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

	registrasiResponse := controller.TodoService.Registrasi(request.Context(), registrasiRequest)
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

	LoginResponse := controller.TodoService.Login(request.Context(), loginRequest)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   LoginResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *TodoControllerImpl) GetAllTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var todoRequest model.GetAllTodoRequest
	helper.ReadFromRequestBody(request, &todoRequest)

	todoResponse := controller.TodoService.GetAllTodo(request.Context(), todoRequest)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   todoResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *TodoControllerImpl) AddTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var todoRequest model.AddNewTodoRequest
	helper.ReadFromRequestBody(request, &todoRequest)

	todoResponse := controller.TodoService.AddTodo(request.Context(), todoRequest)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   todoResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *TodoControllerImpl) GetTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	useridStr := params.ByName("userid")
	userid, err := strconv.Atoi(useridStr)
	helper.IfError(err)

	idStr := params.ByName("id")
	id, err := strconv.Atoi(idStr)
	helper.IfError(err)

	todoRequest := model.GetorDeleteTodoRequest{
		Userid: int32(userid),
		ID:     int32(id),
	}
	helper.ReadFromRequestBody(request, &todoRequest)

	todoResponse := controller.TodoService.GetTodo(request.Context(), todoRequest)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   todoResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *TodoControllerImpl) UpdateStatusTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var todoRequest model.UpdateStatusTodoRequest
	helper.ReadFromRequestBody(request, &todoRequest)

	todoResponse := controller.TodoService.UpdateStatusTodo(request.Context(), todoRequest)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   todoResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *TodoControllerImpl) DeleteTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	useridStr := params.ByName("userid")
	userid, err := strconv.Atoi(useridStr)
	helper.IfError(err)

	idStr := params.ByName("id")
	id, err := strconv.Atoi(idStr)
	helper.IfError(err)

	todoRequest := model.GetorDeleteTodoRequest{
		Userid: int32(userid),
		ID:     int32(id),
	}
	helper.ReadFromRequestBody(request, &todoRequest)

	todoResponse := controller.TodoService.DeleteTodo(request.Context(), todoRequest)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   todoResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *TodoControllerImpl) GetRandomTodo(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoResponse := controller.TodoService.GetRandomTodo(request.Context())
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   todoResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}

func (controller *TodoControllerImpl) GetTodoFilter(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	useridStr := params.ByName("userid")
	userid, err := strconv.Atoi(useridStr)
	helper.IfError(err)

	limitStr := params.ByName("limit")
	limit, err := strconv.Atoi(limitStr)
	helper.IfError(err)

	offsetStr := params.ByName("offset")
	offset, err := strconv.Atoi(offsetStr)
	helper.IfError(err)

	todoRequest := model.GetTodoFilterRequest{
		Userid: int32(userid),
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	helper.ReadFromRequestBody(request, &todoRequest)

	todoResponse := controller.TodoService.GetTodoFilter(request.Context(), todoRequest)
	webResponse := model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   todoResponse,
	}
	helper.WriteToResponse(writer, webResponse)
}
