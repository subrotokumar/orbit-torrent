package app

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/subrotokumar/orbit-torrent/internal/bencode"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run() {
	command := os.Args[1]
	param := os.Args[2]
	switch command {
	case "decode":
		app.Decode(param)
	case "info":
		app.Info(param)
	case "peers":
		app.peers(param)
	default:
		fmt.Println("Unknown command: " + command)
	}
}

func (app *App) Decode(input string) {
	bencodedString := os.Args[2]
	decoded, _, err := bencode.Decode(bencodedString, 0)
	if err != nil {
		fmt.Println("Error decoding input")
		fmt.Println(err.Error())
		return
	}
	jsonOutput, _ := json.MarshalIndent(decoded, "", " ")
	fmt.Println(string(jsonOutput))
}
