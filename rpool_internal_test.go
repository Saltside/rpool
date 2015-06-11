package rpool

import (
	"testing"
	"time"

	"github.com/facebookgo/clock"
	"github.com/facebookgo/ensure"
)

func TestSentinelCloserPanic(t *testing.T) {
	defer ensure.PanicDeepEqual(t, "should never get called")
	sentinelCloser(0).Close()
}

func TestRequestExpired(t *testing.T) {
	klock := clock.NewMock()

	r := request{
		made:    klock.Now(),
		timeout: time.Second,
		clock:   klock,
	}

	ensure.False(t, r.expired())
	klock.Add(time.Second)
	ensure.True(t, r.expired())
}
