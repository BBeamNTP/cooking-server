package Auth

import (
	"bitbucket.org/BBeamnantapong/cooking-server/models"
	"github.com/labstack/echo/v4"
)

type AuthInterface interface {
	CreateUserSrv(UserData *models.Userdata, c echo.Context, Xtoken string) interface{}
	CheckUserSrv(UserData *models.Userdata, c echo.Context) interface{}
	LogoutSrv(Xtoken string, c echo.Context) interface{}
	SendOTPSrv(postAuth *models.Userdata, c echo.Context) interface{}
	ResetPasswordSrv(postAuth *models.ResetPassword, c echo.Context) interface{}
}
