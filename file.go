package fileapi

import (
	"net/url"
	"os"
	"strconv"
)

type File struct {
}

func (f *File) Open(path *url.URL) (FileInterface, error) {
	settings := path.Query()
	tmp := settings.Get("flag")
	if tmp == "" {
		tmp = "66"
	}
	fileFlag, err := strconv.Atoi(tmp)
	if err != nil {
		return nil, err
	}
	tmp = settings.Get("perm")
	if tmp == "" {
		tmp = "0755"
	}
	filePerm, err := strconv.Atoi(tmp)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(path.Path, fileFlag, os.FileMode(filePerm))
}
