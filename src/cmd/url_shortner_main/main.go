package main

import (
	"flag"
	"urlshortner/src/pkg/api"
	"urlshortner/src/pkg/webserver"
)

func main() {
	port := flag.Int("port", 8080, "Port to run the server on")
	range_min := flag.Int64("range_min", 100000, "Min value of range")
	range_max := flag.Int64("range_max", 2500000, "Max value of range")
	flag.Parse()
	api.Init(*range_min, *range_max)
	r := api.NewRouter()
	s := webserver.NewServer(*port)
	s.Start(r)
	s.WaitTillSignaled()
}
