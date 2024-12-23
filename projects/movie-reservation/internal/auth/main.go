package auth

type Tokens struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type UserData struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
