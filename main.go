package main

import (
	"log"
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
		URL: "http://10.50.0.4:8765",
	}

	server := loadtest.NewServer(m)

	panic(loadtest.StartListener(server))
}
