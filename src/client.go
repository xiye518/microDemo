package main

import (
	micro "github.com/xiaoenai/tp-micro"
	tp "github.com/henrylee2cn/teleport"
)

func main() {
	cli := micro.NewClient(
		micro.CliConfig{},
		micro.NewStaticLinker(":9090"),
	)
	defer cli.Close()
	
	type Arg struct {
		A int
		B int
	}
	
	var result int
	rerr := cli.Call("/p/divide", &Arg{
		A: 10,
		B: 2,
	}, &result).Rerror()
	if rerr != nil {
		tp.Fatalf("%v", rerr)
	}
	tp.Infof("10/2=%d", result)
	rerr = cli.Call("/p/divide", &Arg{
		A: 10,
		B: 0,
	}, &result).Rerror()
	if rerr == nil {
		tp.Fatalf("%v", rerr)
	}
	tp.Infof("test binding error: ok: %v", rerr)
}