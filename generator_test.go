package uid64

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGen(t *testing.T) {
	// Just generate UID.
	g, err := NewGenerator(0)
	assert.Nil(t, err)
	i := 0
	for i < 1024 {
		_, err := g.Gen()
		assert.Nil(t, err)
		i++
	}
}

func TestNewGenerator(t *testing.T) {
	var err error
	// NewGenerator needs GeneratorID as int
	// the id must be in 0 ~ 3
	_, err = NewGenerator(0)
	assert.Nil(t, err)
	_, err = NewGenerator(1)
	assert.Nil(t, err)
	_, err = NewGenerator(2)
	assert.Nil(t, err)
	_, err = NewGenerator(3)
	assert.Nil(t, err)
}

func TestNewGeneratorFail(t *testing.T) {
	var err error
	// GeneratorID cannot be less than 0 or greater than 3
	_, err = NewGenerator(-1)
	assert.NotNil(t, err)
	_, err = NewGenerator(4)
	assert.NotNil(t, err)
}
