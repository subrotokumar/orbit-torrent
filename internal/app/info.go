package app

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/subrotokumar/orbit-torrent/internal/console"
	"github.com/subrotokumar/orbit-torrent/internal/peers"
)

func InfoCommandRun(file string) {
	result := GetTorrentMeta(file)

	result.DisplayAsJson()
	result.DisplayInfoHash()
	result.DisplayPieceHash()

	infoHash := result.GetInfoHash()
	peers, err := peers.DiscoverPeers(peers.Param{
		Tracker:    result.Announce,
		InfoHash:   infoHash,
		Port:       "6881",
		FileLength: result.Info.Length,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Peers:")
	for _, val := range peers {
		fmt.Println("\t", val)
	}
}

func (app *App) RegisterInfoCmd() {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Get info about torrent file",
		Run: func(cmd *cobra.Command, args []string) {
			file, err := cmd.Flags().GetString("file")
			if err != nil {
				console.ErrorFatal(err.Error())
			}
			InfoCommandRun(file)
		},
	}
	app.cmd.AddCommand(cmd)
	cmd.Flags().StringP("file", "f", "", "Path of the torrent file")
}
