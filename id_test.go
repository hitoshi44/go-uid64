package uid64

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var TestIDs = []struct {
	timestamp   int64
	entropy     uint8
	counter     uint8
	generatorID uint8
}{
	{time.Now().Unix(), 0xff, 0, 0},
	{0, 0x0f, 0, 0},
	{1, 0xf0, 0, 0},
	{0xfffffffffffffff, 255, 255, 255},
	{0x000000000000000, 0, 0, 0},
	{0xffffffff0000000, 0, 0, 0},
	{0x10000000fffffff, 0, 0, 0},
	{0x01000000fffffff, 0, 0, 0},
	{0x00100000fff00ff, 0, 0, 1},
	{0x00010000fffff00, 0, 1, 1},
	{0x000010001234567, 1, 1, 1},
	{0x0000010089abcde, 255, 255, 255},
	{0x00000010fffffff, 255, 255, 255},
	{0x0000000100fffff, 15, 15, 15},
}

func TestInitID(t *testing.T) {
	for _, fields := range TestIDs {
		c := fields
		id := InitUID(c.timestamp, c.entropy, c.counter, c.generatorID)
		expectedTS := c.timestamp & 0x000000ffffffffff
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
		original := InitUID(f.timestamp, f.entropy, f.counter, f.generatorID)

		integer := original.ToInt()
		uid := FromInt(integer)

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

		original := InitUID(f.timestamp, f.entropy, f.counter, f.generatorID)
		parsed, err := Parse(original.String())

		if !assert.Nil(t, err) {
			t.FailNow()
		}
		assert.Equal(t, original.String(), parsed.String())
		assert.Equal(t, original.Timestamp(), parsed.Timestamp())
		assert.Equal(t, original.Entropy(), parsed.Entropy())
		assert.Equal(t, original.Counter(), parsed.Counter())
		assert.Equal(t, original.GeneratorID(), parsed.GeneratorID())
	}
}
