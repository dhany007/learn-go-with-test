/*
	You have been asked to make a function called WebsiteRacer which takes two URLs and "races" them
	by hitting them with an HTTP GET and returning the URL which returned first.
	If none of them return within 10 seconds then it should return an error.
	For this, we will be using
	- net/http to make the HTTP calls.
	- net/http/httptest to help us test them.
	- goroutines.
	- select to synchronise processes.
*/

package fundamentalstest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var tenSecondTimeOut = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeOut)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL

	got, err := Racer(slowURL, fastURL)

	if err != nil {
		t.Error("did not expect an error but got one", err)
	}

	if want != got {
		t.Error("got:", got, "want:", want)
	}
}

func TestRacerErrorTimeout(t *testing.T) {
	server := makeDelayedServer(25 * time.Millisecond)

	defer server.Close()

	_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

	if err == nil {
		t.Error("expected an error but didn't get one")
	}
}
