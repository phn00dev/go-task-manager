package admintoken

import "github.com/golang-jwt/jwt"

const SecretKeyAdmin = "JYUIiyiuhiJKH^$M)&NKJHjug(JL_)GR#%)MLKHjjsaQED"

type AdminClaims struct {
	UUID      string `json:"uuid"`
	AdminID   int    `json:"admin_id"`
	AdminRole string `json:"admin_role"`
	jwt.StandardClaims
}

func GenerateAdminToken(adminID int, adminRole string) (*string, error) {
	panic("generate admin Token")
}

func ValidateAdminToken(adminToken string) (*AdminClaims, error) {
	panic("validate admin Token")
}
