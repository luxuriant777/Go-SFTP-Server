package auth

import (
	"golang.org/x/crypto/ssh"
	"os"
)

// AuthHandler handles SSH key-based authentication
type AuthHandler struct {
	privateKey ssh.Signer
}

// NewHandler creates a new AuthHandler
func NewHandler() (*AuthHandler, error) {
	privateBytes, err := os.ReadFile("path/to/your/private/key")
	if err != nil {
		return nil, err
	}

	privateKey, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		return nil, err
	}

	return &AuthHandler{privateKey: privateKey}, nil
}

// GetPrivateKey returns the parsed private key
func (h *AuthHandler) GetPrivateKey() ssh.Signer {
	return h.privateKey
}
