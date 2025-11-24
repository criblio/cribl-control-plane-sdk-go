/**
 * On-Prem Authentication Example
 *
 * This example demonstrates how to configure authentication for an on-prem 
 * Cribl instance using username and password credentials.
 *
 * 1. Create an SDK client with username and password credentials using the 
 * BearerAuth security scheme.
 * 2. Automatically handle token exchange and refresh using a callback function.
 * 3. Validate the connection by listing all git branches.
 *
 * Prerequisites: Replace the placeholder values for ONPREM_SERVER_URL
 * ONPREM_USERNAME, and ONPREM_PASSWORD with your server URL and credentials.
 * Your credentials are sensitive information and should be kept private.
 *
 * NOTE: This example is for on-prem deployments only. It does not require .env
 * file configuration.
 */

package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/apierrors"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

// On-prem configuration: Replace the placeholder values
const (
	ONPREM_SERVER_URL = "http://localhost:9000" // Replace with your server URL
	ONPREM_USERNAME   = "admin"                 // Replace with your username
	ONPREM_PASSWORD   = "admin"                 // Replace with your password
)

// Token cache
var (
	cachedToken    string
	tokenExpiresAt time.Time
)

// getJWTExp extracts the expiration time from a JWT token
func getJWTExp(token string) (time.Time, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return time.Time{}, fmt.Errorf("invalid JWT token format")
	}

	// Decode the payload (second part)
	payloadB64 := parts[1]
	// Add padding if needed
	padding := (4 - len(payloadB64)%4) % 4
	payloadB64 += strings.Repeat("=", padding)

	payload, err := base64.URLEncoding.DecodeString(payloadB64)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to decode JWT payload: %w", err)
	}

	// Parse the JSON payload
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return time.Time{}, fmt.Errorf("failed to parse JWT payload: %w", err)
	}

	// Extract the exp field
	exp, ok := claims["exp"]
	if !ok {
		return time.Time{}, fmt.Errorf("token missing 'exp' field")
	}

	// Convert to int64 (exp is typically a float64 from JSON)
	expFloat, ok := exp.(float64)
	if !ok {
		return time.Time{}, fmt.Errorf("exp field is not a number")
	}

	return time.Unix(int64(expFloat), 0), nil
}

func main() {
	ctx := context.Background()
	baseURL := fmt.Sprintf("%s/api/v1", ONPREM_SERVER_URL)

	// Create client for retrieving Bearer token
	tokenClient := criblcontrolplanesdkgo.New(baseURL)

	// Create callback function for automatic token refresh
	securityCallback := func(ctx context.Context) (components.Security, error) {
		// Check cache - use token if it's still valid (with 3-second buffer)
		now := time.Now()
		if cachedToken != "" && now.Add(3*time.Second).Before(tokenExpiresAt) {
			return components.Security{
				BearerAuth: &cachedToken,
			}, nil
		}

		// Retrieve Bearer token initially and refresh automatically when it expires
		loginInfo := components.LoginInfo{
			Username: ONPREM_USERNAME,
			Password: ONPREM_PASSWORD,
		}

		response, err := tokenClient.Auth.Tokens.Get(ctx, loginInfo)
		if err != nil {
			return components.Security{}, err
		}

		if response.AuthToken == nil || response.AuthToken.Token == "" {
			return components.Security{}, fmt.Errorf("no token received from authentication response")
		}

		token := response.AuthToken.Token
		exp, err := getJWTExp(token)
		if err != nil {
			return components.Security{}, fmt.Errorf("failed to parse token expiration: %w", err)
		}

		// Update cache
		cachedToken = token
		tokenExpiresAt = exp

		return components.Security{
			BearerAuth: &token,
		}, nil
	}

	// Create authenticated SDK client
	client := criblcontrolplanesdkgo.New(
		baseURL,
		criblcontrolplanesdkgo.WithSecuritySource(securityCallback),
	)
	fmt.Println("✅ Authenticated SDK client created for on-prem server")

	// Validate connection and list all git branches
	response, err := client.Versions.Branches.List(ctx)
	if err != nil {
		handleError(err)
		return
	}

	if response.CountedBranchInfo == nil || response.CountedBranchInfo.Items == nil {
		fmt.Println("⚠️ No branches found")
		return
	}

	var branches []string
	for _, branch := range response.CountedBranchInfo.Items {
		if branch.ID != "" {
			branches = append(branches, branch.ID)
		}
	}

	fmt.Printf("✅ Client works! Your list of branches:\n\t%s\n", strings.Join(branches, "\n\t"))
}

func handleError(err error) {
	// Check if it's an APIError with a status code
	var apiErr *apierrors.APIError
	if e, ok := err.(*apierrors.APIError); ok {
		apiErr = e
	}

	if apiErr != nil {
		switch apiErr.StatusCode {
		case 401:
			fmt.Println("⚠️ Authentication failed! Check your USERNAME and PASSWORD.")
			return
		case 429:
			fmt.Println("⚠️ Uh oh, you've reached the rate limit! Try again in a few seconds.")
			return
		}
	}

	fmt.Printf("❌ Something went wrong: %v\n", err)
}
