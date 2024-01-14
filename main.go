package main

import (
	"PRACTICE/controller"
	_ "PRACTICE/docs"
	"net/http"

	"github.com/go-playground/validator"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// @title GOLANG API
// @version 1.0

// @host 192.168.1.6:8000
func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/registerUser", controller.RegisterUser)
	e.POST("/login", controller.LoginUser)
	e.GET("/refreshToken", controller.RefreshToken)
	e.GET("/generateCaptcha", controller.GenerateCaptcha)

	e.Logger.Fatal(e.Start(":8000"))

}
