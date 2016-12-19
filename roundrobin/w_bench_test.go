package roundrobin

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func BenchmarkW1_Next(b *testing.B) {
	b.ReportAllocs()
	rand.Seed(time.Now().UnixNano())
	w := NewWeightedRR(RR_NGINX)
	for i := 0; i < 10; i++ {
		w.Add("server"+strconv.Itoa(i), rand.Intn(100))
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w.Next()
	}
}

func BenchmarkW2_Next(b *testing.B) {
	b.ReportAllocs()
	rand.Seed(time.Now().UnixNano())
	w := NewWeightedRR(RR_LVS)
	for i := 0; i < 10; i++ {
		w.Add("server"+strconv.Itoa(i), rand.Intn(100))
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w.Next()
	}
}
