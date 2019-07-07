package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
)

func main() {
	// Wrap the shared transport for use with the integration ID 1 authenticating with installation ID 99.
	//fp, _ := homedir.Expand("~/Downloads/kb.2019-07-07.private-key.pem")
	decoded, _ := base64.StdEncoding.DecodeString(os.Getenv("GHE_APP_KB_B64PEM"))
	fmt.Println("decoded: ", string(decoded))
	tmpPEM := filepath.Join(os.TempDir(), "temp.PEM")
	err := ioutil.WriteFile(tmpPEM, decoded, 0644)
	if err != nil {
		panic(err)
	}
	// defer here in case ghinstall fails
	defer os.Remove(tmpPEM)
	itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, 1, 1, tmpPEM)
	if err != nil {
		// Handle error.
		panic(err)
	}
	os.Remove(tmpPEM)
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
