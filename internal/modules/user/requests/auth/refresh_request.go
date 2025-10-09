package auth

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
	ClientId     string `json:"client_id" form:"client_id"`
	GrantType    string `json:"grant_type" form:"grant_type"`
}
