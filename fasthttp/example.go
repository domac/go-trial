package main

import (
	"log"

	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

func main() {
	// Perpare a client, which fetches webpages via HTTP proxy listening
	// on the localhost:8080.
	c := &fasthttp.HostClient{
		Addr:        "192.168.46.173",
		ReadTimeout: time.Millisecond * 100,
	}

	// Fetch google page via local proxy.
	statusCode, body, err := c.Get(nil, "http://test.goose.sysop.vipshop.com/vicuna/stat?item=400")
	if err != nil {
		log.Fatalf("Error when loading google page through local proxy: %s", err)
	}
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}
	useResponseBody(body)

}

func useResponseBody(body []byte) {
	// Do something with body :)
	fmt.Println(string(body))
}
