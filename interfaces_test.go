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

	// Check 2 ids equal.
	assert.Equal(t, id, selectedID)
	assert.Equal(t, id.String(), selectedID.String())
}

func TestSortInterface(t *testing.T) {
	g, _ := NewGenerator(1)

	// Test IsSorted for sort.Interface
	checkSorted := make(UID64Slice, 0, 256)
	for i := 0; i < 256; i++ {
		// Sleep with each generation for a milli sec for make sequentially.
		time.Sleep(1 * time.Millisecond)
		id, err := g.GenDanger()
		assert.Nil(t, err)
		checkSorted = append(checkSorted, id)
	}
	assert.True(t, sort.IsSorted(checkSorted))

	// Test Sort for sort.Interface
	checkSorted = make(UID64Slice, 0, 128)
	for i := 0; i < 128; i++ {
		// Without a sleep, UIDs are ordered randomly thanks to /dev/urandom
		id, err := g.GenDanger()
		assert.Nil(t, err)
		checkSorted = append(checkSorted, id)
	}
	sort.Sort(checkSorted)
	assert.True(t, sort.IsSorted(checkSorted))
}
