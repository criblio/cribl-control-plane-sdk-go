/*
Cribl.Cloud Lake Dataset with Search Pack Example

This example demonstrates how to install a Search Pack and create a Lake Dataset
using the Control Plane SDK with Cribl.Cloud authentication.

The example:
1. Installs the AWS VPC Flow Logs Search Pack from Cribl Packs Dispensary
2. Creates a Lake Dataset with basic configuration

Prerequisites: Replace the placeholder values for ORG_ID, CLIENT_ID,
CLIENT_SECRET, and WORKSPACE_NAME with your Organization ID, Client ID and
Secret, and Workspace name. To get your CLIENT_ID and CLIENT_SECRET values,
follow the steps at 
https://docs.cribl.io/cribl-as-code/cribl-as-code/authentication/#cloud-auth. 
Your Client ID
and Secret are sensitive information and should be kept private.

NOTE: This example is for Cribl.Cloud deployments only. It does not require
.env file configuration.
*/

package main

import (
	"context"
	"fmt"

	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

// Configuration constants - Replace with your actual values
const (
	ORG_ID        = "your-org-id"         // Replace with your Organization ID
	CLIENT_ID     = "your-client-id"      // Replace with your Client ID
	CLIENT_SECRET = "your-client-secret"  // Replace with your Client Secret
	WORKSPACE     = "your-workspace-name" // Replace with your Workspace name

	PACK_URL   = "https://packs.cribl.io/dl/cribl-search-aws-vpc-flow-logs/0.1.1/cribl-search-aws-vpc-flow-logs-0.1.1.crbl"
	PACK_ID    = "cribl-search-aws-vpc-flow-logs"
	LAKE_ID    = "default"
	DATASET_ID = "aws-vpc-flow-logs-dataset"
)

func main() {
	fmt.Println("Cribl.Cloud Lake Dataset with Search Pack Example")
	fmt.Println("=================================================")
	fmt.Println()

	ctx := context.Background()

	// Create the SDK client with OAuth2 authentication
	s := criblcontrolplanesdkgo.New(
		fmt.Sprintf("https://%s.cribl.cloud", WORKSPACE),
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			ClientOauth: &components.SchemeClientOauth{
				ClientID:     CLIENT_ID,
				ClientSecret: CLIENT_SECRET,
				TokenURL:     "https://login.cribl.cloud/oauth/token",
				Audience:     "https://api.cribl.cloud",
			},
		}),
	)

	// Install Search Pack
	fmt.Printf("📦 Installing Search Pack: %s\n", PACK_ID)
	installReq := components.CreatePackRequestBodyUnionPackRequestBody1(
		components.PackRequestBody1{
			ID:     PACK_ID,
			Source: criblcontrolplanesdkgo.String(PACK_URL),
		},
	)

	searchGroupURL := fmt.Sprintf("https://%s.cribl.cloud/api/v1/search", WORKSPACE)
	s.Packs.Install(ctx, installReq, operations.WithServerURL(searchGroupURL))

	// Create Lake Dataset
	fmt.Printf("Creating Lake Dataset: %s\n", DATASET_ID)
	retention := float64(30)
	httpDAUsed := false
	storageLocationID := "cribl_lake"

	s.LakeDatasets.Create(ctx, LAKE_ID, components.CriblLakeDataset{
		ID:                    DATASET_ID,
		RetentionPeriodInDays: &retention,
		HTTPDAUsed:            &httpDAUsed,
		StorageLocationID:     &storageLocationID,
	})
}
