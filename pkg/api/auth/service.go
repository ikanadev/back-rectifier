package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/vkevv/back-rectifier/pkg/models"
)

// Service represents all handlers for authentication
type Service interface {
	Login(email, password string) (models.User, error)
	Register(email, password, name, lastName string) (models.User, error)
	Me(c *gin.Context) (models.User, error)
	GenerateToken(user models.User) (string, error)
}

// LoadAuthService function which generates a AuthService
func LoadAuthService(db *pg.DB, tokenGen TokenGenerator) Auth {
	return Auth{
		db:        db,
		tokenGen:  tokenGen,
		DBActions: DB{},
	}
}

// Auth struct who implements Service Interface
type Auth struct {
	db        *pg.DB
	tokenGen  TokenGenerator
	DBActions DBActions
}

// TokenGenerator interface which generates a token
type TokenGenerator interface {
	GenerateToken(models.User) (string, error)
}

// DBActions all DB actions
type DBActions interface {
	GetUserByID(db orm.DB, id int) (models.User, error)
	GetUserByEmail(db orm.DB, email string) (models.User, error)
	ExistsEmail(db orm.DB, email string) (bool, error)
	InsertUser(db orm.DB, user *models.User) error
}
