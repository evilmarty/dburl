package dburl

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"testing"
)

func Test_Open(t *testing.T) {
	cases := []string{
		"mysql://user:pass@host/dbname",
		"postgres://user:pass@host/dbname",
	}

	for _, url := range cases {
		db, err := Open(url)
		if err != nil {
			t.Error(err)
		} else if db == nil {
			t.Fatal("Expected *sql.DB but received nil")
		}
	}
}
