package nego

// ElasticTracker struct to talk with an ElasticSearch server
type ElasticTracker struct {
	buckets  map[string]bool
	servAddr string
	queries  chan nEQuery
	exitbot  chan bool
	nbbot    int
}

// NewElasticTracker Get a new instance of ElasticTracker
func NewElasticTracker(servAddr string, nbbot int) (*ElasticTracker, error) {
	ret := new(ElasticTracker)
	ret.servAddr = servAddr
	ret.queries = make(chan nEQuery)
	ret.exitbot = make(chan bool)
	ret.buckets = make(map[string]bool)
	ret.nbbot = nbbot
	for i := 0; i < ret.nbbot; i++ {
		go nEQueryBot(ret.queries, ret.exitbot, i)
	}
	return ret, nil
}

func (that *ElasticTracker) DeleteTracker() error {
	for i := 0; i < that.nbbot; i++ {
		that.exitbot <- true
	}
	return nil
}
