package services

import (
	"fmt"
	"time"
	"web-go-gin/utils"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	CreateToken(name string, admin bool) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "buithean2010@gmail.com",
	}
}

func getSecretKey() string {
	config, err := utils.LoadConfig()

	if err != nil {
		return "secret"
	}

	return config.JwtSecretKey
}

func (jwtSrv *jwtService) CreateToken(username string, admin bool) string {

	// Set custom and standard claims
	claims := &jwtCustomClaims{
		username,
		admin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "buithean2010@gmail.com",
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token using the secret signing key
	ss, err := token.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		panic(err)
	}
	return ss
}

func (jwtSrv *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(jwtSrv.secretKey), nil
	})

	return token, err
}
