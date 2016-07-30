package main

import (
	"flag"
	"fmt"
	"github.com/phillihq/go-trial/heartbeat"
	"time"
)

var addr = flag.String("addr", "127.0.0.1:8888", "hb client addr")

func main() {
	flag.Parse()
	heartbeat.SendHeartBeat(*addr)
	for {
		time.Sleep(1e9)
	}
	fmt.Println("END")
}
