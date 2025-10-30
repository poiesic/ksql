package kduckdb

import (
	"database/sql"
	"io"
	"testing"

	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/sqldialect"
)

func TestAdapter(t *testing.T) {
	ksql.RunTestsForAdapter(t, "duckdb", sqldialect.DuckDbDialect{}, "/tmp/ksql.db", func(t *testing.T) (ksql.DBAdapter, io.Closer) {
		db, err := sql.Open("sqlite3", "/tmp/ksql.db")
		if err != nil {
			t.Fatal(err.Error())
		}
		return SQLAdapter{db}, db
	})
}
