package node

import (
	"context"
	"gx/ipfs/QmNUCLv5fmUBuAcwbkt58NQvMcJgd5FPCYV2yNCXq4Wnd6/go-ipfs/cmd/ipfs/util"
	"gx/ipfs/QmNUCLv5fmUBuAcwbkt58NQvMcJgd5FPCYV2yNCXq4Wnd6/go-ipfs/core"
	"gx/ipfs/QmNUCLv5fmUBuAcwbkt58NQvMcJgd5FPCYV2yNCXq4Wnd6/go-ipfs/repo/fsrepo"
	"gx/ipfs/QmQvJiADDe7JR4m968MwXobTCCzUqQkP87aRHe29MEBGHV/go-logging"
)

var log = logging.MustGetLogger("node")

type Node struct {
	IpfsNode *core.IpfsNode
	RepoPath string
}

func NewNode(ctx context.Context, repoPath string) (*Node, error) {
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err := util.ManageFdLimit(); err != nil {
		log.Error(err)
		return nil, err
	}

	config := &core.BuildCfg{
		Repo:      repo,
		Online:    true,
		ExtraOpts: map[string]bool{"mplex": true},
	}

	ipfsNode, err := core.NewNode(ctx, config)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	ipfsNode.SetLocal(false)

	n := &Node{
		IpfsNode: ipfsNode,
		RepoPath: repoPath,
	}

	return n, nil
}
