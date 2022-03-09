package github

import (
	"context"

	"github.com/cli/oauth"
	"github.com/cli/oauth/api"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func Login() (*api.AccessToken, error) {
	flow := &oauth.Flow{
		Host:     oauth.GitHubHost("https://github.com"),
		ClientID: "Iv1.3db883241a360e1c",
		Scopes:   []string{"repo", "admin:org", "workflow", "read:packages", "write:packages"},
	}

	return flow.DetectFlow()
}

func GetClient(token string) *githubv4.Client {
	src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	httpClient := oauth2.NewClient(context.Background(), src)
	return githubv4.NewClient(httpClient)
}
