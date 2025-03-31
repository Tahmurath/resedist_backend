package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"resedist/pkg/config"
	"resedist/pkg/database"
	"resedist/pkg/redis"

	Authctl "resedist/internal/modules/auth/controllers"
	DepRequest "resedist/internal/modules/department/department/requests/department"

	DepartmentModels "resedist/internal/modules/department/department/models"
	depResponse "resedist/internal/modules/department/department/responses"
	depRoutes "resedist/internal/modules/department/department/routes"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var bearer string
var router *gin.Engine

var title, newtitle, idstr string
var deptypeid, parentid uint

func init() {
	title = "TestDepartment"
	newtitle = "TestDepartment Updated"
	deptypeid = 1
	parentid = 1

	config.Set("./../../../../../config", "config")
	database.Connect()
	redis.Connect()
	token, _ := Authctl.GenerateAccessToken(1)
	bearer = "Bearer " + token
	router = gin.Default()
	depRoutes.Routes(router)

	database.DB.Raw("DELETE FROM departments where title = ?", newtitle).Scan(&DepartmentModels.Department{})
}
func TestStore(t *testing.T) {
	var response depResponse.DepartmentResponse
	request := &DepRequest.AddDepartmentRequest{
		Title:            title,
		DepartmentTypeId: deptypeid,
		ParentID:         parentid,
	}
	requestJson, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/department", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", bearer)

	router.ServeHTTP(w, req)

	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error unmarshaling JSON: %v", err)
	}

	idstr = fmt.Sprintf("%d", response.Data.ID)

	assert.Equal(t, http.StatusCreated, w.Code, "Expected status 200, got %d", w.Code)
	assert.Contains(t, w.Body.String(), `data":{"id":`+idstr)

	fmt.Println("Request Body:", req.Body)
	fmt.Println("Response Body:", w.Body.String())
}

func TestShow(t *testing.T) {

	request := &DepRequest.ShowDepartmentRequest{
		Expand: false,
	}
	requestJson, _ := json.Marshal(request)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/v1/department/"+idstr, bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", bearer)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200, got %d", w.Code)
	assert.Contains(t, w.Body.String(), `data":{"id":`+idstr)
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

func TestUpdate(t *testing.T) {

	request := &DepRequest.EditDepartmentRequest{
		Title:            newtitle,
		DepartmentTypeId: deptypeid,
		ParentID:         parentid,
	}
	requestJson, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/department/"+idstr, bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", bearer)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200, got %d", w.Code)
	assert.Contains(t, w.Body.String(), `data":{"id":`+idstr)
	assert.Contains(t, w.Body.String(), `title":"`+newtitle)

	fmt.Println("Request Body:", req.Body)
	fmt.Println("Response Body:", w.Body.String())
}

func TestRemove(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/department/"+idstr, nil)
	req.Header.Set("Content-type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", bearer)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code, "Expected status 200, got %d", w.Code)

	fmt.Println("Request Body:", req.Body)
	fmt.Println("Response Body:", w.Body.String())
}
