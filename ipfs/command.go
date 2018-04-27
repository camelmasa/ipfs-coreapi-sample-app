package ipfs

import (
	"context"
	"github.com/camelmasa/ipfs-coreapi-sample-app/node"
	"gx/ipfs/QmNUCLv5fmUBuAcwbkt58NQvMcJgd5FPCYV2yNCXq4Wnd6/go-ipfs/core/coreapi"
	coreiface "gx/ipfs/QmNUCLv5fmUBuAcwbkt58NQvMcJgd5FPCYV2yNCXq4Wnd6/go-ipfs/core/coreapi/interface"
	"io"
)

type Command struct {
	Node *node.Node
}

func (c *Command) Add(content io.Reader) (coreiface.Path, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	path, err := coreapi.NewCoreAPI(c.Node.IpfsNode).Unixfs().Add(ctx, content)

	if err != nil {
		return nil, err
	}

	return path, nil
}
