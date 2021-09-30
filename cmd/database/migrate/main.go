package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/rubenv/sql-migrate"

	"github.com/ianyong/todo-backend/internal/adapters/infrastructure/database"
	"github.com/ianyong/todo-backend/internal/config"
)

const (
	directionEnvVar = "MIGRATION_DIRECTION"
	stepsEnvVar     = "MIGRATION_STEPS"
	stepsNoLimit    = 0
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	db, err := database.SetUp(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	direction := migrate.Up
	if dirEnvVarValue := os.Getenv(directionEnvVar); dirEnvVarValue != "" {
		switch strings.ToLower(dirEnvVarValue) {
		case "up":
			direction = migrate.Up
		case "down":
			direction = migrate.Down
		default:
			log.Fatalf("unrecognised value for %s: '%s'\n", directionEnvVar, dirEnvVarValue)
		}
	}

	maxSteps := stepsNoLimit
	if stepsEnvVarValue := os.Getenv(stepsEnvVar); stepsEnvVarValue != "" {
		maxSteps, err = strconv.Atoi(stepsEnvVarValue)
		if err != nil {
			log.Fatalf("unrecognised value for %s: '%s'", stepsEnvVar, stepsEnvVarValue)
		}
	}

	steps, err := migrate.ExecMax(db.DB, "postgres", migrations, direction, maxSteps)
	if err != nil {
		log.Fatalf("failed to migrate database: %v\n", err)
	}

	if direction == migrate.Up {
		log.Printf("Applied %d migrations!\n", steps)
	} else {
		log.Printf("Reverted %d migrations!\n", steps)
	}
}
