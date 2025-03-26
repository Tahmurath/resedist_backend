package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	//depRoutes "resedist/internal/modules/department/department/routes"

	authRoutes "resedist/internal/modules/auth/routes"

	userModels "resedist/internal/modules/user/models"
	// "resedist/pkg/database"

	//"strings"
	"testing"

	//"github.com/gin-gonic/gin"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	// articleModels "resedist/internal/modules/article/models"
	// contactModels "resedist/internal/modules/contact/models"
	// departmentModels "resedist/internal/modules/department/department/models"
	// orderModels "resedist/internal/modules/order/models"
	// tenantModels "resedist/internal/modules/tenant/models"
	// userModels "resedist/internal/modules/user/models"
	// "gorm.io/driver/sqlite"

	"resedist/pkg/config"
	"resedist/pkg/database"
	"resedist/pkg/redis"
)

var user userModels.User

func TestLogin(t *testing.T) {

	config.Set("./../../../../config")
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

	// fmt.Println("Request Body:", formBody)
	// fmt.Println("Response Body:", w.Body.String())
}
