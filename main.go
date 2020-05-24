package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"bitbucket.org/BBeamnantapong/cooking-server/middlewares"
	Authorization "bitbucket.org/BBeamnantapong/cooking-server/Auth/http"
	Profile "bitbucket.org/BBeamnantapong/cooking-server/profile/http"
	CMS "bitbucket.org/BBeamnantapong/cooking-server/cms/http"
	Menu "bitbucket.org/BBeamnantapong/cooking-server/foodsmenu/http"
	"log"
)

func main()  {
	e := echo.New()
	e.Use(middlewares.Core)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowOrigins:     []string{"*"},
	}))
	Authorization.AuthRoute(e)
	Profile.ProfileRoute(e)
	Menu.MenuRoute(e)
	CMS.CMSRoute(e)

	e.Use(middleware.Static(""))
	e.Use(middleware.Static("img"))
	log.Fatal(e.Start(":9000"))
}
