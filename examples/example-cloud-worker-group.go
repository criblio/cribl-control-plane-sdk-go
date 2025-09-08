/*
Cloud Worker Group Management Example

This example demonstrates how to create and manage a cloud-based worker group
in Cribl Stream using the Control Plane SDK. It shows:

1. Creates a cloud worker group with AWS provider configuration with
   24MB/s max estimated ingest rate / 9 worker processes
2. Verifies the worker group doesn't already exist
3. Scales the worker group up to 48MB/s max estimated ingest rate / 21 worker processes
4. Provisions the worker group to activate cloud resources

Check this documentation for more details about ingest rates and provisioning:
  https://docs.cribl.io/stream/cloud-workers/#create-a-cloud-worker-group

Prerequisites: Configure cloud environment variables: ORG_ID, CLIENT_ID, CLIENT_SECRET, WORKSPACE_NAME
How to get these values: https://docs.cribl.io/api/#criblcloud

Note: This example is for cloud deployments only and does not require a .env
file configuration to run.
*/

package main

import (
	"context"
	"fmt"
	"log"

	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

const (
	ORG_ID          = "your-org-id"
	CLIENT_ID       = "your-client-id"
	CLIENT_SECRET   = "your-client-secret"
	WORKSPACE_NAME  = "your-workspace-name"
	WORKER_GROUP_ID = "your-cloud-worker-group-id"
)

func main() {
	ctx := context.Background()
	baseURL := fmt.Sprintf("https://%s-%s.cribl.cloud/api/v1", WORKSPACE_NAME, ORG_ID)

	// Create authenticated client
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

	// Verify worker group doesn't already exist
	getResponse, err := client.Groups.Get(ctx, components.ProductsCoreStream, WORKER_GROUP_ID, nil)
	if err != nil {
		log.Fatalf("Error checking for existing worker group: %v", err)
	}

	if getResponse.Object != nil &&
		getResponse.Object.Items != nil &&
		len(getResponse.Object.Items) > 0 {
		fmt.Printf("⚠️ Worker Group already exists: %s. Try different group id.\n", WORKER_GROUP_ID)
		return
	}

	// Create the worker group
	awsProvider := components.CloudProviderAws
	group := components.ConfigGroup{
		ID:                  WORKER_GROUP_ID,
		Name:                criblcontrolplanesdkgo.String("my-aws-worker-group"),
		OnPrem:              criblcontrolplanesdkgo.Bool(false),
		WorkerRemoteAccess:  criblcontrolplanesdkgo.Bool(true),
		Provisioned:         criblcontrolplanesdkgo.Bool(false),
		IsFleet:             criblcontrolplanesdkgo.Bool(false),
		IsSearch:            criblcontrolplanesdkgo.Bool(false),
		EstimatedIngestRate: criblcontrolplanesdkgo.Float64(2048), // 24MB/s Max est ingest rate / 9 Worker Processes
		Cloud: &components.ConfigGroupCloud{
			Provider: &awsProvider,
			Region:   "us-east-1",
		},
	}

	createResponse, err := client.Groups.Create(ctx, components.ProductsCoreStream, group)
	if err != nil {
		log.Fatalf("Error creating worker group: %v", err)
	}

	if createResponse.Object == nil {
		log.Fatal("No response received when creating worker group")
	}

	fmt.Printf("✅ Worker Group created: %s\n", WORKER_GROUP_ID)

	// Scale and provision the worker group
	group.EstimatedIngestRate = criblcontrolplanesdkgo.Float64(4096) // 48MB/s Max est ingest rate / 21 Worker Processes
	group.Provisioned = criblcontrolplanesdkgo.Bool(true)

	updateResponse, err := client.Groups.Update(ctx, components.ProductsCoreStream, WORKER_GROUP_ID, group)
	if err != nil {
		log.Fatalf("Error updating worker group: %v", err)
	}

	if updateResponse.Object == nil {
		log.Fatal("No response received when updating worker group")
	}

	fmt.Printf("✅ Worker Group scaled and provisioned: %s\n", WORKER_GROUP_ID)
}
