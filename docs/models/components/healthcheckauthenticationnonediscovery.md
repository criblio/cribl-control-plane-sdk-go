# HealthCheckAuthenticationNoneDiscovery


## Supported Types

### HealthCheckAuthenticationNoneHealthCheckDiscoveryDiscoverTypeHTTP

```go
healthCheckAuthenticationNoneDiscovery := components.CreateHealthCheckAuthenticationNoneDiscoveryHTTP(components.HealthCheckAuthenticationNoneHealthCheckDiscoveryDiscoverTypeHTTP{/* values here */})
```

### HealthCheckAuthenticationNoneHealthCheckDiscoveryDiscoverTypeJSON

```go
healthCheckAuthenticationNoneDiscovery := components.CreateHealthCheckAuthenticationNoneDiscoveryJSON(components.HealthCheckAuthenticationNoneHealthCheckDiscoveryDiscoverTypeJSON{/* values here */})
```

### HealthCheckAuthenticationNoneHealthCheckDiscoveryDiscoverTypeList

```go
healthCheckAuthenticationNoneDiscovery := components.CreateHealthCheckAuthenticationNoneDiscoveryList(components.HealthCheckAuthenticationNoneHealthCheckDiscoveryDiscoverTypeList{/* values here */})
```

### HealthCheckAuthenticationNoneHealthCheckDiscoveryDiscoverTypeNone

```go
healthCheckAuthenticationNoneDiscovery := components.CreateHealthCheckAuthenticationNoneDiscoveryNone(components.HealthCheckAuthenticationNoneHealthCheckDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckAuthenticationNoneDiscovery.Type {
	case components.HealthCheckAuthenticationNoneDiscoveryTypeHTTP:
		// healthCheckAuthenticationNoneDiscovery.HealthCheckAuthenticationNoneHealthCheckDiscoveryDiscoverTypeHTTP is populated
	case components.HealthCheckAuthenticationNoneDiscoveryTypeJSON:
		// healthCheckAuthenticationNoneDiscovery.HealthCheckAuthenticationNoneHealthCheckDiscoveryDiscoverTypeJSON is populated
	case components.HealthCheckAuthenticationNoneDiscoveryTypeList:
		// healthCheckAuthenticationNoneDiscovery.HealthCheckAuthenticationNoneHealthCheckDiscoveryDiscoverTypeList is populated
	case components.HealthCheckAuthenticationNoneDiscoveryTypeNone:
		// healthCheckAuthenticationNoneDiscovery.HealthCheckAuthenticationNoneHealthCheckDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use healthCheckAuthenticationNoneDiscovery.GetUnknownRaw() for raw JSON
}
```
