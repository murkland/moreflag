// //go:build js

package moreflag_test

import (
	"flag"
	"syscall/js"
	"testing"

	"github.com/yumland/moreflag"
)

func TestParseFlags(t *testing.T) {
	global := js.Global()
	global.Set("location", js.ValueOf(map[string]interface{}{
		"search": "?a=1&b=hello",
	}))

	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	a := fs.Int("a", 2, "test")
	b := fs.String("b", "bye", "test")
	c := fs.String("c", "donttouchthis", "test")
	moreflag.ParseFlagSet(fs)

	if *a != 1 {
		t.Fatalf("expected a = 1, got %d", *a)
	}
	if *b != "hello" {
		t.Fatalf("expected b = \"hello\", got %v", *b)
	}
	if *c != "donttouchthis" {
		t.Fatalf("expected c = \"donttouchthis\", got %v", *b)
	}
}
