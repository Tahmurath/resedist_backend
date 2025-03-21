package rest

import (
	"net/http"
	"resedist/pkg/config"

	configStruct "resedist/config"

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

type BadrequestConfig struct {
	Status        string
	Error_message interface{}
	Error_code    string
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

func (j *Jsonresponse) Badrequest(c *gin.Context, config BadrequestConfig) {

	if config.Status == "" {
		config.Status = j.rest.Failed
	}

	if config.Error_code == "" {
		config.Error_code = j.rest.Bind_error
	}

	c.JSON(http.StatusBadRequest, gin.H{
		j.Status:        config.Status,
		j.Error_message: config.Error_message,
		j.Error_code:    config.Error_code,
	})
}
