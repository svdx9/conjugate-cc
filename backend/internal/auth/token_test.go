package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	t.Parallel()

	pair, err := newTokenPair()
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	// Token should be non-empty
	if pair.token == "" {
		t.Error("Token is empty")
	}

	// TokenHash should be non-empty and match SHA-256 hash
	if len(pair.tokenHash) != 32 { // SHA-256 produces 32 bytes
		t.Errorf("TokenHash length = %d, want 32", len(pair.tokenHash))
	}

	// Verify the hash is correct
	decodedToken, _ := base64.URLEncoding.DecodeString(string(pair.token))
	expectedHash := sha256.Sum256(decodedToken)
	for i, b := range expectedHash[:] {
		if pair.tokenHash[i] != b {
			t.Errorf("TokenHash mismatch at position %d", i)
			break
		}
	}

	// Two generated tokens should be different
	pair2, _ := newTokenPair()
	if pair.token == pair2.token {
		t.Error("Two generated tokens are identical, expected unique tokens")
	}
}
