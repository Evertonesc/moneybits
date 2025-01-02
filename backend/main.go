package main

import "moneybits/core"

func main() {
	app := core.NewAppContainer()

	app.StartHTTPServer()
}
