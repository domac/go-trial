package roundrobin

// RR: 基于 权重round robin算法的接口
type RR interface {
	Next() interface{}
	Add(node interface{}, weight int)
	RemoveAll()
	Reset()
}

const (
	RR_NGINX = 0 //Nginx算法
	RR_LVS   = 1 //LVS算法
)

//算法实现工厂类
func NewWeightedRR(rtype int) RR {
	if rtype == RR_NGINX {
		return &WNGINX{}
	} else if rtype == RR_LVS {
		return &WLVS{}
	}
	return nil
}
