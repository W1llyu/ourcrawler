package downloader

import (
	"net/http"
	"net"
	"time"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	DefaultDownloader = NewNetHttpDownloader()
)

type Downloader interface {
	Download(request *http.Request) *Response
}

type NetHttpDownloader struct {
	HttpClient *http.Client
}

func NewNetHttpDownloader() *NetHttpDownloader {
	return &NetHttpDownloader{
		HttpClient: &http.Client{
			Transport: &http.Transport{
				Dial: func(netw, addr string) (net.Conn, error) {
					conn, err := net.DialTimeout(netw, addr, time.Second * 30)    // 建立连接超时
					if err != nil {
						return nil, err
					}
					conn.SetDeadline(time.Now().Add(time.Second * 30))   // 发送接收数据超时
					return conn, nil
				},
				ResponseHeaderTimeout: time.Second * 30,
			},
		},
	}
}

func (n *NetHttpDownloader) Download(request *http.Request) *Response {
	response, err := n.HttpClient.Do(request)
	if err != nil {
		failOnError(err, fmt.Sprintf("Failed to download %s", request.URL.String()))
	}
	body := ""
	resBody, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err == nil {
		body = string(resBody)
		infof("Download %s Ok", request.URL.String())
	}
	return &Response{
		Request: request,
		Body: body,
		StatusCode: uint(response.StatusCode),
	}
}

func infof(format string, a ...interface {}) {
	log.Printf("[DOWNLOADER][INFO] %s", fmt.Sprintf(format, a...))
}

func failOnError(err error, msg string) {
	log.Printf("[DOWNLOADER][ERROR] %s %s", msg, err)
}
