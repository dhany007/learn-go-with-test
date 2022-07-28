/*
We want to make a counter which is safe to use concurrently.
We'll start with an unsafe counter and verify its behaviour works in a single-threaded environment.
Then we'll exercise it's unsafeness with multiple goroutines trying to use it via a test and fix it.
*/

package fundamentalstest

import (
	"sync"
	"testing"
)

type Counter struct {
	sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it run safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		wg := sync.WaitGroup{}

		counter := NewCounter()

		wg.Add(wantedCount)
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Error("got", got.Value(), "want", want)
	}
}
