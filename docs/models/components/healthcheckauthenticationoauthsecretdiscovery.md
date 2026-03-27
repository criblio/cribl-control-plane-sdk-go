# HealthCheckAuthenticationOauthSecretDiscovery


## Supported Types

### HealthCheckAuthenticationOauthSecretHealthCheckDiscoveryDiscoverTypeHTTP

```go
healthCheckAuthenticationOauthSecretDiscovery := components.CreateHealthCheckAuthenticationOauthSecretDiscoveryHTTP(components.HealthCheckAuthenticationOauthSecretHealthCheckDiscoveryDiscoverTypeHTTP{/* values here */})
```

### HealthCheckAuthenticationOauthSecretHealthCheckDiscoveryDiscoverTypeJSON

```go
healthCheckAuthenticationOauthSecretDiscovery := components.CreateHealthCheckAuthenticationOauthSecretDiscoveryJSON(components.HealthCheckAuthenticationOauthSecretHealthCheckDiscoveryDiscoverTypeJSON{/* values here */})
```

### HealthCheckAuthenticationOauthSecretHealthCheckDiscoveryDiscoverTypeList

```go
healthCheckAuthenticationOauthSecretDiscovery := components.CreateHealthCheckAuthenticationOauthSecretDiscoveryList(components.HealthCheckAuthenticationOauthSecretHealthCheckDiscoveryDiscoverTypeList{/* values here */})
```

### HealthCheckAuthenticationOauthSecretHealthCheckDiscoveryDiscoverTypeNone

```go
healthCheckAuthenticationOauthSecretDiscovery := components.CreateHealthCheckAuthenticationOauthSecretDiscoveryNone(components.HealthCheckAuthenticationOauthSecretHealthCheckDiscoveryDiscoverTypeNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckAuthenticationOauthSecretDiscovery.Type {
	case components.HealthCheckAuthenticationOauthSecretDiscoveryTypeHTTP:
		// healthCheckAuthenticationOauthSecretDiscovery.HealthCheckAuthenticationOauthSecretHealthCheckDiscoveryDiscoverTypeHTTP is populated
	case components.HealthCheckAuthenticationOauthSecretDiscoveryTypeJSON:
		// healthCheckAuthenticationOauthSecretDiscovery.HealthCheckAuthenticationOauthSecretHealthCheckDiscoveryDiscoverTypeJSON is populated
	case components.HealthCheckAuthenticationOauthSecretDiscoveryTypeList:
		// healthCheckAuthenticationOauthSecretDiscovery.HealthCheckAuthenticationOauthSecretHealthCheckDiscoveryDiscoverTypeList is populated
	case components.HealthCheckAuthenticationOauthSecretDiscoveryTypeNone:
		// healthCheckAuthenticationOauthSecretDiscovery.HealthCheckAuthenticationOauthSecretHealthCheckDiscoveryDiscoverTypeNone is populated
	default:
		// Unknown type - use healthCheckAuthenticationOauthSecretDiscovery.GetUnknownRaw() for raw JSON
}
```
