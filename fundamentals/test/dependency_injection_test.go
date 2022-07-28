package fundamentalstest

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "hello, %s", name)
}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "dhany")

	got := buffer.String()
	want := "hello, dhany"

	if got != want {
		t.Error("want:", want, "got:", got)
	}
}
