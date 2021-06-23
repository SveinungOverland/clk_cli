package clockify

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

const api_base = "https://api.clockify.me/api/v1/"

type Workspace struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Project struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Archived bool   `json:"archived"`
}

func Api(path string) string {
	return fmt.Sprint(api_base, path)
}

func Call(method, url string, body io.Reader) (resp *http.Response, err error) {
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
