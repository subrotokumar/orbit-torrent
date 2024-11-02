package types

import (
	"crypto/sha1"
	"fmt"

	"github.com/jackpal/bencode-go"
	"github.com/subrotokumar/orbit-torrent/internal/console"
	"github.com/subrotokumar/orbit-torrent/internal/utility"
)

type Meta struct {
	Announce string   `bencode:"announce"`
	Info     MetaInfo `bencode:"info"`
}

type MetaInfo struct {
	Name        string `bencode:"name"`
	Length      int    `bencode:"length"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
}

func (m *Meta) GetInfoHash() []byte {
	h := sha1.New()
	if err := bencode.Marshal(h, m.Info); err != nil {
		console.ErrorFatal(err.Error())
	}
	return h.Sum(nil)
}

func (meta *Meta) GetPieceHashes() [][20]byte {
	hashLen := 20
	buf := []byte(meta.Info.Pieces)
	numHashes := len(buf) / hashLen
	hashes := make([][20]byte, numHashes)
	for i := 0; i < numHashes; i++ {
		copy(hashes[i][:], buf[i*hashLen:(i+1)*hashLen])
	}
	return hashes
}

func (m *Meta) DisplayAsJson() {
	utility.DisplayAsJson(m)
}

func (m *Meta) DisplayInfoHash() {
	fmt.Println("Tracker URL:", m.Announce)
	fmt.Println("Length:", m.Info.Length)

	infoHash := m.GetInfoHash()
	fmt.Printf("Info Hash: %x\n", infoHash)
}

func (m *Meta) DisplayPieceHash() {
	pieceHash := m.GetPieceHashes()
	fmt.Println("Piece Length:", m.Info.PieceLength)
	fmt.Println("Piece Hashes:")
	for _, value := range pieceHash {
		fmt.Printf("\t%x\n", value)
	}
}
