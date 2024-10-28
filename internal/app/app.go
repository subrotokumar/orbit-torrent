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
	switch command {
	case "decode":
		bencodedValue := os.Args[2]
		app.Decode(bencodedValue)
	case "info":
		path := os.Args[2]
		data, err := os.ReadFile(path)
		fmt.Println(string(data))
		if err != nil {
			fmt.Println("unable to read file")
			fmt.Println(err)
			return
		}
		app.Decode(string(data))
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
	jsonOutput, _ := json.Marshal(decoded)
	fmt.Println(string(jsonOutput))
}

func (app *App) Encode() {

}
