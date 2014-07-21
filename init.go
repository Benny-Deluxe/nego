package nego

import (
	"net/http"
)

// ElasticTracker struct to talk with an ElasticSearch server
type ElasticTracker struct {
	buckets  map[string]bool
	servAddr string
	client *http.Client
}

// NewElasticTracker Get a new instance of ElasticTracker
func NewElasticTracker(client *http.Client, servAddr string) (*ElasticTracker, error) {
	ret := new(ElasticTracker)
	ret.servAddr = servAddr
	ret.client = client
	ret.buckets = make(map[string]bool)
	return ret, nil
}

func (that *ElasticTracker) DeleteTracker() error {
	return nil
}
