package datamodel

// LoginRequest is the login API request
type LoginRequest struct {
	ClientID string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse is the login API response
type LoginResponse struct {
	Expire string `json:"expire"`
	Token  string `json:"token"`
}
