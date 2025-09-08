/*
On-Premises Authentication Example

This example demonstrates how to authenticate with an on-premises Cribl instance
using username and password credentials. It shows the two-step authentication process:

1. Authenticates with username/password to obtain a bearer token
2. Creates a new SDK client using the obtained token for API calls
3. Validates the connection by checking server health status

Prerequisites: Configure the on-premises server URL, username, and password

Note: This example is for onprem deployments only and does not require a .env
file configuration to run.
*/

package main

import (
	"context"
	"fmt"
	"strings"

	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

const (
	ONPREM_SERVER_URL = "http://localhost:9000"
	ONPREM_USERNAME   = "admin"
	ONPREM_PASSWORD   = "admin"
)

func main() {
	ctx := context.Background()
	baseURL := fmt.Sprintf("%s/api/v1", ONPREM_SERVER_URL)

	// Retrieve authentication token
	tokenClient := criblcontrolplanesdkgo.New(baseURL)

	loginInfo := components.LoginInfo{
		Username: ONPREM_USERNAME,
		Password: ONPREM_PASSWORD,
	}

	response, err := tokenClient.Auth.Tokens.Get(ctx, loginInfo)
	if err != nil {
		handleError(err)
		return
	}

	if response.AuthToken == nil || response.AuthToken.Token == "" {
		fmt.Println("❌ No token received from authentication response")
		return
	}

	token := response.AuthToken.Token
	fmt.Printf("✅ Authenticated with on-premises server, token: %s\n", token)

	// Create authenticated client
	client := criblcontrolplanesdkgo.New(
		baseURL,
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: &token,
		}),
	)
	fmt.Println("✅ Cribl SDK client created for on-premises server")

	// Validate connection, try to list all git branches
	branchResponse, err := client.Versions.Branches.List(ctx)
	if err != nil {
		handleError(err)
		return
	}

	if branchResponse.Object == nil || branchResponse.Object.Items == nil {
		fmt.Println("⚠️ No branches found")
		return
	}

	var branches []string
	for _, branch := range branchResponse.Object.Items {
		if branch.ID != "" {
			branches = append(branches, branch.ID)
		}
	}

	fmt.Printf("✅ Client works! Your list of branches:\n\t%s\n", strings.Join(branches, "\n\t"))
}

func handleError(err error) {
	errStr := err.Error()
	if strings.Contains(errStr, "429") || strings.Contains(errStr, "rate limit") {
		fmt.Println("⚠️ Uh oh, you've hit the rate limit! Try again in a few seconds.")
	} else {
		fmt.Printf("❌ Something went wrong: %v\n", err)
	}
}
