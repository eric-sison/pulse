package main

import (
	i "github.com/eric-sison/pulse/internal"
)

func main() {
	app := i.CreateApp()

	app.StartServer()
}
