/*
Software often kicks off long-running, resource-intensive processes (often in goroutines).
If the action that caused this gets cancelled or fails for some reason you need
to stop these processes in a consistent way through your application.
If you don't manage this your snappy Go application that you're so proud of could start having difficult
to debug performance problems.
In this chapter we'll use the package context to help us manage long-running processes.
We're going to start with a classic example of a web server that when hit kicks off
a potentially long-running process to fetch some data for it to return in the response.
We will exercise a scenario where a user cancels the request before the data can be retrieved
and we'll make sure the process is told to give up.
*/

package fundamentalstest

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you like
		}

		fmt.Fprint(w, data)
	}
}

type Store interface {
	Fetch(context context.Context) (string, error)
}

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string

		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store get cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}

		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// ini sama seperti http.ResponseWriter
type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}
func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}
func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestStore(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"

		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Error("want", data, "got", response.Body.String())
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"

		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
