/*
Cribl Stream Configuration Example

This example demonstrates how to programmatically create and configure a complete
data pipeline in Cribl Stream using the Control Plane SDK. It creates:

1. A worker group to manage the configuration
2. A TCP JSON source to receive data on port 9020
3. A filesystem destination to output processed data
4. A pipeline that filters events to keep only the "name" field
5. A route that connects the source to the pipeline and destination
6. Deploys the configuration to the worker group to make it active

Data flow: TCP JSON Source → Route → Pipeline → File Destination

The example includes proper error handling, checks for existing resources,
and actually invokes all the corresponding APIs.

Prerequisites:
- Configure your .env file
- Requires Enterprise License on the server
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

	// Check if worker group already exists
	getResponse, err := client.Groups.Get(ctx, components.ProductsCoreStream, WORKER_GROUP_ID, nil)
	if err != nil {
		log.Printf("Worker group doesn't exist yet, will create: %v", err)
	}

	if getResponse != nil && getResponse.Object != nil &&
		getResponse.Object.Items != nil &&
		len(getResponse.Object.Items) > 0 {
		fmt.Printf("⚠️ Worker Group already exists: %s. Using existing group.\n", WORKER_GROUP_ID)
	} else {
		// Create worker group
		myWorkerGroup := components.ConfigGroup{
			ID:          WORKER_GROUP_ID,
			Description: criblcontrolplanesdkgo.String("My Worker Group"),
			OnPrem:      criblcontrolplanesdkgo.Bool(true),
		}

		createResponse, err := client.Groups.Create(ctx, components.ProductsCoreStream, myWorkerGroup)
		if err != nil {
			log.Printf("Error creating worker group: %v", err)
		} else if createResponse != nil && createResponse.Object != nil {
			fmt.Printf("✅ Worker Group created: %s\n", WORKER_GROUP_ID)
		}
	}

	// Create TCP JSON source
	tcpJSONConfig := map[string]interface{}{
		"id":        "my-tcp-json",
		"type":      "tcpjson",
		"port":      PORT,
		"authType":  "manual",
		"authToken": AUTH_TOKEN,
	}

	// Convert to components.Input using JSON marshaling/unmarshaling (pattern from tests)
	sourceBytes, err := json.Marshal(tcpJSONConfig)
	if err != nil {
		log.Printf("Error marshaling TCP JSON source config: %v", err)
	} else {
		var tcpJSONSource components.Input
		err = json.Unmarshal(sourceBytes, &tcpJSONSource)
		if err != nil {
			log.Printf("Error unmarshaling TCP JSON source config: %v", err)
		} else {
			_, err = client.Sources.Create(ctx, tcpJSONSource, operations.WithServerURL(groupURL))
			if err != nil {
				log.Printf("Error creating TCP JSON source: %v", err)
			} else {
				fmt.Printf("✅ TCP JSON Source created: my-tcp-json\n")
			}
		}
	}

	// Create filesystem destination
	fileSystemConfig := map[string]interface{}{
		"id":             "my-fs-destination",
		"type":           "filesystem",
		"destPath":       "/tmp/my-output",
		"fileNameSuffix": "\".log\"", // JavaScript expression that returns ".log"
	}

	// Convert to components.Output using JSON marshaling/unmarshaling (pattern from tests)
	destBytes, err := json.Marshal(fileSystemConfig)
	if err != nil {
		log.Printf("Error marshaling filesystem destination config: %v", err)
	} else {
		var fileSystemDestination components.Output
		err = json.Unmarshal(destBytes, &fileSystemDestination)
		if err != nil {
			log.Printf("Error unmarshaling filesystem destination config: %v", err)
		} else {
			_, err = client.Destinations.Create(ctx, fileSystemDestination, operations.WithServerURL(groupURL))
			if err != nil {
				log.Printf("Error creating filesystem destination: %v", err)
			} else {
				fmt.Printf("✅ Filesystem Destination created: my-fs-destination\n")
			}
		}
	}

	// Create pipeline
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

	// Convert to components.Conf using JSON marshaling/unmarshaling (pattern from tests)
	confBytes, err := json.Marshal(pipelineConf)
	if err != nil {
		log.Printf("Error marshaling pipeline config: %v", err)
	} else {
		var conf components.Conf
		err = json.Unmarshal(confBytes, &conf)
		if err != nil {
			log.Printf("Error unmarshaling pipeline config: %v", err)
		} else {
			pipeline := components.Pipeline{
				ID:   "my-pipeline",
				Conf: conf,
			}

			_, err = client.Pipelines.Create(ctx, pipeline, operations.WithServerURL(groupURL))
			if err != nil {
				log.Printf("Error creating pipeline: %v", err)
			} else {
				fmt.Printf("✅ Pipeline created: my-pipeline\n")
			}
		}
	}

	// Get existing routes and add new route (pattern from tests)
	routesListResponse, err := client.Routes.List(ctx, operations.WithServerURL(groupURL))
	if err != nil {
		log.Printf("Error listing routes: %v", err)
	} else if routesListResponse.Object != nil && routesListResponse.Object.Items != nil && len(routesListResponse.Object.Items) > 0 {
		// Get the first routes configuration
		existingRoutes := routesListResponse.Object.Items[0]

		// Create new route
		newRoute := components.RoutesRoute{
			Final:                  criblcontrolplanesdkgo.Bool(false),
			ID:                     criblcontrolplanesdkgo.String("my-route"),
			Name:                   "my-route",
			Pipeline:               "my-pipeline",
			Output:                 "my-fs-destination",
			EnableOutputExpression: criblcontrolplanesdkgo.Bool(true), // Allow custom output destinations
			Filter:                 criblcontrolplanesdkgo.String("__inputId=='tcpjson:my-tcp-json'"),
			Description:            criblcontrolplanesdkgo.String("This is my new route"),
		}

		// Add new route to existing routes (prepend to match TypeScript behavior)
		updatedRoutes := append([]components.RoutesRoute{newRoute}, existingRoutes.Routes...)

		// Update routes configuration
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

	// Get current version for deployment (pattern from worker group tests)
	versionResponse, err := client.Groups.Configs.Versions.Get(ctx, components.ProductsCoreStream, WORKER_GROUP_ID)
	if err != nil {
		log.Printf("Error getting version for deployment: %v", err)
	} else if versionResponse.Object != nil && versionResponse.Object.Items != nil && len(versionResponse.Object.Items) > 0 {
		// Get the current version (it's a string)
		version := versionResponse.Object.Items[0]

		// Deploy the configuration using DeployRequest
		deployRequest := components.DeployRequest{
			Version: version,
		}

		_, err = client.Groups.Deploy(ctx, components.ProductsCoreStream, WORKER_GROUP_ID, deployRequest)
		if err != nil {
			log.Printf("Error deploying worker group configuration: %v", err)
		} else {
			fmt.Printf("✅ Worker Group changes deployed: %s\n", WORKER_GROUP_ID)
		}
	}

	fmt.Println("ℹ️ Complete data pipeline created and deployed!")
	fmt.Printf("ℹ️ Worker group URL: %s\n", groupURL)
}
