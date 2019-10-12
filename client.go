package main

import (
	"context"
	"github.com/machinebox/graphql"
)

type Client struct {
	client *graphql.Client
	Token  string
}

type NotesResponse struct {
	Notes struct {
		Nodes []struct {
			ID    string
			Title string
		}
	}
}

type SearchResultResponse struct {
	Search struct {
		Nodes []struct {
			Title    string
			Document struct {
				ID      string
				Title   string
				Content string
			}
		}
	}
}

func NewClient(endpoint string, token string) *Client {
	return &Client{
		client: graphql.NewClient(endpoint),
		Token:  token,
	}
}

func (c *Client) applyHeader(req *graphql.Request) {
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}

func (c *Client) Like(id string) error {
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

func (c *Client) GetNotes() (NotesResponse, error) {
	req := graphql.NewRequest(`
		query ($first: Int!) {
			notes(first: $first) {
				nodes {
					id
					title
				}
			}
		}
	`)
	req.Var("first", 5)
	c.applyHeader(req)

	ctx := context.Background()
	resp := NotesResponse{}
	err := c.client.Run(ctx, req, &resp)
	return resp, err
}

func (c *Client) Search(query string) (SearchResultResponse, error) {
	req := graphql.NewRequest(`
		query ($first: Int!, $query: String!) {
			search(first: $first, query: $query) {
				nodes {
					document {
						__typename
						... on Comment {
							id
							content
						}
						... on Note {
							id
							title
						}
					}
					title
				}
			}
		}
	`)
	req.Var("first", 5)
	req.Var("query", query)
	c.applyHeader(req)

	ctx := context.Background()
	resp := SearchResultResponse{}
	err := c.client.Run(ctx, req, &resp)
	return resp, err
}
