package console

import (
	"fmt"
	"os"
	"strings"

	"github.com/subrotokumar/orbit-torrent/pkg/styles"
)

func Error(a ...string) {
	message := strings.Join(a, " ")
	fmt.Println(styles.TextRed.Render(message))
}

func ErrorFatal(a ...string) {
	message := strings.Join(a, " ")
	fmt.Println(styles.TextRed.Render(message))
	os.Exit(0)
}

func Log(a ...string) {
	message := strings.Join(a, " ")
	fmt.Println(styles.TextVoilet.Render(message))
}
