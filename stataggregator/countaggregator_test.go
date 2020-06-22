package stataggregator

import (
	"sync"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCountAggregator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStater := NewMockStat(ctrl)
	tests := []struct {
		name            string
		ExecuteStatting func(stater CountAggregator, mockStater *MockStat)
	}{
		{
			name: "simple aggregation",
			ExecuteStatting: func(stater CountAggregator, mockStater *MockStat) {
				mockStater.EXPECT().Count("stat", float64(15), []string{"scoob"})
				go stater.Count("stat", 9, "scoob")
				go stater.Count("stat", 6, "scoob")
				time.Sleep(200 * time.Millisecond)
			},
		},
		{
			name: "tag ordering aggregation",
			ExecuteStatting: func(stater CountAggregator, mockStater *MockStat) {
				mockStater.EXPECT().Count("triple", float64(10), []string{"double", "single"})
				go func() {
					stater.Count("triple", 9, "double", "single")
					stater.Count("triple", 1, "single", "double")
				}()
				time.Sleep(200 * time.Millisecond)
			},
		},
		{
			name: "complex stat aggregation",
			ExecuteStatting: func(stater CountAggregator, mockStater *MockStat) {
				mockStater.EXPECT().Count("stat", float64(15), []string{"scoob"})
				mockStater.EXPECT().Count("yakitori", float64(4), []string{"book"})
				go func() {
					stater.Count("stat", 9, "scoob")
					stater.Count("yakitori", 3, "book")
					stater.Count("stat", 6, "scoob")
					stater.Count("yakitori", 1, "book")
				}()
				time.Sleep(200 * time.Millisecond)
			},
		},
		{
			name: "multiple stat aggregation",
			ExecuteStatting: func(stater CountAggregator, mockStater *MockStat) {
				mockStater.EXPECT().Count("stat", float64(15), []string{"scoob"})
				mockStater.EXPECT().Count("yakitori", float64(4), []string{"book"})
				go func() {
					stater.Count("stat", 9, "scoob")
					stater.Count("yakitori", 3, "book")
					stater.Count("stat", 6, "scoob")
					stater.Count("yakitori", 1, "book")
				}()
				time.Sleep(200 * time.Millisecond)
				mockStater.EXPECT().Count("stat", float64(2), []string{"scoob"})
				mockStater.EXPECT().Count("yakitori", float64(2), []string{"book"})
				go func() {
					stater.Count("stat", 1, "scoob")
					stater.Count("yakitori", 1, "book")
					stater.Count("stat", 1, "scoob")
					stater.Count("yakitori", 1, "book")
				}()
				time.Sleep(200 * time.Millisecond)
			},
		},
	}
	for _, test := range tests {

		t.Run(test.name, func(tt *testing.T) {
			countAggregator := CountAggregator{
				Stater:        mockStater,
				Bucket:        make(map[StatTagKey]float64),
				lock:          &sync.Mutex{},
				flushJob:      make(chan int, 1),
				FlushInterval: 100 * time.Millisecond,
			}
			test.ExecuteStatting(countAggregator, mockStater)
		})

	}
}
func TestGauge(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStater := NewMockStat(ctrl)
	stater := CountAggregator{
		Stater: mockStater,
	}
	mockStater.EXPECT().Gauge("yakitori", float64(1), []string{"book"})
	stater.Gauge("yakitori", 1, "book")
}

func TestHistogram(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStater := NewMockStat(ctrl)
	stater := CountAggregator{
		Stater: mockStater,
	}
	mockStater.EXPECT().Histogram("yakitori", float64(1), []string{"book"})
	stater.Histogram("yakitori", 1, "book")
}
func TestTiming(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStater := NewMockStat(ctrl)
	stater := CountAggregator{
		Stater: mockStater,
	}
	mockStater.EXPECT().Timing("yakitori", 1*time.Millisecond, []string{"book"})
	stater.Timing("yakitori", 1*time.Millisecond, "book")
}

func TestAddTags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStater := NewMockStat(ctrl)
	stater := CountAggregator{
		Stater: mockStater,
	}
	mockStater.EXPECT().AddTags([]string{"tag1", "tag2"})
	stater.AddTags("tag1", "tag2")
}

func TestGetTags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStater := NewMockStat(ctrl)
	stater := CountAggregator{
		Stater: mockStater,
	}
	mockStater.EXPECT().GetTags().Return([]string{"tag1", "tag2"})
	tags := stater.GetTags()
	assert.Equal(t, []string{"tag1", "tag2"}, tags)
}
