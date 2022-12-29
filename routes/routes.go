package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/danielsugianto/pre-test/test-talenta-nusantara-berkarya/controllers"
)

func New() *echo.Echo {
	e := echo.New()

	//general
	e.GET("/", controllers.GetHello)
	e.GET("/palindrome", controllers.GetPalindrome)

	//language Hard Coded
	e.GET("/language", controllers.GetLanguageHardCodedController)

	//CRUD Languages
	e.GET("/languages", controllers.GetLanguagesController)
	e.POST("/language", controllers.CreateLanguageController)
	e.GET("/language/:id", controllers.GetLanguageController)
	e.PATCH("/language/:id", controllers.UpdateLanguageController)
	e.DELETE("/language/:id", controllers.DeleteLanguageController)

	return e
}
