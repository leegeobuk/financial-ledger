package resource

// ReqAccessToken binds header for authorized access.
type ReqAccessToken struct {
	Authorization string `header:"Authorization" binding:"required"`
}
