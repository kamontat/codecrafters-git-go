package actions

import (
	"flag"
	"fmt"
)

func NewUnknown(cmd string) *Unknown {
	return &Unknown{cmd: cmd}
}

type Unknown struct {
	cmd string
}

func (a *Unknown) Name() string {
	return fmt.Sprintf("Unknown: %s", a.cmd)
}

func (a *Unknown) Setup(fs *flag.FlagSet) {
}

func (a *Unknown) Exec(args []string) error {
	return fmt.Errorf("invalid command %s", a.cmd)
}
