package moreflag

import "flag"

func Parse() {
	ParseFlagSet(flag.CommandLine)
}
