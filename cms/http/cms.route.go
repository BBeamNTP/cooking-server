package http

import "github.com/labstack/echo/v4"

func CMSRoute(e *echo.Echo) {
	controller := NewCMSController()
	e.GET("/CMS/GetIngredients", controller.GetIngredients)
	e.POST("/CMS/CreateMenu", controller.CMSCreateMenu)
	e.POST("/CMS/DeleteMenu", controller.CMSDeleteMenu)
	e.POST("/CMS/GetDetailUpdateMenu", controller.CMSGetDetailUpdateMenu)
	e.POST("/CMS/Upload", controller.UploadFile)
	e.POST("/CMS/ManageMenu", controller.CMSManageMenu)


}
