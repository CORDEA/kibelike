package main

import (
	"context"
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

func (c *Client) Like(id string) error  {
	req := graphql.NewRequest(`
		mutation Like($id: ID!) {
			like(input: {likableId: $id}) {
				clientMutationId
			}
		}
	`)
	req.Var("id", id)
	c.applyHeader(req)

	ctx := context.Background()
	return c.client.Run(ctx, req, nil)
}