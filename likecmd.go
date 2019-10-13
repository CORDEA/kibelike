package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
)

type LikeCmd struct {
	client *Client
}

func (*LikeCmd) Name() string {
	return "like"
}

func (*LikeCmd) Synopsis() string {
	return ""
}

func (*LikeCmd) Usage() string {
	return ""
}

func (*LikeCmd) SetFlags(f *flag.FlagSet) {
}

func (l *LikeCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return subcommands.ExitSuccess
}
