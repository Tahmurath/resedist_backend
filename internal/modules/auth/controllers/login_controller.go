package controllers

import (
	"log"
	"net/http"
	"resedist/internal/modules/user/requests/auth"
	UserResponse "resedist/internal/modules/user/responses"
	UserService "resedist/internal/modules/user/services"
	"resedist/pkg/config"
	"resedist/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	// UserService "resedist/internal/modules/user/services"
	"time"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {

	return &Controller{
		userService: UserService.New(),
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

func (controller *Controller) HandleLogin(c *gin.Context) {
	var request auth.LoginRequest
	cfg := config.Get().Rest

	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromError(err)

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			//"message": "Opps, there is an error with ShouldBind",
			cfg.Status:        "failed",
			cfg.Error_message: "Opps, there is an error with ShouldBind", //errors.Get(),
			cfg.Error_code:    "",
			cfg.Data:          "",
		})
		return
	}

	user, err := controller.userService.HandleUserLogin(request)
	if err != nil {
		errors.Init()
		errors.Add("email", err.Error())

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			//"message":         "Opps, there is an error to find user",
			cfg.Status:        "failed",
			cfg.Error_message: "Opps, there is an error to find user", //errors.Get(),
			cfg.Error_code:    "",
			cfg.Data:          "",
		})
		return
	}

	token, err := createJwt(user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			//"message": "Opps, there is an error",
			cfg.Status:        "failed",
			cfg.Error_message: "Opps, there is an error", //errors.Get(),
			cfg.Error_code:    "",
			cfg.Data:          "",
		})
		return
	}

	log.Printf("The user logged in successfully with a name %s \n", user.Name)
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

func (controller *Controller) User(c *gin.Context) {
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
