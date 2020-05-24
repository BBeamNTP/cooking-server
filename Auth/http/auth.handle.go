package http

import (
	"bitbucket.org/BBeamnantapong/cooking-server/Auth/service"
	"bitbucket.org/BBeamnantapong/cooking-server/core"
	"bitbucket.org/BBeamnantapong/cooking-server/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}
func (*AuthController) CreateUser(c echo.Context) error {
	cc := c.(core.IContext)
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	AuthSrv := service.NewAuthrService(cc)
	postAuth := models.Userdata{}
	Xtoken := c.Request().Header.Get("x-token")

	json.Unmarshal(b, &postAuth)
	log.Println("PostAuth :", postAuth)
	AuthSrv.CreateUserSrv(&postAuth, c, Xtoken)
	return nil
}
func (*AuthController) CheckUser(c echo.Context) error {
	cc := c.(core.IContext)
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	postAuth := models.Userdata{}
	json.Unmarshal(b, &postAuth)
	AuthSrv := service.NewAuthrService(cc)
	AuthSrv.CheckUserSrv(&postAuth, c)
	return nil
}
func (*AuthController) Logout(c echo.Context) error {
	cc := c.(core.IContext)
	Xtoken := c.Request().Header.Get("x-token")
	AuthSrv := service.NewAuthrService(cc)
	AuthSrv.LogoutSrv(Xtoken, c)
	return nil
}
func (*AuthController) ResetPassword(c echo.Context) error {
	cc := c.(core.IContext)
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	postAuth := models.ResetPassword{}
	json.Unmarshal(b, &postAuth)
	AuthSrv := service.NewAuthrService(cc)
	AuthSrv.ResetPasswordSrv(&postAuth, c)
	return nil
}
func (*AuthController) SendOTP(c echo.Context) error {
	cc := c.(core.IContext)
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	postAuth := models.Userdata{}
	json.Unmarshal(b, &postAuth)
	AuthSrv := service.NewAuthrService(cc)
	AuthSrv.SendOTPSrv(&postAuth, c)
	return nil
}
