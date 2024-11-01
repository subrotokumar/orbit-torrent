package utility

import (
	"encoding/json"
	"fmt"
)

func DisplayAsJson(value any) {
	jsonResp, err := json.MarshalIndent(value, "", "    ")
	fmt.Println()
	if err == nil {
		fmt.Println(string(jsonResp))
	}
	fmt.Println()
}
