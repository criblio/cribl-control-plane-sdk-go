/**
 * Cribl.Cloud Worker Group Management Example
 *
 * This example demonstrates how to create and manage a Cribl.Cloud-based
 * Worker Group in Cribl Stream using the Control Plane SDK.
 *
 * 1. Create a Cribl.Cloud Worker Group using the AWS provider configuration.
 * Set the maximum estimated ingest rate to 24 MB/s and configure 9
 * Worker Processes.
 * 2. Verify that the Worker Group doesn't already exist.
 * 3. Scale up the Worker Group to 48 MB/s maximum estimated ingest rate and
 * 21 Worker Processes.
 * 4. Provision the Worker Group to activate Cribl.Cloud resources.
 *
 * The Cribl documentation provides more information about ingest rates and
 * provisioning:
 * https://docs.cribl.io/stream/cloud-workers/#create-a-cloud-worker-group
 *
 * Prerequisites: Replace the placeholder values for ORG_ID, CLIENT_ID,
 * CLIENT_SECRET, and WORKSPACE_NAME with your Organization ID, Client ID and
 * Secret, and Workspace name. To get your CLIENT_ID and CLIENT_SECRET values,
 * follow the steps at
 * https://docs.cribl.io/cribl-as-code/authentication/#cloud-auth.
 * Your Client ID and Secret are sensitive information and should be kept private.
 *
 * NOTE: This example is for Cribl.Cloud deployments only. It does not require
 * .env file configuration.
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

	// Create the Worker Group
	awsProvider := components.CloudProviderAws
	group := components.ConfigGroup{
		ID:                  WORKER_GROUP_ID,
		Name:                criblcontrolplanesdkgo.String("my-aws-worker-group"),
		OnPrem:              criblcontrolplanesdkgo.Bool(false),
		WorkerRemoteAccess:  criblcontrolplanesdkgo.Bool(true),
		Provisioned:         criblcontrolplanesdkgo.Bool(false),
		IsFleet:             criblcontrolplanesdkgo.Bool(false),
		IsSearch:            criblcontrolplanesdkgo.Bool(false),
		EstimatedIngestRate: criblcontrolplanesdkgo.Float64(2048), // Equivalent to 24 MB/s maximum estimated ingest rate with 9 Worker Processes
		Cloud: &components.ConfigGroupCloud{
			Provider: &awsProvider,
			Region:   "us-east-1",
		},
	}

	createResponse, err := client.Groups.Create(ctx, components.ProductsCoreStream, group)
	if err != nil {
		log.Fatalf("Error creating Worker Group: %v", err)
	}

	if createResponse.Object == nil {
		log.Fatal("No response received when creating Worker Group")
	}

	fmt.Printf("✅ Worker Group created: %s\n", WORKER_GROUP_ID)

	// Scale and provision the Worker Group
	group.EstimatedIngestRate = criblcontrolplanesdkgo.Float64(4096) // Equivalent to 48 MB/s maximum estimated ingest rate with 21 Worker Processes
	group.Provisioned = criblcontrolplanesdkgo.Bool(true)

	updateResponse, err := client.Groups.Update(ctx, components.ProductsCoreStream, WORKER_GROUP_ID, group)
	if err != nil {
		log.Fatalf("Error updating Worker Group: %v", err)
	}

	if updateResponse.Object == nil {
		log.Fatal("No response received when updating Worker Group")
	}

	fmt.Printf("✅ Worker Group scaled and provisioned: %s\n", WORKER_GROUP_ID)
}
