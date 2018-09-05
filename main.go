package main

import (
	"net/http"

	client "github.com/go-lo/fasthttp-client"
	"github.com/go-lo/go-lo"
)

var (
	endpoint = "http://localhost"
)

type API struct {
	req *http.Request
}

func NewAPI(url string) (a API, err error) {
	a.req, err = http.NewRequest("GET", url, nil)

	return
}

func (a API) Run() {
	seq := golo.NewSequenceID()

	golo.DoRequest(seq, a.req)
}

func main() {
	golo.Client = client.New()

	m, err := NewAPI(endpoint)
	if err != nil {
		panic(err)
	}

	server := golo.New(m)

	panic(golo.Start(server))
}
