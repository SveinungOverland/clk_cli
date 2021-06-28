package util

import (
	"clk/clockify"
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
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

	var user clockify.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return false
	}

	viper.Set("user_id", user.ID)
	viper.Set("user_name", user.Name)
	viper.Set("user_timezone", user.TimeZone)

	err = viper.WriteConfig()
	if err != nil {
		// TODO, potentially do some error handling
		return false
	}

	return true
}
