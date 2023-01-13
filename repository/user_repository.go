package repositories

import (
	" hery-ciaputra/demo-gin/models"
	"gorm.io/gorm"
)

// todo: define user repo interface, struct, and config
type UserRepository interface {
	MatchingCredential(email, password string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) UserRepository {
	return &userRepository{db: c.DB}
}

func (u *userRepository) MatchingCredential(email, password string) (*models.User, error) {
	// todo: query user by email and pwd
	var user *models.User

	res := u.db.Where("email = ?", email).Where("password = ?", password).First(&user)
	return user, res.Error
}
