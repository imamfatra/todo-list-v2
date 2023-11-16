package controller_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"todo-api/app"
	"todo-api/controller"
	"todo-api/helper"
	"todo-api/middleware"
	"todo-api/model"
	"todo-api/repository"
	"todo-api/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type todoControllerSuite struct {
	suite.Suite
	repository      *repository.Queries
	service         service.TodoService
	controller      controller.TodoController
	httpTest        *httptest.Server
	cleanUpDatabase model.TruncateTableExecutor
}

func (suite *todoControllerSuite) SetupTest() http.Handler {
	validate := validator.New()
	config, err := model.LoadConfig("../.")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	repository := repository.NewTodoRepository()
	service := service.NewTodoService(repository, validate, db)
	controller := controller.NewTodoController(service)
	routerApp := app.NewRouter(controller)

	suite.repository = repository
	suite.service = service
	suite.controller = controller
	suite.cleanUpDatabase = model.InitTruncateTableExecutor(db)

	return middleware.Auth(routerApp)
}

func (suite *todoControllerSuite) TearDownTest() {
	defer suite.cleanUpDatabase.TruncateTable([]string{"users", "todos"})
}

func (suite *todoControllerSuite) registrasiTodo() model.RegistrasiResponse {
	router := suite.SetupTest()

	arg := model.RegistrasiRequest{
		Email:    model.RandomMail(),
		Username: model.RandomString(10),
		Password: "secret123",
	}
	requestBody, err := json.Marshal(arg)
	helper.IfError(err)
	request := httptest.NewRequest(http.MethodPost, "/api/user/registrasi", bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)

	response := record.Result()
	suite.Equal(200, response.StatusCode)

	body, _ := io.ReadAll(record.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	// fmt.Println(responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 200)
	suite.Equal(responseBody["status"], "SUCCESS")
	suite.Equal(responseBody["data"].(map[string]interface{})["username"], arg.Username)
	suite.Equal(responseBody["data"].(map[string]interface{})["email"], arg.Email)
	suite.NotZero(responseBody["data"].(map[string]interface{})["userid"])

	return model.RegistrasiResponse{
		Email:    responseBody["data"].(map[string]interface{})["email"].(string),
		Username: responseBody["data"].(map[string]interface{})["username"].(string),
		Userid:   int32(responseBody["data"].(map[string]interface{})["userid"].(float64)),
	}
}

func (suite *todoControllerSuite) TestRegistrasi_Positive() {
	user := suite.registrasiTodo()
	suite.NotZero(user.Userid)
}

func (suite *todoControllerSuite) TestRegistrasi_NilPointer_Negative() {
	router := suite.SetupTest()

	arg := model.RegistrasiRequest{}
	requestBody, err := json.Marshal(arg)
	helper.IfError(err)
	request := httptest.NewRequest(http.MethodPost, "/api/user/registrasi", bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)

	response := record.Result()
	suite.Equal(http.StatusBadRequest, response.StatusCode)
}

func (suite *todoControllerSuite) TestRegistrasi_ValueFailed_Negative() {
	router := suite.SetupTest()

	arg := model.RegistrasiRequest{
		Email:    model.RandomMail(),
		Username: "abc",
		Password: "secret1234",
	}
	requestBody, err := json.Marshal(arg)
	helper.IfError(err)
	request := httptest.NewRequest(http.MethodPost, "/api/user/registrasi", bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)

	response := record.Result()
	suite.Equal(http.StatusBadRequest, response.StatusCode)
}

func (suite *todoControllerSuite) TestRegistrasi_DuplicateData_Negative() {
	user := suite.registrasiTodo()

	router := suite.SetupTest()

	arg := model.RegistrasiRequest{
		Email:    user.Email,
		Username: user.Username,
		Password: "secret1234",
	}
	requestBody, err := json.Marshal(arg)
	helper.IfError(err)
	request := httptest.NewRequest(http.MethodPost, "/api/user/registrasi", bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)

	response := record.Result()
	suite.Equal(http.StatusConflict, response.StatusCode)
}

func (suite *todoControllerSuite) TestLoginAccout_Positive() {
	user := suite.registrasiTodo()
	router := suite.SetupTest()

	arg := model.LoginRequest{
		Username: user.Username,
		Password: "secret123",
	}
	requestBody, err := json.Marshal(arg)
	helper.IfError(err)
	request := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)

	respones := record.Result()
	suite.Equal(http.StatusOK, respones.StatusCode)

	body, _ := io.ReadAll(respones.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	suite.Equal(int(responseBody["code"].(float64)), 200)
	suite.Equal(responseBody["status"], "SUCCESS")
	suite.Equal(responseBody["data"].(map[string]interface{})["username"], arg.Username)
	suite.NotEmpty(responseBody["data"].(map[string]interface{})["email"])
	suite.NotEmpty(responseBody["data"].(map[string]interface{})["token"])
	suite.NotZero(responseBody["data"].(map[string]interface{})["userid"])

}

func (suite *todoControllerSuite) TestLoginAccout_WrongPassword_Negative() {
	user := suite.registrasiTodo()
	router := suite.SetupTest()

	arg := model.LoginRequest{
		Username: user.Username,
		Password: "secret12",
	}
	requestBody, err := json.Marshal(arg)
	helper.IfError(err)
	request := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)

	respones := record.Result()
	suite.Equal(http.StatusUnauthorized, respones.StatusCode)

	body, _ := io.ReadAll(respones.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(responseBody["status"], "UNAUTHORIZED")
	suite.Equal(int(responseBody["code"].(float64)), 401)
}

func (suite *todoControllerSuite) TestLoginAccout_WrongUsername_Negative() {
	suite.registrasiTodo()
	router := suite.SetupTest()

	arg := model.LoginRequest{
		Username: model.RandomString(10),
		Password: "secret122",
	}
	requestBody, err := json.Marshal(arg)
	helper.IfError(err)
	request := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)

	respones := record.Result()
	suite.Equal(http.StatusUnauthorized, respones.StatusCode)

	body, _ := io.ReadAll(respones.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(responseBody["status"], "UNAUTHORIZED")
	suite.Equal(int(responseBody["code"].(float64)), 401)
}

func (suite *todoControllerSuite) TestLoginAccount_NilPointer_Negative() {
	router := suite.SetupTest()

	arg := model.LoginRequest{}
	requestBody, err := json.Marshal(arg)
	helper.IfError(err)
	request := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)

	respones := record.Result()
	suite.Equal(http.StatusBadRequest, respones.StatusCode)

	body, _ := io.ReadAll(respones.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(responseBody["status"], "BAD REQUEST")
	suite.Equal(int(responseBody["code"].(float64)), 400)
}

func (suite *todoControllerSuite) loginAccount() model.LoginResponse {
	user := suite.registrasiTodo()
	router := suite.SetupTest()

	arg := model.LoginRequest{
		Username: user.Username,
		Password: "secret123",
	}
	requestBody, err := json.Marshal(arg)
	helper.IfError(err)
	request := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)

	respones := record.Result()
	suite.Equal(http.StatusOK, respones.StatusCode)
	body, _ := io.ReadAll(respones.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(responseBody["data"].(map[string]interface{})["username"], arg.Username)
	suite.NotEmpty(responseBody["data"].(map[string]interface{})["email"])
	suite.NotEmpty(responseBody["data"].(map[string]interface{})["token"])
	suite.NotZero(responseBody["data"].(map[string]interface{})["userid"])

	return model.LoginResponse{
		Email:    responseBody["data"].(map[string]interface{})["email"].(string),
		Username: responseBody["data"].(map[string]interface{})["username"].(string),
		Userid:   int32(responseBody["data"].(map[string]interface{})["userid"].(float64)),
		Token:    responseBody["data"].(map[string]interface{})["token"].(string),
	}
}

func (suite *todoControllerSuite) addOneTodo() model.LoginResponse {
	router := suite.SetupTest()
	user := suite.loginAccount()

	arg := model.AddNewTodoRequest{
		Todo:      "Create API todolist",
		Complated: false,
		Userid:    user.Userid,
	}
	requestBody, err := json.Marshal(arg)
	helper.IfError(err)
	request := httptest.NewRequest(http.MethodPost, "/api/todo", bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("authorization", user.Token)

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)

	response := record.Result()
	suite.Equal(200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 200)
	suite.Equal(responseBody["status"].(string), "SUCCESS")
	suite.Equal(responseBody["data"].(map[string]interface{})["todo"].(string), arg.Todo)
	suite.Equal(responseBody["data"].(map[string]interface{})["userid"].(float64), float64(arg.Userid))

	return user
}

func (suite *todoControllerSuite) addManyTodo() model.LoginResponse {
	router := suite.SetupTest()
	user := suite.loginAccount()

	record := httptest.NewRecorder()
	for i := 0; i < 100; i++ {
		arg := model.AddNewTodoRequest{
			Todo:      "Create API todolist",
			Complated: false,
			Userid:    user.Userid,
		}
		requestBody, err := json.Marshal(arg)
		helper.IfError(err)
		request := httptest.NewRequest(http.MethodPost, "/api/todo", bytes.NewBuffer(requestBody))
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("authorization", user.Token)
		router.ServeHTTP(record, request)
		response := record.Result()
		suite.Equal(200, response.StatusCode)

	}

	return user
}

func (suite *todoControllerSuite) TestAddTodoMany() {
	suite.addManyTodo()
}

func (suite *todoControllerSuite) TestAddTodo_Positive() {
	suite.addOneTodo()
	// suite.addManyTodo()
}

func (suite *todoControllerSuite) TestAddTodo_Authorization_Negative() {
	router := suite.SetupTest()
	user := suite.loginAccount()

	arg := model.AddNewTodoRequest{
		Todo:      "Create API todolist",
		Complated: false,
		Userid:    user.Userid,
	}
	requestBody, err := json.Marshal(arg)
	helper.IfError(err)
	request := httptest.NewRequest(http.MethodPost, "/api/todo", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("authorization", model.RandomString(50))

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)
	response := record.Result()
	suite.Equal(401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), http.StatusUnauthorized)
	suite.Equal(responseBody["status"].(string), "UNAUTHORIZED")
}

