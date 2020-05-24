package http

import "github.com/labstack/echo/v4"

func MenuRoute(e *echo.Echo) {
	controller := NewMenuController()
	e.POST("/GetTopMenu", controller.GetDetailMenu)
	e.POST("/GetDetailFood", controller.GetDetailFood)
	e.POST("/GetPointMenu", controller.GetPointMenu)
	e.POST("/GetCategoryMenu", controller.GetCategoryMenu)
	e.POST("/Search", controller.Search)

}
