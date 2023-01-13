package tokenutil

import (
	"crypto/ed25519"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Claims is a custom claims for jwt token.
type Claims struct {
	Type string `json:"type"`
	jwt.RegisteredClaims
}

// IssueAccessToken issues signed string of access token.
func IssueAccessToken(subject, signingKey string, days int) (string, error) {
	return issueToken("access", subject, signingKey, days)
}

// IssueRefreshToken issues signed string of refresh token.
func IssueRefreshToken(subject, signingKey string, days int) (string, error) {
	return issueToken("refresh", subject, signingKey, days)
}

// issueToken issues signed string of jwt token using EdDSA signing method.
// EdDSA was chosen considering both performance and security.
func issueToken(tokenType, subject, signingKey string, days int) (string, error) {
	now, dur := time.Now(), time.Duration(24*days)*time.Hour
	claims := Claims{
		Type: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(now.Add(dur)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        subject,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	signedStr, err := token.SignedString(ed25519.PrivateKey(signingKey))
	if err != nil {
		return "", fmt.Errorf("signed string: %w", err)
	}

	return signedStr, nil
}
