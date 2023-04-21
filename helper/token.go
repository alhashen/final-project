package helper

import (
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "MHcCAQEEIGaozMA951amsyyAjz/C3FUhdspS1Kqi3s5EdbJeop0boAoGCCqGSM49AwEHoUQDQgAEPvB35tXsy4P4Z"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":         id,
		"user_email": email,
	}

	parsedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parsedToken.SignedString([]byte(secretKey))
	if err != nil {
		log.Println(err)
	}

	return signedToken
}

func ValidateToken(c *gin.Context) (any, error) {
	errResponse := errors.New("Sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil

}

func GetTokenValue(c *gin.Context, key string) (any, error) {
	userData := c.MustGet("token").(jwt.MapClaims)

	val, ok := userData[key]
	if !ok {
		return val, errors.New("invalid key")
	}

	return val, nil
}
