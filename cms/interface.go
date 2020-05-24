package cms

import (
	"bitbucket.org/BBeamnantapong/cooking-server/models"
	"github.com/labstack/echo/v4"
)


type CMSInterface interface {
	GetIngredients (Xtoken string, c echo.Context ) interface{}
	CMSCreateMenu (Xtoken string, c echo.Context, postMenu *models.DataIngredients) interface{}
	CMSDeleteMenu (Xtoken string, c echo.Context, postMenu *models.DataIngredients) interface{}
	CMSGetDetailUpdateMenu (Xtoken string, c echo.Context, postMenu *models.DataIngredients) interface{}
	CMSManageMenu (Xtoken string, c echo.Context,  postMenu *models.Menu) interface{}
	UploadFile(Xtoken string, c echo.Context) interface{}

}

