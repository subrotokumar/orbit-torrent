package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type App struct {
	version string
	cmd     *cobra.Command
}

func NewApp(version string) *App {
	var rootCmd *cobra.Command = &cobra.Command{
		Short: "Orbit - A lightweight BitTorrent client",
		Long:  banner(),
	}
	return &App{cmd: rootCmd, version: version}
}

func (app *App) Run() {
	app.RegisterVersionCmd()
	app.RegisterDecodeCmd()
	app.RegisterInfoCmd()
	app.RegisterPeersCmd()
	app.RegisterHandshakeCmd()
	if err := app.cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}
