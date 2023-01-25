package uid64

import (
	"database/sql"
	"sort"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestSQLInterface(t *testing.T) {
	db, _ := sql.Open("sqlite3", "file::memory:?cache=shared")
	db.Exec("CREATE TABLE IF NOT EXISTS user (id INT PRIMARY KEY, name VARCHAR(250));")

	// prepare data
	g, _ := NewGenerator(0)
	id, _ := g.GenDanger()
	name := "test"

	// Insertion Test for sql.Valuer
	stmt, _ := db.Prepare("INSERT INTO user VALUES(?,?);")
	defer stmt.Close()
	_, err := stmt.Exec(id, name)
	assert.Nil(t, err)

	// Selection Test for sql.Scanner
	rows, err := db.Query("SELECT id FROM user WHERE name = 'test'")
	assert.Nil(t, err)
	var selectedID UID
	rows.Next()
	err = rows.Scan(&selectedID)
	assert.Nil(t, err)

	// Compare 2 ids.
	assert.Equal(t, id, selectedID)
	assert.Equal(t, id.String(), selectedID.String())
}

func TestSortInterface(t *testing.T) {
	g, _ := NewGenerator(1)

	// Test IsSorted for sort.Interface
	check_sorted := make(UID64Slice, 0, 8*128)
	for i := 0; i < 256; i++ {
		// Sleep to generate uids at each milli sec for sequentiality.
		time.Sleep(1 * time.Millisecond)
		id, err := g.GenDanger()
		assert.Nil(t, err)
		check_sorted = append(check_sorted, id)
	}
	assert.True(t, sort.IsSorted(check_sorted))

	// Test Sort for sort.Interface
	check_sorted = make(UID64Slice, 0)
	for i := 0; i < 128; i++ {
		// Without time sleep, it ordered randomly thanks to /dev/urandom
		id, err := g.GenDanger()
		assert.Nil(t, err)
		check_sorted = append(check_sorted, id)
	}
	sort.Sort(check_sorted)
	assert.True(t, sort.IsSorted(check_sorted))
}
