package database

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"

	"cloud.google.com/go/firestore"
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

var reUnique = regexp.MustCompile(`duplicate key value \(([a-z0-9_]+)\)=\('([^']+)'\) violates unique constraint "([a-z0-9_]+)"`)

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
	// try to match the regex to get more info
	if reUnique.MatchString(err.Message) {
		matches := reUnique.FindStringSubmatch(err.Message)

		column, value, constraint := matches[1], matches[2], matches[3]

		return terrors.BadRequest(fmt.Sprintf("bad_param.%s.unique", column), fmt.Sprintf("%s '%s' already exists", column, value), map[string]string{
			"key":        column,
			"value":      value,
			"constraint": constraint,
		})
	}

	return terrors.BadRequest("bad_param.unique", "Unable to create record; key already exists", nil)
}

func NotFound(ref *firestore.DocumentSnapshot) bool {
	if ref == nil {
		return false
	}

	return !ref.Exists()
}
