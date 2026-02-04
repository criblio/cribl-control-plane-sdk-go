/**
 * On-Prem/Hybrid Worker Group Process Optimization Example
 *
 * This example demonstrates how to optimize Worker Process settings for an
 * existing on-prem or hybrid Worker Group in Cribl Stream using the Control Plane SDK,
 * following the scaling guidelines from the Cribl documentation.
 *
 * This example performs the following operations:
 *
 * 1. Connects to an existing on-prem or hybrid Worker Group.
 * 2. Retrieves the current system settings for the Worker Group.
 * 3. Optimizes Worker Process settings following the scaling documentation:
 *    - Worker Process count: -2 (to reserve two CPU cores for system/API overhead)
 *    - Memory: 2048 MB (2 GB per Worker Process)
 *    - Minimum Worker Process count: 2 (to spawn at least two Worker Processes)
 * 4. Updates the Worker Group's system settings with the optimized configuration.
 * 5. Commits the configuration changes to the Worker Group.
 * 6. Deploys the configuration changes to the Worker Group.
 * 7. Restarts the Worker Group to apply the changes.
 *
 * The Cribl documentation provides more information about optimizing Worker Processes:
 * https://docs.cribl.io/stream/scaling/#optimize-a-distributed-deployment-or-hybrid-group
 *
 * Prerequisites:
 * - An .env file configured with your authentication credentials (which must have Permissions that include updating Worker Groups and system settings).
 * - A Worker Group whose ID matches the configured WORKER_GROUP_ID value.
 */

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

const (
	WORKER_GROUP_ID = "your-worker-group-id" // Replace with your actual Worker Group ID
)

func main() {
	fmt.Println("On-Prem/Hybrid Worker Group Process Optimization Example")
	fmt.Println("======================================================")
	fmt.Println()

	ctx := context.Background()

	// Initialize Cribl client
	client, err := CreateCriblClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Cribl client: %v", err)
	}

	// Construct the base URL for the Worker Group
	groupURL := fmt.Sprintf("%s/m/%s", BaseURL, WORKER_GROUP_ID)

	// Verify that Worker Group exists
	workerGroupResponse, err := client.Groups.Get(ctx, components.ProductsCoreStream, WORKER_GROUP_ID, nil)
	if err != nil {
		log.Fatalf("Error checking for Worker Group: %v", err)
	}

	if workerGroupResponse.CountedConfigGroup == nil ||
		workerGroupResponse.CountedConfigGroup.Items == nil ||
		len(workerGroupResponse.CountedConfigGroup.Items) == 0 {
		fmt.Printf("⚠️ Worker Group not found: %s. Please create it first.\n", WORKER_GROUP_ID)
		return
	}
	fmt.Printf("✅ Found Worker Group: %s\n", WORKER_GROUP_ID)

	// Get current system settings to preserve existing configuration
	currentSettings, err := client.System.Settings.Cribl.List(ctx, operations.WithServerURL(groupURL))
	if err != nil {
		log.Fatalf("Failed to retrieve current system settings: %v", err)
	}

	if currentSettings.CountedSystemSettingsConf == nil ||
		currentSettings.CountedSystemSettingsConf.Items == nil ||
		len(currentSettings.CountedSystemSettingsConf.Items) == 0 {
		log.Fatal("Failed to retrieve current system settings - no items returned")
	}

	currentConf := currentSettings.CountedSystemSettingsConf.Items[0]

	fmt.Printf("📊 Current Worker Process settings:\n")
	fmt.Printf("   Worker Process count: %.0f\n", currentConf.Workers.Count)
	fmt.Printf("   Memory: %.0f MB\n", currentConf.Workers.Memory)
	fmt.Printf("   Minimum Worker Processes: %.0f\n", currentConf.Workers.Minimum)

	// Configure Worker Process settings following scaling documentation
	// For x86 hyperthreaded CPUs: Process count = -2 (default; reserves 2 CPU cores for system/API overhead)
  // Memory: 2048 MB (default; 2 GB per Worker Process)
  // Minimum: 2 (spawn at least two Worker Processes)
	workersConfig := components.WorkersTypeSystemSettingsConf{
		Count:   -2,  // Negative number specifies the number of CPU cores to reserve for system/API overhead
		Memory:  2048, // Amount of heap memory available to each Worker Process, in MB
		Minimum: 2,    // Minimum number of Worker Processes to spawn
	}

	// Update system settings with the optimized configuration for Worker Processes
	// Preserve other settings from the current configuration
	updatedConf := currentConf
	updatedConf.Workers = workersConfig

	_, err = client.System.Settings.Cribl.Update(ctx, updatedConf, operations.WithServerURL(groupURL))
	if err != nil {
		log.Fatalf("Failed to update system settings: %v", err)
	}

	fmt.Printf("\n✅ Worker Process settings optimized:\n")
	fmt.Printf("   Worker Process count: %.0f\n", workersConfig.Count)
	fmt.Printf("   Memory: %.0f MB per Worker Process\n", workersConfig.Memory)
	fmt.Printf("   Minimum Worker Processes: %.0f\n", workersConfig.Minimum)

	// Commit configuration changes
	effective := true
	commitParams := components.GitCommitParams{
		Message:   "Optimize Worker Process settings",
		Effective: &effective,
		Files:     []string{"."},
	}

	commitResponse, err := client.Versions.Commits.Create(ctx, commitParams, operations.WithServerURL(groupURL))
	if err != nil {
		log.Fatalf("Failed to commit configuration changes: %v", err)
	}

	if commitResponse.CountedGitCommitSummary == nil ||
		commitResponse.CountedGitCommitSummary.Items == nil ||
		len(commitResponse.CountedGitCommitSummary.Items) == 0 {
		log.Fatal("Failed to commit configuration changes - no response received")
	}

	version := commitResponse.CountedGitCommitSummary.Items[0].Commit
	fmt.Printf("✅ Committed configuration changes to the Worker Group: %s, commit ID: %s\n", WORKER_GROUP_ID, version)

	// Deploy configuration changes to the Worker Group
	deployRequest := components.DeployRequest{
		Version: version,
	}

	_, err = client.Groups.Deploy(ctx, components.ProductsCoreStream, WORKER_GROUP_ID, deployRequest)
	if err != nil {
		log.Fatalf("Failed to deploy Worker Group changes: %v", err)
	}
	fmt.Printf("✅ Worker Group changes deployed: %s\n", WORKER_GROUP_ID)

	// Restart the Worker Group to apply the changes
	fmt.Printf("\n🔄 Restarting Worker Group to apply changes...\n")
	_, err = client.System.Settings.Restart(ctx, operations.WithServerURL(groupURL))
	if err != nil {
		log.Fatalf("Failed to restart Worker Group: %v", err)
	}
	fmt.Printf("✅ Worker Group restarted successfully\n")
}
