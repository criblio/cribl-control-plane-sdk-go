# HealthCheckCollectMethodPostDiscovery


## Supported Types

### HealthCheckCollectMethodPostHealthCheckDiscoveryDiscoverTypeHTTP

```go
healthCheckCollectMethodPostDiscovery := components.CreateHealthCheckCollectMethodPostDiscoveryHTTP(components.HealthCheckCollectMethodPostHealthCheckDiscoveryDiscoverTypeHTTP{/* values here */})
```

### HealthCheckCollectMethodPostHealthCheckDiscoveryDiscoverTypeJSON

```go
healthCheckCollectMethodPostDiscovery := components.CreateHealthCheckCollectMethodPostDiscoveryJSON(components.HealthCheckCollectMethodPostHealthCheckDiscoveryDiscoverTypeJSON{/* values here */})
```

### HealthCheckCollectMethodPostHealthCheckDiscoveryDiscoverTypeList

```go
healthCheckCollectMethodPostDiscovery := components.CreateHealthCheckCollectMethodPostDiscoveryList(components.HealthCheckCollectMethodPostHealthCheckDiscoveryDiscoverTypeList{/* values here */})
```

### HealthCheckCollectMethodPostHealthCheckDiscoveryDiscoverTypeNone

```go
healthCheckCollectMethodPostDiscovery := components.CreateHealthCheckCollectMethodPostDiscoveryNone(components.HealthCheckCollectMethodPostHealthCheckDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckCollectMethodPostDiscovery.Type {
	case components.HealthCheckCollectMethodPostDiscoveryTypeHTTP:
		// healthCheckCollectMethodPostDiscovery.HealthCheckCollectMethodPostHealthCheckDiscoveryDiscoverTypeHTTP is populated
	case components.HealthCheckCollectMethodPostDiscoveryTypeJSON:
		// healthCheckCollectMethodPostDiscovery.HealthCheckCollectMethodPostHealthCheckDiscoveryDiscoverTypeJSON is populated
	case components.HealthCheckCollectMethodPostDiscoveryTypeList:
		// healthCheckCollectMethodPostDiscovery.HealthCheckCollectMethodPostHealthCheckDiscoveryDiscoverTypeList is populated
	case components.HealthCheckCollectMethodPostDiscoveryTypeNone:
		// healthCheckCollectMethodPostDiscovery.HealthCheckCollectMethodPostHealthCheckDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use healthCheckCollectMethodPostDiscovery.GetUnknownRaw() for raw JSON
}
```
