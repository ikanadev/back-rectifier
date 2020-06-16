package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	// DB adapter
	_ "github.com/lib/pq"
	"github.com/vkevv/back-rectifier/pkg/config"
)

type dbLogger struct{}

// BeforeQuery hooks before pg queries
func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

// AfterQuery hooks after pg queries
func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	query, err := q.FormattedQuery()
	fmt.Println(query)
	return err
}

// New creates new database connection to a postgres database
func New(dbConf config.DB) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     dbConf.Host + ":5432",
		User:     dbConf.Username,
		Password: dbConf.Password,
		Database: dbConf.DBName,
	})

	_, err := db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	if dbConf.TimeoutSeconds > 0 {
		db = db.WithTimeout(time.Second * time.Duration(dbConf.TimeoutSeconds))
	}

	if dbConf.LogQueries {
		db.AddQueryHook(dbLogger{})
	}

	return db, nil
}

// CreateTables function which create Tables
func CreateTables(db *pg.DB, models ...interface{}) error {
	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
			IfNotExists:   true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
