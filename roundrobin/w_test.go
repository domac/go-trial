package roundrobin

import "testing"

//测试nginx风格的rr
func TestNginxWeight(t *testing.T) {
	w := NewWeightedRR(RR_NGINX)
	w.Add("a", 4)
	w.Add("b", 2)
	w.Add("c", 1)

	results := make(map[string]int)

	for i := 0; i < 70; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if results["a"] != 40 || results["b"] != 20 || results["c"] != 10 {
		t.Error("the algorithm is wrong")
	}

	w.Reset()
	results = make(map[string]int)

	for i := 0; i < 70; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if results["a"] != 40 || results["b"] != 20 || results["c"] != 10 {
		t.Error("the algorithm is wrong")
	}

	w.RemoveAll()
	w.Add("a", 7)
	w.Add("b", 9)
	w.Add("c", 13)

	results = make(map[string]int)

	for i := 0; i < 29000; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if results["a"] != 7000 || results["b"] != 9000 || results["c"] != 13000 {
		t.Error("the algorithm is wrong")
	}
}

//测试lvs风格的rr
func TestLVSWeight(t *testing.T) {
	w := NewWeightedRR(RR_LVS)
	w.Add("a", 5)
	w.Add("b", 2)
	w.Add("c", 3)
	w.Add("d", 10)

	results := make(map[string]int)

	for i := 0; i < 200; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if results["a"] != 50 || results["b"] != 20 || results["c"] != 30 || results["d"] != 100 {
		t.Error("the algorithm is wrong")
	}

	w.Reset()
	results = make(map[string]int)

	for i := 0; i < 200; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if results["a"] != 50 || results["b"] != 20 || results["c"] != 30 || results["d"] != 100 {
		t.Error("the algorithm is wrong")
	}

	w.RemoveAll()
	w.Add("a", 7)
	w.Add("b", 9)
	w.Add("c", 13)
	w.Add("d", 29)

	results = make(map[string]int)

	for i := 0; i < 58000; i++ {
		s := w.Next().(string)
		results[s]++
	}

	if results["a"] != 7000 || results["b"] != 9000 || results["c"] != 13000 || results["d"] != 29000 {
		t.Error("the algorithm is wrong")
	}
}
