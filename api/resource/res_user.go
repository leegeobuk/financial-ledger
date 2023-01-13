package resource

// ResSignUp binds response body for sign up.
type ResSignUp struct {
	UserID string `json:"user_id"`
}

// ResSignIn binds response body for sign in.
type ResSignIn struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
