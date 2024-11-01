package main

import (
	"github.com/subrotokumar/orbit-torrent/internal/app"
)

func main() {
	app := app.NewApp("v0.0.1")
	app.Run()
}
