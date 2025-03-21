package controllers

import (
	"log"
	"net/http"
	"resedist/internal/modules/user/requests/auth"
	UserResponse "resedist/internal/modules/user/responses"
	UserService "resedist/internal/modules/user/services"
	"resedist/pkg/applog"
	"resedist/pkg/config"
	"resedist/pkg/errors"
	"resedist/pkg/rest"

	configStruct "resedist/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	// UserService "resedist/internal/modules/user/services"
	"time"
)

type Controller struct {
	userService UserService.UserServiceInterface
	errFmt      *errors.ErrorFormat
	json        *rest.Jsonresponse
	rest        configStruct.Rest
}

func New() *Controller {

	return &Controller{
		userService: UserService.New(),
		rest:        config.Get().Rest,
		errFmt:      errors.New(),
		json:        rest.New(),
	}
}

func (controller *Controller) Login(c *gin.Context) {
	//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	c.JSON(http.StatusOK, gin.H{
		"sdfsdfs": "sdfsdfs",
		"sdsda":   "sdfsdfs",
		"sdfsd":   "sdfsdfs",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	var request auth.RegisterRequest

	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromError(err)

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Opps, there is an error with bind",
			"errors":  errors.Get(),
		})
		return
	}

	if controller.userService.CheckUserExist(request.Email) {
		errors.Init()
		errors.Add("Email", "Email address already exists")

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Opps, there is an error with email",
			"errors":  errors.Get(),
		})
		return
	}

	// Create the user
	user, err := controller.userService.Create(request)

	// Check if there is any error on the user creation
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Opps, there is an error with user creation",
		})
		return
	}

	log.Printf("The user created successfully with a name %s \n", user.Name)
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

// @Summary Login and get JWT token
// @Description Authenticate user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param user query auth.LoginRequest true "User data"
// @Success 200 {object} map[string]string "Token"
// @Router /api/v1/auth/login [post]
func (ctl *Controller) HandleLogin(c *gin.Context) {
	var request auth.LoginRequest
	cfg := config.Get().Rest

	if err := c.ShouldBind(&request); err != nil {
		ctl.json.Badrequest(c, rest.RestConfig{
			Error_message: ctl.errFmt.SetFromError(err),
		})
		return
	}

	user, err := ctl.userService.HandleUserLogin(request)
	if err != nil {
		ctl.json.NotFound(c, rest.RestConfig{
			Error_message: err.Error(),
		})
		return
	}

	token, err := createJwt(user)
	if err != nil {
		ctl.json.NotFound(c, rest.RestConfig{
			Error_message: err.Error(),
			Http:          http.StatusInternalServerError,
		})
		return
	}

	applog.Info("The user logged in successfully")

	c.JSON(http.StatusOK, gin.H{
		//"message": "User logged in successfully",
		cfg.Status:        "success",
		cfg.Error_message: "User logged in successfully", //errors.Get(),
		cfg.Error_code:    "",
		//"user":            user,
		"token": token,
		"user": map[string]interface{}{
			"info":  user,
			"token": token,
		},
	})
}

func (ctl *Controller) User(c *gin.Context) {
	user, _ := c.Get("auth")
	c.JSON(http.StatusOK, gin.H{"message": "Authenticated", "user": user})
}

func createJwt(user UserResponse.User) (string, error) {
	cfg := config.Get()

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user,
		"ExpireAt": time.Now().Add(cfg.Jwt.Duration).Unix(), //jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		"IssuedAt": jwt.NewNumericDate(time.Now()),
	})

	// Sign the claim with a secret key
	token, err := claims.SignedString([]byte(cfg.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
