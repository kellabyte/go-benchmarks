package calls

import (
	"plugin"
	"testing"

	"github.com/kellabyte/go-benchmarks/calls/asm"
)

var pluginNop func()

func init() {
	plug, err := plugin.Open("func.plugin.so")
	if err != nil {
		panic(err)
	}

	symNop, err := plug.Lookup("Nop")
	if err != nil {
		panic(err)
	}

	pluginNop = symNop.(func())
}

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

func BenchmarkPlugin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pluginNop()
		b.SetBytes(1)
	}
}

func BenchmarkAsm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		asm.Nop()
		b.SetBytes(1)
	}
}
