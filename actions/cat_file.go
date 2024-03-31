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
	"strconv"
	"strings"

	"github.com/kamontat/gogit/git"
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
	var hash = args[0]
	log.Printf("input hash is '%s' with pretty=%t\n", hash, a.pretty)

	var blobFolder = hash[:2]
	var blobFile = hash[2:]

	var blobPath = path.Join(".git", "objects", blobFolder, blobFile)
	log.Printf("looking blob at '%s'", blobPath)

	var blobBuffer, err = os.Open(blobPath)
	if err != nil {
		return fmt.Errorf("open blob object '%s'", err)
	}

	src, err := zlib.NewReader(blobBuffer)
	if err != nil {
		return fmt.Errorf("read blob object '%s'", err)
	}
	defer src.Close()

	raw, err := io.ReadAll(src)
	if err != nil {
		return fmt.Errorf("copy blob object '%s'", err)
	}

	var indexContent = bytes.IndexByte(raw, 0)
	if indexContent == -1 {
		return fmt.Errorf("corrupted file at %s", blobPath)
	}

	var metadata = strings.Split(string(raw[:indexContent]), " ")
	size, err := strconv.Atoi(metadata[1])
	if err != nil {
		return fmt.Errorf("convert content size %s", err)
	}

	myObject := git.GitObject{
		Hash:    hash,
		Type:    metadata[0],
		Size:    size,
		Content: raw[indexContent+1:],
	}

	fmt.Print(string(myObject.Content))
	return nil
}
