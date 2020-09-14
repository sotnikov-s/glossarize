package security

import (
	"github.com/sotnikov-s/glossarize-backend/backend/model"
	"github.com/sotnikov-s/glossarize-backend/backend/repository"
)

// Identity represents a wrapper around a user.
type Identity struct {
	user *model.User
}

// NewIdentityBuilder creates a new instance of the IdentityBuilder.
func NewIdentityBuilder(userRepo *repository.UserRepository) *IdentityBuilder {
	return &IdentityBuilder{userRepo: userRepo}
}

// IdentityBuilder handles identity creation.
type IdentityBuilder struct {
	userRepo *repository.UserRepository
}

// Build returns an Identity in accordance with the passed userID.
func (b *IdentityBuilder) Build(userID int) (*Identity, error) {
	user, err := b.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return &Identity{user: user}, nil
}
