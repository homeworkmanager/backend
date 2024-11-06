package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PGConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	SSLMode  string `json:"sslMode"`
}

func Connect(cfg *PGConfig) *sqlx.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.DBName,
		cfg.SSLMode,
	)

	if cfg.Password != "" {
		dsn += fmt.Sprintf(" %s", cfg.Password)
	}

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
