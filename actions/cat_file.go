package actions

import (
	"bytes"
	"compress/zlib"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func NewCatFile() *CatFile {
	return &CatFile{
		pretty: false,
	}
}

type CatFile struct {
	pretty bool
}

func (a *CatFile) Name() string {
	return "cat-file"
}

func (a *CatFile) Setup(fs *flag.FlagSet) {
	fs.BoolVar(&(a.pretty), "p", false, "pretty print content")
}

func (a *CatFile) Exec(args []string) error {
	var blob = args[0]
	log.Printf("input blob is '%s' with pretty=%t\n", blob, a.pretty)

	var blobFolder = blob[:2]
	var blobFile = blob[2:]

	var blobPath = path.Join(".git", "objects", blobFolder, blobFile)
	log.Printf("looking blob at %s", blobPath)

	var blobBuffer, err = os.Open(blobPath)
	if err != nil {
		return fmt.Errorf("open blob object '%s'", err)
	}

	src, err := zlib.NewReader(blobBuffer)
	if err != nil {
		return fmt.Errorf("read blob object '%s'", err)
	}
	defer src.Close()

	var content bytes.Buffer
	_, err = io.Copy(&content, src)
	if err != nil {
		return fmt.Errorf("copy blob object '%s'", err)
	}

	fmt.Printf("%s", content.String())

	return nil
}
