package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/household-ledger/api/resource"
	"github.com/leegeobuk/household-ledger/cfg"
	"github.com/leegeobuk/household-ledger/util/hashutil"
	"github.com/leegeobuk/household-ledger/util/tokenutil"
)

//	@Tags			User
//	@Summary		Sign up
//	@Description	Signs user up.
//	@Accept			json
//	@Produce		json
//	@Param			request	body		resource.ReqSignUp	true	"Sign up request body"
//	@Success		201		{object}	resource.ResSignUp
//	@Failure		400		{object}	resource.ResErr
//	@Failure		409		{object}	resource.ResErr
//	@Failure		500		{object}	resource.ResErr
//	@Router			/api/household-ledger/user/signup [post]
func (s *Server) SignUp(c *gin.Context) {
	req := c.MustGet("req").(resource.ReqSignUp)
	email, passwd := req.Email, req.Password

	// find if user id exists
	_, noRows, err := s.db.FindUserLogIn(email)
	if err != nil {
		resource.Error(c, fmt.Errorf("db error: %w", err))
		return
	}

	if !noRows {
		resource.Conflict(c, fmt.Errorf("email already exists: %s", email))
		return
	}

	// hash passwd
	hashedPW, err := hashutil.HashPassword(passwd, 11)
	if err != nil {
		resource.Error(c, fmt.Errorf("hash password: %w", err))
		return
	}

	// insert to db
	if err = s.db.InsertUserAccount(email); err != nil {
		resource.Error(c, fmt.Errorf("db error: %w", err))
		return
	}

	if err = s.db.InsertUserLogIn(email, hashedPW); err != nil {
		resource.Error(c, fmt.Errorf("db error: %w", err))
		return
	}

	res := resource.ResSignUp{UserID: email}

	resource.Created(c, res)
}

//	@Tags			User
//	@Summary		Sign in
//	@Description	Signs user in.
//	@Accept			json
//	@Produce		json
//	@Param			request	body		resource.ReqSignIn	true	"Sign in request body"
//	@Success		200		{object}	resource.ResSignIn
//	@Failure		400		{object}	resource.ResErr
//	@Failure		404		{object}	resource.ResErr
//	@Failure		500		{object}	resource.ResErr
//	@Router			/api/household-ledger/user/signin [post]
func (s *Server) SignIn(c *gin.Context) {
	req := c.MustGet("req").(resource.ReqSignIn)
	email, passwd := req.Email, req.Password

	// find email, passwd from db
	userLogIn, noRows, err := s.db.FindUserLogIn(email)
	if err != nil {
		resource.Error(c, fmt.Errorf("db error: %w", err))
		return
	}

	if noRows {
		resource.NotFound(c, fmt.Errorf("user id %s doesn't exist", email))
		return
	}

	// check pw
	if err = hashutil.CompareHashAndPassword(userLogIn.Passwd, passwd); err != nil {
		resource.Error(c, fmt.Errorf("compare password: %w", err))
		return
	}

	// issue access, refresh tokens
	issuer, privateKey := cfg.Env.Token.Issuer, cfg.Env.Token.PrivateKey
	accessDur := cfg.Env.Token.AccessDuration
	accessToken, err := tokenutil.IssueAccessToken(issuer, privateKey, accessDur)
	if err != nil {
		resource.Error(c, fmt.Errorf("access token: %w", err))
		return
	}

	refreshDur := cfg.Env.Token.RefreshDuration
	refreshToken, err := tokenutil.IssueRefreshToken(issuer, privateKey, refreshDur)
	if err != nil {
		resource.Error(c, fmt.Errorf("refresh token: %w", err))
		return
	}

	res := resource.ResSignIn{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	resource.OK(c, res)
}
