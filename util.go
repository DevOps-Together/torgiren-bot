package main

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(data interface{}) (string, error) {
	var p []byte
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {

		return "", err
	}
	return fmt.Sprintf("%s", p), nil
}
