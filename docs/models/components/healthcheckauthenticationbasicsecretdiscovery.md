# HealthCheckAuthenticationBasicSecretDiscovery


## Supported Types

### HealthCheckAuthenticationBasicSecretHealthCheckDiscoveryDiscoverTypeHTTP

```go
healthCheckAuthenticationBasicSecretDiscovery := components.CreateHealthCheckAuthenticationBasicSecretDiscoveryHTTP(components.HealthCheckAuthenticationBasicSecretHealthCheckDiscoveryDiscoverTypeHTTP{/* values here */})
```

### HealthCheckAuthenticationBasicSecretHealthCheckDiscoveryDiscoverTypeJSON

```go
healthCheckAuthenticationBasicSecretDiscovery := components.CreateHealthCheckAuthenticationBasicSecretDiscoveryJSON(components.HealthCheckAuthenticationBasicSecretHealthCheckDiscoveryDiscoverTypeJSON{/* values here */})
```

### HealthCheckAuthenticationBasicSecretHealthCheckDiscoveryDiscoverTypeList

```go
healthCheckAuthenticationBasicSecretDiscovery := components.CreateHealthCheckAuthenticationBasicSecretDiscoveryList(components.HealthCheckAuthenticationBasicSecretHealthCheckDiscoveryDiscoverTypeList{/* values here */})
```

### HealthCheckAuthenticationBasicSecretHealthCheckDiscoveryDiscoverTypeNone

```go
healthCheckAuthenticationBasicSecretDiscovery := components.CreateHealthCheckAuthenticationBasicSecretDiscoveryNone(components.HealthCheckAuthenticationBasicSecretHealthCheckDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckAuthenticationBasicSecretDiscovery.Type {
	case components.HealthCheckAuthenticationBasicSecretDiscoveryTypeHTTP:
		// healthCheckAuthenticationBasicSecretDiscovery.HealthCheckAuthenticationBasicSecretHealthCheckDiscoveryDiscoverTypeHTTP is populated
	case components.HealthCheckAuthenticationBasicSecretDiscoveryTypeJSON:
		// healthCheckAuthenticationBasicSecretDiscovery.HealthCheckAuthenticationBasicSecretHealthCheckDiscoveryDiscoverTypeJSON is populated
	case components.HealthCheckAuthenticationBasicSecretDiscoveryTypeList:
		// healthCheckAuthenticationBasicSecretDiscovery.HealthCheckAuthenticationBasicSecretHealthCheckDiscoveryDiscoverTypeList is populated
	case components.HealthCheckAuthenticationBasicSecretDiscoveryTypeNone:
		// healthCheckAuthenticationBasicSecretDiscovery.HealthCheckAuthenticationBasicSecretHealthCheckDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use healthCheckAuthenticationBasicSecretDiscovery.GetUnknownRaw() for raw JSON
}
```
