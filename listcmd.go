package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
)

type ListCmd struct {
	client *Client
}

func (*ListCmd) Name() string {
	return "list"
}

func (*ListCmd) Synopsis() string {
	return ""
}

func (*ListCmd) Usage() string {
	return ""
}

func (*ListCmd) SetFlags(f *flag.FlagSet) {
}

func (s *ListCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	response, err := s.client.GetNotes()
	if err != nil {
		return subcommands.ExitFailure
	}
	nodes := response.Notes.Nodes
	for e := range nodes {
		node := nodes[e]
		fmt.Println(node.ID)
		fmt.Println("\t" + node.Title)
	}
	return subcommands.ExitSuccess
}
