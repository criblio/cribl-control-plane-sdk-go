# HealthCheckCollectMethodPostWithBodyRetryRules


## Supported Types

### HealthCheckCollectMethodPostWithBodyHealthCheckRetryRulesTypeNone

```go
healthCheckCollectMethodPostWithBodyRetryRules := components.CreateHealthCheckCollectMethodPostWithBodyRetryRulesNone(components.HealthCheckCollectMethodPostWithBodyHealthCheckRetryRulesTypeNone{/* values here */})
```

### HealthCheckCollectMethodPostWithBodyHealthCheckRetryRulesTypeStatic

```go
healthCheckCollectMethodPostWithBodyRetryRules := components.CreateHealthCheckCollectMethodPostWithBodyRetryRulesStatic(components.HealthCheckCollectMethodPostWithBodyHealthCheckRetryRulesTypeStatic{/* values here */})
```

### HealthCheckCollectMethodPostWithBodyHealthCheckRetryRulesTypeBackoff

```go
healthCheckCollectMethodPostWithBodyRetryRules := components.CreateHealthCheckCollectMethodPostWithBodyRetryRulesBackoff(components.HealthCheckCollectMethodPostWithBodyHealthCheckRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckCollectMethodPostWithBodyRetryRules.Type {
	case components.HealthCheckCollectMethodPostWithBodyRetryRulesTypeNone:
		// healthCheckCollectMethodPostWithBodyRetryRules.HealthCheckCollectMethodPostWithBodyHealthCheckRetryRulesTypeNone is populated
	case components.HealthCheckCollectMethodPostWithBodyRetryRulesTypeStatic:
		// healthCheckCollectMethodPostWithBodyRetryRules.HealthCheckCollectMethodPostWithBodyHealthCheckRetryRulesTypeStatic is populated
	case components.HealthCheckCollectMethodPostWithBodyRetryRulesTypeBackoff:
		// healthCheckCollectMethodPostWithBodyRetryRules.HealthCheckCollectMethodPostWithBodyHealthCheckRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use healthCheckCollectMethodPostWithBodyRetryRules.GetUnknownRaw() for raw JSON
}
```
