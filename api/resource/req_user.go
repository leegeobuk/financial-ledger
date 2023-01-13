package resource

// ReqSignUp binds request body for sign up.
type ReqSignUp struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required,min=8,max=18"`
}

// ReqSignIn binds request body for sign in.
type ReqSignIn struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required,min=8,max=18"`
}
