/**
 * Cribl Worker Group Replication Example
 *
 * This example demonstrates how to programmatically replicate an existing
 * Worker Group configuration using the Control Plane SDK.
 *
 * This example performs the following operations:
 *
 * 1. Retrieves a list of all Worker Groups in Cribl Stream.
 * 2. Selects the first Worker Group in the list as the source Worker Group
 * to replicate.
 * 3. Retrieves the complete configuration of the source Worker Group.
 * 4. Creates a new Worker Group that uses the same configuration as the source
 * Worker Group. The replica Worker Group has a unique ID and a name and
 * description that identify it as a replica.
 *
 * Data flow: Source Worker Group → Configuration Extraction → New Worker Group Creation
 *
 * Prerequisites:
 * - An .env file configured with your authentication credentials.
 * - At least one existing Worker Group in Cribl Stream.
 * - API Bearer token with Permissions that include creating Worker Groups.
 */

package main

import (
	"context"
	"fmt"
	"log"

	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

func main() {
	fmt.Println("Cribl Worker Group Replication Example")
	fmt.Println("======================================")
	fmt.Println()

	ctx := context.Background()

	// Initialize Cribl client
	client, err := CreateCriblClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Cribl client: %v", err)
	}

	fmt.Println("🔍 Searching for existing Worker Groups...")

	// Get the first available Worker Group
	workerGroupsList, err := client.Groups.List(ctx, components.ProductsCoreStream, nil)
	if err != nil {
		log.Fatalf("Error listing Worker Groups: %v", err)
	}

	if workerGroupsList.CountedListConfigGroup == nil ||
		workerGroupsList.CountedListConfigGroup.Items == nil ||
		len(workerGroupsList.CountedListConfigGroup.Items) == 0 {
		fmt.Println("❌ No Worker Groups found. Please create at least one Worker Group first.")
		return
	}

	firstWorkerGroup := workerGroupsList.CountedListConfigGroup.Items[0]
	fmt.Printf("📋 Found Worker Group to replicate: %s\n", firstWorkerGroup.ID)

	// Replicate the first Worker Group
	replicatedGroup, err := replicateWorkerGroup(ctx, client, firstWorkerGroup.ID)
	if err != nil {
		log.Fatalf("Failed to replicate Worker Group: %v", err)
	}

	if replicatedGroup != nil {
		fmt.Println("🎉 Worker Group replication completed successfully!")
		fmt.Printf("ℹ️ Original: %s\n", firstWorkerGroup.ID)
		fmt.Printf("ℹ️ Replica: %s\n", replicatedGroup.ID)
	}
}

/**
 * Replicates a Worker Group with a unique ID
 *
 * - Retrieve the source Worker Group configuration
 * - Generate a unique ID and name to use for the replica Worker Group
 * - Filter out read-only fields from the source Worker Group configuration
 * - Create the replica Worker Group with proper metadata
 */
func replicateWorkerGroup(ctx context.Context, client *criblcontrolplanesdkgo.CriblControlPlane, sourceID string) (*components.ConfigGroup, error) {
	fmt.Printf("📖 Retrieving configuration for Worker Group: %s\n", sourceID)

	// Retrieve the source Worker Group configuration
	sourceResponse, err := client.Groups.Get(ctx, components.ProductsCoreStream, sourceID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get Worker Group '%s': %w", sourceID, err)
	}

	if sourceResponse.CountedListConfigGroup == nil || len(sourceResponse.CountedListConfigGroup.Items) == 0 {
		return nil, fmt.Errorf("Worker Group %q not found", sourceID)
	}

	source := sourceResponse.CountedListConfigGroup.Items[0]

	// Generate a unique ID and name for the replica Worker Group
	replicaID := fmt.Sprintf("%s-replica", sourceID)
	replicaName := fmt.Sprintf("%s-replica", getStringValue(source.Name, sourceID))
	replicaDescription := fmt.Sprintf("Replica of '%s'", sourceID)

	fmt.Printf("Creating replica with ID: %s\n", replicaID)

	// Convert ConfigGroup to GroupCreateRequest for Create operation
	// Note: ConfigVersion is omitted as it's a read-only field
	var estimatedIngestRate *components.GroupCreateRequestEstimatedIngestRate
	if source.EstimatedIngestRate != nil {
		rate := components.GroupCreateRequestEstimatedIngestRate(*source.EstimatedIngestRate)
		estimatedIngestRate = &rate
	}

	var git *components.GroupCreateRequestGit
	if source.Git != nil {
		git = &components.GroupCreateRequestGit{
			Commit:       source.Git.Commit,
			LocalChanges: source.Git.LocalChanges,
			Log:          source.Git.Log,
		}
	}

	var groupType *components.GroupCreateRequestType
	if source.Type != nil {
		t := components.GroupCreateRequestType(*source.Type)
		groupType = &t
	}

	replicaGroup := components.GroupCreateRequest{
		Cloud:                   source.Cloud,
		DeployingWorkerCount:    source.DeployingWorkerCount,
		Description:             criblcontrolplanesdkgo.String(replicaDescription),
		EstimatedIngestRate:     estimatedIngestRate,
		Git:                     git,
		ID:                      replicaID,
		IncompatibleWorkerCount: source.IncompatibleWorkerCount,
		Inherits:                source.Inherits,
		IsFleet:                 source.IsFleet,
		IsSearch:                source.IsSearch,
		LookupDeployments:       source.LookupDeployments,
		MaxWorkerAge:            source.MaxWorkerAge,
		Name:                    criblcontrolplanesdkgo.String(replicaName),
		OnPrem:                  source.OnPrem,
		Provisioned:             source.Provisioned,
		Streamtags:              source.Streamtags,
		Tags:                    source.Tags,
		Type:                    groupType,
		UpgradeVersion:          source.UpgradeVersion,
		WorkerCount:             source.WorkerCount,
		WorkerRemoteAccess:      source.WorkerRemoteAccess,
	}

	// Create the replica Worker Group
	result, err := client.Groups.Create(ctx, components.ProductsCoreStream, replicaGroup)
	if err != nil {
		return nil, fmt.Errorf("Failed to create replica Worker Group: %w", err)
	}

	if result.CountedListConfigGroup == nil || len(result.CountedListConfigGroup.Items) == 0 {
		return nil, fmt.Errorf("Failed to create replica - no response received")
	}

	created := result.CountedListConfigGroup.Items[0]
	fmt.Printf("Worker Group replicated: %s\n", created.ID)

	return &created, nil
}

// getStringValue safely extracts string value from pointer, with fallback
func getStringValue(ptr *string, fallback string) string {
	if ptr != nil {
		return *ptr
	}
	return fallback
}
