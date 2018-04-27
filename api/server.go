package api

import (
	"fmt"
	"github.com/camelmasa/ipfs-coreapi-sample-app/ipfs"
	"github.com/camelmasa/ipfs-coreapi-sample-app/node"
	"gx/ipfs/QmQvJiADDe7JR4m968MwXobTCCzUqQkP87aRHe29MEBGHV/go-logging"
	manet "gx/ipfs/QmRK2LxanhK2gZq6k6R7vk5ZoYZk8ULSSTB7FzDsMUX6CB/go-multiaddr-net"
	ma "gx/ipfs/QmWWQ2Txc2c6tqjsBpzg5Ar652cHPGNsQQp2SejkNmkUMb/go-multiaddr"
	"net/http"
	"net/url"
)

var log = logging.MustGetLogger("api")

type Server struct{ IpfsCommand *ipfs.Command }

func NewServer(n *node.Node) Server {
	command := &ipfs.Command{Node: n}

	return Server{IpfsCommand: command}
}

func (s *Server) Serve() error {
	gatewayMaddr, err := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/4002")
	if err != nil {
		return err
	}

	gwLis, err := manet.Listen(gatewayMaddr)
	if err != nil {
		return err
	}

	topMux := http.NewServeMux()
	topMux.Handle("/api/", s)

	http.Serve(gwLis.NetListener(), topMux)

	return nil
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.Path)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")

	endpoint := NewEndpoint(s.IpfsCommand)

	switch r.Method {
	case "POST":
		err = endpoint.post(u.String(), w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found")
		err = nil
	}

	if err != nil {
		log.Error(err)
		panic(err)
	}
}
