package controller

import (
	"kulina-interview-test/domain/model"
	"kulina-interview-test/domain/services"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type SupplierController struct {
	supplierService services.SupplierService
}

func (controller SupplierController) CreateStore(c echo.Context) error {
	result := controller.supplierService.CreateStore(model.Supplier{})
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "store created successfully",
		"data":    result,
	})
}

func (controller SupplierController) CreateProduct(c echo.Context) error {
	result := controller.supplierService.CreateProduct(model.Product{})
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "product created successfully",
		"data":    result,
	})
}

func NewSupplierController(db *gorm.DB) SupplierController {
	service := services.NewSupplierService(db)
	controller := SupplierController{
		supplierService: service,
	}

	return controller
}
