package git

type GitObject struct {
	Hash    string
	Type    string
	Size    int
	Content []byte
}
