/**
 * On-Prem Authentication Example
 *
 * This example demonstrates the authentication process for an on-prem Cribl
 * instance using username and password credentials.
 *
 * 1. Authenticate with your username and password to obtain a Bearer token.
 * 2. Create an SDK client that uses the Bearer token for API calls.
 * 3. Validate the connection by checking the server health status and listing
 * all git branches.
 *
 * Prerequisites: Replace the placeholder values for ONPREM_SERVER_URL
 * ONPREM_USERNAME, and ONPREM_PASSWORD with your server URL and credentials.
 * Your credentials are sensitive information and should be kept private.
 *
 * NOTE: This example is for on-prem deployments only. It does not require .env
 * file configuration.
 */

package main

import (
	"context"
	"fmt"
	"strings"

	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

// On-prem configuration: Replace the placeholder values
const (
	ONPREM_SERVER_URL = "http://localhost:9000" // Replace with your server URL
	ONPREM_USERNAME   = "admin"                 // Replace with your username
	ONPREM_PASSWORD   = "admin"                 // Replace with your password
)

func main() {
	ctx := context.Background()
	baseURL := fmt.Sprintf("%s/api/v1", ONPREM_SERVER_URL)

	// Retrieve Bearer token for authentication
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
	fmt.Printf("✅ Authenticated with on-prem server, Token: %s\n", token)

	// Create authenticated SDK client
	client := criblcontrolplanesdkgo.New(
		baseURL,
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: &token,
		}),
	)
	fmt.Println("✅ Cribl SDK client created for on-prem server")

	// Validate connection, and list all git branches
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
		fmt.Println("⚠️ Uh oh, you've reached the rate limit! Try again in a few seconds.")
	} else {
		fmt.Printf("❌ Something went wrong: %v\n", err)
	}
}
