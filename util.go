package main

import (
	"encoding/json"
)

func PrettyPrint(data interface{}) (string, error) {
	var p []byte
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {

		return "", err
	}
	return string(p[:]), nil
}
