package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ismoilroziboyev/bookshelf/internal/config"
	"github.com/ismoilroziboyev/bookshelf/internal/repository/postgres/sqlc"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Database interface {
	sqlc.Querier
	StartTX(ctx context.Context) (pgx.Tx, *sqlc.Queries, error)
	Close()
}

type db struct {
	*sqlc.Queries
	Conn *pgxpool.Pool
	m    migrate.Migrate
}

func New(cfg *config.Config) Database {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	dbConn, err := pgxpool.Connect(ctx, cfg.PSQLUri)

	if err != nil {
		logrus.Fatalf("error occurred while connecting database: %s", err.Error())
	}

	if err := dbConn.Ping(ctx); err != nil {
		logrus.Fatalf("error occurred while sending ping signal to the database: %s", err.Error())
	}

	logrus.Info("database connection established successfully...")

	sqlDB, err := sql.Open("postgres", cfg.PSQLUri)

	if err != nil {
		logrus.Fatal("cannot open connection for migration", err)
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})

	if err != nil {
		logrus.Fatal("cannot make driver", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)

	if err != nil {
		logrus.Fatal("cannot make migration instance", err)
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			logrus.Fatalf("cannot migrate up database: %s", err)
		}
	}

	logrus.Info("database migration was successfully.....")

	return &db{
		Queries: sqlc.New(dbConn),
		Conn:    dbConn,
		m:       *m,
	}
}

func (d *db) Close() {

	// if err := d.m.Drop(); err != nil {
	// 	logrus.Errorf("cannnot drop database: %s", err.Error())
	// }

	d.Conn.Close()
}

func (d *db) StartTX(ctx context.Context) (pgx.Tx, *sqlc.Queries, error) {
	tx, err := d.Conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return tx, nil, err
	}

	q := sqlc.New(tx)
	return tx, q, nil
}
