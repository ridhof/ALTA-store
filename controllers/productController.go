package controllers

import (
	"net/http"
	"project/lib/database"
	"project/models"
	"github.com/labstack/echo"
)

func GetProductsControllers(c echo.Context) error {
	kategori := c.QueryParam("kategori")
	prods, e := database.GetProducts(kategori)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  prods,
	})
}

func GetProductControllers(c echo.Context) error {
	id := c.Param("id")
	prods, e := database.GetProduct(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  prods,
	})
}

func CreateProductsController(c echo.Context) error {
	prod := models.Products{}
	c.Bind(&prod)
	prods, e := database.CreateProducts(&prod)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  prods,
	})
}