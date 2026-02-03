/**
 * Cribl Edge Configuration Example
 *
 * This example demonstrates how to programmatically create and configure a
 * complete data pipeline in Cribl Edge using the Control Plane SDK.
 *
 * This example creates:
 *
 * 1. A Fleet to manage the configuration.
 * 2. A Syslog Source to receive data on port 9021.
 * 3. An Amazon S3 Destination to store processed data.
 * 4. A Pipeline that filters events and keeps only data in the "eventSource"
 * and "eventID" fields.
 * 5. A Route that connects the Source to the Pipeline and Destination.
 *
 * This example also deploys the configuration to the Fleet to make it active.
 *
 * Data flow: Syslog Source → Route → Pipeline → Amazon S3 Destination
 *
 * This example includes error handling and checks for existing resources.
 *
 * Prerequisites:
 * - An .env file that is configured with your credentials.
 * - Your AWS S3 values for AWS_API_KEY, AWS_SECRET_KEY, AWS_BUCKET_NAME, and
 * AWS_REGION.
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
	FLEET_ID = "my-fleet"

	// Syslog Source configuration
	SYSLOG_PORT = 9021

	// Amazon S3 Destination configuration: Replace the placeholder values
	AWS_API_KEY     = "your-aws-api-key"     // Replace with your AWS Access Key ID
	AWS_SECRET_KEY  = "your-aws-secret-key"  // Replace with your AWS Secret Access Key
	AWS_BUCKET_NAME = "your-aws-bucket-name" // Replace with your S3 bucket name
	AWS_REGION      = "us-east-2"            // Replace with your S3 bucket region
)

func main() {
	ctx := context.Background()

	// Initialize Cribl client
	client, err := CreateCriblClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Cribl client: %v", err)
	}

	groupURL := fmt.Sprintf("%s/m/%s", BaseURL, FLEET_ID)

	// Check if Fleet already exists
	getResponse, err := client.Groups.Get(ctx, components.ProductsCoreEdge, FLEET_ID, nil)
	if err != nil {
		log.Printf("Fleet doesn't exist yet, will create: %v", err)
	}

	if getResponse != nil && getResponse.CountedConfigGroup != nil &&
		getResponse.CountedConfigGroup.Items != nil &&
		len(getResponse.CountedConfigGroup.Items) > 0 {
		fmt.Printf("⚠️ Fleet already exists: %s. Using existing fleet.\n", FLEET_ID)
	} else {
		// Create Fleet
		myFleetCreateRequest := components.GroupCreateRequest{
			ID:                 FLEET_ID,
			OnPrem:             criblcontrolplanesdkgo.Bool(true),
			WorkerRemoteAccess: criblcontrolplanesdkgo.Bool(true),
			IsFleet:            criblcontrolplanesdkgo.Bool(true),
			IsSearch:           criblcontrolplanesdkgo.Bool(false),
		}
		createResponse, err := client.Groups.Create(ctx, components.ProductsCoreEdge, myFleetCreateRequest)
		if err != nil {
			log.Printf("Error creating fleet: %v", err)
		} else if createResponse != nil && createResponse.CountedConfigGroup != nil {
			fmt.Printf("✅ Fleet created: %s\n", FLEET_ID)
		}
	}

	// Create Syslog Source
	udpPort := float64(SYSLOG_PORT)
	syslogSource := operations.CreateCreateInputInputSyslogUnionCreateInputInputSyslogSyslog2(operations.CreateInputInputSyslogSyslog2{
		ID:      "my-syslog-source",
		Type:    operations.CreateInputInputSyslogType2Syslog,
		Host:    "0.0.0.0",
		TCPPort: float64(SYSLOG_PORT),
		UDPPort: &udpPort,
		TLS: &components.TLSSettingsServerSideType{
			Disabled: criblcontrolplanesdkgo.Bool(true),
		},
	})

	createRequest := operations.CreateCreateInputRequestSyslog(syslogSource)
	_, err = client.Sources.Create(ctx, createRequest, operations.WithServerURL(groupURL))
	if err != nil {
		log.Printf("Error creating Syslog source: %v", err)
	} else {
		fmt.Printf("✅ Syslog source created: my-syslog-source\n")
	}

	// Create Amazon S3 Destination
	fileNameSuffix := "\".log\"" // JavaScript expression that returns ".log"
	region := AWS_REGION
	secretKey := AWS_SECRET_KEY
	apiKey := AWS_API_KEY
	stagePath := "/tmp/cribl-stage"
	s3Destination := operations.CreateOutputOutputS3{
		ID:             "my-s3-destination",
		Type:           operations.CreateOutputTypeS3S3,
		Bucket:         AWS_BUCKET_NAME,
		StagePath:      stagePath,
		Region:         &region,
		AwsSecretKey:   &secretKey,
		AwsAPIKey:      &apiKey,
		Compress:       components.CompressionOptions2Gzip.ToPointer(),
		FileNameSuffix: &fileNameSuffix,
	}

	createOutputRequest := operations.CreateCreateOutputRequestS3(s3Destination)
	_, err = client.Destinations.Create(ctx, createOutputRequest, operations.WithServerURL(groupURL))
	if err != nil {
		log.Printf("Error creating Amazon S3 destination: %v", err)
	} else {
		fmt.Printf("✅ Amazon S3 destination created: my-s3-destination\n")
	}

	// Create pipeline
	asyncFuncTimeout := int64(1000)
	filter := "true"
	final := true
	evalFunction := components.CreatePipelineFunctionConfInputEval(components.PipelineFunctionEval{
		Filter: &filter,
		ID:     components.PipelineFunctionEvalIDEval,
		Final:  &final,
		Conf: components.FunctionConfSchemaEval{
			Remove: []string{"*"},
			Keep:   []string{"eventSource", "eventID"},
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
		log.Printf("Error creating pipeline: %v", err)
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
			Output:                 "my-s3-destination",
			EnableOutputExpression: criblcontrolplanesdkgo.Bool(true), // Allow custom output destinations
			Filter:                 criblcontrolplanesdkgo.String("__inputId=='syslog:my-syslog-source'"),
			Description:            criblcontrolplanesdkgo.String("This is my new Route"),
		}

		// Add new Route to existing Routes
		updatedRoutes := append([]components.RoutesRoute{newRoute}, existingRoutes.Routes...)

		// Update Routes configuration
		if existingRoutes.ID != nil {
			_, err = client.Routes.Update(ctx, *existingRoutes.ID, components.Routes{
				ID:     existingRoutes.ID,
				Routes: updatedRoutes,
			}, operations.WithServerURL(groupURL))

			if err != nil {
				log.Printf("Error updating Routes: %v", err)
			} else {
				fmt.Printf("✅ Route inserted: my-route\n")
			}
		}
	}

	// Commit configuration changes
	effective := true
	commitParams := components.GitCommitParams{
		Message:   "Commit for Edge example",
		Effective: &effective,
		Files:     []string{"."},
	}

	commitResponse, err := client.Versions.Commits.Create(ctx, commitParams, operations.WithServerURL(groupURL))
	if err != nil {
		log.Printf("Error creating commit: %v", err)
	} else if commitResponse.CountedGitCommitSummary != nil && commitResponse.CountedGitCommitSummary.Items != nil && len(commitResponse.CountedGitCommitSummary.Items) > 0 {
		version := commitResponse.CountedGitCommitSummary.Items[0].Commit
		fmt.Printf("✅ Committed configuration changes to the fleet: %s, commit ID: %s\n", FLEET_ID, version)

		// Deploy the configuration using DeployRequest
		deployRequest := components.DeployRequest{
			Version: version,
		}

		_, err = client.Groups.Deploy(ctx, components.ProductsCoreEdge, FLEET_ID, deployRequest)
		if err != nil {
			log.Printf("Error deploying fleet configuration: %v", err)
		} else {
			fmt.Printf("✅ Fleet changes deployed: %s\n", FLEET_ID)
		}
	}
}
