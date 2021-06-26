package main

import "github.com/agstrc/yadb/internal/setup"

func main() {
	app := setup.App()
	app.Listen("0.0.0.0:8000")
}
