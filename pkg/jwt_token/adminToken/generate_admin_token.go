package admintoken

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const SecretKeyAdmin = "JYUIiyiuhiJKH^$M)&NKJHjug(JL_)GR#%)MLKHjjsaQED"

type AdminClaims struct {
	UUID      string `json:"uuid"`
	AdminID   int    `json:"admin_id"`
	AdminRole string `json:"admin_role"`
	jwt.StandardClaims
}

func GenerateAdminToken(adminID int, adminRole string) (*string, error) {
	expirationTime := time.Now().Add(3 * time.Hour)
	newUUID := uuid.New().String()
	adminClaims := AdminClaims{
		UUID:      newUUID,
		AdminID:   adminID,
		AdminRole: adminRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, adminClaims)

	tokenString, err := token.SignedString([]byte(SecretKeyAdmin))
	if err != nil {
		// Return the error properly
		return nil, err
	}
	return &tokenString, nil
}

func ValidateAdminToken(adminTokenString string) (*AdminClaims, error) {
	token, err := jwt.ParseWithClaims(adminTokenString, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKeyAdmin), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*AdminClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
