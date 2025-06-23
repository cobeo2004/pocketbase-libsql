package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

// register the libsql driver to use the same query builder
// implementation as the already existing sqlite3 builder
func init() {
	dbx.BuilderFuncMap["libsql"] = dbx.BuilderFuncMap["sqlite3"]
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file, but continuing: %s", err);
	}

	dbPath := os.Getenv("DB_PATH")
	dbAuthToken := os.Getenv("DB_AUTH_TOKEN")
	var dbUrl string;

	if dbPath == "" {
		log.Fatalf("DB_PATH is not set")
	} else if strings.HasPrefix(dbPath, "http://") || strings.HasPrefix(dbPath, "https://") {
		if dbAuthToken == "" {
			dbUrl = dbPath
		} else {
			dbUrl = fmt.Sprintf("%s?authToken=%s", dbPath, dbAuthToken)
		}
	} else if dbAuthToken == "" {
		dbUrl = fmt.Sprintf("libsql://%s", dbPath)
	} else {
		dbUrl = fmt.Sprintf("libsql://%s?authToken=%s", dbPath, dbAuthToken)
	}

	app := pocketbase.NewWithConfig(pocketbase.Config{
		DBConnect: func(dbPath string) (*dbx.DB, error) {
			if strings.Contains(dbPath, "data.db") {
				return dbx.Open("libsql", dbUrl)
			}

			// optionally for the logs (aka. pb_data/auxiliary.db) use the default local filesystem driver
			return core.DefaultDBConnect(dbPath)
		},
	})

	// any custom hooks or plugins...

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
