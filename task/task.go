package task

import (
	"log"
	"os"

	"github.com/panjf2000/ants/v2"
)

// Tasks -  control currently running goroutines
type Tasks struct {
	WaitGroupCount
	*ants.Pool
}

// New - New currently running goroutines
func New(woker int, options ...ants.Option) (*Tasks, error) {
	t := &Tasks{
		WaitGroupCount: WaitGroupCount{},
	}

	defaultOptions := []ants.Option{
		ants.WithPanicHandler(panicHendler(t)),
		ants.WithLogger(log.New(os.Stdout, "task logger", log.Flags())),
		ants.WithPreAlloc(true),
	}

	options = append(defaultOptions, options...)
	pool, err := ants.NewPool(woker, options...)
	t.Pool = pool
	return t, err
}

func (t *Tasks) Summit(fn func()) {
	t.Add()
	go func(t *Tasks, fn func()) {
		_ = t.Submit(func() {
			defer t.Done()
			fn()
		})
	}(t, fn)
}
