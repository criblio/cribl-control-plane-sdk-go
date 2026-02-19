# HealthCheckCollectMethodPostWithBodyDiscovery


## Supported Types

### HealthCheckCollectMethodPostWithBodyHealthCheckDiscoveryDiscoverTypeHTTP

```go
healthCheckCollectMethodPostWithBodyDiscovery := components.CreateHealthCheckCollectMethodPostWithBodyDiscoveryHTTP(components.HealthCheckCollectMethodPostWithBodyHealthCheckDiscoveryDiscoverTypeHTTP{/* values here */})
```

### HealthCheckCollectMethodPostWithBodyHealthCheckDiscoveryDiscoverTypeJSON

```go
healthCheckCollectMethodPostWithBodyDiscovery := components.CreateHealthCheckCollectMethodPostWithBodyDiscoveryJSON(components.HealthCheckCollectMethodPostWithBodyHealthCheckDiscoveryDiscoverTypeJSON{/* values here */})
```

### HealthCheckCollectMethodPostWithBodyHealthCheckDiscoveryDiscoverTypeList

```go
healthCheckCollectMethodPostWithBodyDiscovery := components.CreateHealthCheckCollectMethodPostWithBodyDiscoveryList(components.HealthCheckCollectMethodPostWithBodyHealthCheckDiscoveryDiscoverTypeList{/* values here */})
```

### HealthCheckCollectMethodPostWithBodyHealthCheckDiscoveryDiscoverTypeNone

```go
healthCheckCollectMethodPostWithBodyDiscovery := components.CreateHealthCheckCollectMethodPostWithBodyDiscoveryNone(components.HealthCheckCollectMethodPostWithBodyHealthCheckDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckCollectMethodPostWithBodyDiscovery.Type {
	case components.HealthCheckCollectMethodPostWithBodyDiscoveryTypeHTTP:
		// healthCheckCollectMethodPostWithBodyDiscovery.HealthCheckCollectMethodPostWithBodyHealthCheckDiscoveryDiscoverTypeHTTP is populated
	case components.HealthCheckCollectMethodPostWithBodyDiscoveryTypeJSON:
		// healthCheckCollectMethodPostWithBodyDiscovery.HealthCheckCollectMethodPostWithBodyHealthCheckDiscoveryDiscoverTypeJSON is populated
	case components.HealthCheckCollectMethodPostWithBodyDiscoveryTypeList:
		// healthCheckCollectMethodPostWithBodyDiscovery.HealthCheckCollectMethodPostWithBodyHealthCheckDiscoveryDiscoverTypeList is populated
	case components.HealthCheckCollectMethodPostWithBodyDiscoveryTypeNone:
		// healthCheckCollectMethodPostWithBodyDiscovery.HealthCheckCollectMethodPostWithBodyHealthCheckDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use healthCheckCollectMethodPostWithBodyDiscovery.GetUnknownRaw() for raw JSON
}
```
