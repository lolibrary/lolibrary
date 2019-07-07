package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	"github.com/monzo/terrors"
	_ "github.com/proullon/ramsql/driver"
)

const (
	defaultAddr   = "running-kitten-cockroachdb-public"
	defaultUser   = "service"
	defaultPort   = 26257
	defaultDB     = "lolibrary"
	defaultRootCA = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
)

// Connect takes a service and returns a database connection.
func Connect() *gorm.DB {
	dsn := getDSN()

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db
}

// NewMock spins up a copy of cockroachdb in insecure mode for the duration of this test.
func NewMock() *gorm.DB {
	db, err := sql.Open("postgres", "root@localhost:26257?sslmode=disable")
	if err != nil {
		panic(err)
	}

	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		panic(err)
	}

	// disable logging for tests. we're using ramsql, the index errors are going to get noisy.
	// TODO: figure out a better way to do in-memory testing, probably with cockroachdb itself!
	gormDB.LogMode(true)

	return gormDB
}

// getDSN gets the configured connection string for use with this database.
func getDSN() string {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = defaultDSN()
	}

	return dsn
}

// defaultDSN returns a default DSN that works as long as you're within a cluster pod.
// Service accounts should ideally be unique per-service, too.
func defaultDSN() string {
	// postgres://service@running-kitten-cockroachdb-public:26257/lolibrary
	connection := fmt.Sprintf("postgres://%s@%s:%d/%s",
		defaultUser,
		defaultAddr,
		defaultPort,
		defaultDB,
	)

	options := fmt.Sprintf("?ssl=true&sslmode=verify-full&sslrootcert=%s&sslkey=%s&sslcert=%s",
		defaultRootCA,
		"/certs/client.service.key",
		"/certs/client.service.crt",
	)

	return connection + options
}

// DuplicateRecord returns an error if the given error is a unique constraint violation from libpq.
func DuplicateRecord(err error) error {
	if err == nil {
		return nil
	}

	// try and cast to a pq error.
	if perr, ok := err.(*pq.Error); ok {
		if perr.Code == "23505" {
			return DuplicateRecordError(perr)
		}
	}

	return nil
}

// DuplicateRecordError casts a pq error to a terror with more details.
func DuplicateRecordError(err *pq.Error) *terrors.Error {
	return terrors.BadRequest(fmt.Sprintf("unique.%s", err.Column), "Key already exists for that table:", map[string]string{
		"column":     err.Column,
		"table":      err.Table,
		"constraint": err.Constraint,
		"hint":       err.Hint,
		"message":    err.Message,
		"detail":     err.Detail,
	})
}
