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

	opt := &github.NotificationListOptions{
		All:           true,
		Participating: true,
	}
	notifications, _, err := client.Activity.ListNotifications(ctx, opt)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < len(notifications); i++ {
		fmt.Println(*notifications[i].URL)
	}
}
