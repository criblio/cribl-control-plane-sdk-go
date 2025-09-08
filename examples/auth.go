/*
Authentication Helper Module

This helper module handles authentication for all SDK examples, supporting both
cloud (OAuth2) and on-premises (username/password) deployments. It automatically
detects the deployment type, loads environment variables, validates credentials,
and provides authenticated SDK client instances.

Used by: Example files that that can run on cloud or on-premises
*/

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"github.com/joho/godotenv"
)

const domain = "cribl-staging.cloud"

type OnpremConfiguration struct {
	ServerURL string
	Username  string
	Password  string
}

type CloudConfiguration struct {
	OrgID         string
	ClientID      string
	ClientSecret  string
	WorkspaceName string
}

var (
	isOnprem      bool
	configuration interface{}
	BaseURL       string
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	isOnprem = os.Getenv("DEPLOYMENT_ENV") == "onprem"

	if isOnprem {
		config := getOnpremConfiguration()
		configuration = config
		BaseURL = fmt.Sprintf("%s/api/v1", config.ServerURL)
	} else {
		config := getCloudConfiguration()
		configuration = config
		BaseURL = fmt.Sprintf("https://%s-%s.%s/api/v1", config.WorkspaceName, config.OrgID, domain)
	}
}

// CreateCriblClient creates an authenticated Cribl Control Plane client
// Automatically detects deployment type and uses appropriate authentication method
func CreateCriblClient(ctx context.Context) (*criblcontrolplanesdkgo.CriblControlPlane, error) {
	if isOnprem {
		return createOnpremClient(ctx)
	}
	return createCloudClient(ctx)
}

func getOnpremConfiguration() OnpremConfiguration {
	serverURL := os.Getenv("ONPREM_SERVER_URL")
	username := os.Getenv("ONPREM_USERNAME")
	password := os.Getenv("ONPREM_PASSWORD")

	if serverURL == "" {
		log.Fatal("ONPREM_SERVER_URL is required for on-premises deployment")
	}
	if username == "" {
		log.Fatal("ONPREM_USERNAME is required for on-premises deployment")
	}
	if password == "" {
		log.Fatal("ONPREM_PASSWORD is required for on-premises deployment")
	}

	return OnpremConfiguration{
		ServerURL: serverURL,
		Username:  username,
		Password:  password,
	}
}

func getCloudConfiguration() CloudConfiguration {
	orgID := os.Getenv("ORG_ID")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	workspaceName := os.Getenv("WORKSPACE_NAME")

	if orgID == "" {
		log.Fatal("ORG_ID is required for cloud deployment")
	}
	if clientID == "" {
		log.Fatal("CLIENT_ID is required for cloud deployment")
	}
	if clientSecret == "" {
		log.Fatal("CLIENT_SECRET is required for cloud deployment")
	}
	if workspaceName == "" {
		log.Fatal("WORKSPACE_NAME is required for cloud deployment")
	}

	return CloudConfiguration{
		OrgID:         orgID,
		ClientID:      clientID,
		ClientSecret:  clientSecret,
		WorkspaceName: workspaceName,
	}
}

// createOnpremClient creates an authenticated client for on-premises deployment
func createOnpremClient(ctx context.Context) (*criblcontrolplanesdkgo.CriblControlPlane, error) {
	config := configuration.(OnpremConfiguration)

	// First, create a client to get the token
	tokenClient := criblcontrolplanesdkgo.New(BaseURL)

	// Get authentication token
	loginInfo := components.LoginInfo{
		Username: config.Username,
		Password: config.Password,
	}

	response, err := tokenClient.Auth.Tokens.Get(ctx, loginInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate with on-premises server: %w", err)
	}

	if response.AuthToken == nil || response.AuthToken.Token == "" {
		return nil, fmt.Errorf("no token received from authentication response")
	}

	// Create authenticated client
	client := criblcontrolplanesdkgo.New(
		BaseURL,
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			BearerAuth: &response.AuthToken.Token,
		}),
	)

	return client, nil
}

// createCloudClient creates an authenticated client for cloud deployment
func createCloudClient(ctx context.Context) (*criblcontrolplanesdkgo.CriblControlPlane, error) {
	config := configuration.(CloudConfiguration)

	tokenURL := fmt.Sprintf("https://login.%s/oauth/token", domain)
	audience := fmt.Sprintf("https://api.%s", domain)

	client := criblcontrolplanesdkgo.New(
		BaseURL,
		criblcontrolplanesdkgo.WithSecurity(components.Security{
			ClientOauth: &components.SchemeClientOauth{
				ClientID:     config.ClientID,
				ClientSecret: config.ClientSecret,
				TokenURL:     tokenURL,
				Audience:     audience,
			},
		}),
	)

	return client, nil
}

// RetryWithBackoff retries an operation with exponential backoff for rate limiting
func RetryWithBackoff(ctx context.Context, maxRetries int, operation func() error) error {
	for attempt := 0; attempt < maxRetries; attempt++ {
		err := operation()
		if err == nil {
			return nil
		}

		// Check if it's a rate limit error (you might need to adjust this based on the actual error structure)
		if attempt < maxRetries-1 {
			fmt.Printf("⚠️ Rate limit exceeded, retrying in %d seconds...\n", attempt+1)
			time.Sleep(time.Duration(attempt+1) * time.Second)
			continue
		}
		return err
	}
	return fmt.Errorf("max retries exceeded")
}
