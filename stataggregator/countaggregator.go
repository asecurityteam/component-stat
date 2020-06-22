package stataggregator

import (
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/rs/xstats"
)

// CountAggregator
type CountAggregator struct {
	Stater        xstats.XStater
	Bucket        map[StatTagKey]float64
	lock          *sync.Mutex
	flushJob      chan int
	FlushInterval time.Duration
}

// StatTagKey is a struct that represents a composite key
// for storing aggregate
type StatTagKey struct {
	StatKey string
	TagsKey string
}

func (ca *CountAggregator) insertStat(stat string, count float64, tags ...string) {
	sort.Strings(tags)
	tagsKey := strings.Join(tags[:], " ")
	compositeKey := StatTagKey{
		StatKey: stat,
		TagsKey: tagsKey,
	}
	ca.lock.Lock()
	defer ca.lock.Unlock()
	if currentCount, ok := ca.Bucket[compositeKey]; ok {
		ca.Bucket[compositeKey] = currentCount + count
		return
	}
	ca.Bucket[compositeKey] = count
	return
}

func (ca *CountAggregator) clearBucket() {
	for key := range ca.Bucket {
		delete(ca.Bucket, key)
	}
}

func (ca *CountAggregator) retrieveAndConvertStat(compositeKey StatTagKey) (string, float64, []string) {
	aggregateCount := ca.Bucket[compositeKey]
	tags := strings.Split(compositeKey.TagsKey, " ")
	stat := compositeKey.StatKey
	return stat, aggregateCount, tags
}
func (ca *CountAggregator) flush() {
	time.Sleep(ca.FlushInterval)
	ca.lock.Lock()
	defer ca.lock.Unlock()
	for compositeKey := range ca.Bucket {
		aggregatedStat, aggregatedCount, aggregatedTags := ca.retrieveAndConvertStat(compositeKey)
		ca.Stater.Count(aggregatedStat, aggregatedCount, aggregatedTags...)
	}
	ca.clearBucket()
	<-ca.flushJob
}

// Count implements XStater interface
func (ca *CountAggregator) Count(stat string, count float64, tags ...string) {
	ca.insertStat(stat, count, tags...)

	select {
	case ca.flushJob <- 1:
		go ca.flush()
	default:
		return
	}
}

// Gauge implements XStater interface
func (ca *CountAggregator) Gauge(stat string, value float64, tags ...string) {
	ca.Stater.Gauge(stat, value, tags...)
}

// Histogram implements XStater interface
func (ca *CountAggregator) Histogram(stat string, value float64, tags ...string) {
	ca.Stater.Histogram(stat, value, tags...)
}

// Timing implements XStater interface
func (ca *CountAggregator) Timing(stat string, duration time.Duration, tags ...string) {
	ca.Stater.Timing(stat, duration, tags...)
}

// AddTags implements XStater interface
func (ca *CountAggregator) AddTags(tags ...string) {
	ca.Stater.AddTags(tags...)
}

// GetTags implements XStater interface
func (ca *CountAggregator) GetTags() []string {
	return ca.Stater.GetTags()
}
