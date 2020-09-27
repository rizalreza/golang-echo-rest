package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/rizalreza/golang-echo-rest/models"
)

func GetAllProduct(c echo.Context) error {
	result, err := models.GetAllProduct()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func GetProductById(c echo.Context) error {
	id := c.Param("id")
	convID, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	result, err := models.GetProductById(convID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateProduct(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	category := c.FormValue("category")
	price := c.FormValue("price")

	convID, _ := strconv.Atoi(id)
	convPrice, _ := strconv.Atoi(price)

	result, err := models.UpdatePoduct(convID, convPrice, name, category)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func DeleteProduct(c echo.Context) error {
	id := c.FormValue("id")
	convID, _ := strconv.Atoi(id)

	result, err := models.DeleteProduct(convID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
