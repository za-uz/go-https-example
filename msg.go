package main

import (
	"./types"
	"encoding/json"
	"net/http"
	"strings"
)

type Messaging struct {
	Txnnel     chan TxnnelMsg
	Inforeqnel chan InforeqnelMsg
	Authennel  chan AuthennelMsg
}

type TxnnelMsg struct {
	types.Transaction
	Returnel chan bool
}

type InforeqnelMsg struct {
	string
	Infonnel chan types.Information
}

type AuthennelMsg struct {
	types.Authentication
	Returnel chan bool
}

func (m *Messaging) rHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		m.rTxHandler(w, r)
	} else if r.Method == "GET" {
		if len(r.URL.Path) <= 3 {
			m.rAllHandler(w, r)
		} else {
			m.rInfoHandler(w, r)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (m *Messaging) rTxHandler(w http.ResponseWriter, r *http.Request) {
	a := types.Authentication{}
	jd := json.NewDecoder(r.Body)
	err := jd.Decode(&a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rc := make(chan bool)
	am := AuthennelMsg{a, rc}
	m.Authennel <- am
	if <-rc != true {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tm := TxnnelMsg{a.Transaction, rc}
	m.Txnnel <- tm
	if <-rc != true {
		w.WriteHeader(http.StatusNotAcceptable)
	}
	w.WriteHeader(http.StatusOK)
}

func (m *Messaging) rInfoHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/r/")
	ic := make(chan types.Information)
	im := InforeqnelMsg{name, ic}
	m.Inforeqnel <- im
	info, ok := <-ic
	if ok {
		e := json.NewEncoder(w)
		err := e.Encode(info)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (m *Messaging) rAllHandler(w http.ResponseWriter, r *http.Request) {
	ic := make(chan types.Information, 50)
	im := InforeqnelMsg{"", ic}
	m.Inforeqnel <- im
	all := map[string]int{}
	for {
		info, more := <-ic
		if !more {
			break
		}
		all[info.Name] = info.Balance
	}
	e := json.NewEncoder(w)
	err := e.Encode(all)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
