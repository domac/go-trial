package main

import (
	"fmt"
	"github.com/phillihq/go-trial/heartbeat"
)

func main() {

	doctor := heartbeat.NewDoctor(":8888")
	notify, err := doctor.Watch()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		host := <-notify
		fmt.Println(string(doctor.JSONMessage()))

		if !host.Alive {
			fmt.Printf("% is dead \n", host.Ip)
		}
	}
}
