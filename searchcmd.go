package main

import (
	"context"
	"flag"
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

func (*SearchCmd) SetFlags(f *flag.FlagSet)  {
}

func (s *SearchCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	arg := f.Arg(0)
	if arg == "" {
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}