package security

import (
	"fmt"
	"sync"
)

// MemoryTokenService represents a simple in-memory token storage.
type MemoryTokenService struct {
	mu     *sync.Mutex
	tokens map[string]*Token
}

// Create registers and returns a new token.
func (s *MemoryTokenService) Create(userID int) (*Token, error) {
	id, err := generateRandomString(randomStringLength)
	if err != nil {
		return nil, fmt.Errorf("failed to create a token id: %v", err)
	}
	token := &Token{ID: id, UserID: userID}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tokens[id] = token
	return token, nil
}

// GetByID returns a stored token by id if it exists.
func (s *MemoryTokenService) GetByID(id string) (*Token, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	token, ex := s.tokens[id]
	if !ex {
		return nil, fmt.Errorf("no token stored by id %s", id)
	}
	return token, nil
}
