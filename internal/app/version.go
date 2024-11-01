package app

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/subrotokumar/orbit-torrent/internal/console"
	"github.com/subrotokumar/orbit-torrent/pkg/styles"
)

func banner() string {
	banner := `                                                                                                                                                 
     .oooooo.              .o8          o8o       .   
    d8P'  'Y8b             "888         '"'     .o8   
   888      888  oooo d8b   888oooo.   oooo   .o888oo 
   888      888  '888""8P   d88' '88b  '888     888   
   888      888   888       888   888   888     888   
   '88b    d88'   888       888   888   888     888 . 
    'Y8bood8P'   d888b      'Y8bod8P'  o888o    "888" 
	
   A lightweight BitTorrent client written in Go designed for command-line interface (CLI) enthusiasts.
	`
	return banner
}

func (app *App) RegisterVersionCmd() {
	app.cmd.AddCommand(&cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			console.Log("Orbit Torrent CLI")
			fmt.Println("Version:", styles.TextYellow.Render(app.version))
		},
	})
}
