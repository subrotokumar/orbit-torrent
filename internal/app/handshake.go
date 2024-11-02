package app

import (
	"github.com/spf13/cobra"
	"github.com/subrotokumar/orbit-torrent/internal/console"
)

func HandshakeCommandRun(file string, peerAddress string) {
	result := GetTorrentMeta(file)
	infoHash := result.GetInfoHash()
	RunHandshake(peerAddress, infoHash)
}

func (app *App) RegisterHandshakeCmd() {
	cmd := &cobra.Command{
		Use:   "handshake",
		Short: "Get peers into of torrent file",
		Run: func(cmd *cobra.Command, args []string) {
			file, err := cmd.Flags().GetString("file")
			if err != nil {
				console.ErrorFatal(err.Error())
			}
			peerAddress, err := cmd.Flags().GetString("peerAddress")
			if err != nil {
				console.ErrorFatal(err.Error())
			}
			HandshakeCommandRun(file, peerAddress)
		},
	}
	app.cmd.AddCommand(cmd)
	cmd.Flags().StringP("file", "f", "", "Path of the torrent file")
	cmd.Flags().StringP("peerAddress", "p", "", "Peer ID")

}
