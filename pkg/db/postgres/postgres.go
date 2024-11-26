package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PGConfig struct {
	Host           string `envconfig:"POSTGRES_HOST"`
	Port           string `envconfig:"POSTGRES_PORT"`
	User           string `envconfig:"POSTGRES_USER"`
	Password       string `envconfig:"POSTGRES_PASSWORD"`
	DBName         string `envconfig:"POSTGRES_DB_NAME"`
	SSLMode        string `envconfig:"POSTGRES_SSLMODE"`
	ConnectTimeout string `envconfig:"POSTGRES_CONNECT_TIMEOUT"`
}

func Connect(cfg *PGConfig) *sqlx.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s connect_timeout=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DBName,
		cfg.SSLMode,
		cfg.ConnectTimeout,
	)

	if cfg.Password != "" {
		dsn += fmt.Sprintf(" password=%s", cfg.Password)
	}

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
