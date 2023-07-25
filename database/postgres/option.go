package postgres

import "fmt"

type PostgresOption struct {
	User     string
	Password string
	Database string
	Host     string
	Post     uint32
	SSLMode  string
}

var DefaultPostgresOption = &PostgresOption{}

type PostgresOptionFn func(opt *PostgresOption)

func WithUser(user string) PostgresOptionFn {
	return func(opt *PostgresOption) {
		opt.User = user
	}
}

func WithPassword(password string) PostgresOptionFn {
	return func(opt *PostgresOption) {
		opt.Password = password
	}
}

func WithDatabase(database string) PostgresOptionFn {
	return func(opt *PostgresOption) {
		opt.Database = database
	}
}

func WithHost(host string) PostgresOptionFn {
	return func(opt *PostgresOption) {
		opt.Host = host
	}
}

func WithSSLMode(mode string) PostgresOptionFn {
	return func(opt *PostgresOption) {
		opt.SSLMode = mode
	}
}

func BuildURI(opts *PostgresOption) string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		opts.User, opts.Password, opts.Host, opts.Database, opts.SSLMode)
}
