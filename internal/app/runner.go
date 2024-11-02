package app

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net"
	"os"

	bencode "github.com/jackpal/bencode-go"
	"github.com/subrotokumar/orbit-torrent/internal/console"
	"github.com/subrotokumar/orbit-torrent/internal/types"
)

func GetTorrentMeta(file string) *types.Meta {
	if file == "" {
		console.ErrorFatal("Invalid file path input")
	}

	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("unable to read file")
		fmt.Println(err)
		return nil
	}

	var result types.Meta
	err = bencode.Unmarshal(bytes.NewReader(content), &result)
	if err != nil {
		fmt.Println("unable to unmarshal file")
		fmt.Println(err)
		return nil
	}

	return &result
}

func RunHandshake(peerAddress string, infoHash []byte) {
	conn, err := net.Dial("tcp", peerAddress)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	pstrlen := byte(19) // The length of the string "BitTorrent protocol"
	pstr := []byte("BitTorrent protocol")
	reserved := make([]byte, 8) // Eight zeros
	handshake := append([]byte{pstrlen}, pstr...)
	handshake = append(handshake, reserved...)
	handshake = append(handshake, infoHash...)
	handshake = append(handshake, []byte{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}...)
	_, err = conn.Write(handshake)
	buf := make([]byte, 68)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("failed:", err)
		return
	}
	fmt.Printf("Peer ID: %s\n", hex.EncodeToString(buf[48:]))
}
