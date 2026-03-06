# HealthCheckAuthenticationLoginSecretDiscovery


## Supported Types

### HealthCheckAuthenticationLoginSecretHealthCheckDiscoveryDiscoverTypeHTTP

```go
healthCheckAuthenticationLoginSecretDiscovery := components.CreateHealthCheckAuthenticationLoginSecretDiscoveryHTTP(components.HealthCheckAuthenticationLoginSecretHealthCheckDiscoveryDiscoverTypeHTTP{/* values here */})
```

### HealthCheckAuthenticationLoginSecretHealthCheckDiscoveryDiscoverTypeJSON

```go
healthCheckAuthenticationLoginSecretDiscovery := components.CreateHealthCheckAuthenticationLoginSecretDiscoveryJSON(components.HealthCheckAuthenticationLoginSecretHealthCheckDiscoveryDiscoverTypeJSON{/* values here */})
```

### HealthCheckAuthenticationLoginSecretHealthCheckDiscoveryDiscoverTypeList

```go
healthCheckAuthenticationLoginSecretDiscovery := components.CreateHealthCheckAuthenticationLoginSecretDiscoveryList(components.HealthCheckAuthenticationLoginSecretHealthCheckDiscoveryDiscoverTypeList{/* values here */})
```

### HealthCheckAuthenticationLoginSecretHealthCheckDiscoveryDiscoverTypeNone

```go
healthCheckAuthenticationLoginSecretDiscovery := components.CreateHealthCheckAuthenticationLoginSecretDiscoveryNone(components.HealthCheckAuthenticationLoginSecretHealthCheckDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckAuthenticationLoginSecretDiscovery.Type {
	case components.HealthCheckAuthenticationLoginSecretDiscoveryTypeHTTP:
		// healthCheckAuthenticationLoginSecretDiscovery.HealthCheckAuthenticationLoginSecretHealthCheckDiscoveryDiscoverTypeHTTP is populated
	case components.HealthCheckAuthenticationLoginSecretDiscoveryTypeJSON:
		// healthCheckAuthenticationLoginSecretDiscovery.HealthCheckAuthenticationLoginSecretHealthCheckDiscoveryDiscoverTypeJSON is populated
	case components.HealthCheckAuthenticationLoginSecretDiscoveryTypeList:
		// healthCheckAuthenticationLoginSecretDiscovery.HealthCheckAuthenticationLoginSecretHealthCheckDiscoveryDiscoverTypeList is populated
	case components.HealthCheckAuthenticationLoginSecretDiscoveryTypeNone:
		// healthCheckAuthenticationLoginSecretDiscovery.HealthCheckAuthenticationLoginSecretHealthCheckDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use healthCheckAuthenticationLoginSecretDiscovery.GetUnknownRaw() for raw JSON
}
```
