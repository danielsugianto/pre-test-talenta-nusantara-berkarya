package main

import "github.com/danielsugianto/pre-test/test-talenta-nusantara-berkarya/routes"

func main() {
	e := routes.New()

	// start the server, and log if it fails
	e.Start(":8000")
}
