# HealthCheckAuthenticationOauthDiscovery


## Supported Types

### HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeHTTP

```go
healthCheckAuthenticationOauthDiscovery := components.CreateHealthCheckAuthenticationOauthDiscoveryHTTP(components.HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeHTTP{/* values here */})
```

### HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeJSON

```go
healthCheckAuthenticationOauthDiscovery := components.CreateHealthCheckAuthenticationOauthDiscoveryJSON(components.HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeJSON{/* values here */})
```

### HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeList

```go
healthCheckAuthenticationOauthDiscovery := components.CreateHealthCheckAuthenticationOauthDiscoveryList(components.HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeList{/* values here */})
```

### HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeNone

```go
healthCheckAuthenticationOauthDiscovery := components.CreateHealthCheckAuthenticationOauthDiscoveryNone(components.HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckAuthenticationOauthDiscovery.Type {
	case components.HealthCheckAuthenticationOauthDiscoveryTypeHTTP:
		// healthCheckAuthenticationOauthDiscovery.HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeHTTP is populated
	case components.HealthCheckAuthenticationOauthDiscoveryTypeJSON:
		// healthCheckAuthenticationOauthDiscovery.HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeJSON is populated
	case components.HealthCheckAuthenticationOauthDiscoveryTypeList:
		// healthCheckAuthenticationOauthDiscovery.HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeList is populated
	case components.HealthCheckAuthenticationOauthDiscoveryTypeNone:
		// healthCheckAuthenticationOauthDiscovery.HealthCheckAuthenticationOauthHealthCheckDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use healthCheckAuthenticationOauthDiscovery.GetUnknownRaw() for raw JSON
}
```
