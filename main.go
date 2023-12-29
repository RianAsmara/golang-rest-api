package main

import (
	"kulina-interview-test/config"
	"kulina-interview-test/domain/controller"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	db := config.InitDB()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})
	supplierController := controller.NewSupplierController(db)

	// userRoute := e.Group("/user")
	// userRoute.POST("/register", controller.UserRegister)

	supplierRoute := e.Group("/supplier")
	supplierRoute.POST("/store/create", supplierController.CreateStore)
	supplierRoute.POST("/product/create", supplierController.CreateProduct)

	algoRoute := e.Group("/algorithm")
	algoRoute.POST("/sock-merchant", controller.PairColor)
	algoRoute.POST("/counting-valleys", controller.CountValley)
	algoRoute.POST("/process-number", controller.ProcessNumber)
	algoRoute.POST("/count-lamps", controller.CountLamps)

	e.Logger.Fatal(e.Start(":8080"))
}
