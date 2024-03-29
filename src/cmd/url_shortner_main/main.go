package main

import (
	"urlshortner/src/pkg/api"
	"urlshortner/src/pkg/webserver"
)

func main() {
	api.Init()

	r := api.NewRouter()
	s := webserver.NewServer(8080)
	s.Start(r)
	s.WaitTillSignaled()
}

