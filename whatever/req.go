package whatever

import (
	"log"
)

type Arg struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Result int

type Req struct{}

func (*Req) Multiply(arg Arg, result *Result) error {
	log.Printf("Multiply %d with %d", arg.A, arg.B)
	*result = Result(arg.A * arg.B)
	return nil
}
