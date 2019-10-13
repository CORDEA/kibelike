package main

import (
	"context"
	"flag"
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
	return subcommands.ExitSuccess
}
