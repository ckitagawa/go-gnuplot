go-gnuplot 
[![GoDoc](https://godoc.org/github.com/ckitagawa/go-gnuplot?status.svg)](http://godoc.org/github.com/ckitagawa/go-gnuplot)
[![Build Status](https://travis-ci.org/ckitagawa/go-gnuplot.svg?branch=master)](https://travis-ci.org/ckitagawa/go-gnuplot)
[![Go Report Card](https://goreportcard.com/badge/github.com/ckitagawa/go-gnuplot)](https://goreportcard.com/report/github.com/ckitagawa/go-gnuplot)
==========

Simple-minded functions to work with ``gnuplot``.
``go-gnuplot`` runs ``gnuplot`` as a subprocess and pushes commands
via the ``STDIN`` of that subprocess.

See http://www.gnuplot.info for more informations on the
exact semantics of these commands.

This is a fork of
[sbinet/go-gnuplot](https://www.github.com/sbinet/go-gnuplot). The fork is
motivated by the lack of maintenance to the original repo and the need for 3D
and quiver plotting in scientific computing not yet provided by native Golang
plotting libraries. This version will aim to extend upon and modernize the
original implementation without changing the base API so that this can remain a
drop-in replacement. See the issues for specific planned changes or if you want
to see a change file an issue against the repo.

Installation
------------

The ``go-gnuplot`` package is ``go get`` installable:

```sh
$ go get github.com/ckitagawa/go-gnuplot
```

Example
--------

```go
package main

import (
  "github.com/ckitagawa/go-gnuplot"
  "fmt"
)

func main() {
	fname := ""
	persist := false
	debug := true

	p,err := gnuplot.NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	p.PlotX([]float64{0,1,2,3,4,5,6,7,8,9,10}, "some data")
	p.CheckedCmd("set terminal pdf")
	p.CheckedCmd("set output 'plot002.pdf'")
	p.CheckedCmd("replot")

	p.CheckedCmd("q")
	return
}
```

![plot-t-002](https://github.com/ckitagawa/go-gnuplot/raw/master/examples/imgs/plot002.png)


Motivation
----------

[Gonum](https://github.com/gonum/gonum) now makes it feasible to author
scientific and engineering computing code in Go. This is great news because it
enables more efficient, compiled and type-safe code that is not nearly as easy
to achieve in Python. Furthermore, it is less of a headache to get a Go program
running than C or C++ and is usually more portable. Moreover, it has great
concurrency primitives which are a real boon to doing numeric computing more
efficiently.

That being said, this project is slightly counter-intuitive in the sense it
limits portability to platforms supporting gnuplot and currently doesn't play
nice with goroutines. Additionally, native graphing packages like
[wcharczuk/go-chart](https://github.com/wcharczuk/go-chart) and
[gonum/plot](https://github.com/gonum/plot) already exist; however, none of
them have mature 3D plotting capabilities nor strong support for quivers and
streamlines. Unfortunately, these are features required in engineering and
scientific computing on a fairly regular basis. As a result, this package needs
to live on a bit longer to service those in need of these specialized charts.
That said, it is my sincere hope that this library will not be necessary for
much longer.

Plans
----

- Add unit tests, a challenge which will probably require mocking out gnuplot
- Extend the original API to allow saving to files without `Cmd` or
  `CheckedCmd`. Also create methods for common actions such as `replot`.
- Support gnuplot > 5 as the original implementation was written for gnuplot 4 

