package api

import (
	"github.com/vkevv/back-rectifier/pkg/api/auth"
	"github.com/vkevv/back-rectifier/pkg/api/common"
	"github.com/vkevv/back-rectifier/pkg/api/project"
	"github.com/vkevv/back-rectifier/pkg/config"
	"github.com/vkevv/back-rectifier/pkg/jwt"
	"github.com/vkevv/back-rectifier/pkg/models"
	"github.com/vkevv/back-rectifier/pkg/postgres"
	"github.com/vkevv/back-rectifier/pkg/server"
)

// StartAPI starts a rest api service
func StartAPI(conf config.Config) error {
	models := []interface{}{
		&models.User{},
		&models.Project{},
		&models.Document{},
		&models.Observation{},
	}
	db, err := postgres.New(conf.DB)
	if err != nil {
		return err
	}
	if err := postgres.CreateTables(db, models); err != nil {
		return err
	}

	jwt, err := jwt.New(conf.JWT.SigningAlgorithm, conf.JWT.Key, conf.JWT.DurationMinutes, conf.JWT.MinSecretLength)
	if err != nil {
		return err
	}
	authMD := server.Auth(conf.JWT.Key, jwt)
	dbActions := common.NewDB(db)
	gin := server.New(conf.Server)
	authGin := gin.Group("/v1")
	authGin.Use(authMD)

	authService := auth.LoadAuthService(dbActions, jwt)
	auth.ServeHTTP(&authService, gin, authMD)

	projectService := project.LoadProjectService(dbActions)
	project.ServeHTTP(&projectService, authGin)

	if err := server.Start(conf.Server, gin); err != nil {
		return err
	}
	return nil
}
