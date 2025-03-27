package controllers_test

import (
	//"encoding/json"
	// "net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	//depRoutes "resedist/internal/modules/department/department/routes"

	"resedist/pkg/config"
	"resedist/pkg/database"
	"resedist/pkg/redis"

	Authctl "resedist/internal/modules/auth/controllers"
	DepRequest "resedist/internal/modules/department/department/requests/department"
	depRoutes "resedist/internal/modules/department/department/routes"

	//"strings"
	"testing"

	//"github.com/gin-gonic/gin"
	// "resedist/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var bearer string
var router *gin.Engine

func init() {
	config.Set("./../../../../../config", "config")
	database.Connect()
	redis.Connect()
	token, _ := Authctl.GenerateAccessToken(1)
	bearer = "Bearer " + token
	router = gin.Default()
	depRoutes.Routes(router)
}

func TestShow(t *testing.T) {

	request := &DepRequest.ShowDepartmentRequest{
		Expand: false,
	}
	requestJson, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/department/1", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", bearer)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200, got %d", w.Code)
	assert.Contains(t, w.Body.String(), `data":{"id":1`)
}

func TestSearch(t *testing.T) {

	request := &DepRequest.ListDepartmentRequest{
		Page:     1,
		PageSize: 1,
		Sort:     "id",
		Order:    "asc",
	}
	requestJson, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/department", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", bearer)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200, got %d", w.Code)
	assert.Contains(t, w.Body.String(), `pagination`)
	assert.Contains(t, w.Body.String(), `"data":[{"id":`)

	fmt.Println("Request Body:", req.Header)
	fmt.Println("Request Body:", req.Body)
	fmt.Println("Response Body:", w.Body.String())
}

func TestStore(t *testing.T) {

	request := &DepRequest.AddDepartmentRequest{
		Title:            "depone",
		DepartmentTypeId: 1,
		ParentID:         1,
	}
	requestJson, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/department", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", bearer)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Expected status 200, got %d", w.Code)
	assert.Contains(t, w.Body.String(), `data":{"id"`)

	fmt.Println("Request Body:", req.Body)
	fmt.Println("Response Body:", w.Body.String())
}

// func TestGenerateAccessToken(t *testing.T) {
// 	token, err := generateAccessToken(1)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, token)

// 	claims := &Claims{}
// 	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	assert.NoError(t, err)
// 	assert.True(t, parsedToken.Valid)
// 	assert.Equal(t, uint(1), claims.UserID)
// 	assert.Equal(t, "access", claims.Type)
// 	assert.WithinDuration(t, time.Now().Add(15*time.Minute), claims.ExpiresAt.Time, time.Second)
// }
