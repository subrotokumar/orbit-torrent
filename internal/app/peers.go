package app

import (
	"bytes"
	"fmt"
	"os"

	bencode "github.com/jackpal/bencode-go"
	"github.com/subrotokumar/orbit-torrent/internal/peers"
	"github.com/subrotokumar/orbit-torrent/internal/types"
)

func (app *App) peers(path string) {
	content, err := os.ReadFile(path)
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
