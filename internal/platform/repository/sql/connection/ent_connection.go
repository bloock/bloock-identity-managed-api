package connection

import (
	"bloock-identity-managed-api/internal/platform/repository/sql/ent"
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"strings"
)

const (
	Mysql    = "mysql"
	Postgres = "postgres"
	Sqlite   = "sqlite3"
)

type EntConnection struct {
	db     *ent.Client
	logger zerolog.Logger
}

func NewEntConnection(connectionURL string, connector SQLConnector, logger zerolog.Logger) (*EntConnection, error) {
	if connectionURL == "" {
		return &EntConnection{}, errors.New("connectionURL cannot be empty")
	}

	if strings.Contains(connectionURL, "file") {
		client, err := open(connector, Sqlite, connectionURL)
		if err != nil {
			return nil, err
		}
		return &EntConnection{
			db: client,
		}, nil
	}
	if strings.Contains(connectionURL, "mysql") {
		client, err := open(connector, Mysql, strings.Replace(connectionURL, "mysql://", "", 1))
		if err != nil {
			return nil, err
		}
		return &EntConnection{
			db: client,
		}, nil
	}
	if strings.Contains(connectionURL, "postgres") {
		db, err := sql.Open("pgx", connectionURL)
		if err != nil {
			return nil, err
		}

		drv := entsql.OpenDB(dialect.Postgres, db)
		return &EntConnection{
			db: ent.NewClient(ent.Driver(drv)),
		}, nil
	}

	err := errors.New("unsupported database")
	logger.Error().Err(err).Msgf(" with url: %s", connectionURL)
	return nil, err

}

func (c *EntConnection) DB() *ent.Client {
	return c.db
}

func open(connector SQLConnector, driver string, connectionURL string) (*ent.Client, error) {
	client, err := connector.Connect(driver, connectionURL)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *EntConnection) Migrate(ctx context.Context) error {
	if err := c.db.Schema.Create(ctx); err != nil {
		c.logger.Err(err).Msg("")
		return err
	}
	return nil
}
