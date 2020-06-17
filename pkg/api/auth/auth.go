package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/vkevv/back-rectifier/pkg"
	"github.com/vkevv/back-rectifier/pkg/api/common"
	"github.com/vkevv/back-rectifier/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

const (
	registerEmailExists = "REGISTER_EMAIL_EXISTS"
	loginBadCredentials = "LOGIN_BAD_CREDENTIALS"
)

var (
	errExistEmail     = pkg.NewAPIErr("Provided email is already in use", registerEmailExists)
	errBadCredentials = pkg.NewAPIErr("Email or password doesn't match", loginBadCredentials)
)

// Login login
func (a *Auth) Login(email, password string) (models.User, error) {
	user, err := a.DBActions.GetUserByEmail(email)
	if err != nil {
		return models.User{}, errBadCredentials
	}
	if !hashMatchesPassword(user.Password, password) {
		return models.User{}, errBadCredentials
	}
	return user, nil
}

// Register register an user, unique email
func (a *Auth) Register(email, password, name, lastName string) (models.User, error) {
	user := models.User{
		Email:    email,
		Name:     name,
		LastName: lastName,
	}
	exist, err := a.DBActions.ExistsEmail(email)
	if err != nil {
		return user, err
	}
	if exist {
		return user, errExistEmail
	}
	hashedPasswd, err := hash(password)
	if err != nil {
		return user, err
	}
	user.Password = hashedPasswd
	if err := a.DBActions.InsertUser(&user); err != nil {
		return user, err
	}
	return user, nil
}

// Me returns the user from a valid token
func (a *Auth) Me(c *gin.Context) (models.User, error) {
	userID, exist := c.Get("id")
	if !exist {
		return models.User{}, common.ErrNoContextID
	}
	userIDInt, ok := userID.(int)
	if !ok {
		return models.User{}, common.ErrContextIDInvalid
	}
	user, err := a.DBActions.GetUserByID(userIDInt)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GenerateToken generates a token
func (a *Auth) GenerateToken(user models.User) (string, error) {
	return a.tokenGen.GenerateToken(user)
}

// HELPER FUNCTIONS
func hash(password string) (string, error) {
	hashedPW, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPW), err
}
func hashMatchesPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
