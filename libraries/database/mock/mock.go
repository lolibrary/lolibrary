package mock

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/ory/dockertest"
	// "github.com/ory/dockertest/docker"
)

func spawnCockroachDB() (*sql.DB, func()) {
	var mockDB *sql.DB

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %v", err)
	}

	opts := &dockertest.RunOptions{
		Repository: "cockroachdb/cockroach",
		Tag:        "v19.1.2",
		Cmd: []string{
			"cockroach", "start", "--insecure",
		},
		// PortBindings: map[docker.Port][]docker.PortBinding{
		// 	"26257": {
		// 		{HostPort: "26257"},
		// 	},
		// },
	}

	resource, err := pool.RunWithOptions(opts)
	if err != nil {
		log.Fatalf("Could not start container: %v", err)
	}
	_ = resource.Expire(30)

	cleanup := func() {
		err := pool.Purge(resource)
		if err != nil {
			log.Fatalf("Failed to purge docker image: %v", err)
		}
	}

	endpoint := fmt.Sprintf("root@127.0.0.1:%s", resource.GetPort("26257/tcp"))

	if err := pool.Retry(func() error {
		db, err := sql.Open("postgres", endpoint)
		if err != nil {
			return err
		}

		err = db.Ping()
		if err == nil {
			mockDB = db
		}

		return err

	}); err != nil {
		cleanup()
		log.Fatalf("Failed to connect to cockroachdb: %v", err)
	}

	return mockDB, cleanup
}

func NewDatabase() (*gorm.DB, func()) {
	baseDB, cleanup := spawnCockroachDB()

	gormDB, err := gorm.Open("postgres", baseDB)
	if err != nil {
		cleanup()
		log.Fatalf("Failed to connect gorm: %v", err)
	}

	return gormDB, cleanup
}
