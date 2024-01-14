package models

import "github.com/golang-jwt/jwt"

type RegisterReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRes struct {
	ApiResponseCode    string `json:"apiResponseCode"`
	ApiResponseMessage string `json:"apiResponseMessage"`
}

type LoginUserReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUserRes struct {
	ApiResponseCode    string `json:"apiResponseCode"`
	ApiResponseMessage string `json:"apiResponseMessage"`
	AccessToken        string `json:"accessToken"`
}
type JwtCustomClaims struct {
	UserID int
	Name   string
	jwt.StandardClaims
}

type RedisClaims struct {
	UserID int
	Name   string
}

type GenerateCaptchaResponse struct {
	ApiResponseCode    string `json:"apiResponseCode"`
	ApiResponseMessage string `json:"apiResponseMessage"`
	CaptchaCode        string `json:"captchaCode"`
}
