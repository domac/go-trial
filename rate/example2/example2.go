package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	rate  int
	begin time.Time
	count int
	lock  sync.Mutex
}

func (self *RateLimiter) Limit() bool {
	result := true
	self.lock.Lock()

	if self.count == self.rate {

		if time.Now().Sub(self.begin) >= time.Second {
			self.begin = time.Now()
			self.count = 0
		} else {
			result = false
		}
	} else {
		self.count++
	}
	self.lock.Unlock()

	return result
}

func (self *RateLimiter) SetRate(r int) {
	self.rate = r
	self.begin = time.Now()
}

func (self *RateLimiter) GetRate() int {
	return self.rate
}

func main() {

	var wg sync.WaitGroup
	var limiter RateLimiter
	limiter.SetRate(3)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			if limiter.Limit() {
				fmt.Println("Go function")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
