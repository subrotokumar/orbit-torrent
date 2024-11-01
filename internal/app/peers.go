package app

import (
	"bytes"
	"fmt"
	"os"

	bencode "github.com/jackpal/bencode-go"
	"github.com/spf13/cobra"
	"github.com/subrotokumar/orbit-torrent/internal/console"
	"github.com/subrotokumar/orbit-torrent/internal/peers"
	"github.com/subrotokumar/orbit-torrent/internal/types"
)

func PeersCommandRun(cmd *cobra.Command, args []string) {
	file, err := cmd.Flags().GetString("file")
	if err != nil {
		console.ErrorFatal(err.Error())
	}
	if file == "" {
		console.ErrorFatal("Invalid file path input")
	}
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("unable to read file")
		fmt.Println(err)
		return
	}

	var result types.Meta
	err = bencode.Unmarshal(bytes.NewReader(content), &result)
	if err != nil {
		fmt.Println("unable to unmarshal file")
		fmt.Println(err)
		return
	}
	infoHash, err := result.GetInfoHash()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
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

func (app *App) RegisterPeersCmd() {
	cmd := &cobra.Command{
		Use:   "peers",
		Short: "Get peers into of torrent file",
		Run:   PeersCommandRun,
	}
	app.cmd.AddCommand(cmd)
	cmd.Flags().StringP("file", "f", "", "path of the torrent file")
}
