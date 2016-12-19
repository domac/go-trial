package roundrobin

//节点结构
type WeightNginx struct {
	Node            interface{}
	Weight          int
	CurrentWeight   int
	EffectiveWeight int
}

func (ww *WeightNginx) fail() {
	ww.EffectiveWeight -= ww.Weight
	if ww.EffectiveWeight < 0 {
		ww.EffectiveWeight = 0
	}
}

//nginx算法实现类
type WNGINX struct {
	nodes []*WeightNginx
	n     int
}

//增加权重节点
func (w *WNGINX) Add(node interface{}, weight int) {
	weighted := &WeightNginx{
		Node:            node,
		Weight:          weight,
		EffectiveWeight: weight}
	w.nodes = append(w.nodes, weighted)
	w.n++
}

func (w *WNGINX) RemoveAll() {
	w.nodes = w.nodes[:0]
	w.n = 0
}

//下次轮询事件
func (w *WNGINX) Next() interface{} {
	if w.n == 0 {
		return nil
	}
	if w.n == 1 {
		return w.nodes[0].Node
	}

	return nextWeightedNode(w.nodes).Node
}

func nextWeightedNode(nodes []*WeightNginx) (best *WeightNginx) {
	total := 0

	for i := 0; i < len(nodes); i++ {
		w := nodes[i]

		if w == nil {
			continue
		}

		w.CurrentWeight += w.EffectiveWeight
		total += w.EffectiveWeight
		if w.EffectiveWeight < w.Weight {
			w.EffectiveWeight++
		}

		if best == nil || w.CurrentWeight > best.CurrentWeight {
			best = w
		}
	}

	if best == nil {
		return nil
	}
	best.CurrentWeight -= total
	return best
}

func (w *WNGINX) Reset() {
	for _, s := range w.nodes {
		s.EffectiveWeight = s.Weight
		s.CurrentWeight = 0
	}
}
