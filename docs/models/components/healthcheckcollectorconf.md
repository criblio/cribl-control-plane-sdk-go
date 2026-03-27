# HealthCheckCollectorConf


## Supported Types

### HealthCheckAuthenticationNone

```go
healthCheckCollectorConf := components.CreateHealthCheckCollectorConfNone(components.HealthCheckAuthenticationNone{/* values here */})
```

### HealthCheckAuthenticationBasic

```go
healthCheckCollectorConf := components.CreateHealthCheckCollectorConfBasic(components.HealthCheckAuthenticationBasic{/* values here */})
```

### HealthCheckAuthenticationBasicSecret

```go
healthCheckCollectorConf := components.CreateHealthCheckCollectorConfBasicSecret(components.HealthCheckAuthenticationBasicSecret{/* values here */})
```

### HealthCheckAuthenticationLogin

```go
healthCheckCollectorConf := components.CreateHealthCheckCollectorConfLogin(components.HealthCheckAuthenticationLogin{/* values here */})
```

### HealthCheckAuthenticationLoginSecret

```go
healthCheckCollectorConf := components.CreateHealthCheckCollectorConfLoginSecret(components.HealthCheckAuthenticationLoginSecret{/* values here */})
```

### HealthCheckAuthenticationOauth

```go
healthCheckCollectorConf := components.CreateHealthCheckCollectorConfOauth(components.HealthCheckAuthenticationOauth{/* values here */})
```

### HealthCheckAuthenticationOauthSecret

```go
healthCheckCollectorConf := components.CreateHealthCheckCollectorConfOauthSecret(components.HealthCheckAuthenticationOauthSecret{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckCollectorConf.Type {
	case components.HealthCheckCollectorConfTypeNone:
		// healthCheckCollectorConf.HealthCheckAuthenticationNone is populated
	case components.HealthCheckCollectorConfTypeBasic:
		// healthCheckCollectorConf.HealthCheckAuthenticationBasic is populated
	case components.HealthCheckCollectorConfTypeBasicSecret:
		// healthCheckCollectorConf.HealthCheckAuthenticationBasicSecret is populated
	case components.HealthCheckCollectorConfTypeLogin:
		// healthCheckCollectorConf.HealthCheckAuthenticationLogin is populated
	case components.HealthCheckCollectorConfTypeLoginSecret:
		// healthCheckCollectorConf.HealthCheckAuthenticationLoginSecret is populated
	case components.HealthCheckCollectorConfTypeOauth:
		// healthCheckCollectorConf.HealthCheckAuthenticationOauth is populated
	case components.HealthCheckCollectorConfTypeOauthSecret:
		// healthCheckCollectorConf.HealthCheckAuthenticationOauthSecret is populated
	default:
		// Unknown type - use healthCheckCollectorConf.GetUnknownRaw() for raw JSON
}
```
