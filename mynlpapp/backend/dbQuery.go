package backend

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	_ "github.com/glebarez/go-sqlite"
)

const dbDirName = "sqldb"
const dbFileName = "QandA.db" // this should agree with the golang embed

// Setup to match the dir and filename

//go:embed sqldb/QandA.db
var database embed.FS

func TermPresentEmbed(term string) (string, bool, error) {
	dbData, err := fs.ReadFile(database, filepath.Join(dbDirName, dbFileName))
	if err != nil {
		return "", false, fmt.Errorf("No database present. If forgot to build, check the following:\n1 Go to ./setup and run \"createDB.go\"")
	}
	// Use the in-memory database driver of SQLite
	db, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		return "", false, fmt.Errorf("sql.Open failed: %s", err)
	}
	defer db.Close()
	// Create a temporary file in the system's temp directory
	tempFile, err := os.CreateTemp(os.TempDir(), "tempdb-*.sqlite")
	if err != nil {
		return "", false, fmt.Errorf("failed to create temp database file: %s", err)
	}
	defer os.Remove(tempFile.Name())

	// Write the embedded database data to the temporary file
	_, err = tempFile.Write(dbData)
	if err != nil {
		return "", false, fmt.Errorf("failed to write to temp database file: %s", err)
	}
	tempFile.Close()

	// Attach the temporary file to the in-memory database
	_, err = db.Exec("ATTACH DATABASE ? AS memdb", tempFile.Name())
	if err != nil {
		return "", false, err
	}

	var answer string
	err = db.QueryRow("SELECT answer FROM qat WHERE question = ?", term).Scan(&answer)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", false, nil
		}
		return "", false, err
	}
	return answer, true, nil
}
