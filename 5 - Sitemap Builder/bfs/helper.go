package main

import (
	"encoding/json"
	"fmt"
)

func prettyPrint(d Link) {
	if res, err := json.MarshalIndent(d, "", "    "); err == nil {
		fmt.Println(string(res))
	}
}
