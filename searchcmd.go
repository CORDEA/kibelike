package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
)

type SearchCmd struct {
	client *Client
}

func (*SearchCmd) Name() string {
	return "search"
}

func (*SearchCmd) Synopsis() string {
	return ""
}

func (*SearchCmd) Usage() string {
	return ""
}

func (*SearchCmd) SetFlags(f *flag.FlagSet) {
}

func (s *SearchCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	arg := f.Arg(0)
	if arg == "" {
		return subcommands.ExitFailure
	}
	response, err := s.client.Search(arg)
	if err != nil {
		return subcommands.ExitFailure
	}
	nodes := response.Search.Nodes
	for e := range response.Search.Nodes {
		node := nodes[e]
		fmt.Println(node.Document.ID)
		if node.Document.Title != "" {
			fmt.Println("\t" + node.Document.Title)
		}
		if node.Document.Content != "" {
			fmt.Println("\t" + node.Document.Content)
		}
	}
	return subcommands.ExitSuccess
}
