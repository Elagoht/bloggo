package services

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Elagoht/bloggo/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var (
	username   string
	passphrase string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	username = os.Getenv("BLOGGO_USERNAME")
	passphrase = os.Getenv("BLOGGO_PASSPHRASE")
}

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (service *AuthService) Login(
	givenUsername, givenPassphrase string,
) (string, *utils.AppError) {
	log.Println("givenUsername", givenUsername)
	log.Println("givenPassphrase", givenPassphrase)
	log.Println("username", username)
	log.Println("passphrase", passphrase)

	if givenUsername != username || givenPassphrase != passphrase {
		return "", utils.NewAppError(
			http.StatusUnauthorized,
			"Invalid username or passphrase",
			nil,
		)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": givenUsername,
		"exp": time.Now().Add(time.Hour).Unix(), // 1 hour validity
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", utils.NewAppError(
			http.StatusInternalServerError,
			"Failed to generate token",
			err,
		)
	}

	return tokenString, nil
}
