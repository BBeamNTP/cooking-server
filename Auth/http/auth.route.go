package http

import "github.com/labstack/echo/v4"

func AuthRoute(e *echo.Echo) {
	controller := NewAuthController()
	e.POST("/SignUp", controller.CreateUser)
	e.POST("/SignIn", controller.CheckUser)
	e.POST("/SignOut", controller.Logout)
	e.POST("/ForgetPassword/SendOTP", controller.SendOTP)
	e.POST("/ForgetPassword/ResetPassword", controller.ResetPassword)

}
