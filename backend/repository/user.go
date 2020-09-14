package repository

import (
	"fmt"
	"time"

	"github.com/sotnikov-s/glossarize-backend/backend/model"

	"github.com/jinzhu/gorm"
)

// UserRepository handles db operations regarding the users table.
type UserRepository struct {
	db *gorm.DB
}

// GetByID finds a single user by id.
func (r *UserRepository) GetByID(id int) (*model.User, error) {
	var user model.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user with id %d doesn't exits: %v", id, err)
	}
	return &user, nil
}

// GetByEmail finds a single user by email.
func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user with email %s doesn't exits: %v", email, err)
	}
	return &user, nil
}

// Create creates a single user.
func (r *UserRepository) Create(user *model.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to create a user: %v", err)
	}
	return nil
}

// Update updates a single user.
func (r *UserRepository) Update(user *model.User) error {
	user.UpdatedAt = time.Now().UTC()
	if err := r.db.Save(user).Error; err != nil {
		return fmt.Errorf("failed to update a user: %v", err)
	}
	return nil
}

// Delete deletes a single user.
func (r *UserRepository) Delete(user *model.User) error {
	if err := r.db.Delete(&user).Error; err != nil {
		return fmt.Errorf("failed to delete a user: %v", err)
	}
	return nil
}
