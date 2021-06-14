package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

const api_base = "https://api.clockify.me/api/v1/"

func api(path string) string {
	return fmt.Sprint(api_base, path)
}

type Workspace struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetWorkSpaces() ([]Workspace, error) {
	resp, err := call("GET", api("workspaces"), nil)
	if err != nil {
		return []Workspace{}, err
	}
	var workspaces []Workspace
	err = json.NewDecoder(resp.Body).Decode(&workspaces)
	if err != nil {
		fmt.Printf(err.Error())
		return []Workspace{}, err
	}

	return workspaces, nil
}

func IsApiKeyUseable() bool {
	resp, err := call("GET", api("user"), nil)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false
	}

	response_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Unvalid response")
	}

	fmt.Println(string(response_bytes))

	return true
}

func call(method, url string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	req.Header.Set("X-Api-Key", viper.GetString("api_key"))
	req.Header.Set("Content-Type", "application/json")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	return
}
