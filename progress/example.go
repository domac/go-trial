package main

import (
	"fmt"
	"github.com/gosuri/uiprogress"
	"github.com/gosuri/uiprogress/util/strutil"
	"math/rand"
	"sync"
	"time"
)

var tasks = []string{
	"127.0.0.1",
	"127.0.0.2",
	"127.0.0.3",
	"127.0.0.4",
	"127.0.0.5",
	"127.0.0.6",
	"127.0.0.7",
	"127.0.0.8",
	"127.0.0.9",
	"127.0.1.10",
	"127.0.1.1",
	"127.0.1.2",
	"127.0.1.3",
	"127.0.1.4",
	"127.0.1.5",
	"127.0.1.6",
	"127.0.1.7",
	"127.0.1.8",
	"127.0.1.9",
	"192.168.1.9",
}

var done = []string{}
var doneCount = 0

var sendChan chan string = make(chan string, len(tasks))
var bar *uiprogress.Bar

func main() {
	fmt.Println("apps: deployment started")
	p := uiprogress.New()
	p.Start()

	var wg sync.WaitGroup
	wg.Add(len(tasks))
	bar = p.AddBar(len(tasks) - 1).AppendCompleted()
	deploy("mysoft", &wg)
	wg.Wait()
	time.Sleep(1 * time.Second)
	fmt.Fprintln(p.Bypass(), "install finished task count:", doneCount)
}

func deploy(app string, wg *sync.WaitGroup) {

	var receiveChan chan string = make(chan string, len(tasks))

	bar.Width = 50
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		res := <-sendChan
		return strutil.Resize(app+" setup : "+res, 28)
	})

	//模拟处理

	rand.Seed(500)

	go func() {
		for _, s := range tasks {
			go func(server string) {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
				receiveChan <- server
			}(s)
		}
	}()

	//报告输出
	go func() {
		for {
			select {
			case s := <-receiveChan:
				//异步执行
				go doJob(s, wg)
			}
		}
	}()
}

func doJob(s string, wg *sync.WaitGroup) {
	time.Sleep(10 * time.Millisecond * time.Duration(rand.Intn(2000)))
	bar.Incr()
	sendChan <- s
	doneCount++
	wg.Done()
}
