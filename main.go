package main

import (
	"./config"
	"net/http"
	"strings"
	"os"
	"fmt"
)

var port, tlsport, address string

func main() {
	if len(os.Args) == 1 {
		port = "8080"
		tlsport = "4430"
		address = "0.0.0.0"
	} else if len(os.Args) == 3 {
		port = os.Args[1]
		tlsport = os.Args[2]
		address = "0.0.0.0"
	} else if len(os.Args) == 4 {
		port = os.Args[1]
		tlsport = os.Args[2]
		address = os.Args[3]
	} else {
		fmt.Fprintln(os.Stderr, "Usage: http [PORT TLSPORT [ADDRESS]]")
		os.Exit(-1)
	}
	fmt.Println("Listening on http://"+ address +":"+ port, "and https://"+ address + ":"+ tlsport)
	m := Messaging{}
	m.Txnnel = make(chan TxnnelMsg)
	m.Inforeqnel = make(chan InforeqnelMsg)
	m.Authennel = make(chan AuthennelMsg)
	go ServeDb(m, config.GetMap())
	go http.ListenAndServe(address + ":" + port, http.HandlerFunc(redirectToHttps))
	http.HandleFunc("/", FileHandler)
	http.HandleFunc("/r", m.rHandler)
	http.HandleFunc("/r/", m.rHandler)
	panic(http.ListenAndServeTLS(address + ":" + tlsport, "cert/cert.pem", "cert/key.pem", nil))
}


func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request. Note that "127.0.0.1:8081"
	// will only work if you are accessing the server from your local
	// machine.
	// from https://www.kaihag.com
	s := "https://" + strings.Split(r.Host, ":")[0] + ":" + tlsport
	http.Redirect(w, r, s+r.RequestURI, http.StatusMovedPermanently)
}

