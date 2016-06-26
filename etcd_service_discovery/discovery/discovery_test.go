package discovery

import (
	"log"
	"testing"
	"time"

	"github.com/coreos/etcd/client"
)

func TestBasic(t *testing.T) {

	cfg := client.Config{
		Endpoints: []string{
			"http://192.168.139.134:2379",
			"http://192.168.139.134:2479",
			"http://192.168.139.134:2579",
		},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)

	if err != nil {
		log.Panicln(err)
	}

	kapi := client.NewKeysAPI(c)

	client := EtcdReigistryClient{
		EtcdRegistryConfig{
			ServiceName:  "test",
			InstanceName: "test1",
			BaseURL:      "127.0.0.1:8080",
		},
		kapi,
	}

	client.Register()
	response, _ := client.ServicesByName("test")
	if len(response) == 0 {
		t.Error("No service registered")
	}
	client.Unregister()
	response, _ = client.ServicesByName("test")
	if len(response) != 0 {
		t.Error("Service not  unregistered")
	}
}

func TestKeepAlive(t *testing.T) {

	cfg := client.Config{
		Endpoints:               []string{"http://127.0.0.1:2379"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)

	if err != nil {
		log.Panicln(err)
	}

	kapi := client.NewKeysAPI(c)

	client := EtcdReigistryClient{
		EtcdRegistryConfig{
			ServiceName:  "test",
			InstanceName: "test1",
			BaseURL:      "127.0.0.1:8080",
		},
		kapi,
	}

	client.Register()

	time.Sleep(50 * time.Second)

	response, _ := client.ServicesByName("test")
	log.Println(response)

	// Ahoj
}
