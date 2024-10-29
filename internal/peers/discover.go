package peers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jackpal/bencode-go"
)

type Param struct {
	Tracker    string
	InfoHash   []byte
	Port       string
	FileLength int
}

func DiscoverPeers(req Param) (Peers, error) {
	params := url.Values{}
	params.Add("info_hash", string(req.InfoHash))
	params.Add("peer_id", "00112233445566778899")
	params.Add("port", "6881")
	params.Add("uploaded", "0")
	params.Add("downloaded", "0")
	params.Add("left", fmt.Sprint(req.FileLength))
	params.Add("compact", "1")

	// Construct the final URL with query parameters
	finalURL := fmt.Sprintf("%s?%s", req.Tracker, params.Encode())

	// Making the GET request
	response, err := http.Get(finalURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var trackerResponse TorrentPeers
	err = bencode.Unmarshal(response.Body, &trackerResponse)
	if err != nil {
		return nil, err
	}
	var peers = make([]string, 0)

	//Each peer is represented using 6 bytes. The first 4 bytes are the peer's IP address and the last 2 bytes are the peer's port number.
	for i := 0; i < len(trackerResponse.Peers); i += 6 {
		port := int(trackerResponse.Peers[i+4])<<8 + int(trackerResponse.Peers[i+5])
		peer := fmt.Sprintf("%d.%d.%d.%d:%s",
			trackerResponse.Peers[i],
			trackerResponse.Peers[i+1],
			trackerResponse.Peers[i+2],
			trackerResponse.Peers[i+3],
			strconv.Itoa(port))
		peers = append(peers, peer)
	}
	return peers, nil
}
