package fileapi

import (
	"errors"
	"io"
	"net/url"
)

type FileApi interface {
	Open(uri *url.URL) (io.ReadWriteCloser, error)
}

var schemeMapper map[string]FileApi

func init() {
	schemeMapper = make(map[string]FileApi)
}

func AddScheme(name string, api FileApi) {
	schemeMapper[name] = api
}

func Open(uri *url.URL) (io.ReadWriteCloser, error) {
	if x, ok := schemeMapper[uri.Scheme]; ok {
		return x.Open(uri)
	}
	return nil, errors.New("This scheme is not supported")
}
