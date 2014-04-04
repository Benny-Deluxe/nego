package nextelasticgo

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type nEAnswer struct {
	Response   []byte
	StatusCode int
	Error      error
}

type nEQuery struct {
	query   string
	method  string // 1 GET 2 PUT 3 POST
	payload []byte
	tries   int64
	answer  chan nEAnswer
}

func (that *ElasticTracker) nEQuery(method, query string, payload []byte) nEAnswer {
	a := make(chan nEAnswer)
	q := nEQuery{
		query:   query,
		payload: payload,
		method:  method,
		tries:   3,
		answer:  a,
	}
	that.queries <- q
	ret := <-a
	return ret
}

func nESetBody(r *http.Request, body io.Reader) {
	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = ioutil.NopCloser(body)
	}
	r.Body = rc
	if body != nil {
		switch v := body.(type) {
		case *strings.Reader:
			r.ContentLength = int64(v.Len())
		case *bytes.Buffer:
			r.ContentLength = int64(v.Len())
		}
	}

}

func nEDoHTTP(query nEQuery) ([]byte, int, error) {
	req, err := http.NewRequest(query.method, query.query, nil)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Add("Accept", "application/json")
	if query.payload != nil {
		nESetBody(req, bytes.NewReader(query.payload))
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return body, resp.StatusCode, err
}

func nEReputQuery(query nEQuery, queries chan nEQuery, freeretry bool, lastError error, lastStatus int) {
	if !freeretry {
		query.tries--
	}
	if query.tries <= 0 {
		query.answer <- nEAnswer{nil, lastStatus, lastError}
		return
	}
	time.Sleep(100 * time.Millisecond)
	queries <- query
}

func nEQueryBot(queries chan nEQuery, num int) {
	for {
		//fmt.Println("Bot", num, "is waiting")
		select {
		case query := <-queries:
			ret, status, err := nEDoHTTP(query)
			if ret == nil {
				go nEReputQuery(query, queries, false, err, status)
			} else {
				query.answer <- nEAnswer{ret, status, err}
			}
		}
	}
}
