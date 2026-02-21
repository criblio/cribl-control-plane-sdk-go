# HealthCheckCollectMethodGetRetryRules


## Supported Types

### HealthCheckCollectMethodGetHealthCheckRetryRulesTypeNone

```go
healthCheckCollectMethodGetRetryRules := components.CreateHealthCheckCollectMethodGetRetryRulesNone(components.HealthCheckCollectMethodGetHealthCheckRetryRulesTypeNone{/* values here */})
```

### HealthCheckCollectMethodGetHealthCheckRetryRulesTypeStatic

```go
healthCheckCollectMethodGetRetryRules := components.CreateHealthCheckCollectMethodGetRetryRulesStatic(components.HealthCheckCollectMethodGetHealthCheckRetryRulesTypeStatic{/* values here */})
```

### HealthCheckCollectMethodGetHealthCheckRetryRulesTypeBackoff

```go
healthCheckCollectMethodGetRetryRules := components.CreateHealthCheckCollectMethodGetRetryRulesBackoff(components.HealthCheckCollectMethodGetHealthCheckRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckCollectMethodGetRetryRules.Type {
	case components.HealthCheckCollectMethodGetRetryRulesTypeNone:
		// healthCheckCollectMethodGetRetryRules.HealthCheckCollectMethodGetHealthCheckRetryRulesTypeNone is populated
	case components.HealthCheckCollectMethodGetRetryRulesTypeStatic:
		// healthCheckCollectMethodGetRetryRules.HealthCheckCollectMethodGetHealthCheckRetryRulesTypeStatic is populated
	case components.HealthCheckCollectMethodGetRetryRulesTypeBackoff:
		// healthCheckCollectMethodGetRetryRules.HealthCheckCollectMethodGetHealthCheckRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use healthCheckCollectMethodGetRetryRules.GetUnknownRaw() for raw JSON
}
```
