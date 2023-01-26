package uid64

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var TestIDs = []struct {
	timestamp   int64
	entropy     uint8
	generatorID uint8
	counter     uint8
}{
	{time.Now().UnixMilli(), 0xff, 0, 0},
	{0, 0x0f, 0, 1},
	{1, 0xf0, 1, 2},
	{0x7ffffffffffffff, 255, 2, 3},
	{0x000000000000000, 0, 3, 4},
	{0x8fffffff0000000, 0, 0, 7},
	{0x10000000fffffff, 0, 1, 8},
	{0x01000000fffffff, 0, 2, 15},
	{0x00100000fff00ff, 0, 3, 16},
	{0x00010000fffff00, 0, 0, 31},
	{0x000010001234567, 1, 1, 32},
	{0x0000010089abcde, 255, 2, 63},
	{0x00000010fffffff, 255, 3, 63},
	{0x0000000100fffff, 15, 0, 63},
}

func TestInitUID(t *testing.T) {
	for _, fields := range TestIDs {
		c := fields

		// InitUID and prepare expected timestamp
		id, err := InitUID(c.timestamp, c.entropy, c.generatorID, c.counter)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}
		expectedTS := c.timestamp & 0x0000ffffffffffff

		// Confirm value with UID.methods
		assert.Equal(t, expectedTS, id.Timestamp())
		assert.Equal(t, c.entropy, id.Entropy())
		assert.Equal(t, c.counter, id.Counter())
		assert.Equal(t, c.generatorID, id.GeneratorID())
	}
}

func TestIntConversion(t *testing.T) {
	for _, fields := range TestIDs {
		f := fields

		// original UID
		// interger: originals' integer representaion
		// uid: recovered UID from integer
		original, _ := InitUID(f.timestamp, f.entropy, f.generatorID, f.counter)
		integer := original.ToInt()
		uid, err := FromInt(integer)
		assert.Nil(t, err)

		// Confirm restored from int is same to the original one.
		assert.Equal(t, integer, uid.ToInt())
		assert.Equal(t, original.Timestamp(), uid.Timestamp())
		assert.Equal(t, original.Entropy(), uid.Entropy())
		assert.Equal(t, original.Counter(), uid.Counter())
		assert.Equal(t, original.GeneratorID(), uid.GeneratorID())
	}
}

func TestStringConversion(t *testing.T) {
	for _, fields := range TestIDs {
		f := fields

		// prepare values
		original, _ := InitUID(f.timestamp, f.entropy, f.generatorID, f.counter)
		parsed, err := Parse(original.String())
		if !assert.Nil(t, err) {
			t.FailNow()
		}

		// Confirm restored form int is same to the original one.
		assert.Equal(t, original.String(), parsed.String())
		assert.Equal(t, original.Timestamp(), parsed.Timestamp())
		assert.Equal(t, original.Entropy(), parsed.Entropy())
		assert.Equal(t, original.Counter(), parsed.Counter())
		assert.Equal(t, original.GeneratorID(), parsed.GeneratorID())
	}
}
