package model

import "todo-api/repository"

type LoginResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Userid   int32  `json:"userid"`
	Token    string `json:"token"`
}

type RegistrasiResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Userid   int32  `json:"userid"`
}

type GetAllTodosResponse struct {
	Todos []repository.GetAllTodosRow `json:"todos"`
	Total int64                       `json:"total"`
}

type GetTodosGFilterResponse struct {
	Todos []repository.GetAllTodosRow `json:"todos"`
	Total int32                       `json:"total"`
	Skip  int32                       `json:"skip"`
	Limit int32                       `json:"limit"`
}

type GetTodoFilterResponse struct {
	Todos []repository.GetSomeTodosRow `json:"todos"`
	Total int32                        `json:"total"`
	Skip  int32                        `json:"skip"`
	Limit int32                        `json:"limit"`
}

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