func (suite *todoControllerSuite) TestGetAllTodo_Positive() {
	suite.registrasiTodo()
	user := suite.addManyTodo()
	router := suite.SetupTest()

	url := fmt.Sprintf("/api/todos/" + strconv.Itoa(int(user.Userid)))
	request := httptest.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("authorization", user.Token)

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)
	response := record.Result()
	suite.Equal(200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 200)
	suite.Equal(responseBody["status"].(string), "SUCCESS")
	suite.Equal(int(responseBody["data"].(map[string]interface{})["total"].(float64)), 100)
}

func (suite *todoControllerSuite) TestGetAllTOdo_TokenWrong_Negative() {
	user := suite.addManyTodo()
	router := suite.SetupTest()

	url := fmt.Sprintf("/api/todos/" + strconv.Itoa(int(user.Userid)))
	request := httptest.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("authorization", user.Token+"1bc")

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)
	response := record.Result()
	suite.Equal(http.StatusUnauthorized, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 401)
	suite.Equal(responseBody["status"].(string), "UNAUTHORIZED")
}

func (suite *todoControllerSuite) TestGetAllTOdo_TokenWrongWithAnotherAccount_Negative() {
	user := suite.addManyTodo()
	router := suite.SetupTest()
	user2 := suite.loginAccount()

	url := fmt.Sprintf("/api/todos/" + strconv.Itoa(int(user.Userid)))
	request := httptest.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("authorization", user2.Token)

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)
	response := record.Result()
	suite.Equal(http.StatusNotFound, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 404)
	suite.Equal(responseBody["status"].(string), "Page Not Found")
}

