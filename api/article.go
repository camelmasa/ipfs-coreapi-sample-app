package api

import (
	"fmt"
	"github.com/camelmasa/ipfs-coreapi-sample-app/ipfs"
	"net/http"
	"strings"
)

type Article struct{ IpfsCommand *ipfs.Command }

func (a *Article) Create(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()

	path, err := a.IpfsCommand.Add(strings.NewReader(r.PostForm["content"][0]))
	if err != nil {
		return err
	}

	fmt.Fprint(w, "Added! You can check on gateway. https://ipfs.io/ipfs/"+path.Cid().String())

	return nil
}
