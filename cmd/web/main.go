package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"snippetbox/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

// Add a snippets field to the application struct. This will allow us to make the SnippetModel object Available to our handlers.
type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "Http netwrok address")

	// Define a new command-line flag for the MYSQL DSN string.
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MYSQL data source name")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// To keep the main() function tidy I've put the code for creating a connection pool into the separate openDB() function below. We pass openDB() the DSN from the command-line flag.
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// We also defer a call to db.Close(), so that the connection pool is closed before the main() function exists.
	defer db.Close()

	//Initialize a models.SnippetModel instance containing the connec
	app := &application{
		logger: logger,
	}

	logger.Info("Starting server", "addr", *addr)

	// Because the error variable is now already declared in the code above, we need to use the assignment operator = here, instead of the := 'declare and assign' operator.
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//To verify everything is setup correctly we use this method to create a connection and check for any errors
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
