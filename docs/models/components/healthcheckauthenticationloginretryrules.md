# HealthCheckAuthenticationLoginRetryRules


## Supported Types

### HealthCheckAuthenticationLoginHealthCheckRetryRulesTypeNone

```go
healthCheckAuthenticationLoginRetryRules := components.CreateHealthCheckAuthenticationLoginRetryRulesNone(components.HealthCheckAuthenticationLoginHealthCheckRetryRulesTypeNone{/* values here */})
```

### HealthCheckAuthenticationLoginHealthCheckRetryRulesTypeStatic

```go
healthCheckAuthenticationLoginRetryRules := components.CreateHealthCheckAuthenticationLoginRetryRulesStatic(components.HealthCheckAuthenticationLoginHealthCheckRetryRulesTypeStatic{/* values here */})
```

### HealthCheckAuthenticationLoginHealthCheckRetryRulesTypeBackoff

```go
healthCheckAuthenticationLoginRetryRules := components.CreateHealthCheckAuthenticationLoginRetryRulesBackoff(components.HealthCheckAuthenticationLoginHealthCheckRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckAuthenticationLoginRetryRules.Type {
	case components.HealthCheckAuthenticationLoginRetryRulesTypeNone:
		// healthCheckAuthenticationLoginRetryRules.HealthCheckAuthenticationLoginHealthCheckRetryRulesTypeNone is populated
	case components.HealthCheckAuthenticationLoginRetryRulesTypeStatic:
		// healthCheckAuthenticationLoginRetryRules.HealthCheckAuthenticationLoginHealthCheckRetryRulesTypeStatic is populated
	case components.HealthCheckAuthenticationLoginRetryRulesTypeBackoff:
		// healthCheckAuthenticationLoginRetryRules.HealthCheckAuthenticationLoginHealthCheckRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use healthCheckAuthenticationLoginRetryRules.GetUnknownRaw() for raw JSON
}
```
