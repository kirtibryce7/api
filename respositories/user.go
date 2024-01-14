package respositories

import (
	"PRACTICE/db"
	"PRACTICE/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func RegisterUser(request *models.RegisterReq) models.RegisterRes {
	con := db.CreateCon()
	response := models.RegisterRes{}
	recordExist := CheckRecordExists(" db_users", "email", request.Email, con)
	if !recordExist {
		sqlStatement := "INSERT INTO db_users(name,email,password) VALUES(?,?,MD5(?))"
		_, err := con.Exec(sqlStatement, request.Name, request.Email, request.Password)
		if err != nil {
			response.ApiResponseCode = "1003"
			response.ApiResponseMessage = "Unable to register your account.Please try again later! "
		} else {
			response.ApiResponseCode = "1001"
			response.ApiResponseMessage = "Your account is registered successfully"
		}
	} else {
		response.ApiResponseCode = "1003"
		response.ApiResponseMessage = "Email already exists.Please try again with another email!"
	}
	defer con.Close()
	return response
}

func LoginUser(request *models.LoginUserReq) models.LoginUserRes {
	con := db.CreateCon()
	response := models.LoginUserRes{}
	count := 0
	var userID int
	var name string
	sqlStatement := "SELECT count(userID),userID,name FROM db_users WHERE email=? AND password=MD5(?) AND status='ACTIVE'"
	err := con.QueryRow(sqlStatement, request.Email, request.Password).Scan(&count, &userID, &name)
	if err != nil {
		fmt.Println(err)
	}
	if count > 0 {
		response.ApiResponseCode = "1001"
		response.ApiResponseMessage = "User logged in successfully"
		response.AccessToken = CreateToken(userID, name)

	} else {
		response.ApiResponseCode = "1003"
		response.ApiResponseMessage = "Invalid credentials.Please try again with valid username and password"
	}

	defer con.Close()
	return response
}

func CreateToken(userID int, name string) string {
	expiresAt := time.Now().Add(time.Hour * 72).Unix()
	claims := &models.JwtCustomClaims{
		Name:   name,
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := os.Getenv("jwt_key")
	jwtToken, err := token.SignedString([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("jwttoken", jwtToken)
	client := db.RedisCon()
	err = client.SetKey("jwtToken_"+jwtToken, claims, time.Hour*72)
	if err != nil {
		fmt.Println(err)
	}
	return jwtToken

}

func RefreshToken(token string) models.LoginUserRes {
	claims := models.RedisClaims{}
	returns := models.LoginUserRes{}
	client := db.RedisCon()
	err := client.GetKey("jwtToken_"+token, &claims)
	if err != nil {
		fmt.Println(err)
		returns.ApiResponseCode = "1002"
		returns.ApiResponseMessage = "Invalid Token"
	} else {
		jwtToken := CreateToken(claims.UserID, claims.Name)
		returns.ApiResponseCode = "1001"
		returns.ApiResponseMessage = "Token refresh successfully"
		returns.AccessToken = jwtToken
		client.DelKey("jwtTokens_" + token)

	}
	return returns

}

func GenerateCaptcha() models.GenerateCaptchaResponse {
	response := models.GenerateCaptchaResponse{}
	captchaCode := GenerateCaptchaCode(2, "ALPHA") + GenerateCaptchaCode(2, "NUM") + GenerateCaptchaCode(2, "ALPHA")
	response.ApiResponseCode = "1001"
	response.ApiResponseMessage = "Captcha code generated successfully"
	response.CaptchaCode = captchaCode
	return response
}
