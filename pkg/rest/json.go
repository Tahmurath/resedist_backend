package rest

import (
	"net/http"
	"resedist/pkg/config"

	configStruct "resedist/config"

	"resedist/pkg/pagination"

	"github.com/gin-gonic/gin"
)

type Jsonresponse struct {
	Status        string
	Error_message string
	Error_code    string
	Pagination    string
	Data          string
	rest          configStruct.Rest
	//errFmt *errors.ErrorFormat
}

type RestConfig struct {
	Status        string
	Error_message interface{}
	Error_code    string
	Data          interface{}
	Pagination    pagination.PagePack
	Paged         bool
	NoContent     bool
	Http          int
	// AccessToken string
	// RefreshToken string
}

func New() *Jsonresponse {

	rest := config.Get().Rest

	return &Jsonresponse{

		Status:        rest.Status,
		Error_message: rest.Error_message,
		Error_code:    rest.Error_code,
		Pagination:    rest.Pagination,
		Data:          rest.Data,
		rest:          rest,
		//errFmt:        errors.New(),
	}
}

func (j *Jsonresponse) Badrequest(c *gin.Context, config RestConfig) {

	if config.Status == "" {
		config.Status = j.rest.Failed
	}

	if config.Error_code == "" {
		config.Error_code = j.rest.Bind_error
	}
	if config.Http < 1 {
		config.Http = http.StatusBadRequest
	}

	c.JSON(config.Http, gin.H{
		j.Status:        config.Status,
		j.Error_message: config.Error_message,
		j.Error_code:    config.Error_code,
	})
}

func (j *Jsonresponse) NotFound(c *gin.Context, config RestConfig) {

	if config.Status == "" {
		config.Status = j.rest.Failed
	}

	if config.Error_code == "" {
		config.Error_code = j.rest.Not_found
	}
	if config.Http < 1 {
		config.Http = http.StatusNotFound
	}

	c.JSON(config.Http, gin.H{
		j.Status:        config.Status,
		j.Error_message: config.Error_message,
		j.Error_code:    config.Error_code,
	})
}

func (j *Jsonresponse) ServerError(c *gin.Context, config RestConfig) {

	if config.Status == "" {
		config.Status = j.rest.Failed
	}

	if config.Error_code == "" {
		config.Error_code = j.rest.Not_found
	}
	if config.Http < 1 {
		config.Http = http.StatusInternalServerError
	}

	c.JSON(config.Http, gin.H{
		j.Status:        config.Status,
		j.Error_message: config.Error_message,
		j.Error_code:    config.Error_code,
	})
}

func (j *Jsonresponse) Success(c *gin.Context, config RestConfig) {

	if config.Status == "" {
		config.Status = j.rest.Success
	}
	if config.Http < 1 {
		config.Http = http.StatusOK
	}

	res := gin.H{
		j.Status:        config.Status,
		j.Error_message: config.Error_message,
		j.Error_code:    config.Error_code,
	}

	if config.Paged {
		res[j.Pagination] = config.Pagination
	}

	if !config.NoContent {
		res[j.Data] = config.Data
	}

	c.JSON(config.Http, res)
}
