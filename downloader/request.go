package downloader

import (
	"bytes"
	"net/http"
	"net/url"
)

func BuildFormPostRequest(requestUrl string, params map[string]string) *http.Request {
	bodyForms := url.Values{}
	for key, value := range params {
		bodyForms.Add(key, value)
	}
	req, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer([]byte(bodyForms.Encode())))
	return req
}
