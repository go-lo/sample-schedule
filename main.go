package main

import (
	"net/http"

	"github.com/jspc/loadtest"
)

type MagnumAPI struct {
	URL string
}

func (m MagnumAPI) Run() {
	req, err := http.NewRequest("GET", m.URL, nil)
	if err != nil {
		panic(err)
	}

	seq := loadtest.NewSequenceID()

	loadtest.DoRequest(seq, req)
}

func main() {
	m := MagnumAPI{
		URL: "http://magnum.jspc.pw",
	}

	server := loadtest.NewServer(m)

	panic(loadtest.StartListener(server))
}
