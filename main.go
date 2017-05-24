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
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opts := &github.SearchOptions{}
	results, _, err := client.Search.Issues(ctx, "is:pr+commenter:ojiry", opts)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(*results.Total)
}
