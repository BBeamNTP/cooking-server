package foodsmenu

import (
	"bitbucket.org/BBeamnantapong/cooking-server/models"
	"github.com/labstack/echo/v4"
)

type MenuInterface interface {
	GetMenu(Xtoken string, c echo.Context) interface{}
	GetFood(Xtoken string, c echo.Context, postMenu *models.DataIngredients) interface{}
	GetPoint(Xtoken string, c echo.Context, postMenu *models.Menu) interface{}
	GetCategory(Xtoken string, c echo.Context,  postMenu *models.Menu) interface{}
	Search (Xtoken string, c echo.Context,  postIngredients *models.ArrayIngredients) interface{}
}
