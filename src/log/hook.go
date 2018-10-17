// lachesis_log.go
// lachesis hook to logrus
//
package lachesis_log

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type Hook struct {
	mu   sync.RWMutex
	stat [6]int64 // 6 is current value of len(logrus.AllLevels)
	// must be adjusted if changed in future in logrus
	startTime time.Time
}

// NewLocal installs a test hook for a given local logger.
func NewLocal(logger *logrus.Logger) {
	h := new(Hook)
	h.startTime = time.Now()
	logger.Hooks.Add(h)
}

func (t *Hook) Fire(e *logrus.Entry) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if e.Time.Sub(t.startTime).Seconds() > 20 { // 20 seconds in nanoseconds
		// NOTE: we can not use logrus logger here as it seems is not re-entrant
		// and we are inside log entry processing here
		fmt.Println(
			"LOGSTAT debug:", t.stat[logrus.DebugLevel],
			" info:", t.stat[logrus.InfoLevel],
			" warn:", t.stat[logrus.WarnLevel],
			" error:", t.stat[logrus.ErrorLevel],
			" fatal:", t.stat[logrus.FatalLevel],
			" panic:", t.stat[logrus.PanicLevel] )
		t.stat = [6]int64{} // 6 is current value of len(logrus.AllLevels)
		// must be adjusted if changed in future in logrus
		t.startTime = e.Time
	}
	switch e.Level {
	case logrus.PanicLevel:
		fallthrough
	case logrus.FatalLevel:
		fallthrough
	case logrus.ErrorLevel:
		debug.PrintStack()
		fallthrough
	default:
		t.stat[e.Level]++
	}
	return nil
}

func (t *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}
