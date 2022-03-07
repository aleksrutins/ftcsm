package github

import (
	"github.com/cli/oauth"
	"github.com/cli/oauth/api"
)

func Login() (*api.AccessToken, error) {
	flow := &oauth.Flow{
		Host:     oauth.GitHubHost("https://github.com"),
		ClientID: "Iv1.3db883241a360e1c",
		Scopes:   []string{"repo", "admin:org", "workflow"},
	}

	return flow.DetectFlow()
}
