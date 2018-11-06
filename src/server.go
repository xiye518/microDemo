package main

import (
	micro "github.com/xiaoenai/tp-micro"
	tp "github.com/henrylee2cn/teleport"
)

// Arg arg
type Arg struct {
	A int
	B int `param:"<range:1:>"`
}

// P handler
type P struct {
	tp.CallCtx
}

// Divide divide API
func (p *P) Divide(arg *Arg) (int, *tp.Rerror) {
	return arg.A / arg.B, nil
}

func main() {
	srv := micro.NewServer(micro.SrvConfig{
		ListenAddress: ":9090",
	})
	srv.RouteCall(new(P))
	srv.ListenAndServe()
}