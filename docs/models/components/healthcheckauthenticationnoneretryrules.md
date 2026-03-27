# HealthCheckAuthenticationNoneRetryRules


## Supported Types

### HealthCheckAuthenticationNoneHealthCheckRetryRulesTypeNone

```go
healthCheckAuthenticationNoneRetryRules := components.CreateHealthCheckAuthenticationNoneRetryRulesNone(components.HealthCheckAuthenticationNoneHealthCheckRetryRulesTypeNone{/* values here */})
```

### HealthCheckAuthenticationNoneHealthCheckRetryRulesTypeStatic

```go
healthCheckAuthenticationNoneRetryRules := components.CreateHealthCheckAuthenticationNoneRetryRulesStatic(components.HealthCheckAuthenticationNoneHealthCheckRetryRulesTypeStatic{/* values here */})
```

### HealthCheckAuthenticationNoneHealthCheckRetryRulesTypeBackoff

```go
healthCheckAuthenticationNoneRetryRules := components.CreateHealthCheckAuthenticationNoneRetryRulesBackoff(components.HealthCheckAuthenticationNoneHealthCheckRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckAuthenticationNoneRetryRules.Type {
	case components.HealthCheckAuthenticationNoneRetryRulesTypeNone:
		// healthCheckAuthenticationNoneRetryRules.HealthCheckAuthenticationNoneHealthCheckRetryRulesTypeNone is populated
	case components.HealthCheckAuthenticationNoneRetryRulesTypeStatic:
		// healthCheckAuthenticationNoneRetryRules.HealthCheckAuthenticationNoneHealthCheckRetryRulesTypeStatic is populated
	case components.HealthCheckAuthenticationNoneRetryRulesTypeBackoff:
		// healthCheckAuthenticationNoneRetryRules.HealthCheckAuthenticationNoneHealthCheckRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use healthCheckAuthenticationNoneRetryRules.GetUnknownRaw() for raw JSON
}
```
