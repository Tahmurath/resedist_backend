package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	authRoutes "resedist/internal/modules/auth/routes"
	"resedist/internal/modules/user/requests/auth"
	"resedist/pkg/config"
	"resedist/pkg/database"
	"resedist/pkg/redis"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	config.Set("./../../../../config", "config")
	database.Connect()
	redis.Connect()

	form := url.Values{}
	form.Add("name", "hooman3@test.com")
	form.Add("email", "hooman4@test.com")
	form.Add("password", "hooman4@test.com")
	formBody := form.Encode()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(formBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router := gin.Default()

	authRoutes.Routes(router)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Expected status 200, got %d", w.Code)
	assert.Contains(t, w.Body.String(), `User registered successfully`)

	fmt.Println("Request Body:", formBody)
	fmt.Println("Response Body:", w.Body.String())
}
func TestLoginForm(t *testing.T) {

	config.Set("./../../../../config", "config")
	database.Connect()
	redis.Connect()

	form := url.Values{}
	form.Add("email", "hooman@test.com")
	form.Add("password", "hooman@test.com")
	formBody := form.Encode()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/auth/login", strings.NewReader(formBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router := gin.Default()

	authRoutes.Routes(router)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Expected status 200, got %d", w.Code)
	assert.Contains(t, w.Body.String(), `access_token`)

	fmt.Println("Request Body:", formBody)
	fmt.Println("Response Body:", w.Body.String())
}
func TestLoginJson(t *testing.T) {

	config.Set("./../../../../config", "config")
	database.Connect()
	redis.Connect()

	request := &auth.LoginRequest{
		Email:    "hooman@test.com",
		Password: "hooman@test.com",
	}
	requestJson, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json; charset=UTF-8")

	router := gin.Default()

	authRoutes.Routes(router)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Expected status 200, got %d", w.Code)
	assert.Contains(t, w.Body.String(), `access_token`)

	fmt.Println("Request Body:", req.Body)
	fmt.Println("Response Body:", w.Body.String())
}
