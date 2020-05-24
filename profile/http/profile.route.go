package http

import "github.com/labstack/echo/v4"

func ProfileRoute(e *echo.Echo) {
	controller := NewProfileController()
	e.GET("/Profile", controller.GetUserProfile)
	e.POST("/Profile/EditProfile", controller.UpdateUserProfile)
	e.POST("/Profile/UploadAvatar", controller.UploadAvatar)
}
