# HealthCheckCollectMethodGetDiscovery


## Supported Types

### HealthCheckCollectMethodGetHealthCheckDiscoveryDiscoverTypeHTTP

```go
healthCheckCollectMethodGetDiscovery := components.CreateHealthCheckCollectMethodGetDiscoveryHTTP(components.HealthCheckCollectMethodGetHealthCheckDiscoveryDiscoverTypeHTTP{/* values here */})
```

### HealthCheckCollectMethodGetHealthCheckDiscoveryDiscoverTypeJSON

```go
healthCheckCollectMethodGetDiscovery := components.CreateHealthCheckCollectMethodGetDiscoveryJSON(components.HealthCheckCollectMethodGetHealthCheckDiscoveryDiscoverTypeJSON{/* values here */})
```

### HealthCheckCollectMethodGetHealthCheckDiscoveryDiscoverTypeList

```go
healthCheckCollectMethodGetDiscovery := components.CreateHealthCheckCollectMethodGetDiscoveryList(components.HealthCheckCollectMethodGetHealthCheckDiscoveryDiscoverTypeList{/* values here */})
```

### HealthCheckCollectMethodGetHealthCheckDiscoveryDiscoverTypeNone

```go
healthCheckCollectMethodGetDiscovery := components.CreateHealthCheckCollectMethodGetDiscoveryNone(components.HealthCheckCollectMethodGetHealthCheckDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckCollectMethodGetDiscovery.Type {
	case components.HealthCheckCollectMethodGetDiscoveryTypeHTTP:
		// healthCheckCollectMethodGetDiscovery.HealthCheckCollectMethodGetHealthCheckDiscoveryDiscoverTypeHTTP is populated
	case components.HealthCheckCollectMethodGetDiscoveryTypeJSON:
		// healthCheckCollectMethodGetDiscovery.HealthCheckCollectMethodGetHealthCheckDiscoveryDiscoverTypeJSON is populated
	case components.HealthCheckCollectMethodGetDiscoveryTypeList:
		// healthCheckCollectMethodGetDiscovery.HealthCheckCollectMethodGetHealthCheckDiscoveryDiscoverTypeList is populated
	case components.HealthCheckCollectMethodGetDiscoveryTypeNone:
		// healthCheckCollectMethodGetDiscovery.HealthCheckCollectMethodGetHealthCheckDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use healthCheckCollectMethodGetDiscovery.GetUnknownRaw() for raw JSON
}
```
