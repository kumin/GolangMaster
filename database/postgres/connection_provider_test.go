package postgres

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:staticcheck
func TestPostgresConnection(t *testing.T) {
	assert := assert.New(t)
	db, err := GetPostgresConnection(
		WithUser("whatsapp"),
		WithPassword("whatsapp_pass"),
		WithDatabase("whatsapp"),
		WithHost("localhost:5432"),
		WithSSLMode("verify-ca"),
	)
	defer db.Close()
	assert.Nil(err)
	row, err := db.Query("SELECT full_name FROM whatsmeow_contacts")
	defer row.Close()
	assert.Nil(err)
	for row.Next() {
		var fullName sql.NullString
		err := row.Scan(&fullName)
		assert.Nil(err)
		fmt.Println(fullName.String)
	}
	assert.Nil(err)
}
