package models

/* LoginResponse have the token which return in the login */
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
