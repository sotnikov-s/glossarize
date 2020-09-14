package security

// Token represents a simple session token.
type Token struct {
	ID     string `json:"id"`
	UserID int    `json:"user_id"`
}

// TokenService represents an interface to operate with session tokens.
type TokenService interface {
	Create(userID int) (*Token, error)
	GetByID(id string) (*Token, error)
}
