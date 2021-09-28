package main

import (
	"log"
	"os"

	"github.com/ianyong/todo-backend/internal/config"
	"github.com/ianyong/todo-backend/internal/database"
)

const overrideEnvVar = "OVERRIDE_DATABASE_DELETE_SAFEGUARD"

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	if cfg.Environment != "development" && cfg.Environment != "test" {
		if os.Getenv(overrideEnvVar) != "1" {
			log.Fatalf("Unable to drop database in a non-development and non-test environment\n"+
				"Re-run this command with the prefix '%s=1' to override this safeguard\n", overrideEnvVar)
			return
		}
	}

	// Connect to the database server without a database specified since we want to drop it.
	dbName := cfg.DBName
	cfg.DBName = ""
	db, err := database.SetUp(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}

	_, err = db.Exec("DROP DATABASE " + dbName)
	if err != nil {
		log.Fatalf("failed to drop database: %v\n", err)
	}

	log.Printf("Successfully dropped database '%s'", dbName)
}
