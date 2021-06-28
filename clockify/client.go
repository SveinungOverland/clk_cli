package clockify

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const api_base = "https://api.clockify.me/api/v1/"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	TimeZone string `json:"timeZone"`
}

type Workspace struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Project struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Archived bool   `json:"archived"`
}

type TimeEntry struct {
	ID           string `json:"id"`
	TimeInterval struct {
		Duration *string `json:"duration"`
		End      *string `json:"end"`
		Start    string  `json:"start"`
	} `json:"timeInterval"`
	Description string `json:"description"`
}

func (timeEntry TimeEntry) String() string {
	startTime, _ := time.Parse(time.RFC3339, timeEntry.TimeInterval.Start)
	endTime, _ := time.Parse(time.RFC3339, *timeEntry.TimeInterval.End)
	return fmt.Sprintf("[%s %s] %s", startTime.Local().Format("2006-01-02 15:04:05"), endTime.Sub(startTime), timeEntry.Description)
}

type UrlBuilder struct {
	url   []string
	query []string
}

/// Path should be of form 'api' with no prepending or trailing '/'
func (builder *UrlBuilder) Path(path string) *UrlBuilder {
	builder.url = append(builder.url, path)
	return builder
}

func (builder *UrlBuilder) PathVar(path, variable string) *UrlBuilder {
	builder.url = append(builder.url, path, variable)
	return builder
}

func (builder *UrlBuilder) Query(key, value string) *UrlBuilder {
	builder.query = append(builder.query, fmt.Sprint(key, "=", value))
	return builder
}

func (builder *UrlBuilder) Build() (url string) {
	url = strings.Join(builder.url, "/")
	if len(builder.query) > 0 {
		query := strings.Join(builder.query, "&")
		url = fmt.Sprint(url, "?", query)
	}
	return
}

/// Base should be in form https://domain.com/* <- with no trailing '/'
func NewUrlBuilder(base string) *UrlBuilder {
	return &UrlBuilder{url: []string{base}, query: []string{}}
}

func ApiBuilder() *UrlBuilder {
	return NewUrlBuilder(api_base)
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
