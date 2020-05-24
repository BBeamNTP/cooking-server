package http

import (
	"bitbucket.org/BBeamnantapong/cooking-server/core"
	"bitbucket.org/BBeamnantapong/cooking-server/models"
	"bitbucket.org/BBeamnantapong/cooking-server/profile/service"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
)

type ProfileController struct {
}

func NewProfileController() *ProfileController {
	return &ProfileController{}
}
func (*ProfileController) GetUserProfile(c echo.Context) error {
	cc := c.(core.IContext)
	ProfileSrv := service.NewUserService(cc)
	Xtoken := c.Request().Header.Get("x-token")
	ProfileSrv.GetUserProfile(Xtoken, c)
	return nil
}
func (*ProfileController) UpdateUserProfile(c echo.Context) error {
	cc := c.(core.IContext)
	postData := models.Data{}
	ProfileSrv := service.NewUserService(cc)
	Xtoken := c.Request().Header.Get("x-token")
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	json.Unmarshal(b, &postData)
	ProfileSrv.UpdateUserProfile(Xtoken, &postData, c)
	return nil
}
func (*ProfileController) UploadAvatar(c echo.Context) error {
	cc := c.(core.IContext)
	UploadSrv := service.NewUserService(cc)
	Xtoken := c.Request().Header.Get("x-token")
	UploadSrv.Upload(Xtoken, c)
	return nil
}
