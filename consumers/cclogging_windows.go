//go:build windows

package consumers

import (
	"os"
	"sync"
	"time"
)

type ConcurrentLoggingConsumer struct {
	w     *ConcurrentLogWriter
	Fname string
	Hour  bool
}
type ConcurrentLogWriter struct {
	rec chan string

	fname string
	file  *os.File

	day  int
	hour int

	hourRotate bool

	wg sync.WaitGroup
}

func InitConcurrentLoggingConsumer(fname string, hour bool) (*ConcurrentLoggingConsumer, error) {
	w, err := InitConcurrentLogWriter(fname, hour)
	if err != nil {
		return nil, err
	}

	c := &ConcurrentLoggingConsumer{Fname: fname, Hour: hour, w: w}
	return c, nil
}
func InitConcurrentLogWriter(fname string, hourRotate bool) (*ConcurrentLogWriter, error) {
	w := &ConcurrentLogWriter{
		fname:      fname,
		day:        time.Now().Day(),
		hour:       time.Now().Hour(),
		hourRotate: hourRotate,
		rec:        make(chan string, CHANNEL_SIZE),
	}

	return w, nil
}