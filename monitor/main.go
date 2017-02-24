package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/vrischmann/go-metrics-influxdb"
)

//使用内存
func useMemory() {
	a := rand.Intn(100)
	m := make([]int, a*1e6)

	d := time.Second * time.Duration(rand.Intn(15))
	fmt.Printf("Generated something with size %dMB, sleeping for %s\n", len(m)/1e6, d)
	time.Sleep(d)
	fmt.Printf("Done using %dMB\n", len(m)/1e6)
}

func main() {

	r := metrics.NewRegistry()
	metrics.RegisterDebugGCStats(r)
	metrics.RegisterRuntimeMemStats(r)

	go metrics.CaptureDebugGCStats(r, time.Second*5)
	go metrics.CaptureRuntimeMemStats(r, time.Second*5)

	go influxdb.InfluxDB(
		r,                         // metrics registry
		time.Second*5,             // 时间间隔
		"http://10.197.32.2:8086", // InfluxDB url
		"helloWorld",              // InfluxDB 数据库名
		"",                        // InfluxDB user
		"",                        // InfluxDB password
	)

	time.Sleep(time.Second * 10)

	for {
		go useMemory()
		time.Sleep(time.Duration(rand.Intn(15)) * time.Second)
	}
}
