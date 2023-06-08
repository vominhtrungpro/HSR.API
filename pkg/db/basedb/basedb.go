package basedb

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/friendsofgo/errors"
	_ "github.com/lib/pq"
	"github.com/vominhtrungpro/config"
)

func Connect(driverName string, c config.MySQLConfig) (*sql.DB, error) {
	// database` host is container which already defined in docker compose
	conn, err := sql.Open(driverName, c.MYSQLURI)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)

		return nil, errors.WithStack(fmt.Errorf("connect DB failed. Err: %w", err))
	}

	// Ping the db to check connection with context timeout
	if err = conn.Ping(); err != nil {

		return nil, errors.WithStack(err)
	}

	fmt.Fprintln(os.Stderr, "connect to DB successfully")

	return conn, nil
}
