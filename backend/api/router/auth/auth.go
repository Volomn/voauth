package auth

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/Volomn/voauth/backend/api/app"
	"github.com/Volomn/voauth/backend/domain"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

func createAccessToken(secretKey string, user domain.User) (string, error) {
	tokenAuth := jwtauth.New("HS256", []byte(secretKey), nil)
	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{"authId": user.UUID.String(), "exp": time.Now().UTC().Add(24 * time.Hour)})
	return tokenString, err
}

func EmailAndPasswordAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	data := &EmailAndPasswordAuthRequestModel{}
	if err := render.Bind(r, data); err != nil {
		slog.Info("binding input data failed", "error", err.Error())
		render.Status(r, 422)
		render.JSON(w, r, map[string]string{"msg": fmt.Sprintf("Invalid request payload, %s", err.Error())})
		return
	}

	application := r.Context().Value("app").(app.Application)
	user, err := application.AuthenticateWithEmailAndPassword(data.Email, data.Password)
	if err != nil {
		slog.Info("Unable to sign up user", "error", err.Error())
		render.Status(r, 400)
		render.JSON(w, r, map[string]string{"msg": err.Error()})
		return
	}
	accessToken, err := createAccessToken(application.GetAuthSecretKey(), user)
	if err != nil {
		slog.Error("Unable to create user access token", "error", err.Error())
		render.Status(r, 500)
		render.JSON(w, r, map[string]string{"msg": "Something went wrong"})
		return
	}
	render.Status(r, 200)
	render.JSON(w, r, map[string]string{"accessToken": accessToken})
}
