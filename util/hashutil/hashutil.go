package hashutil

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes given password with difficulty of cost.
func HashPassword(password string, cost int) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(hashed), err
}

// CompareHashAndPassword compares and checks
// if given hash and password are equal.
func CompareHashAndPassword(hashedPW, plainPW string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPW), []byte(plainPW)) //nolint:wrapcheck
}
