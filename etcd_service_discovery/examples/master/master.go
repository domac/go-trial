package main

import (
	"log"
	"time"

	sd "github.com/phillihq/go-trial/etcd_service_discovery"
)

func main() {

	m, err := sd.NewMaster("sd-test", []string{
		"http://192.168.139.134:2079",
		"http://192.168.139.134:2179",
		"http://192.168.139.134:2279",
		"http://192.168.139.134:2379",
		"http://192.168.139.134:2479",
	})
	if err != nil {
		log.Fatal(err)
	}
	for {

		if len(m.GetNodes()) > 0 {
			log.Printf("all -> %v \n", m.GetNodes())
		}

		if len(m.GetNodesStrictly()) > 0 {
			log.Printf("all(strictly) -> %v \n", m.GetNodesStrictly())
		}

		time.Sleep(time.Second * 2)
	}
}
