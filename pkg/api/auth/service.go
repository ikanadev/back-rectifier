package auth

import (
	"github.com/gin-gonic/gin"
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
func LoadAuthService(dbActions DBActions, tokenGen TokenGenerator) Auth {
	return Auth{
		tokenGen:  tokenGen,
		DBActions: dbActions,
	}
}

// Auth struct who implements Service Interface
type Auth struct {
	tokenGen  TokenGenerator
	DBActions DBActions
}

// TokenGenerator interface which generates a token
type TokenGenerator interface {
	GenerateToken(models.User) (string, error)
}

// DBActions all DB actions
type DBActions interface {
	GetUserByID(id int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	ExistsEmail(email string) (bool, error)
	InsertUser(user *models.User) error
}
