package auth

import "net/http"

type EmailAndPasswordAuthRequestModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *EmailAndPasswordAuthRequestModel) Bind(r *http.Request) error {
	return nil
}
