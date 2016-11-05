package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var lasttime int64 = time.Now().UnixNano()
	time.Sleep(5 * time.Second)
	now := time.Now()

	cur := time.Unix(0, atomic.LoadInt64(&lasttime))

	subtime := now.Sub(cur)

	fmt.Printf("cur: %v \n", cur)
	fmt.Printf("now: %v \n", now)
	fmt.Printf("sub: %v \n", subtime)
}
