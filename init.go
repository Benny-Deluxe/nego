package nextelastic

// ElasticTracker struct to talk with an ElasticSearch server
type ElasticTracker struct {
	buckets  map[string]bool
	servAddr string
	queries  chan nEQuery
}

// NewElasticTracker Get a new instance of ElasticTracker
func NewElasticTracker(servAddr string) (*ElasticTracker, error) {
	ret := new(ElasticTracker)
	ret.servAddr = servAddr
	ret.queries = make(chan nEQuery)
	ret.buckets = make(map[string]bool)
	for i := 0; i < 20; i++ {
		go nEQueryBot(ret.queries, i)
	}
	return ret, nil
}
