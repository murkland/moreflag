//go:build !js

package moreflag

import (
	"flag"
	"os"
)

func ParseFlagSet(fs *flag.FlagSet) {
	fs.Parse(os.Args[1:])
}
