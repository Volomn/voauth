package auth

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Volomn/voauth/backend/api/app"
	"github.com/Volomn/voauth/backend/api/router/util"
	"github.com/go-chi/render"
)

func EmailAndPasswordAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	data := &EmailAndPasswordAuthRequestModel{}
	if err := render.Bind(r, data); err != nil {
		slog.Info("binding input data failed", "error", err.Error())
		render.Status(r, 422)
		render.JSON(w, r, map[string]string{"msg": fmt.Sprintf("Invalid request payload, %s", err.Error())})
		return
	}

	application := r.Context().Value("app").(app.Application)

	ctx := context.Background()

	user, err := application.AuthenticateWithEmailAndPassword(ctx, data.Email, data.Password)
	if err != nil {
		slog.Info("Unable to authenticate user with email and password", "error", err.Error())
		render.Status(r, 400)
		render.JSON(w, r, map[string]string{"msg": err.Error()})
		return
	}
	accessToken, err := util.CreateUserAccessToken(application.GetAuthSecretKey(ctx), user)
	if err != nil {
		slog.Error("Unable to create user access token", "error", err.Error())
		render.Status(r, 500)
		render.JSON(w, r, map[string]string{"msg": "Something went wrong"})
		return
	}
	render.Status(r, 200)
	render.JSON(w, r, map[string]string{"accessToken": accessToken})
}
