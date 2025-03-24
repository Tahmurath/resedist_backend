package controllers_test

import (
	//"encoding/json"
	// "net/http"
	"net/http"
	"net/http/httptest"

	depRoutes "resedist/internal/modules/department/department/routes"

	//"strings"
	"testing"

	//"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestShow(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://127.0.0.1:4000/api/v1/department/1", nil)

	//router := routing.GetRouter()
	router := gin.Default()
	depRoutes.Routes(router)

	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, `{"error":"Authorization header missing"}`, w.Body.String())
}
