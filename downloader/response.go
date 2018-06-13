package downloader

import "net/http"

type Response struct {
	Request    *http.Request
	Body       string
	StatusCode uint
}
