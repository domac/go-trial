package proxy

import (
	"fmt"
	"net/http"
)

// 查询监控信息的接口
func statsHandler(w http.ResponseWriter, r *http.Request) {
	_str := ""
	for _, v := range pBackendSvrs {
		_str += fmt.Sprintf("Server:%s FailTimes:%d isUp:%t\n", v.svrStr, v.failTimes, v.isUp)
	}
	fmt.Fprintf(w, "%s", _str)
}

func InitStats(pconfig *ProxyConfig) {

	fmt.Printf("Stat monitor on addr %s", pconfig.Stats)

	go func() {
		http.HandleFunc("/stats", statsHandler)
		http.ListenAndServe(pconfig.Stats, nil)
	}()
}
