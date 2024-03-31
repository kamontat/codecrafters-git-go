package actions

import "flag"

type Action interface {
	Name() string
	Setup(fs *flag.FlagSet)
	Exec(args []string) error
}
