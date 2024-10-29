package utility

import (
	"encoding/json"
	"fmt"
)

func DisplayAsJson(value any) {
	jsonResp, err := json.MarshalIndent(value, "", "    ")
	if err == nil {
		fmt.Println(string(jsonResp))
	}
}
