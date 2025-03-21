package controllers

import (
	"net/http"
	DepScopes "resedist/internal/modules/department/department/scopes"

	//DepScopes "resedist/internal/modules/department/department/scopes"
	DepTypeRequest "resedist/internal/modules/department/department_type/requests/department_type"
	DepTypeScopes "resedist/internal/modules/department/department_type/scopes"
	"resedist/pkg/config"
	"resedist/pkg/pagination"

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

// @Summary Get Department types
// @Description Returns a list of Deaprtment types
// @Tags department-type
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/v1/department-type/ [get]
func (controller *Controller) Search(c *gin.Context) {
	var request DepTypeRequest.ListDepartmentTypeRequest

	cfg := config.Get().Rest

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"request":         request,
			cfg.Status:        "failed",
			cfg.Error_message: "Opps, there is an error with Query bind",
			cfg.Error_code:    "",
		})
		return
	}

	page := pagination.NewPagePack(request.Page, request.PageSize)

	depTypes := controller.departmentTypeService.SearchScope(
		page,
		DepTypeScopes.TitleLike(request.Title),
		DepScopes.IdsOr(request.DepartmentType),
		DepTypeScopes.Sort(request.Sort, request.Order),
	)

	c.JSON(http.StatusOK, gin.H{
		cfg.Status:        "",
		cfg.Error_message: "",
		cfg.Error_code:    "",
		cfg.Pagination:    page,
		cfg.Data:          depTypes.Data,
	})
}

//func (controller *Controller) Search(c *gin.Context) {
//
//	depTypeTitle := c.DefaultQuery("query", "")
//	depTypes := controller.departmentTypeService.Search(depTypeTitle)
//
//	c.JSON(http.StatusOK, depTypes.Data)
//}

// func (controller *Controller) Search2(c *gin.Context) {
// 	var results []DepTypeModels.DepartmentType

// 	firstname := c.DefaultQuery("query", "")
// 	dep := DepTypeRepo.New()

// 	dep.DB.Limit(10).Where("title like ?", "%"+firstname+"%").Find(&results)
// 	// دریافت داده‌ها از دیتابیس

// 	// ارسال داده‌ها به صورت مستقیم بدون استفاده از کلید "Data"
// 	c.JSON(200, results)
// }
