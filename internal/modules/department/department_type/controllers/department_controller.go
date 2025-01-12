package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	//articleRepository "resedist/internal/modules/article/repositories"

	DepTypeService "resedist/internal/modules/department/department_type/services"
)

type Controller struct {
	departmentTypeService DepTypeService.DepartmentTypeServiceInterface
}

func New() *Controller {

	return &Controller{
		departmentTypeService: DepTypeService.New(),
	}
}

func (controller *Controller) Search(c *gin.Context) {

	depTypeTitle := c.DefaultQuery("query", "")
	depTypes := controller.departmentTypeService.Search(depTypeTitle)

	c.JSON(http.StatusOK, depTypes.Data)
}

// func (controller *Controller) Search2(c *gin.Context) {
// 	var results []DepTypeModels.DepartmentType

// 	firstname := c.DefaultQuery("query", "")
// 	dep := DepTypeRepo.New()

// 	dep.DB.Limit(10).Where("title like ?", "%"+firstname+"%").Find(&results)
// 	// دریافت داده‌ها از دیتابیس

// 	// ارسال داده‌ها به صورت مستقیم بدون استفاده از کلید "Data"
// 	c.JSON(200, results)
// }
