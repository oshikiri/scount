package main

import "flag"

func main() {
	var topnPrint = flag.Int("n", 10, "Print the top N items")
	var approximateCounting = flag.Bool("a", false, "Use approximate counting algorithm (default is naive counting)")
	var approximateCountingThreshold = flag.Float64("t", 1e-5, "theta of KSP algorithm")
	var quietMode = flag.Bool("q", false, "Turn on quiet mode")
	//	var separator = flag.String("s", " ", "[WIP] Separator for tuple counting")
	separator := " "
	var flushMilliSec = flag.Int64("f", 100, "Flush counting progress every X [msec]")
	flag.Parse()

	loop(separator, *approximateCounting, *approximateCountingThreshold, *topnPrint, *flushMilliSec, *quietMode)
}
