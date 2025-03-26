package controllers_test

import (
	//"encoding/json"
	// "net/http"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	//depRoutes "resedist/internal/modules/department/department/routes"

	"resedist/pkg/config"
	"resedist/pkg/database"
	"resedist/pkg/redis"

	Authctl "resedist/internal/modules/auth/controllers"
	depRoutes "resedist/internal/modules/department/department/routes"

	//"strings"
	"testing"

	//"github.com/gin-gonic/gin"
	// "resedist/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestListDepartments(t *testing.T) {

	config.Set("./../../../../../config", "config")
	database.Connect()
	redis.Connect()

	form := url.Values{}
	form.Add("page", "1")
	form.Add("page_size", "5")
	form.Add("sort", "id")
	form.Add("order", "asc")
	formBody := form.Encode()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/department", strings.NewReader(formBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	token, _ := Authctl.GenerateAccessToken(1)
	b := "Bearer " + token
	req.Header.Set("Authorization", b)

	router := gin.Default()

	depRoutes.Routes(router)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Expected status 200, got %d", w.Code)
	assert.Contains(t, w.Body.String(), `pagination`)
	assert.Contains(t, w.Body.String(), `"data":[{"id":`)

	fmt.Println("Request Body:", req.Header)
	fmt.Println("Request Body:", formBody)
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
