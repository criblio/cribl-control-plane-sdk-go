# HealthCheckAuthenticationBasicRetryRules


## Supported Types

### HealthCheckAuthenticationBasicHealthCheckRetryRulesTypeNone

```go
healthCheckAuthenticationBasicRetryRules := components.CreateHealthCheckAuthenticationBasicRetryRulesNone(components.HealthCheckAuthenticationBasicHealthCheckRetryRulesTypeNone{/* values here */})
```

### HealthCheckAuthenticationBasicHealthCheckRetryRulesTypeStatic

```go
healthCheckAuthenticationBasicRetryRules := components.CreateHealthCheckAuthenticationBasicRetryRulesStatic(components.HealthCheckAuthenticationBasicHealthCheckRetryRulesTypeStatic{/* values here */})
```

### HealthCheckAuthenticationBasicHealthCheckRetryRulesTypeBackoff

```go
healthCheckAuthenticationBasicRetryRules := components.CreateHealthCheckAuthenticationBasicRetryRulesBackoff(components.HealthCheckAuthenticationBasicHealthCheckRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckAuthenticationBasicRetryRules.Type {
	case components.HealthCheckAuthenticationBasicRetryRulesTypeNone:
		// healthCheckAuthenticationBasicRetryRules.HealthCheckAuthenticationBasicHealthCheckRetryRulesTypeNone is populated
	case components.HealthCheckAuthenticationBasicRetryRulesTypeStatic:
		// healthCheckAuthenticationBasicRetryRules.HealthCheckAuthenticationBasicHealthCheckRetryRulesTypeStatic is populated
	case components.HealthCheckAuthenticationBasicRetryRulesTypeBackoff:
		// healthCheckAuthenticationBasicRetryRules.HealthCheckAuthenticationBasicHealthCheckRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use healthCheckAuthenticationBasicRetryRules.GetUnknownRaw() for raw JSON
}
```
