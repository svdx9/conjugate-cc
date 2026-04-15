package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
)

const (
	// TokenSize is the size of the random token in bytes
	TokenSize = 32
)

type Token string
type TokenHash []byte

type tokenPair struct {
	token     Token     // Plain text token (sent to user)
	tokenHash TokenHash // SHA-256 hash (stored in database)
}

func (t tokenPair) String() string {
	return string(t.token)
}

func verifyToken(token Token, tokenHash TokenHash) error {
	// VerifyToken checks if a plain token matches a stored hash using constant-time comparison
	if token == "" {
		return ErrInvalidToken
	}
	// Decode the plain token from base64
	decodedToken, err := base64.URLEncoding.DecodeString(string(token))
	if err != nil {
		return ErrInvalidToken
	}
	// Hash the decoded token
	hash := sha256.Sum256(decodedToken)
	// Compare using constant-time comparison to prevent timing attacks
	if subtle.ConstantTimeCompare(hash[:], tokenHash) != 1 {
		return ErrTokenHashMismatch
	}
	return nil
}

// GenerateToken creates a new random token and returns both the plain token and its hash
func newTokenPair() (tokenPair, error) {
	// Generate random bytes
	randomBytes := make([]byte, TokenSize)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return tokenPair{}, err
	}

	// Encode to base64 URL-safe format (sent to user)
	token := base64.URLEncoding.EncodeToString(randomBytes)

	// Hash the token (stored in database)
	hash := sha256.Sum256(randomBytes)

	return tokenPair{
		token:     Token(token),
		tokenHash: hash[:],
	}, nil
}
