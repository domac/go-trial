package main

import (
	"flag"
	"fmt"
	proxy "github.com/phillihq/go-trial/tcp/simple_proxy/proxy"
	"os"
	"os/signal"
	"syscall"
)

var (
	configFile = flag.String("c", "etc/conf.yaml", "配置文件，默认etc/conf.yaml")
)

func exitSignal() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGUSR1, syscall.SIGTERM, syscall.SIGINT)
L:
	for {
		sig := <-sigChan
		switch sig {
		case syscall.SIGUSR1:
			fmt.Errorf("Reopen log file\n")
		case syscall.SIGTERM, syscall.SIGINT:
			fmt.Errorf("Catch SIGTERM singal, exit\n")
			break L
		}
	}
}

func main() {
	pconfig, err := proxy.ParseConfigFile(*configFile)
	if nil != err {
		return
	}

	//初始化后端服务器
	proxy.InitBackendServers(pconfig)

	go exitSignal()

	proxy.InitStats(pconfig)

	proxy.InitProxy(pconfig)

}
