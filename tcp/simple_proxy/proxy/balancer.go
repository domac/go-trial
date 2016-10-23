package proxy

import (
	"math/rand"
	"net"
	"stathat.com/c/consistent"
	"time"
)

// BackendSvr Type
type BackendSvr struct {
	svrStr    string
	isUp      bool // is Up or Down
	failTimes int
}

var (
	pConsisthash *consistent.Consistent
	pBackendSvrs map[string]*BackendSvr
)

func InitBackendServers(pconfig *ProxyConfig) {

	servers := pconfig.Backend
	pConsisthash = consistent.New()
	pBackendSvrs = make(map[string]*BackendSvr)

	for _, svr := range servers {
		pConsisthash.Add(svr)
		pBackendSvrs[svr] = &BackendSvr{
			svrStr:    svr,
			isUp:      true,
			failTimes: 0,
		}
	}
	go checkBackendServers(pconfig)
}

func getBackendSvr(conn net.Conn) (*BackendSvr, bool) {
	remoteAddr := conn.RemoteAddr().String()
	svr, _ := pConsisthash.Get(remoteAddr)

	bksvr, ok := pBackendSvrs[svr]
	return bksvr, ok
}

func checkBackendServers(pconfig *ProxyConfig) {
	// scheduler every 10 seconds
	rand.Seed(time.Now().UnixNano())
	t := time.Tick(time.Duration(10)*time.Second + time.Duration(rand.Intn(100))*time.Millisecond*100)

	for _ = range t {
		for _, v := range pBackendSvrs {
			if v.failTimes >= pconfig.FailOver && v.isUp == true {
				v.isUp = false
				pConsisthash.Remove(v.svrStr)
			}
		}

	}
}
