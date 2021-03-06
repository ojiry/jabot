package main

import (
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func main() {
	accessToken := os.Getenv("ACCESS_TOKEN")
	query := os.Getenv("QUERY")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	results, _, err := client.Search.Issues(ctx, query, nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(*results.Total)
}
