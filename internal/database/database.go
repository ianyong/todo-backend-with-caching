package database

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/ianyong/todo-backend/internal/config"
)

// SetUp sets up a sqlx.DB database connection and returns it.
func SetUp(cfg *config.Config) (*sqlx.DB, error) {
	dsn := buildDSN(cfg)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	db.MapperFunc(convertCamelToSnake)
	return db, nil
}

// buildDSN builds the data source name that is used to connect to a database.
func buildDSN(cfg *config.Config) string {
	dsn := ""
	if cfg.DBName != "" {
		dsn += fmt.Sprintf("dbname=%v", cfg.DBName)
	}
	if cfg.DBHost != "" {
		dsn += fmt.Sprintf(" host=%v", cfg.DBHost)
	}
	if cfg.DBPort != 0 {
		dsn += fmt.Sprintf(" port=%v", cfg.DBPort)
	}
	if cfg.DBUser != "" {
		dsn += fmt.Sprintf(" user=%v", cfg.DBUser)
	}
	if cfg.DBPassword != "" {
		dsn += fmt.Sprintf(" password=%v", cfg.DBPassword)
	}
	if cfg.DBSSLMode != "" {
		dsn += fmt.Sprintf(" sslmode=%v", cfg.DBSSLMode)
	}
	return dsn
}

func convertCamelToSnake(s string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
