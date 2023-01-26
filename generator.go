package uid64

import (
	"bufio"
	"crypto/rand"
	"errors"
	"io"
	"sync"
	"time"
)

// UID64 generator with generator-id
type Generator struct {
	lock      sync.Mutex
	timestamp int64     // Unix Timestamp Milli of a last generation
	counter   uint8     // Counter to increment in each milli time
	rng       io.Reader // Random Number Generator as io.Reader
	genid     uint8     // Generator ID: should be 0,1,2 or 3
}

// NewGenerator creates new *Generator with generator-id
// generator-id is IDentifier of Generator, for ditributed usage.
// It should be 0 <= id <= 3.
func NewGenerator(id int) (*Generator, error) {
	if id > 3 || id < 0 {
		return nil, errors.New("ID for Generator should be 0 ~ 3")
	}
	return &Generator{
		rng:   bufio.NewReaderSize(rand.Reader, 16),
		genid: uint8(id),
	}, nil
}

// Gen generates new UID, can be called thread safely.
func (g *Generator) Gen() (UID, error) {
	g.lock.Lock()
	defer g.lock.Unlock()
	return g.gen()
}

// GenDanger generates new UID without using Mutex. (thus, Not-threadsafe)
func (g *Generator) GenDanger() (UID, error) {
	return g.gen()
}

// ID shows generator's id
func (g *Generator) ID() int {
	return int(g.genid)
}

// SetID can set generator id to the generator.
func (g *Generator) SetID(newid int) error {
	if newid < 0 || 3 < newid {
		return errors.New("gen-id must be in 0 ~ 3")
	}
	g.genid = uint8(newid)
	return nil
}

func (g *Generator) gen() (UID, error) {
	now := time.Now().UnixMilli()
	if now == g.timestamp {
		if g.counter == 0xff {
			return 0, errors.New("exceed gen max ratio")
		}
		g.counter++
	} else {
		g.counter = 0
	}
	entropy, err := read(g.rng)
	if err != nil {
		return 0, err
	}
	return initUID(now, entropy, g.counter, g.genid), nil
}

func read(rng io.Reader) (uint8, error) {
	b := [1]byte{}
	_, err := rng.Read(b[:])
	return b[0], err
}
