package auth

import (
	"crypto/rand"
	"crypto/sha256"
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

func decodeToken(token Token) (TokenHash, error) {
	// Decode and hash the token
	decodedToken, err := base64.URLEncoding.DecodeString(string(token))
	if err != nil {
		return TokenHash{}, ErrInvalidToken
	}
	buf := sha256.Sum256(decodedToken)
	return buf[:], nil
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
