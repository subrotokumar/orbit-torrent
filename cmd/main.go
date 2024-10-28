package main

import (
	"github.com/subrotokumar/orbit-torrent/internal/app"
)

func main() {
	app := &app.App{}
	app.Logo()
	app.Run()
}
