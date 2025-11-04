/**
 * Cribl Stream Configuration Example
 *
 * This example demonstrates how to programmatically create and configure a
 * complete data pipeline in Cribl Stream using the Control Plane SDK.
 *
 * This example creates:
 *
 * 1. A Worker Group to manage the configuration.
 * 2. A TCP JSON source to receive data on port 9020.
 * 3. A Filesystem Destination to store processed data.
 * 4. A Pipeline that filters events and keeps only data in the "name"
 * field.
 * 5. A Route that connects the Source to the Pipeline and Destination.
 *
 * This example also deploys the configuration to the Worker Group to make it
 * active.
 *
 * Data flow: TCP JSON Source → Route → Pipeline → Filesystem Destination
 *
 * The example includes error handling and checks for existing resources.
 *
 * Prerequisites:
 * - An .env file that is configured with your credentials.
 * - An Enterprise License on the server.
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

const (
	PORT            = 9020
	AUTH_TOKEN      = "4a4b3663-7a57-7369-7632-795553573668"
	WORKER_GROUP_ID = "my-worker-group"
)

func main() {
	ctx := context.Background()

	// Initialize Cribl client
	client, err := CreateCriblClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Cribl client: %v", err)
	}

	groupURL := fmt.Sprintf("%s/m/%s", BaseURL, WORKER_GROUP_ID)

	// Check if Worker Group already exists
	getResponse, err := client.Groups.Get(ctx, components.ProductsCoreStream, WORKER_GROUP_ID, nil)
	if err != nil {
		log.Printf("Worker Group doesn't exist yet, will create: %v", err)
	}

	if getResponse != nil && getResponse.Object != nil &&
		getResponse.Object.Items != nil &&
		len(getResponse.Object.Items) > 0 {
		fmt.Printf("⚠️ Worker Group already exists: %s. Using existing group.\n", WORKER_GROUP_ID)
	} else {
		// Create Worker Group
		myWorkerGroup := components.GroupCreateRequest{
			ID:          WORKER_GROUP_ID,
			Description: criblcontrolplanesdkgo.String("My Worker Group"),
			OnPrem:      criblcontrolplanesdkgo.Bool(true),
		}

		createResponse, err := client.Groups.Create(ctx, components.ProductsCoreStream, myWorkerGroup)
		if err != nil {
			log.Printf("Error creating Worker Group: %v", err)
		} else if createResponse != nil && createResponse.Object != nil {
			fmt.Printf("✅ Worker Group created: %s\n", WORKER_GROUP_ID)
		}
	}

	// Create TCP JSON Source
	tcpJSONConfig := map[string]interface{}{
		"id":        "my-tcp-json",
		"type":      "tcpjson",
		"port":      PORT,
		"authType":  "manual",
		"authToken": AUTH_TOKEN,
	}

	// Convert to components.Input using JSON marshaling/unmarshaling
	sourceBytes, err := json.Marshal(tcpJSONConfig)
	if err != nil {
		log.Printf("Error marshaling TCP JSON Source config: %v", err)
	} else {
		var tcpJSONSource components.Input
		err = json.Unmarshal(sourceBytes, &tcpJSONSource)
		if err != nil {
			log.Printf("Error unmarshaling TCP JSON Source config: %v", err)
		} else {
			_, err = client.Sources.Create(ctx, tcpJSONSource, operations.WithServerURL(groupURL))
			if err != nil {
				log.Printf("Error creating TCP JSON Source: %v", err)
			} else {
				fmt.Printf("✅ TCP JSON Source created: my-tcp-json\n")
			}
		}
	}

	// Create Filesystem Destination
	fileSystemConfig := map[string]interface{}{
		"id":             "my-fs-destination",
		"type":           "filesystem",
		"destPath":       "/tmp/my-output",
		"fileNameSuffix": "\".log\"", // JavaScript expression that returns ".log"
	}

	// Convert to components.Output using JSON marshaling/unmarshaling
	destBytes, err := json.Marshal(fileSystemConfig)
	if err != nil {
		log.Printf("Error marshaling Filesystem Destination config: %v", err)
	} else {
		var fileSystemDestination components.Output
		err = json.Unmarshal(destBytes, &fileSystemDestination)
		if err != nil {
			log.Printf("Error unmarshaling Filesystem Destination config: %v", err)
		} else {
			_, err = client.Destinations.Create(ctx, fileSystemDestination, operations.WithServerURL(groupURL))
			if err != nil {
				log.Printf("Error creating Filesystem Destination: %v", err)
			} else {
				fmt.Printf("✅ Filesystem Destination created: my-fs-destination\n")
			}
		}
	}

	// Create Pipeline
	pipelineConf := map[string]interface{}{
		"asyncFuncTimeout": 1000,
		"functions": []map[string]interface{}{
			{
				"filter": "true",
				"conf": map[string]interface{}{
					"remove": []string{"*"},
					"keep":   []string{"name"},
				},
				"id":    "eval",
				"final": true,
			},
		},
	}

	// Convert to components.PipelineConf using JSON marshaling/unmarshaling
	confBytes, err := json.Marshal(pipelineConf)
	if err != nil {
		log.Printf("Error marshaling Pipeline config: %v", err)
	} else {
		var conf components.PipelineConf
		err = json.Unmarshal(confBytes, &conf)
		if err != nil {
			log.Printf("Error unmarshaling Pipeline config: %v", err)
		} else {
			pipeline := components.Pipeline{
				ID:   "my-pipeline",
				Conf: conf,
			}

			_, err = client.Pipelines.Create(ctx, pipeline, operations.WithServerURL(groupURL))
			if err != nil {
				log.Printf("Error creating Pipeline: %v", err)
			} else {
				fmt.Printf("✅ Pipeline created: my-pipeline\n")
			}
		}
	}

	// Get existing Routes and add new Route
	routesListResponse, err := client.Routes.List(ctx, operations.WithServerURL(groupURL))
	if err != nil {
		log.Printf("Error listing routes: %v", err)
	} else if routesListResponse.Object != nil && routesListResponse.Object.Items != nil && len(routesListResponse.Object.Items) > 0 {
		// Get the first Routes configuration
		existingRoutes := routesListResponse.Object.Items[0]

		// Create new Route
		newRoute := components.RoutesRoute{
			Final:                  criblcontrolplanesdkgo.Bool(false),
			ID:                     criblcontrolplanesdkgo.String("my-route"),
			Name:                   "my-route",
			Pipeline:               "my-pipeline",
			Output:                 "my-fs-destination",
			EnableOutputExpression: criblcontrolplanesdkgo.Bool(true), // Allow custom output destinations
			Filter:                 criblcontrolplanesdkgo.String("__inputId=='tcpjson:my-tcp-json'"),
			Description:            criblcontrolplanesdkgo.String("This is my new Route"),
		}

		// Add new route to existing Routes
		updatedRoutes := append([]components.RoutesRoute{newRoute}, existingRoutes.Routes...)

		// Update Routes configuration
		if existingRoutes.ID != nil {
			_, err = client.Routes.Update(ctx, *existingRoutes.ID, components.Routes{
				ID:     existingRoutes.ID,
				Routes: updatedRoutes,
			}, operations.WithServerURL(groupURL))

			if err != nil {
				log.Printf("Error updating routes: %v", err)
			} else {
				fmt.Printf("✅ Route inserted: my-route\n")
			}
		}
	}

	// Commit configuration changes
	effective := true
	commitParams := components.GitCommitParams{
		Message:   "Commit for Stream example",
		Effective: &effective,
		Files:     []string{"."},
	}

	workerGroupID := WORKER_GROUP_ID
	commitResponse, err := client.Versions.Commits.Create(ctx, commitParams, &workerGroupID)
	if err != nil {
		log.Printf("Error creating commit: %v", err)
	} else if commitResponse.Object != nil && commitResponse.Object.Items != nil && len(commitResponse.Object.Items) > 0 {
		version := commitResponse.Object.Items[0].Commit
		fmt.Printf("✅ Committed configuration changes to the group: %s, commit ID: %s\n", WORKER_GROUP_ID, version)

		// Deploy the configuration using DeployRequest
		deployRequest := components.DeployRequest{
			Version: version,
		}

		_, err = client.Groups.Deploy(ctx, components.ProductsCoreStream, WORKER_GROUP_ID, deployRequest)
		if err != nil {
			log.Printf("Error deploying Worker Group configuration: %v", err)
		} else {
			fmt.Printf("✅ Worker Group changes deployed: %s\n", WORKER_GROUP_ID)
		}
	}
}
