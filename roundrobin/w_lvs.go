package roundrobin

//节点结构
type WeightLvs struct {
	Node   interface{}
	Weight int
}

//lvs算法实现类
type WLVS struct {
	nodes []*WeightLvs
	n     int
	gcd   int //通用的权重因子
	maxW  int //最大权重
	i     int //被选择的次数
	cw    int //当前的权重值
}

//下次轮询事件
func (w *WLVS) Next() interface{} {
	if w.n == 0 {
		return nil
	}

	if w.n == 1 {
		return w.nodes[0].Node
	}

	for {
		w.i = (w.i + 1) % w.n
		if w.i == 0 {
			w.cw = w.cw - w.gcd
			if w.cw <= 0 {
				w.cw = w.maxW
				if w.cw == 0 {
					return nil
				}
			}
		}
		if w.nodes[w.i].Weight >= w.cw {
			return w.nodes[w.i].Node
		}
	}
}

//增加权重节点
func (w *WLVS) Add(node interface{}, weight int) {
	weighted := &WeightLvs{Node: node, Weight: weight}
	if weight > 0 {
		if w.gcd == 0 {
			w.gcd = weight
			w.maxW = weight
			w.i = -1
			w.cw = 0
		} else {
			w.gcd = gcd(w.gcd, weight)
			if w.maxW < weight {
				w.maxW = weight
			}
		}
	}
	w.nodes = append(w.nodes, weighted)
	w.n++
}

func gcd(x, y int) int {
	var t int
	for {
		t = (x % y)
		if t > 0 {
			x = y
			y = t
		} else {
			return y
		}
	}
}
func (w *WLVS) RemoveAll() {
	w.nodes = w.nodes[:0]
	w.n = 0
	w.gcd = 0
	w.maxW = 0
	w.i = -1
	w.cw = 0
}
func (w *WLVS) Reset() {
	w.i = -1
	w.cw = 0
}
