package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// get hello
func GetHello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello Go Developers")
}

// get palindrome
func GetPalindrome(c echo.Context) error {
	input := c.QueryParam("text")

	isPalindrome := palindrome(input)

	response := "Palindrome"
	statusCode := http.StatusOK
	if !isPalindrome {
		response = "Not palindrome"
		statusCode = http.StatusBadRequest
	}
	return c.JSON(statusCode, response)
}

//palindrome function
func palindrome(text string) bool {
	for i := 0; i < len(text); i++ {
		j := len(text) - 1 - i
		if text[i] != text[j] {
			return false
		}
	}
	return true
}
