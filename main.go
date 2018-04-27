package main

import (
	"context"
	"fmt"
	"github.com/camelmasa/ipfs-coreapi-sample-app/api"
	"github.com/camelmasa/ipfs-coreapi-sample-app/node"
	"os"
	"os/signal"
)

func main() {
	repoPath := os.Getenv("HOME") + "/.ipfs"

	// Interrupt signal for exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c

		os.Exit(1)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start Node daemon
	n, err := node.NewNode(ctx, repoPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Peer ID: ", n.IpfsNode.Identity.Pretty())

	server := api.NewServer(n)
	err = server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
}
