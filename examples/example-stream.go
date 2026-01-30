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

	if getResponse != nil && getResponse.CountedConfigGroup != nil &&
		getResponse.CountedConfigGroup.Items != nil &&
		len(getResponse.CountedConfigGroup.Items) > 0 {
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
		} else if createResponse != nil && createResponse.CountedConfigGroup != nil {
			fmt.Printf("✅ Worker Group created: %s\n", WORKER_GROUP_ID)
		}
	}

	// Create TCP JSON Source
	authType := components.AuthenticationMethodOptionsAuthTokensItemsManual
	authToken := AUTH_TOKEN
	sendToRoutes := true
	tcpJSONSource := operations.CreateInputInputTcpjson{
		ID:           "my-tcp-json",
		Type:         operations.CreateInputTypeTcpjsonTcpjson,
		Host:         "0.0.0.0",
		Port:         float64(PORT),
		AuthType:     &authType,
		AuthToken:    &authToken,
		SendToRoutes: &sendToRoutes,
	}
	createInputRequest := operations.CreateCreateInputRequestTcpjson(tcpJSONSource)
	_, err = client.Sources.Create(ctx, createInputRequest, operations.WithServerURL(groupURL))
	if err != nil {
		log.Printf("Error creating TCP JSON Source: %v", err)
	} else {
		fmt.Printf("✅ TCP JSON Source created: my-tcp-json\n")
	}

	// Create Filesystem Destination
	fileNameSuffix := "\".log\"" // JavaScript expression that returns ".log"
	fileSystemDestination := operations.CreateOutputOutputFilesystem{
		ID:             "my-fs-destination",
		Type:           operations.CreateOutputTypeFilesystemFilesystem,
		DestPath:       "/tmp/my-output",
		FileNameSuffix: &fileNameSuffix,
	}

	createOutputRequest := operations.CreateCreateOutputRequestFilesystem(fileSystemDestination)
	_, err = client.Destinations.Create(ctx, createOutputRequest, operations.WithServerURL(groupURL))
	if err != nil {
		log.Printf("Error creating Filesystem Destination: %v", err)
	} else {
		fmt.Printf("✅ Filesystem Destination created: my-fs-destination\n")
	}

	// Create Pipeline
	asyncFuncTimeout := int64(1000)
	filter := "true"
	final := true
	evalFunction := components.CreatePipelineFunctionConfInputEval(components.PipelineFunctionEval{
		Filter: &filter,
		ID:     components.PipelineFunctionEvalIDEval,
		Final:  &final,
		Conf: components.FunctionConfSchemaEval{
			Remove: []string{"*"},
			Keep:   []string{"name"},
		},
	})

	conf := components.ConfInput{
		AsyncFuncTimeout: &asyncFuncTimeout,
		Functions:        []components.PipelineFunctionConfInput{evalFunction},
	}

	pipeline := components.PipelineInput{
		ID:   "my-pipeline",
		Conf: conf,
	}

	_, err = client.Pipelines.Create(ctx, pipeline, operations.WithServerURL(groupURL))
	if err != nil {
		log.Printf("Error creating Pipeline: %v", err)
	} else {
		fmt.Printf("✅ Pipeline created: my-pipeline\n")
	}

	// Get existing Routes and add new Route
	routesListResponse, err := client.Routes.List(ctx, operations.WithServerURL(groupURL))
	if err != nil {
		log.Printf("Error listing routes: %v", err)
	} else if routesListResponse.CountedRoutes != nil && routesListResponse.CountedRoutes.Items != nil && len(routesListResponse.CountedRoutes.Items) > 0 {
		// Get the first Routes configuration
		existingRoutes := routesListResponse.CountedRoutes.Items[0]

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
	} else if commitResponse.CountedGitCommitSummary != nil && commitResponse.CountedGitCommitSummary.Items != nil && len(commitResponse.CountedGitCommitSummary.Items) > 0 {
		version := commitResponse.CountedGitCommitSummary.Items[0].Commit
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
