package main

import (
	"log"
	"time"

	sd "github.com/phillihq/go-trial/etcd_service_discovery"
)

func main() {
	m, err := sd.NewMaster("sd-test", []string{
		"http://192.168.139.134:2379",
		"http://192.168.139.134:2479",
		"http://192.168.139.134:2579",
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		log.Println("all ->", m.GetNodes())
		log.Println("all(strictly) ->", m.GetNodesStrictly())
		time.Sleep(time.Second * 2)
	}
}
