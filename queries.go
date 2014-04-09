package nego

import (
	"encoding/json"
	"fmt"
)

// DeleteFromElastic Help you insert your data at the right place
func (that *ElasticTracker) DeleteFromElastic(bucket string, media string, index string) error {
	rep := that.nEQuery("DELETE", that.servAddr+"/"+bucket+"/"+media+"/"+index, nil)
	if rep.Error != nil {
		return rep.Error
	}
	if rep.StatusCode < 200 || rep.StatusCode >= 300 {
		return fmt.Errorf(" DeleteFromElastic %d status code returned  %s", rep.StatusCode, rep.Response)
	}

	return nil
}

// InsertIntoElastic Help you insert your data at the right place
func (that *ElasticTracker) InsertIntoElastic(bucket string, media string, index string, data []byte) error {
	rep := that.nEQuery("PUT", that.servAddr+"/"+bucket+"/"+media+"/"+index, data)
	if rep.Error != nil {
		return rep.Error
	}
	if rep.StatusCode < 200 || rep.StatusCode >= 300 {
		return fmt.Errorf(" InsertIntoElastic %d status code returned  %s", rep.StatusCode, rep.Response)
	}

	return nil
}

// GetFromElastic Help you Recover your data from the right place
func (that *ElasticTracker) GetFromElastic(bucket string, media string, index string) ([]byte, error) {
	rep := that.nEQuery("GET", that.servAddr+"/"+bucket+"/"+media+"/"+index, nil)
	if rep.Error != nil {
		return nil, rep.Error
	}
	if rep.StatusCode < 200 || rep.StatusCode >= 300 {
		return nil, fmt.Errorf(" GetFromElastic %d status code returned %s", rep.StatusCode, rep.Response)
	}
	f := GetFromElasticStruct{}
	err := json.Unmarshal(rep.Response, &f)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	return f.Source, nil
}

// SearchFromElastic Return A Search from elastic
func (that *ElasticTracker) SearchFromElastic(bucket string, media string, query []byte) (*SearchHitsStruct, error) {
	//fmt.Printf("QUERY\n%s\n", query)
	rep := that.nEQuery("POST", that.servAddr+"/"+bucket+"/"+media+"/_search", query)
	if rep.Error != nil {
		return nil, rep.Error
	}
	if rep.StatusCode < 200 || rep.StatusCode >= 300 {
		return nil, fmt.Errorf(" CompleteFromElastic %d status code returned %s", rep.StatusCode, rep.Response)
	}
	f := SearchFromElasticStruct{}
	err := json.Unmarshal(rep.Response, &f)
	if err != nil {
		fmt.Printf("Error %s!\n", err.Error())
		return nil, err
	}
	return &f.Hits, nil
}

// ForceRefreshElastic Force refresh of that bucket
func (that *ElasticTracker) ForceRefreshElastic(bucket string) error {
	rep := that.nEQuery("POST", that.servAddr+"/"+bucket+"/_refresh", nil)
	if rep.Error != nil {
		return rep.Error
	}
	if rep.StatusCode < 200 || rep.StatusCode >= 300 {
		return fmt.Errorf(" CompleteFromElastic %d status code returned %s", rep.StatusCode, rep.Response)
	}
	return nil
}
