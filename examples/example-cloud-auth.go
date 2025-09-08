/*
Cloud Authentication Example

This example demonstrates how to authenticate with Cribl Cloud using OAuth2
client credentials flow. It shows the authentication process:

1. Creates an SDK client with OAuth2 client credentials configuration
2. Automatically handles token exchange and refresh
3. Validates the connection by checking server health status

Prerequisites: Configure cloud environment variables: ORG_ID, CLIENT_ID, CLIENT_SECRET, WORKSPACE_NAME, CRIBL_DOMAIN
How to get these values: https://docs.cribl.io/api/#criblcloud

Note: This example is for cloud deployments only and does not require a .env
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

// Cloud configuration - UPDATE THESE VALUES
const (
	ORG_ID         = "your-org-id"        // Replace with your organization ID
	CLIENT_ID      = "your-client-id"     // Replace with your OAuth2 client ID
	CLIENT_SECRET  = "your-client-secret" // Replace with your OAuth2 client secret
	WORKSPACE_NAME = "main"               // Replace with your workspace name
)

func main() {
	ctx := context.Background()
	baseURL := fmt.Sprintf("https://%s-%s.cribl.cloud/api/v1", WORKSPACE_NAME, ORG_ID)

	// Create authenticated client with OAuth2
	client := criblcontrolplanesdkgo.New(
		baseURL,
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			ClientOauth: &components.SchemeClientOauth{
				ClientID:     CLIENT_ID,
				ClientSecret: CLIENT_SECRET,
				TokenURL:     "https://login.cribl.cloud/oauth/token",
				Audience:     "https://api.cribl.cloud",
			},
		}),
	)
	fmt.Println("✅ Cribl SDK client created for cloud deployment")

	// Validate connection, try to list all git branches
	response, err := client.Versions.Branches.List(ctx)
	if err != nil {
		handleError(err)
		return
	}

	if response.Object == nil || response.Object.Items == nil {
		fmt.Println("⚠️ No branches found")
		return
	}

	var branches []string
	for _, branch := range response.Object.Items {
		if branch.ID != "" {
			branches = append(branches, branch.ID)
		}
	}

	fmt.Printf("✅ Client works! Your list of branches:\n\t%s\n", strings.Join(branches, "\n\t"))
}

func handleError(err error) {
	// Note: You might need to adjust error checking based on the actual error types in the Go SDK
	fmt.Printf("❌ Something went wrong: %v\n", err)

	// Basic error message checks - you may need to refine these based on actual SDK error types
	errStr := err.Error()
	if strings.Contains(errStr, "401") || strings.Contains(errStr, "unauthorized") {
		fmt.Println("⚠️ Authentication failed! Check your CLIENT_ID and CLIENT_SECRET.")
	} else if strings.Contains(errStr, "429") || strings.Contains(errStr, "rate limit") {
		fmt.Println("⚠️ Uh oh, you've hit the rate limit! Try again in a few seconds.")
	}
}
