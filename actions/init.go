package actions

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func NewInit() *Init {
	return &Init{}
}

type Init struct {
}

func (a *Init) Name() string {
	return "init"
}

func (a *Init) Setup(fs *flag.FlagSet) {
}

func (a *Init) Exec(args []string) error {
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("creating directory: %s", err)
		}
	}

	headFileContents := []byte("ref: refs/heads/main\n")
	if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
		return fmt.Errorf("writing file: %s", err)
	}

	log.Println("Initialized git directory")
	return nil
}
