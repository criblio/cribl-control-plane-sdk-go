/**
 * Cribl.Cloud and Hybrid Authentication Example
 *
 * This example demonstrates how to configure authentication on Cribl.Cloud
 * and in hybrid deployments using OAuth2 credentials.
 *
 * 1. Create an SDK client with OAuth2 client credentials using the 
 * ClientOauth security scheme.
 * 2. Automatically handle token exchange and refresh.
 * 3. Validate the connection by listing all git branches.
 *
 * Prerequisites: Replace the placeholder values for ORG_ID, CLIENT_ID,
 * CLIENT_SECRET, and WORKSPACE_NAME with your Organization ID, Client ID and
 * Secret, and Workspace name. To get your Client ID and Secret, follow the 
 * steps at https://docs.cribl.io/cribl-as-code/sdks-auth/#sdks-auth-cloud. 
 * Your Client ID and Secret are sensitive information and should be kept private.
 *
 * NOTE: This example is for Cribl.Cloud and hybrid deployments only. 
 * It does not require .env file configuration.
 */

package main

import (
	"context"
	"fmt"
	"strings"

	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

// Cribl.Cloud and hybrid configuration: Replace the placeholder values
const (
	CLIENT_ID      = "your-client-id"     // Replace with your OAuth2 Client ID
	CLIENT_SECRET  = "your-client-secret" // Replace with your OAuth2 Client Secret
	ORG_ID         = "your-org-id"        // Replace with your Organization ID
	WORKSPACE_NAME = "main"               // Replace with your Workspace name
)

func main() {
	ctx := context.Background()
	baseURL := fmt.Sprintf("https://%s-%s.cribl.cloud/api/v1", WORKSPACE_NAME, ORG_ID)

	// Create authenticated SDK client with OAuth2
	client := criblcontrolplanesdkgo.New(
		criblcontrolplanesdkgo.WithServerURL(baseURL),
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			ClientOauth: &components.SchemeClientOauth{
				ClientID:     CLIENT_ID,
				ClientSecret: CLIENT_SECRET,
				TokenURL:     "https://login.cribl.cloud/oauth/token",
				Audience:     "https://api.cribl.cloud",
			},
		}),
	)
	fmt.Println("✅ Cribl SDK client created for Cribl.Cloud deployment")

	// Validate connection and list all git branches
	response, err := client.Versions.Branches.List(ctx)
	if err != nil {
		handleError(err)
		return
	}

	if response.CountedBranchInfo == nil || response.CountedBranchInfo.Items == nil {
		fmt.Println("⚠️ No branches found")
		return
	}

	var branches []string
	for _, branch := range response.CountedBranchInfo.Items {
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
		fmt.Println("⚠️ Uh oh, you've reached the rate limit! Try again in a few seconds.")
	}
}
