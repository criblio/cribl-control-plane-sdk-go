# HealthCheckAuthenticationBasicDiscovery


## Supported Types

### HealthCheckAuthenticationBasicHealthCheckDiscoveryDiscoverTypeHTTP

```go
healthCheckAuthenticationBasicDiscovery := components.CreateHealthCheckAuthenticationBasicDiscoveryHTTP(components.HealthCheckAuthenticationBasicHealthCheckDiscoveryDiscoverTypeHTTP{/* values here */})
```

### HealthCheckAuthenticationBasicHealthCheckDiscoveryDiscoverTypeJSON

```go
healthCheckAuthenticationBasicDiscovery := components.CreateHealthCheckAuthenticationBasicDiscoveryJSON(components.HealthCheckAuthenticationBasicHealthCheckDiscoveryDiscoverTypeJSON{/* values here */})
```

### HealthCheckAuthenticationBasicHealthCheckDiscoveryDiscoverTypeList

```go
healthCheckAuthenticationBasicDiscovery := components.CreateHealthCheckAuthenticationBasicDiscoveryList(components.HealthCheckAuthenticationBasicHealthCheckDiscoveryDiscoverTypeList{/* values here */})
```

### HealthCheckAuthenticationBasicHealthCheckDiscoveryDiscoverTypeNone

```go
healthCheckAuthenticationBasicDiscovery := components.CreateHealthCheckAuthenticationBasicDiscoveryNone(components.HealthCheckAuthenticationBasicHealthCheckDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckAuthenticationBasicDiscovery.Type {
	case components.HealthCheckAuthenticationBasicDiscoveryTypeHTTP:
		// healthCheckAuthenticationBasicDiscovery.HealthCheckAuthenticationBasicHealthCheckDiscoveryDiscoverTypeHTTP is populated
	case components.HealthCheckAuthenticationBasicDiscoveryTypeJSON:
		// healthCheckAuthenticationBasicDiscovery.HealthCheckAuthenticationBasicHealthCheckDiscoveryDiscoverTypeJSON is populated
	case components.HealthCheckAuthenticationBasicDiscoveryTypeList:
		// healthCheckAuthenticationBasicDiscovery.HealthCheckAuthenticationBasicHealthCheckDiscoveryDiscoverTypeList is populated
	case components.HealthCheckAuthenticationBasicDiscoveryTypeNone:
		// healthCheckAuthenticationBasicDiscovery.HealthCheckAuthenticationBasicHealthCheckDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use healthCheckAuthenticationBasicDiscovery.GetUnknownRaw() for raw JSON
}
```
