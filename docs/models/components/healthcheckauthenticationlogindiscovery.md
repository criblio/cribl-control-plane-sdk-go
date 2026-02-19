# HealthCheckAuthenticationLoginDiscovery


## Supported Types

### HealthCheckAuthenticationLoginHealthCheckDiscoveryDiscoverTypeHTTP

```go
healthCheckAuthenticationLoginDiscovery := components.CreateHealthCheckAuthenticationLoginDiscoveryHTTP(components.HealthCheckAuthenticationLoginHealthCheckDiscoveryDiscoverTypeHTTP{/* values here */})
```

### HealthCheckAuthenticationLoginHealthCheckDiscoveryDiscoverTypeJSON

```go
healthCheckAuthenticationLoginDiscovery := components.CreateHealthCheckAuthenticationLoginDiscoveryJSON(components.HealthCheckAuthenticationLoginHealthCheckDiscoveryDiscoverTypeJSON{/* values here */})
```

### HealthCheckAuthenticationLoginHealthCheckDiscoveryDiscoverTypeList

```go
healthCheckAuthenticationLoginDiscovery := components.CreateHealthCheckAuthenticationLoginDiscoveryList(components.HealthCheckAuthenticationLoginHealthCheckDiscoveryDiscoverTypeList{/* values here */})
```

### HealthCheckAuthenticationLoginHealthCheckDiscoveryDiscoverTypeNone

```go
healthCheckAuthenticationLoginDiscovery := components.CreateHealthCheckAuthenticationLoginDiscoveryNone(components.HealthCheckAuthenticationLoginHealthCheckDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckAuthenticationLoginDiscovery.Type {
	case components.HealthCheckAuthenticationLoginDiscoveryTypeHTTP:
		// healthCheckAuthenticationLoginDiscovery.HealthCheckAuthenticationLoginHealthCheckDiscoveryDiscoverTypeHTTP is populated
	case components.HealthCheckAuthenticationLoginDiscoveryTypeJSON:
		// healthCheckAuthenticationLoginDiscovery.HealthCheckAuthenticationLoginHealthCheckDiscoveryDiscoverTypeJSON is populated
	case components.HealthCheckAuthenticationLoginDiscoveryTypeList:
		// healthCheckAuthenticationLoginDiscovery.HealthCheckAuthenticationLoginHealthCheckDiscoveryDiscoverTypeList is populated
	case components.HealthCheckAuthenticationLoginDiscoveryTypeNone:
		// healthCheckAuthenticationLoginDiscovery.HealthCheckAuthenticationLoginHealthCheckDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use healthCheckAuthenticationLoginDiscovery.GetUnknownRaw() for raw JSON
}
```
