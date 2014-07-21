package nego

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
}

func (that *ElasticTracker) nEQuery(method, query string, payload []byte) nEAnswer {
	q := nEQuery{
		query:   query,
		payload: payload,
		method:  method,
		tries:   3,
	}
	ret, status, err := that.nEDoHTTP(&q)
	for ret == nil && q.tries > 0 {
		time.Sleep(100 * time.Millisecond)
		ret, status, err = that.nEDoHTTP(&q)
	}
	return nEAnswer{ret, status, err}
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

func (that *ElasticTracker) nEDoHTTP(query *nEQuery) ([]byte, int, error) {
	query.tries--
	req, err := http.NewRequest(query.method, query.query, nil)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Add("Accept", "application/json")
	if query.payload != nil {
		nESetBody(req, bytes.NewReader(query.payload))
	}
	resp, err := that.client.Do(req)
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