func (suite *todoControllerSuite) TestGetAllTOdo_WrongUserId_Negative() {
	user := suite.addManyTodo()
	router := suite.SetupTest()

	url := fmt.Sprintf("/api/todos/" + "1000")
	request := httptest.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("authorization", user.Token)

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)
	response := record.Result()
	suite.Equal(http.StatusNotFound, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 404)
	suite.Equal(responseBody["status"].(string), "Page Not Found")
}

func (suite *todoControllerSuite) TestGetTodoSingle_Positive() {
	user := suite.addOneTodo()
	router := suite.SetupTest()
	userId := strconv.Itoa(int(user.Userid))

	url := fmt.Sprintf("/api/todo?userId=%s&id=%s", userId, "1")
	request := httptest.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("authorization", user.Token)

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)
	response := record.Result()
	suite.Equal(http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 200)
	suite.Equal(responseBody["status"].(string), "SUCCESS")
	suite.Equal(responseBody["data"].(map[string]interface{})["userid"].(float64), float64(user.Userid))
	suite.Equal(responseBody["data"].(map[string]interface{})["id"].(float64), float64(1))
}

func (suite *todoControllerSuite) TestGetTodoSingle_WrongToken_Negative() {
	user := suite.addOneTodo()
	router := suite.SetupTest()
	userId := strconv.Itoa(int(user.Userid))

	url := fmt.Sprintf("/api/todo?userId=%s&id=%s", userId, "1")
	request := httptest.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("authorization", user.Token+"abc")

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)
	response := record.Result()
	suite.Equal(http.StatusUnauthorized, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 401)
	suite.Equal(responseBody["status"].(string), "UNAUTHORIZED")
}

