package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rizalreza/golang-echo-rest/controllers"
)

func InitialRoute() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World :)")
	})

	e.GET("/products", controllers.GetAllProduct)
	e.GET("/product/:id", controllers.GetProductById)
	e.PUT("/product", controllers.UpdateProduct)
	e.DELETE("/product", controllers.DeleteProduct)

	return e
}
