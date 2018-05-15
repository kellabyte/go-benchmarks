package calls

//void nop() { }
import "C"

//go:noinline
func Nop() {}

func CNop() { C.nop() }
