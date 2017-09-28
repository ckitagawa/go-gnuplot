package main

import (
	"fmt"
	"github.com/ckitagawa/go-gnuplot"
)

func main() {
	fname := ""
	persist := false
	debug := true

	p, err := gnuplot.NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	p.CheckedCmd("plot %f*x", 23.0)
	p.CheckedCmd("plot %f * cos(%f * x)", 32.0, -3.0)
	p.CheckedCmd("set terminal pdf")
	p.CheckedCmd("set output 'plot001.pdf'")
	p.CheckedCmd("replot")

	p.CheckedCmd("q")

	return
}
