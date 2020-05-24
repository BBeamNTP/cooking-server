package http

import (
	"bitbucket.org/BBeamnantapong/cooking-server/cms/service"
	"bitbucket.org/BBeamnantapong/cooking-server/core"
	"bitbucket.org/BBeamnantapong/cooking-server/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
)

type CMSController struct {
}

func NewCMSController() *CMSController {
	return &CMSController{}
}
func (*CMSController) GetIngredients(c echo.Context) error {
	cc := c.(core.IContext)
	CMSSrv := service.NewCMSService(cc)
	Xtoken := c.Request().Header.Get("x-token")
	CMSSrv.GetIngredients(Xtoken, c)
	return nil
}

func (*CMSController) CMSCreateMenu(c echo.Context) error {
	cc := c.(core.IContext)
	CMSSrv := service.NewCMSService(cc)
	postMenu := models.DataIngredients{}
	Xtoken := c.Request().Header.Get("x-token")
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	json.Unmarshal(b, &postMenu)
	CMSSrv.CMSCreateMenu(Xtoken, c, &postMenu)
	return nil
}
func (*CMSController) CMSDeleteMenu(c echo.Context) error {
	cc := c.(core.IContext)
	CMSSrv := service.NewCMSService(cc)
	postMenu := models.DataIngredients{}
	Xtoken := c.Request().Header.Get("x-token")
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	json.Unmarshal(b, &postMenu)
	CMSSrv.CMSDeleteMenu(Xtoken, c, &postMenu)
	return nil
}
func (*CMSController) CMSGetDetailUpdateMenu(c echo.Context) error {
	cc := c.(core.IContext)
	CMSSrv := service.NewCMSService(cc)
	postMenu := models.DataIngredients{}
	Xtoken := c.Request().Header.Get("x-token")
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	json.Unmarshal(b, &postMenu)
	CMSSrv.CMSGetDetailUpdateMenu(Xtoken, c, &postMenu)

	return nil
}

func (*CMSController) CMSManageMenu(c echo.Context) error {
	cc := c.(core.IContext)
	CMSSrv := service.NewCMSService(cc)
	postMenu := models.Menu{}

	Xtoken := c.Request().Header.Get("x-token")
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	json.Unmarshal(b, &postMenu)
	CMSSrv.CMSManageMenu(Xtoken, c, &postMenu)

	return nil

}
func (*CMSController) UploadFile(c echo.Context) error {
	cc := c.(core.IContext)
	CMSSrv := service.NewCMSService(cc)
	Xtoken := c.Request().Header.Get("x-token")
	log.Println("=====================1")
	CMSSrv.UploadFile(Xtoken, c)

	return nil
}
