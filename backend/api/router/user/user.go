package user

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Volomn/voauth/backend/api/app"
	"github.com/go-chi/render"
)

func SignupUserHandler(w http.ResponseWriter, r *http.Request) {
	data := &SignupUserRequestModel{}
	if err := render.Bind(r, data); err != nil {
		slog.Info("binding input data failed", "error", err.Error())
		render.Status(r, 422)
		render.JSON(w, r, map[string]string{"msg": fmt.Sprintf("Invalid request payload, %s", err.Error())})
		return
	}

	application := r.Context().Value("app").(app.Application)
	ctx := context.Background()
	_, err := application.SignupUser(ctx, data.FirstName, data.LastName, data.Email, data.Password)
	if err != nil {
		slog.Info("Unable to sign up user", "error", err.Error())
		render.Status(r, 400)
		render.JSON(w, r, map[string]string{"msg": err.Error()})
		return
	}

	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"msg": "User sign up successful"})
}
