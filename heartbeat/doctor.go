package heartbeat

import (
	"encoding/json"
	"fmt"
	"net"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	DEFAULT_RECYCLE     = 2 * time.Second
	LEVEL_FULL      int = 10
	LEVEL_ALIVE     int = 7
	LEVEL_WARNING   int = 3
	LEVEL_DEAD      int = 0
)

type Host struct {
	Ip    string
	Time  time.Time
	Alive bool
	HP    int
}

type Doctor struct {
	sync.Mutex
	Addr            string
	RecycleDuration time.Duration
	Hosts           map[string]*Host
}

func NewDoctor(addr string) *Doctor {
	return &Doctor{
		Addr:            addr,
		RecycleDuration: DEFAULT_RECYCLE,
		Hosts:           make(map[string]*Host, 100),
	}
}

//更新状态
func (this *Doctor) updateStatus(ip string) {
	this.Lock()
	defer this.Unlock()
	fmt.Printf("oncall : %s \n", ip)
	h, ok := this.Hosts[ip]
	if !ok {
		this.Hosts[ip] = &Host{Ip: ip, Time: time.Now(), Alive: true, HP: LEVEL_FULL}
		return
	}

	h.HP += 1
	if h.HP > LEVEL_FULL {
		h.HP = LEVEL_FULL
	}

	if h.HP < LEVEL_WARNING {
		h.HP = LEVEL_WARNING
	}

	if h.HP > LEVEL_ALIVE {
		h.Alive = true
	}

	if h.HP == LEVEL_DEAD {
		h.Alive = false
	}

	h.Time = time.Now()
}

//工作处理
func (this *Doctor) oncall(notify chan Host) {
	for {
		this.Lock()
		for _, h := range this.Hosts {

			var state = h.Alive

			if h.HP -= 1; h.HP < LEVEL_DEAD {
				h.HP = LEVEL_DEAD
			}

			if h.HP > LEVEL_ALIVE {
				h.Alive = true
			}

			if h.HP == LEVEL_DEAD {
				h.Alive = false
			}

			if h.Alive != state {
				notify <- *h
			}
		}
		this.Unlock()
		time.Sleep(this.RecycleDuration)
	}
}

//心跳诊断
func (this *Doctor) Watch() (chan Host, error) {
	ch, err := this.listenUDP()
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			ip := <-ch
			fmt.Println("read from channel :", ip)
			this.updateStatus(ip)
		}
	}()
	notify := make(chan Host, 50)
	go this.oncall(notify)
	return notify, nil
}

//监听UDP包
func (this *Doctor) listenUDP() (chan string, error) {
	packetConn, err := net.ListenPacket("udp", this.Addr)
	if err != nil {
		return nil, err
	}
	fmt.Printf("start listen packet from : %s \n", this.Addr)
	ch := make(chan string)
	go func() {
		defer packetConn.Close()
		buf := make([]byte, 1000)
		for {
			n, addr, err := packetConn.ReadFrom(buf)
			msg := string(buf[:n])
			fmt.Printf("Receive from %s : %s \n", addr, msg)

			if err != nil {
				fmt.Errorf(err.Error())
				continue
			}

			ip := strings.Split(addr.String(), ":")[0]
			ch <- ip
		}
	}()
	return ch, nil
}

//输出当前监控数据
func (this *Doctor) JSONMessage() []byte {
	this.Lock()
	defer this.Unlock()
	aliveds := make([]string, 0, len(this.Hosts))
	deads := make([]string, 0, len(this.Hosts))

	for ip, host := range this.Hosts {
		if host.Alive {
			aliveds = append(aliveds, LookupName(ip))
		} else {
			deads = append(deads, LookupName(ip))
		}
	}

	sort.Strings(aliveds)
	sort.Strings(deads)

	//输出JSON格式
	data, err := json.Marshal(struct {
		Aliveds []string
		Deads   []string
	}{
		aliveds,
		deads,
	})

	if err != nil {
		panic(err)
	}

	return data
}
