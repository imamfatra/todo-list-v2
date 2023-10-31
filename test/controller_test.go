package test_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegistrasiSuccess(t *testing.T) {
	delTable(testDB)

	requestBody := strings.NewReader((`
		"email": "aku@mail.com",
		"username": "andini",
		"password": "Atlanti322k"`))
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/user/registrasi", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	hendler.ServeHTTP(recorder, request)

	response := recorder.Result()
	// assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)
	// assert.Equal(t, 200, int(responseBody["code"].(float64))) // konversi nilai code ke float64 kemudian baru ke int
	// assert.Equal(t, "OK", responseBody["status"])
	// assert.Equal(t, "agnes", responseBody["data"].(map[string]interface{})["name"])
}
