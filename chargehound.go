// Package chargehound contains Go bindings for the Chargehound API.
package chargehound

import (
	"net/http"
	"time"
)

const (
	basepath = "/v1/"
	host     = "api.chargehound.com"
	protocol = "https://"
	version  = "1.1.0"

	defaultHTTPTimeout = 60 * time.Second
)

type Client struct {
	// The Chargehound API key used to interact with the API.
	ApiKey string
	// The API host.
	Host string
	// The API scheme.
	Protocol string
	// The API base path.
	Basepath string
	// The client version.
	Version string
	// The client http timeout.
	HTTPClient *http.Client
	// The disputes resource.
	Disputes *Disputes
}

// Creates a new chargehound client with the specified api key and the default configuration.
func New(key string) *Client {

	ch := Client{
		ApiKey:     key,
		Basepath:   basepath,
		Host:       host,
		HTTPClient: &http.Client{Timeout: defaultHTTPTimeout},
		Protocol:   protocol,
		Version:    version,
	}

	ch.Disputes = &Disputes{client: &ch}

	return &ch
}
