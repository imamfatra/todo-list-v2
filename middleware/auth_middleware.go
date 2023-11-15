package middleware

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"todo-api/exception"
	"todo-api/helper"
	"todo-api/model"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/user/registrasi" || r.URL.Path == "/api/user/login" {
			next.ServeHTTP(w, r)
			return
		}

		config, err := model.LoadConfig("./")
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			webResponse := model.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Data:   err,
			}
			helper.WriteToResponse(w, webResponse)
			return
		}

		tokenString := r.Header.Get("authorization")
		if tokenString == "" {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			webResponse := model.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Interneal Server Error",
				Data:   err,
			}
			helper.WriteToResponse(w, webResponse)
			return
		}

		maker, err := NewPasetoMaker(config.TokenSymmetricKey)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			webResponse := model.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Interneal Server Error",
				Data:   err,
			}
			helper.WriteToResponse(w, webResponse)
			return
		}

		var userId int
		if r.Method == http.MethodGet || r.Method == http.MethodDelete || r.Method == http.MethodPut {
			var err error
			if len(strings.Split(r.URL.Path, "/")) >= 3 {
				path := strings.Split(r.URL.Path, "/")
				userId, err = strconv.Atoi(path[len(path)-1])
				if err != nil {
					panic(exception.NewNotFoundError(err.Error()))
				}
			} else {
				userIdStr := r.URL.Query().Get("userId")
				userId, err = strconv.Atoi(userIdStr)
				if err != nil {
					panic(exception.NewNotFoundError(err.Error()))
				}
			}

		} else {

			body, err := io.ReadAll(r.Body)
			helper.IfError(err)
			var requestBody map[string]interface{}
			err = json.Unmarshal(body, &requestBody)
			helper.IfError(err)

			todo, _ := requestBody["todo"].(string)
			complated, _ := requestBody["todo"].(bool)
			userId = int(requestBody["userid"].(float64))

			model.AddTodoBuffer = model.AddNewTodoRequest{
				Todo:      todo,
				Complated: complated,
				Userid:    int32(userId),
			}
		}

		_, err = maker.VarifyToken(tokenString, userId)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			webResponse := model.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Data:   err,
			}
			helper.WriteToResponse(w, webResponse)
			return
		}

		next.ServeHTTP(w, r)

	})
}
