package main

import (
	"expvar"
	"fmt"
	"io"
	"net/http"
)

//reference : http://go-wise.blogspot.kr/2011/10/e.html

// Two metrics, these are exposed by "magic" :)
// Number of calls to our server.
var numCalls = expvar.NewInt("num_calls")

// Last user.
var lastUser = expvar.NewString("last_user")

func HelloServer(w http.ResponseWriter, req *http.Request) {
	user := req.FormValue("user")

	// Update metrics
	numCalls.Add(1)
	lastUser.Set(user)

	msg := fmt.Sprintf("G'day %s\n", user)
	io.WriteString(w, msg)
}

func main() {
	http.HandleFunc("/demo", HelloServer)
	http.ListenAndServe(":8080", nil)
}
