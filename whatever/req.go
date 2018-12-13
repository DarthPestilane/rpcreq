package whatever

import (
	"log"
)

type Arg struct {
	A int
	B int
}

type Result int

type Req int

func (*Req) Multiply(arg Arg, result *Result) error {
	log.Printf("Multiply %d with %d\n", arg.A, arg.B)
	*result = Result(arg.A * arg.B)
	return nil
}

func Method(m string) string {
	return "Req." + m
}
