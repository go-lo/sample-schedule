package main

import (
	"net/http"

	"github.com/go-lo/go-lo"
)

type API struct {
	URL string
}

func (m API) Run() {
	req, err := http.NewRequest("GET", m.URL, nil)
	if err != nil {
		panic(err)
	}

	seq := golo.NewSequenceID()

	golo.DoRequest(seq, req)
}

func main() {
	m := API{
		URL: "http://localhost:8765",
	}

	server := golo.New(m)

	panic(golo.Start(server))
}
