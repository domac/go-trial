package heartbeat

import (
	"fmt"
	"net"
	"time"
)

func SendHeartBeat(add string) {
	go func() {
		for {
			Send(add, "over over over")
			time.Sleep(time.Second)
		}
	}()
}

func Send(addr string, msg string) {
	conn, err := net.DialTimeout("udp", addr, time.Second*2)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("send msg: ", msg)
	_, err = conn.Write([]byte(msg))
	return
}
