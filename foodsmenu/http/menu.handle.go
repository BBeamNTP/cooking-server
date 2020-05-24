package http

import (
	"bitbucket.org/BBeamnantapong/cooking-server/core"
	"bitbucket.org/BBeamnantapong/cooking-server/foodsmenu/service"
	"bitbucket.org/BBeamnantapong/cooking-server/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
)

type MenuController struct {
}

func NewMenuController() *MenuController {
	return &MenuController{}
}
func (*MenuController) GetDetailMenu(c echo.Context) error {
	cc := c.(core.IContext)
	MenuSrv := service.NewMenuService(cc)
	Xtoken := c.Request().Header.Get("x-token")
	MenuSrv.GetMenu(Xtoken, c)
	return nil
}

func (*MenuController) GetDetailFood(c echo.Context) error {
	cc := c.(core.IContext)
	postMenu := models.DataIngredients{}
	MenuSrv := service.NewMenuService(cc)
	Xtoken := c.Request().Header.Get("x-token")
	log.Print("Post Menu : ", postMenu)
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	json.Unmarshal(b, &postMenu)
	MenuSrv.GetFood(Xtoken, c, &postMenu)

	log.Print("PostMenu : = ", postMenu)
	return nil
}

func (*MenuController) GetPointMenu(c echo.Context) error {
	cc := c.(core.IContext)
	postMenu := models.Menu{}
	MenuSrv := service.NewMenuService(cc)
	Xtoken := c.Request().Header.Get("x-token")
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	json.Unmarshal(b, &postMenu)
	MenuSrv.GetPoint(Xtoken, c, &postMenu)

	return nil
}
func (*MenuController) GetCategoryMenu(c echo.Context) error {
	cc := c.(core.IContext)
	postMenu := models.Menu{}
	MenuSrv := service.NewMenuService(cc)
	Xtoken := c.Request().Header.Get("x-token")
	log.Print("Post : ", postMenu)
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	json.Unmarshal(b, &postMenu)
	MenuSrv.GetCategory(Xtoken, c, &postMenu)

	return nil
}
func (*MenuController) Search (c echo.Context) error {
	cc := c.(core.IContext)
	postIngredients := models.ArrayIngredients{}
	MenuSrv := service.NewMenuService(cc)
	Xtoken := c.Request().Header.Get("x-token")
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	json.Unmarshal(b, &postIngredients)
	MenuSrv.Search(Xtoken, c, &postIngredients)

	return nil
}
