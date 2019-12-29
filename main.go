package main

import "flag"

func main() {
	var topnPrint = flag.Int("n", 10, "Print the top N items")
	var approximateCounting = flag.Bool("a", false, "Use approximate counting algorithm (default is naive counting)")
	var approximateCountingEpsilon = flag.Float64("ae", 1e-5, "Epsilon of lossy counting algorithm")
	var approximateCountingSupport = flag.Float64("as", 2e-5, "Support of lossy counting algorithm")
	var quietMode = flag.Bool("q", false, "Suppress a progress report")
	var flushMilliSec = flag.Int64("f", 200, "Flush counting progress every X [msec]")
	flag.Parse()

	loop(*approximateCounting, *approximateCountingEpsilon, *approximateCountingSupport, *topnPrint, *flushMilliSec, *quietMode)
}
