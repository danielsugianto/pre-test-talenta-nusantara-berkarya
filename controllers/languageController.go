package controllers

import (
	"net/http"

	"github.com/danielsugianto/pre-test/test-talenta-nusantara-berkarya/lib/database"
	"github.com/danielsugianto/pre-test/test-talenta-nusantara-berkarya/models"
	"github.com/labstack/echo/v4"
)

// get language HardCoded
func GetLanguageHardCodedController(c echo.Context) error {
	language, err := database.GetLanguageHardCoded()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, language)
}

// get all language
func GetLanguagesController(c echo.Context) error {
	languages, err := database.GetLanguages()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, languages)
}

// create new language
func CreateLanguageController(c echo.Context) error {
	language := models.Language{}
	c.Bind(&language)
	newLanguage, err := database.CreateLanguage(language)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, newLanguage)
}

// get language by ID
func GetLanguageController(c echo.Context) error {
	id := c.Param("id")
	var language models.Language

	language, err := database.GetLanguage(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, language)
}

// delete language by ID
func DeleteLanguageController(c echo.Context) error {
	id := c.Param("id")

	err := database.DeleteLanguage(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "success delete language")
}

// update language by ID
func UpdateLanguageController(c echo.Context) error {
	id := c.Param("id")
	LanguageParam := models.Language{}
	c.Bind(&LanguageParam)
	var language models.Language

	language, err := database.UpdateLanguage(id, LanguageParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, language)
}
