/**
 * Cribl Packs Integration Example
 *
 * This example demonstrates how to install and configure a Cribl Pack using the
 * Control Plane SDK. It installs the Palo Alto Networks Pack from GitHub and
 * creates a complete data pipeline within the Pack.
 *
 * 1. Install the Palo Alto Networks Pack from a remote URL.
 * 2. Create a TCP JSON Source to receive data on port 9020.
 * 3. Create an Amazon S3 Destination for data storage.
 * 4. Create a Pipeline that filters events and keeps only data in the "name"
 * field.
 * 5. Create a Route that connects the Source to the Pipeline and Destination.
 *
 * Data flow: TCP JSON Source → Route → Pipeline → Amazon S3 Destination
 *
 * NOTE: This example creates resources within the Pack but does not commit
 * or deploy the configuration to a Worker Group.
 *
 * Prerequisites:
 * - An .env file that is configured with your credentials.
 * - A Worker Group whose ID matches the configured WORKER_GROUP_ID value.
 * - Your AWS S3 values for AWS_API_KEY, AWS_SECRET_KEY, AWS_BUCKET_NAME, and
 * AWS_REGION.
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
	PACK_URL        = "https://github.com/criblpacks/cribl-palo-alto-networks/releases/download/1.1.5/cribl-palo-alto-networks-d6bc6883-1.1.5.crbl"
	PACK_ID         = "cribl-palo-alto-networks"
	WORKER_GROUP_ID = "my-awesome-worker-group" // Replace with your worker group id

	// TCP JSON Source configuration
	AUTH_TOKEN = "4a4b3663-7a57-7369-7632-795553573668"
	PORT       = 9020

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

	groupURL := fmt.Sprintf("%s/m/%s", BaseURL, WORKER_GROUP_ID)
	packURL := fmt.Sprintf("%s/p/%s", groupURL, PACK_ID)

	// Check if Worker Group exists first
	getResponse, err := client.Groups.Get(ctx, components.ProductsCoreStream, WORKER_GROUP_ID, nil)
	if err != nil || getResponse == nil || getResponse.CountedConfigGroup == nil ||
		getResponse.CountedConfigGroup.Items == nil || len(getResponse.CountedConfigGroup.Items) == 0 {
		log.Printf("⚠️ Worker group '%s' does not exist. Please create this worker group first and then run this example.", WORKER_GROUP_ID)
		log.Printf("💡 You can create a Worker Group using the examples from this repository.")
		return
	}
	fmt.Printf("✅ Worker Group '%s' exists, proceeding with Pack installation.\n", WORKER_GROUP_ID)

	// First, check if Pack already exists and clean up
	client.Packs.Delete(ctx, PACK_ID, operations.WithServerURL(groupURL))

	// Install Pack from URL
	fmt.Printf("Installing Pack from: %s\n", PACK_URL)
	installReq := components.CreatePackRequestBodyUnionPackRequestBody1(
		components.PackRequestBody1{
			ID:     PACK_ID,
			Source: criblcontrolplanesdkgo.String(PACK_URL),
		},
	)

	_, err = client.Packs.Install(ctx, installReq, operations.WithServerURL(groupURL))
	if err != nil {
		log.Printf("Error installing Pack - %v", err)
	} else {
		fmt.Printf("✅ Installed Pack \"%s\" from: %s\n", PACK_ID, PACK_URL)
	}

	// Create TCP JSON Source in Pack
	authType := operations.CreateInputAuthenticationMethodTcpjsonManual
	authToken := AUTH_TOKEN
	tcpJSONSource := operations.InputTcpjson{
		ID:        "my-tcp-json",
		Type:      operations.CreateInputTypeTcpjsonTcpjson,
		Port:      float64(PORT),
		AuthType:  &authType,
		AuthToken: &authToken,
	}

	createInputRequest := operations.CreateCreateInputRequestTcpjson(tcpJSONSource)
	_, err = client.Sources.Create(ctx, createInputRequest, operations.WithServerURL(packURL))
	if err != nil {
		log.Printf("Error creating TCP JSON Source in Pack: %v", err)
	} else {
		fmt.Printf("✅ Created TCP JSON Source: my-tcp-json in Pack: \"%s\"\n", PACK_ID)
	}

	// Create Amazon S3 Destination in pack
	fileNameSuffix := "\".log\""
	region := AWS_REGION
	secretKey := AWS_SECRET_KEY
	apiKey := AWS_API_KEY
	s3Destination := operations.OutputS3{
		ID:             "my-s3-destination",
		Type:           operations.CreateOutputTypeS3S3,
		Bucket:         AWS_BUCKET_NAME,
		Region:         &region,
		AwsSecretKey:   &secretKey,
		AwsAPIKey:      &apiKey,
		Compress:       operations.CreateOutputCompressionS3Gzip.ToPointer(),
		FileNameSuffix: &fileNameSuffix,
	}

	createOutputRequest := operations.CreateCreateOutputRequestS3(s3Destination)
	_, err = client.Destinations.Create(ctx, createOutputRequest, operations.WithServerURL(packURL))
	if err != nil {
		log.Printf("Error creating Amazon S3 Destination in Pack: %v", err)
	} else {
		fmt.Printf("✅ Created Amazon S3 Destination: my-s3-destination in Pack: \"%s\"\n", PACK_ID)
	}

	// Create Pipeline in Pack
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

	_, err = client.Pipelines.Create(ctx, pipeline, operations.WithServerURL(packURL))
	if err != nil {
		log.Printf("Error creating Pipeline in Pack: %v", err)
	} else {
		fmt.Printf("✅ Created Pipeline: my-pipeline in Pack: \"%s\"\n", PACK_ID)
	}

	// Get existing Routes and add new Route
	routesListResponse, err := client.Routes.List(ctx, operations.WithServerURL(packURL))
	if err != nil {
		log.Printf("Error listing routes in pack: %v", err)
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
			Filter:                 criblcontrolplanesdkgo.String("__inputId=='tcpjson:my-tcp-json'"),
			Description:            criblcontrolplanesdkgo.String("This is my new route"),
		}

		// Add new route to existing Routes
		updatedRoutes := append([]components.RoutesRoute{newRoute}, existingRoutes.Routes...)

		// Update Routes configuration
		if existingRoutes.ID != nil {
			_, err = client.Routes.Update(ctx, *existingRoutes.ID, components.Routes{
				ID:     existingRoutes.ID,
				Routes: updatedRoutes,
			}, operations.WithServerURL(packURL))

			if err != nil {
				log.Printf("Error updating Routes in Pack: %v", err)
			} else {
				fmt.Printf("✅ Route inserted: my-route in Pack: %s\n", PACK_ID)
			}
		}
	}

	// Note about deployment
	fmt.Println("ℹ️ This example does not commit / deploy the configuration to the Worker Group.")
}
