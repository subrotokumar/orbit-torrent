package main

import (
	"fmt"

	"github.com/subrotokumar/orbit-torrent/internal/app"
)

func main() {
	a := `�v�z*����kg&���-n"u��vfVsn���R��5�`
	for _, v := range a {
		fmt.Println(string(v))
	}
	fmt.Println(len(a))
	app := &app.App{}
	// app.Banner()
	app.Run()
}
