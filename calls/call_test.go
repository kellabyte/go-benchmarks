package calls

import (
	"testing"

	"github.com/kellabyte/go-benchmarks/calls/asm"
)

func BenchmarkCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CNop()
		b.SetBytes(1)
	}
}

func BenchmarkGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Nop()
		b.SetBytes(1)
	}
}

func BenchmarkAsm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		asm.Nop()
		b.SetBytes(1)
	}
}
