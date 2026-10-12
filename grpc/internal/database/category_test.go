package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func newTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("sqlite3", t.TempDir()+"/test.sqlite")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = db.Close() })

	if err := Migrate(db); err != nil {
		t.Fatal(err)
	}

	return db
}

func TestCategoryCreateFindAndList(t *testing.T) {
	categoryStore := NewCategory(newTestDB(t))

	created, err := categoryStore.Create("gRPC", "Remote procedure calls")
	if err != nil {
		t.Fatal(err)
	}
	if created.ID == "" {
		t.Fatal("expected generated ID")
	}

	found, err := categoryStore.Find(created.ID)
	if err != nil {
		t.Fatal(err)
	}
	if found != created {
		t.Fatalf("Find() = %#v, want %#v", found, created)
	}

	categories, err := categoryStore.FindAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(categories) != 1 || categories[0] != created {
		t.Fatalf("FindAll() = %#v, want [%#v]", categories, created)
	}
}