func (suite *todoControllerSuite) TestGetTodoSingle_WrongUserId_Negative() {
	user := suite.addOneTodo()
	router := suite.SetupTest()
	userId := "100"

	url := fmt.Sprintf("/api/todo?userId=%s&id=%s", userId, "1")
	request := httptest.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("authorization", user.Token)

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)
	response := record.Result()
	suite.Equal(http.StatusNotFound, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 404)
	suite.Equal(responseBody["status"].(string), "Page Not Found")
}

func (suite *todoControllerSuite) TestGetTodoSingle_IdNotFound_Negative() {
	user := suite.addOneTodo()
	router := suite.SetupTest()
	userId := strconv.Itoa(int(user.Userid))

	url := fmt.Sprintf("/api/todo?userId=%s&id=%s", userId, "1000")
	request := httptest.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("authorization", user.Token)

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)
	response := record.Result()
	suite.Equal(http.StatusNotFound, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 404)
	suite.Equal(responseBody["status"].(string), "REQUEST NOT FOUND")
}

func (suite *todoControllerSuite) TestGetTodoSingle_NilPointer_Negative() {
	user := suite.addOneTodo()
	router := suite.SetupTest()

	url := fmt.Sprintf("/api/todo?userId=%s&id=%s", "", "")
	request := httptest.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("authorization", user.Token)

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)
	response := record.Result()
	suite.Equal(http.StatusNotFound, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 404)
	suite.Equal(responseBody["status"].(string), "Page Not Found")
}

func (suite *todoControllerSuite) TestUpdateStatus_Positive() {
	user := suite.addOneTodo()
	router := suite.SetupTest()
	userId := strconv.Itoa(int(user.Userid))
	arg := model.UpdateStatusTodoRequest{
		Complated: true,
	}
	requestBody, err := json.Marshal(arg)
	helper.IfError(err)

	url := fmt.Sprintf("/api/todo?userId=%s&id=%s", userId, "1")
	request := httptest.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	request.Header.Set("authorization", user.Token)

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)
	response := record.Result()
	suite.Equal(http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 200)
	suite.Equal(responseBody["status"].(string), "SUCCESS")
	suite.Equal(responseBody["data"].(map[string]interface{})["userid"].(float64), float64(user.Userid))
	suite.Equal(responseBody["data"].(map[string]interface{})["id"].(float64), float64(1))
}

func (suite *todoControllerSuite) TestUpdateStatus_WrongToken_Negative() {
	user := suite.addOneTodo()
	router := suite.SetupTest()
	userId := strconv.Itoa(int(user.Userid))
	arg := model.UpdateStatusTodoRequest{
		Complated: true,
	}
	requestBody, err := json.Marshal(arg)
	helper.IfError(err)

	url := fmt.Sprintf("/api/todo?userId=%s&id=%s", userId, "1")
	request := httptest.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	request.Header.Set("authorization", user.Token+"1")

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)
	response := record.Result()
	suite.Equal(http.StatusUnauthorized, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 401)
	suite.Equal(responseBody["status"].(string), "UNAUTHORIZED")
}

func (suite *todoControllerSuite) TestDeleteTodo_Positive() {
	user := suite.addOneTodo()
	router := suite.SetupTest()
	userId := strconv.Itoa(int(user.Userid))

	url := fmt.Sprintf("/api/todo?userId=%s&id=%s", userId, "1")
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	request.Header.Set("authorization", user.Token)

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)
	response := record.Result()
	suite.Equal(http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 200)
	suite.Equal(responseBody["status"].(string), "SUCCESS")
}

func (suite *todoControllerSuite) TestRandomTodo_Positive() {
	user := suite.addManyTodo()
	router := suite.SetupTest()
	userId := strconv.Itoa(int(user.Userid))

	url := fmt.Sprintf("/api/todo/random/" + userId)
	request := httptest.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("authorization", user.Token)

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)
	response := record.Result()
	suite.Equal(http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 200)
	suite.Equal(responseBody["status"].(string), "SUCCESS")
}

func (suite *todoControllerSuite) TestGetTodoFilter_Positive() {
	user := suite.addManyTodo()
	router := suite.SetupTest()
	userId := strconv.Itoa(int(user.Userid))

	url := fmt.Sprintf("/api/todos?userId=%s&limit=%s&offset=%s", userId, "25", "10")
	request := httptest.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("authorization", user.Token)

	record := httptest.NewRecorder()
	router.ServeHTTP(record, request)
	response := record.Result()
	suite.Equal(http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	suite.Equal(int(responseBody["code"].(float64)), 200)
	suite.Equal(responseBody["status"].(string), "SUCCESS")
}

func TestTodoController(t *testing.T) {
	suite.Run(t, new(todoControllerSuite))
}
