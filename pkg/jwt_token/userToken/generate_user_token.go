package usertoken

import "github.com/golang-jwt/jwt"

const SecretKeyUser = "LJUGBHAUghjb^$%^(mmsdadas)fdfgdfg!$ewr5UT"

type UserClaims struct {
	UUID   string
	UserID int
	jwt.StandardClaims
}

func GenerateUserToken(userId int) (*string, error) {
	panic("generate user token")
}

func ValidateUserToken(userToken string) (*UserClaims, error) {
	panic("validate user token")
}
