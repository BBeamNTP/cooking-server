package profile

import (
	"bitbucket.org/BBeamnantapong/cooking-server/models"
	"github.com/labstack/echo/v4"
)

type ProfileInterface interface {
	GetUserProfile(Xtoken string, c echo.Context) interface{}
	UpdateUserProfile(Xtoken string, postData *models.Data, c echo.Context) interface{}
	Upload(Xtoken string, c echo.Context) interface{}
}
