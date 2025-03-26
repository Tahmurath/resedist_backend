package controllers_test

import (
	//"encoding/json"
	// "net/http"
	"net/http"
	"net/http/httptest"

	//depRoutes "resedist/internal/modules/department/department/routes"
	appRoutes "resedist/internal/providers/routes"

	//"strings"
	"testing"

	//"github.com/gin-gonic/gin"
	// "resedist/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// func Token(){
// 	GenerateAccessToken()
// }

func TestShow(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/department/1", nil)

	router := gin.Default()

	// viper.SetConfigName("config") // name of config file (without extension)
	// viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	// viper.AddConfigPath("./../../../../config")

	//config.Set("./../../../../config")

	// database.Connect()
	appRoutes.RegisterRoutes(router)

	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, `{"error":"Authorization header missing"}`, w.Body.String())
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
