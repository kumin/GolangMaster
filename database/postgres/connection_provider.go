package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type ConnectionProvider struct{}

func GetPostgresConnection(optFns ...PostgresOptionFn) (*sql.DB, error) {
	opts := DefaultPostgresOption
	for _, optFn := range optFns {
		optFn(opts)
	}
	uri := BuildURI(opts)
	log.Print(uri)
	return sql.Open("postgres", uri)
}
