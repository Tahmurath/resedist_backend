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

	userModels "resedist/internal/modules/user/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var name, email, password string
var router *gin.Engine

func init() {
	name = "Test Test"
	email = "test@test.test"
	password = "TestSecretTest"

	config.Set("./../../../../config", "config")
	database.Connect()
	redis.Connect()
	router = gin.Default()
	authRoutes.Routes(router)

	database.DB.Raw("DELETE FROM users where email = ?", email).Scan(&userModels.User{})
}

func TestRegister(t *testing.T) {

	form := url.Values{}
	form.Add("name", name)
	form.Add("email", email)
	form.Add("password", password)
	formBody := form.Encode()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(formBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Expected status 200, got %d", w.Code)
	assert.Contains(t, w.Body.String(), `User registered successfully`)

	fmt.Println("Request Body:", formBody)
	fmt.Println("Response Body:", w.Body.String())
}

func TestLoginForm(t *testing.T) {

	form := url.Values{}
	form.Add("email", email)
	form.Add("password", password)
	formBody := form.Encode()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/auth/login", strings.NewReader(formBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Expected status 200, got %d", w.Code)
	assert.Contains(t, w.Body.String(), `access_token`)

	fmt.Println("Request Body:", formBody)
	fmt.Println("Response Body:", w.Body.String())
}

func TestLoginJson(t *testing.T) {

	request := &auth.LoginRequest{
		Email:    email,
		Password: password,
	}
	requestJson, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json; charset=UTF-8")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Expected status 200, got %d", w.Code)
	assert.Contains(t, w.Body.String(), `access_token`)

	fmt.Println("Request Body:", req.Body)
	fmt.Println("Response Body:", w.Body.String())
}
