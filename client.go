package main

import (
	"github.com/machinebox/graphql"
)

type Client struct {
	client *graphql.Client
	Token string
}

func NewClient(endpoint string, token string) *Client {
	return &Client{
		client:graphql.NewClient(endpoint),
		Token:token,
	}
}

func (c *Client) applyHeader(req *graphql.Request) {
	req.Header.Add("Authorization", "Bearer " + c.Token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}
