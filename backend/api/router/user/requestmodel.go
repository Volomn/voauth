package user

import "net/http"

type SignupUserRequestModel struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (a *SignupUserRequestModel) Bind(r *http.Request) error {
	return nil
}
