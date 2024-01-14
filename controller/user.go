package controller

import (
	"PRACTICE/models"
	"PRACTICE/respositories"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// @Tags User
// @Accept json
// @Produce json
// @Param request body models.RegisterReq true "Send request data"
// @Success 200 {object} models.RegisterRes
// @Router /registerUser [post]
func RegisterUser(c echo.Context) error {
	request := new(models.RegisterReq)
	response := models.RegisterRes{}
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		response.ApiResponseCode = "1004"
		response.ApiResponseMessage = err.Error()
	} else {
		response = respositories.RegisterUser(request)
	}
	return c.JSON(http.StatusOK, response)
}

// @Tags User
// @Accept json
// @Produce json
// @Param request body models.LoginUserReq true "Send request data"
// @Success 200 {object} models.LoginUserRes
// @Router /login [post]
func LoginUser(c echo.Context) error {
	request := new(models.LoginUserReq)
	response := models.LoginUserRes{}
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		response.ApiResponseCode = "1004"
		response.ApiResponseMessage = err.Error()
	} else {
		response = respositories.LoginUser(request)
	}
	return c.JSON(http.StatusOK, response)

}

// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.LoginUserRes
// @Router /refreshToken [get]
func RefreshToken(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	response := models.LoginUserRes{}
	if token != "" {
		bearerToken := strings.Fields(token)
		response = respositories.RefreshToken(bearerToken[1])
	} else {
		response.ApiResponseCode = "1002"
		response.ApiResponseMessage = "Authorization token required"
	}
	return c.JSON(http.StatusOK, response)

}

// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.GenerateCaptchaResponse
// @Router /generateCaptcha [get]
func GenerateCaptcha(c echo.Context) error {
	response := respositories.GenerateCaptcha()
	return c.JSON(http.StatusOK, response)
}
