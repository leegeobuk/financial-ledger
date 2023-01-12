package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leegeobuk/household-ledger/api/resource"
	"github.com/leegeobuk/household-ledger/util/hashutil"
)

//	@Tags			User
//	@Summary		Sign up
//	@Description	Signs user up.
//	@Accept			json
//	@Produce		json
//	@Param			request	body		resource.ReqSignUp	true	"Sign up request body"
//	@Success		201		{object}	resource.ResSignUp
//	@Failure		400		{object}	resource.ResErr
//	@Failure		500		{object}	resource.ResErr
//	@Router			/api/household-ledger/user/signup [post]
func (s *Server) SignUp(c *gin.Context) {
	req := c.MustGet("req").(resource.ReqSignUp)
	email, passwd := req.Email, req.Password

	// find if user id exists
	_, noRows, err := s.db.FindUserLogIn(email)
	if err != nil {
		if !noRows {
			resource.Conflict(c, fmt.Errorf("email already exists: %w", err))
			return
		}

		resource.Error(c, fmt.Errorf("db error: %w", err))
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
