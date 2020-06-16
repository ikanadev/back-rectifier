package auth

import (
	"github.com/go-pg/pg/v10/orm"
	"github.com/vkevv/back-rectifier/pkg/models"
)

// DB all database actions for auth
type DB struct{}

// GetUserByID get by ID
func (a DB) GetUserByID(db orm.DB, id int) (models.User, error) {
	user := models.User{}
	user.ID = id
	err := db.Select(&user)
	return user, err
}

// GetUserByEmail get user by email
func (a DB) GetUserByEmail(db orm.DB, email string) (models.User, error) {
	user := models.User{}
	user.Email = email
	err := db.Model(&user).Where("email = ?email").First()
	return user, err
}

// ExistsEmail check email existence
func (a DB) ExistsEmail(db orm.DB, email string) (bool, error) {
	user := models.User{}
	user.Email = email
	return db.Model(&user).Where("email = ?email").Exists()
}

// InsertUser insert user in db
func (a DB) InsertUser(db orm.DB, user *models.User) error {
	return db.Insert(user)
}
