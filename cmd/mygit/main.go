package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kamontat/gogit/actions"
)

func executor(action actions.Action) {
	var flags = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	log.Printf("Start action: %s\n", action.Name())

	action.Setup(flags)
	var err = flags.Parse(os.Args[2:])
	if err == nil {
		err = action.Exec(flags.Args())
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s", err)
		os.Exit(1)
	}
}

// Usage: your_git.sh <command> <arg1> <arg2> ...
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
		os.Exit(1)
	}

	var action actions.Action
	switch command := os.Args[1]; command {
	case "init":
		action = actions.NewInit()
	case "cat-file":
		action = actions.NewCatFile()
	default:
		action = actions.NewUnknown(command)
	}

	executor(action)
}
