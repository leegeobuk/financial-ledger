package tokenutil

import (
	"crypto"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenExpired      = errors.New("token expired")
	ErrIncorrectIssuedAt = errors.New("token issued at different time")
	ErrTokenIssuedBefore = errors.New("token issued before")
	ErrIncorrectIssuer   = errors.New("incorrect issuer")
)

// Claims is a custom claims for jwt token.
type Claims struct {
	Type string `json:"type"`
	jwt.RegisteredClaims
}

// IssueAccessToken issues signed string of access token.
func IssueAccessToken(issuer string, privateKey crypto.PrivateKey, days int) (string, error) {
	return issueToken("access", issuer, privateKey, days)
}

// IssueRefreshToken issues signed string of refresh token.
func IssueRefreshToken(issuer string, privateKey crypto.PrivateKey, days int) (string, error) {
	return issueToken("refresh", issuer, privateKey, days)
}

// issueToken issues signed string of jwt token using EdDSA signing method.
// EdDSA was chosen considering both performance and security.
func issueToken(tokenType, issuer string, privateKey crypto.PrivateKey, days int) (string, error) {
	now, dur := time.Now(), time.Duration(24*days)*time.Hour
	claims := Claims{
		Type: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(now.Add(dur)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	signedStr, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("signed string: %w", err)
	}

	return signedStr, nil
}

// ValidateAccessToken parses access token.
func ValidateAccessToken(accessToken string, publicKey crypto.PublicKey) error {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("invalid signing method: %s", token.Method.Alg())
		}

		return publicKey, nil
	})
	if err != nil {
		return fmt.Errorf("parse with claims: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return ErrInvalidToken
	}

	now := time.Now()
	if !claims.VerifyExpiresAt(now, true) {
		return ErrTokenExpired
	}

	if !claims.VerifyIssuedAt(now, true) {
		return ErrIncorrectIssuedAt
	}

	if !claims.VerifyNotBefore(now, true) {
		return ErrTokenIssuedBefore
	}

	if !claims.VerifyIssuer("household-ledger", true) {
		return ErrIncorrectIssuer
	}

	return nil
}
