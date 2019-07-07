package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
	"github.com/mitchellh/go-homedir"
)

func main() {
	// Wrap the shared transport for use with the integration ID 1 authenticating with installation ID 99.
	fp, _ := homedir.Expand("~/Downloads/kb.2019-07-07.private-key.pem")
	itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, 1, 1, fp)
	if err != nil {
		// Handle error.
	}
	itr.BaseURL = "https://ghe.metrumrg.dev/api/v3"
	token, err := itr.Token()
	if err != nil {
		fmt.Println(err)
		panic("error with token")
	}
	fmt.Println(token)
	// Use installation transport with client.

	// Use client...
	client, err := github.NewEnterpriseClient(itr.BaseURL, itr.BaseURL, &http.Client{Transport: itr})
	if err != nil {
		panic(err)
	}

	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "mrg", opt)
	if err != nil {
		panic(err)
	}
	fmt.Println(repos)
}
