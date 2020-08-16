package storage

import (
	"fmt"
	"strings"
	"testing"
)

func TestStorage(t *testing.T, databaseURL string) (*Storage, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseUrl = databaseURL
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
