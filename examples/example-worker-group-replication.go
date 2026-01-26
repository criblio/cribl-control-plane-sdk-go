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
	"encoding/json"
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

	if workerGroupsList.CountedConfigGroup == nil ||
		workerGroupsList.CountedConfigGroup.Items == nil ||
		len(workerGroupsList.CountedConfigGroup.Items) == 0 {
		fmt.Println("❌ No Worker Groups found. Please create at least one Worker Group first.")
		return
	}

	firstWorkerGroup := workerGroupsList.CountedConfigGroup.Items[0]
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

	if sourceResponse.CountedConfigGroup == nil || len(sourceResponse.CountedConfigGroup.Items) == 0 {
		return nil, fmt.Errorf("Worker Group %q not found", sourceID)
	}

	source := sourceResponse.CountedConfigGroup.Items[0]

	// Generate a unique ID and name for the replica Worker Group
	replicaID := fmt.Sprintf("%s-replica", sourceID)
	replicaName := fmt.Sprintf("%s-replica", getStringValue(source.Name, sourceID))
	replicaDescription := fmt.Sprintf("Replica of '%s'", sourceID)

	fmt.Printf("Creating replica with ID: %s\n", replicaID)

	// Create the replica Worker Group by copying configuration from source
	// Convert ConfigGroup to GroupCreateRequest using JSON marshaling
	sourceBytes, err := json.Marshal(source)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal source Worker Group: %w", err)
	}

	var replicaGroup components.GroupCreateRequest
	if err := json.Unmarshal(sourceBytes, &replicaGroup); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal replica Worker Group: %w", err)
	}

	// Override with replica-specific values
	replicaGroup.ID = replicaID
	replicaGroup.Name = criblcontrolplanesdkgo.String(replicaName)
	replicaGroup.Description = criblcontrolplanesdkgo.String(replicaDescription)

	// Create the replica Worker Group
	result, err := client.Groups.Create(ctx, components.ProductsCoreStream, replicaGroup)
	if err != nil {
		return nil, fmt.Errorf("Failed to create replica Worker Group: %w", err)
	}

	if result.CountedConfigGroup == nil || len(result.CountedConfigGroup.Items) == 0 {
		return nil, fmt.Errorf("Failed to create replica - no response received")
	}

	created := result.CountedConfigGroup.Items[0]
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
