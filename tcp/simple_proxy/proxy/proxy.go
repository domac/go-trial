package proxy

import (
	"fmt"
	"net"
	"time"
)

func InitProxy(pconfig *ProxyConfig) {
	fmt.Printf("Proxying %s -> %s\n", pconfig.Bind, pconfig.Backend)

	server, err := net.Listen("tcp", pconfig.Bind)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	waitQueue := make(chan net.Conn, pconfig.WaitQueueLen)
	availPools := make(chan bool, pconfig.MaxConn)
	for i := 0; i < pconfig.MaxConn; i++ {
		availPools <- true
	}

	go loop(waitQueue, availPools, pconfig)

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Errorf(err.Error())
		} else {
			fmt.Printf("Received connection from %s.\n", connection.RemoteAddr())
			waitQueue <- connection
		}
	}
}

func loop(waitQueue chan net.Conn, availPools chan bool, pconfig *ProxyConfig) {
	for connection := range waitQueue {
		<-availPools
		go func(connection net.Conn, pconfig *ProxyConfig) {
			handleConnection(connection, pconfig)
			availPools <- true
			fmt.Printf("Closed connection from %s.\n", connection.RemoteAddr())
		}(connection, pconfig)
	}
}

func handleConnection(connection net.Conn, pconfig *ProxyConfig) {
	defer connection.Close()

	bksvr, ok := getBackendSvr(connection)
	if !ok {
		return
	}
	remote, err := net.Dial("tcp", bksvr.svrStr)

	if err != nil {
		fmt.Errorf(err.Error())
		bksvr.failTimes++
		return
	}

	//等待双向连接完成
	complete := make(chan bool, 2)
	oneSide := make(chan bool, 1)
	otherSide := make(chan bool, 1)
	go pass(connection, remote, complete, oneSide, otherSide, pconfig)
	go pass(remote, connection, complete, otherSide, oneSide, pconfig)
	<-complete
	<-complete
	remote.Close()
}

// copy Content two-way
func pass(from net.Conn, to net.Conn, complete chan bool, oneSide chan bool, otherSide chan bool, pconfig *ProxyConfig) {
	var err error
	var read int
	bytes := make([]byte, 256)

	for {
		select {

		case <-otherSide:
			complete <- true
			return

		default:

			from.SetReadDeadline(time.Now().Add(time.Duration(pconfig.Timeout) * time.Second))
			read, err = from.Read(bytes)
			if err != nil {
				complete <- true
				oneSide <- true
				return
			}

			to.SetWriteDeadline(time.Now().Add(time.Duration(pconfig.Timeout) * time.Second))
			_, err = to.Write(bytes[:read])
			if err != nil {
				complete <- true
				oneSide <- true
				return
			}
		}
	}
}
