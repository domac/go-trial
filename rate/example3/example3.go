package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	rate       int
	interval   time.Duration
	lastAction time.Time
	lock       sync.Mutex
}

func (self *RateLimiter) Limit() bool {
	result := true

	for {
		self.lock.Lock()
		if time.Now().Sub(self.lastAction) > self.interval {
			self.lastAction = time.Now()
			result = true
		}
		self.lock.Unlock()
		if result {
			return result
		}
		time.Sleep(self.interval)
	}
}

//SetRate 设置Rate
func (self *RateLimiter) SetRate(r int) {
	self.rate = r
	self.interval = time.Microsecond * time.Duration(1000*1000/self.rate)
}

//GetRate 获取Rate
func (self *RateLimiter) GetRate() int {
	return self.rate
}

func main() {
	var wg sync.WaitGroup
	var limiter RateLimiter
	limiter.SetRate(3)

	b := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			if limiter.Limit() {
				fmt.Println("Got it!")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(time.Since(b))
}
