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
	return subcommands.ExitSuccess
}
