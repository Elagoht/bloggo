package guards

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var (
	username     string
	publicRoutes = []string{
		"/auth/login",
		"/api/auth",
	}
	staticsRoute = "/statics"
	loginRoute   = "/auth/login"
	mainRoute    = "/"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	username = os.Getenv("BLOGGO_USERNAME")
}

func AuthorizationGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		writer http.ResponseWriter,
		request *http.Request,
	) {
		// Static files does not require authorization
		if strings.HasPrefix(request.URL.Path, staticsRoute) {
			next.ServeHTTP(writer, request)
			return
		}

		isPublic := isPublicRoute(request.URL.Path)
		validated := checkToken(request)

		if validated && isPublic {
			http.Redirect(writer, request, mainRoute, http.StatusSeeOther)
			return
		}

		if !validated && !isPublic {
			redirectToLogin(writer, request)
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func redirectToLogin(
	writer http.ResponseWriter,
	request *http.Request,
) {
	// Delete the token cookie
	http.SetCookie(writer, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour * 24),
		HttpOnly: true,
	})
	http.Redirect(writer, request, loginRoute, http.StatusSeeOther)
}

func checkToken(request *http.Request) bool {
	// Check bearer token
	cookie, err := request.Cookie("token")
	if err != nil {
		log.Println("no token", err)
		return false
	}
	authCookie := cookie.Value

	// Verify token
	parsedToken, err := jwt.Parse(
		authCookie,
		func(token *jwt.Token) (any, error) {
			return []byte("secret"), nil
		},
	)

	if err != nil || !parsedToken.Valid {
		return false
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return false
	}

	tokenUsername := claims["sub"].(string)
	return tokenUsername == username
}

func isPublicRoute(path string) bool {
	for _, route := range publicRoutes {
		if strings.HasPrefix(path, route) {
			return true
		}
	}
	return false
}
