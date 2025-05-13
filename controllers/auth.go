package controllers

import (
	"net/http"
	"time"

	"github.com/Elagoht/bloggo/services"
	"github.com/Elagoht/bloggo/utils"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: services.NewAuthService(),
	}
}

func (controller *AuthController) Login(
	writer http.ResponseWriter,
	request *http.Request,
) *utils.AppError {
	givenUsername := request.FormValue("username")
	givenPassphrase := request.FormValue("passphrase")

	token, err := controller.authService.Login(givenUsername, givenPassphrase)
	if err != nil {
		return err
	}

	http.SetCookie(writer, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour), // 1 hour validity
		SameSite: http.SameSiteLaxMode,
	})
	http.Redirect(writer, request, "/", http.StatusSeeOther)
	return nil
}
