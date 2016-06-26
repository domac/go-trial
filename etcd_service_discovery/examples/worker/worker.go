package main

import (
	"flag"
	"fmt"
	sd "github.com/phillihq/go-trial/etcd_service_discovery"
	"log"
	"time"
)

func main() {
	name := flag.String("name", fmt.Sprintf("%d", time.Now().Unix()), "des")
	extInfo := "lhq-demo..."

	flag.Parse()
	w, err := sd.NewWorker("sd-test", *name, extInfo, []string{
		"http://192.168.139.134:2379",
		"http://192.168.139.134:2479",
		"http://192.168.139.134:2579",
	})
	if err != nil {
		log.Fatal(err)
	}
	w.Register()
	log.Println("name ->", *name, "extInfo ->", extInfo)

	go func() {
		time.Sleep(time.Second * 20)
		w.Unregister()
	}()

	for {
		log.Println("isActive ->", w.IsActive())
		log.Println("isStop ->", w.IsStop())
		time.Sleep(time.Second * 2)
		//服务退出
		if w.IsStop() {
			return
		}
	}
}
