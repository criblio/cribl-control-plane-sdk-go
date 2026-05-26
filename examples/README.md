# cribl-control-plane Examples

This directory contains example programs demonstrating how to use the cribl-control-plane Go SDK.

## Prerequisites

- Go 1.22 or higher
- Access to a Cribl instance (cloud or on-prem)

## Setup

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Copy `env_template` to `.env`:
   ```bash
   cp env_template .env
   ```

3. Edit `.env` and add your actual credentials (API keys, tokens, etc.)

## Running the Examples

### Authentication Examples (Standalone - No auth.go required)

These examples handle authentication internally and can be run standalone:

```bash
# Cloud authentication example (uses hardcoded values)
go run example-cloud-auth.go

# On-prem authentication example (uses hardcoded values)  
go run example-onprem-auth.go
```

### Configuration Examples (Require auth.go + .env file)

These examples use the `auth.go` helper and require a properly configured `.env` file:

```bash
# Cloud worker group management
go run example-cloud-worker-group.go auth.go

# Complete Cribl Stream data pipeline
go run example-stream.go auth.go

# Complete Cribl Edge data pipeline  
go run example-edge.go auth.go

# Cribl Packs integration
go run example-packs.go auth.go

# Worker group replication
go run example-worker-group-replication.go auth.go
```

**Important**: Always include `auth.go` when running configuration examples, as they depend on the authentication helper functions.

## Example Descriptions

| Example File | Description | Authentication Method |
|---|---|---|
| `example-cloud-auth.go` | Cloud authentication demo - validates OAuth2 connection | Hardcoded credentials |
| `example-onprem-auth.go` | On-prem authentication demo - validates username/password login | Hardcoded credentials |
| `example-cloud-worker-group.go` | Creates and manages cloud worker groups | Uses `auth.go` + `.env` |
| `example-stream.go` | Complete Cribl Stream pipeline: worker group, source, destination, pipeline, routes, deployment | Uses `auth.go` + `.env` |
| `example-edge.go` | Complete Cribl Edge pipeline: fleet, source, destination, pipeline, routes, deployment | Uses `auth.go` + `.env` |
| `example-packs.go` | Installs Cribl Pack and creates complete data pipeline within the pack | Uses `auth.go` + `.env` |
| `example-worker-group-replication.go` | Replicates an existing worker group with a new unique identity | Uses `auth.go` + `.env` |

## Configuration

Each example can be configured by either:
1. Using a `.env` file (recommended for examples that support both cloud and on-prem)
2. Editing the configuration constants directly in the example files (for simple examples)

### Environment Variables

For cloud deployments:
- `DEPLOYMENT_ENV=cloud` (or omit for cloud as default)
- `ORG_ID` - Your organization ID
- `CLIENT_ID` - Your OAuth2 client ID
- `CLIENT_SECRET` - Your OAuth2 client secret
- `WORKSPACE_NAME` - Your workspace name

For on-prem deployments:
- `DEPLOYMENT_ENV=onprem`
- `ONPREM_SERVER_URL` - Your on-prem server URL
- `ONPREM_USERNAME` - Your username
- `ONPREM_PASSWORD` - Your password

## Building Examples

To build a specific example:

```bash
go build -o example-cloud-auth example-cloud-auth.go auth.go
```
