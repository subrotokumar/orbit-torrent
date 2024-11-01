package app

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/subrotokumar/orbit-torrent/internal/bencode"
	"github.com/subrotokumar/orbit-torrent/internal/console"
)

func DecodeCommand(cmd *cobra.Command, args []string) {
	bencodedString, err := cmd.Flags().GetString("input")
	if err != nil {
		console.ErrorFatal(err.Error())
	}
	if bencodedString == "" {
		console.ErrorFatal("Invalid input")
	}
	decoded, _, err := bencode.Decode(bencodedString, 0)
	if err != nil {
		console.Error("Error decoding input")
		console.ErrorFatal(err.Error())
	}
	jsonOutput, _ := json.MarshalIndent(decoded, "", " ")
	console.Log(string(jsonOutput))
}

func (app *App) RegisterDecodeCmd() {
	cmd := &cobra.Command{
		Use:   "decode",
		Short: "Print information about given torrent file",
		Run:   DecodeCommand,
	}
	app.cmd.AddCommand(cmd)
	cmd.Flags().StringP("input", "i", "", "input bencoded string")

}
