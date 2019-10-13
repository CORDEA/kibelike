package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
)

type UnlikeCmd struct {
	client *Client
}

func (*UnlikeCmd) Name() string {
	return "like"
}

func (*UnlikeCmd) Synopsis() string {
	return ""
}

func (*UnlikeCmd) Usage() string {
	return ""
}

func (*UnlikeCmd) SetFlags(f *flag.FlagSet) {
}

func (u *UnlikeCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	arg := f.Arg(0)
	if arg == "" {
		return subcommands.ExitFailure
	}
	if err := u.client.Unlike(arg); err != nil {
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
