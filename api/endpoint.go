package api

import (
	"fmt"
	"github.com/camelmasa/ipfs-coreapi-sample-app/ipfs"
	"net/http"
	"strings"
)

type Endpoint struct {
	article *Article
}

func NewEndpoint(command *ipfs.Command) Endpoint {
	article := &Article{IpfsCommand: command}

	return Endpoint{article: article}
}

func (e *Endpoint) post(path string, w http.ResponseWriter, r *http.Request) error {
	switch {
	case strings.HasPrefix(path, "/api/articles"):
		return e.article.Create(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found")
		return nil
	}
}
