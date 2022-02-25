//go:build js && wasm

package moreflag

import (
	"flag"
	"fmt"
	"net/url"
	"syscall/js"
)

func ParseFlagSet(fs *flag.FlagSet) {
	global := js.Global()
	location := global.Get("location")
	if location.IsUndefined() {
		return
	}
	rawQS := location.Get("search").String()
	if len(rawQS) == 0 || rawQS[0] != '?' {
		return
	}
	query, err := url.ParseQuery(rawQS[1:])
	if err != nil {
		panic(fmt.Sprintf("moreflag: %s", err))
	}

	fs.VisitAll(func(f *flag.Flag) {
		if !query.Has(f.Name) {
			return
		}
		f.Value.Set(query.Get(f.Name))
	})
}
