package uid64

import (
	"testing"

	"github.com/google/uuid"
)

func BenchmarkGoogleUUIDNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		uuid.New()
	}
}

func BenchmarkUID64Gen(b *testing.B) {
	g, _ := NewGenerator(0)
	for n := 0; n < b.N; n++ {
		g.Gen()
	}
}

func BenchmarkUID64GenDanger(b *testing.B) {
	g, _ := NewGenerator(0)
	for n := 0; n < b.N; n++ {
		g.GenDanger()
	}
}
