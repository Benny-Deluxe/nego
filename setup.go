package nextelastic

import (
	"fmt"
	"sync"
)

var checkBucketMut sync.Mutex

//ExistBucket Check if a bucket Exist
func (that *ElasticTracker) ExistBucket(bucket string) (bool, error) {
	checkBucketMut.Lock()
	defer checkBucketMut.Unlock()
	if i, ok := that.buckets[bucket]; ok == true && i == true {
		return true, nil
	}
	rep := that.nEQuery("HEAD", that.servAddr+"/"+bucket+"/", nil)
	if rep.Error != nil {
		return false, rep.Error
	}
	if rep.StatusCode == 200 {
		that.buckets[bucket] = true
		return true, nil
	} else if rep.StatusCode == 404 {
		return false, nil
	}
	return false, fmt.Errorf(" ExistBucket StatusCode unknow %d (%s)", rep.StatusCode, rep.Response)
}

//CreateBucketIfDoesntExist Create A Bucket Only if it doesnt Exist
func (that *ElasticTracker) CreateBucketIfDoesntExist(bucket string, data []byte) error {
	exist, err := that.ExistBucket(bucket)
	if err != nil {
		return err
	}
	if exist == true {
		return nil
	}
	return that.CreateBucket(bucket, data)
}

// CreateBucket Create a bucket
func (that *ElasticTracker) CreateBucket(bucket string, data []byte) error {
	checkBucketMut.Lock()
	defer checkBucketMut.Unlock()
	rep := that.nEQuery("PUT", that.servAddr+"/"+bucket+"/", data)
	if rep.Error != nil {
		return rep.Error
	}
	if rep.StatusCode < 200 || rep.StatusCode >= 300 {
		return fmt.Errorf(" CreateBucket StatusCode unknow %d (%s)", rep.StatusCode, rep.Response)
	}
	that.buckets[bucket] = true
	return nil
}

// DeleteBucket Delete a bucket
func (that *ElasticTracker) DeleteBucket(bucket string) error {
	checkBucketMut.Lock()
	defer checkBucketMut.Unlock()
	if i, ok := that.buckets[bucket]; ok == true && i == false {
		return nil
	}
	rep := that.nEQuery("DELETE", that.servAddr+"/"+bucket+"/", nil)
	if rep.Error != nil {
		return rep.Error
	}
	that.buckets[bucket] = false
	return nil
}

// SetMapping set the mapping
func (that *ElasticTracker) SetMapping(bucket string, media string, data []byte) error {
	rep := that.nEQuery("PUT", that.servAddr+"/"+bucket+"/"+media+"/_mapping", data)
	if rep.Error != nil {
		return rep.Error
	}
	if rep.StatusCode < 200 || rep.StatusCode >= 300 {
		return fmt.Errorf(" SetMapping StatusCode unknow %d (%s)", rep.StatusCode, rep.Response)
	}
	return nil
}

// ExistMapping Check if the mapping for that bucker/media is Already set
func (that *ElasticTracker) ExistMapping(bucket string, media string) (bool, error) {
	rep := that.nEQuery("GET", that.servAddr+"/"+bucket+"/"+media+"/_mapping", nil)
	if rep.Error != nil {
		return false, rep.Error
	}
	if rep.StatusCode != 200 {
		return false, fmt.Errorf("elasticSearch GetMapping wrong status code (%d) %s", rep.StatusCode, rep.Response)
	}
	if len(rep.Response) <= 2 {
		return false, nil
	}
	return true, nil
}
