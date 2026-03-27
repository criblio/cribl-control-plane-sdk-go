# HealthCheckCollectMethodPostRetryRules


## Supported Types

### HealthCheckCollectMethodPostHealthCheckRetryRulesTypeNone

```go
healthCheckCollectMethodPostRetryRules := components.CreateHealthCheckCollectMethodPostRetryRulesNone(components.HealthCheckCollectMethodPostHealthCheckRetryRulesTypeNone{/* values here */})
```

### HealthCheckCollectMethodPostHealthCheckRetryRulesTypeStatic

```go
healthCheckCollectMethodPostRetryRules := components.CreateHealthCheckCollectMethodPostRetryRulesStatic(components.HealthCheckCollectMethodPostHealthCheckRetryRulesTypeStatic{/* values here */})
```

### HealthCheckCollectMethodPostHealthCheckRetryRulesTypeBackoff

```go
healthCheckCollectMethodPostRetryRules := components.CreateHealthCheckCollectMethodPostRetryRulesBackoff(components.HealthCheckCollectMethodPostHealthCheckRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckCollectMethodPostRetryRules.Type {
	case components.HealthCheckCollectMethodPostRetryRulesTypeNone:
		// healthCheckCollectMethodPostRetryRules.HealthCheckCollectMethodPostHealthCheckRetryRulesTypeNone is populated
	case components.HealthCheckCollectMethodPostRetryRulesTypeStatic:
		// healthCheckCollectMethodPostRetryRules.HealthCheckCollectMethodPostHealthCheckRetryRulesTypeStatic is populated
	case components.HealthCheckCollectMethodPostRetryRulesTypeBackoff:
		// healthCheckCollectMethodPostRetryRules.HealthCheckCollectMethodPostHealthCheckRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use healthCheckCollectMethodPostRetryRules.GetUnknownRaw() for raw JSON
}
```
