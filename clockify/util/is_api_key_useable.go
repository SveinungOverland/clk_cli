package util

import (
	"clk/clockify"
	"fmt"
)

func IsApiKeyUseable() bool {
	resp, err := clockify.Call("GET", clockify.Api("user"), nil)
	if err != nil {
		fmt.Printf("Error when validating api-key towards api: %s\n", err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false
	}

	return true
}
