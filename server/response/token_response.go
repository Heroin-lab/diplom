package response

type TokenResponse struct {
	AccessToken  string `json:"acces_token"`
	RefreshToken string `json:"refresh_token"`
}
