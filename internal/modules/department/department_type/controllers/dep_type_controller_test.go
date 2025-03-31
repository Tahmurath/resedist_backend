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
	DepTypeRequest "resedist/internal/modules/department/department_type/requests/deptype"

	DepTypeModels "resedist/internal/modules/department/department_type/models"
	depTypeResponse "resedist/internal/modules/department/department_type/responses"
	depTypeRoutes "resedist/internal/modules/department/department_type/routes"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var bearer string
var router *gin.Engine

var title, newtitle, idstr string
var deptypeid, parentid uint

func init() {
	title = "TestDepartment type"
	newtitle = "TestDepartment type Updated"
	deptypeid = 1
	parentid = 1

	config.Set("./../../../../../config", "config")
	database.Connect()
	redis.Connect()
	token, _ := Authctl.GenerateAccessToken(1)
	bearer = "Bearer " + token
	router = gin.Default()
	depTypeRoutes.Routes(router)

	database.DB.Raw("DELETE FROM department_types where title = ?", newtitle).Scan(&DepTypeModels.DepartmentType{})
}
func TestStore(t *testing.T) {
	var response depTypeResponse.DepTypeResponse
	request := &DepTypeRequest.AddDepTypeRequest{
		Title:    title,
		IsActive: true,
	}
	requestJson, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/department-type", bytes.NewBuffer(requestJson))
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

	request := &DepTypeRequest.ShowDepTypeRequest{
		Expand: false,
	}
	requestJson, _ := json.Marshal(request)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/v1/department-type/"+idstr, bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", bearer)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200, got %d", w.Code)
	assert.Contains(t, w.Body.String(), `data":{"id":`+idstr)
}

func TestSearch(t *testing.T) {

	request := &DepTypeRequest.ListDepTypeRequest{
		Page:     1,
		PageSize: 1,
		Sort:     "id",
		Order:    "asc",
	}
	requestJson, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/department-type", bytes.NewBuffer(requestJson))
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

	request := &DepTypeRequest.EditDepTypeRequest{
		Title:    newtitle,
		IsActive: false,
	}
	requestJson, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/department-type/"+idstr, bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", bearer)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status 200, got %d", w.Code)
	assert.Contains(t, w.Body.String(), `data":{"id":`+idstr)
	assert.Contains(t, w.Body.String(), `title":"`+newtitle)
	assert.Contains(t, w.Body.String(), `is_active":false`)

	fmt.Println("Request Body:", req.Body)
	fmt.Println("Response Body:", w.Body.String())
}

func TestRemove(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/department-type/"+idstr, nil)
	req.Header.Set("Content-type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", bearer)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code, "Expected status 200, got %d", w.Code)

	fmt.Println("Request Body:", req.Body)
	fmt.Println("Response Body:", w.Body.String())
}
