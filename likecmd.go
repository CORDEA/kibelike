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
	arg := f.Arg(0)
	if arg == "" {
		return subcommands.ExitFailure
	}
	if err := l.client.Like(arg); err != nil {
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
