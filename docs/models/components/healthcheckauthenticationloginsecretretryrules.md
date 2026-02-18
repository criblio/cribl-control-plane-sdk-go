# HealthCheckAuthenticationLoginSecretRetryRules


## Supported Types

### HealthCheckAuthenticationLoginSecretHealthCheckRetryRulesTypeNone

```go
healthCheckAuthenticationLoginSecretRetryRules := components.CreateHealthCheckAuthenticationLoginSecretRetryRulesNone(components.HealthCheckAuthenticationLoginSecretHealthCheckRetryRulesTypeNone{/* values here */})
```

### HealthCheckAuthenticationLoginSecretHealthCheckRetryRulesTypeStatic

```go
healthCheckAuthenticationLoginSecretRetryRules := components.CreateHealthCheckAuthenticationLoginSecretRetryRulesStatic(components.HealthCheckAuthenticationLoginSecretHealthCheckRetryRulesTypeStatic{/* values here */})
```

### HealthCheckAuthenticationLoginSecretHealthCheckRetryRulesTypeBackoff

```go
healthCheckAuthenticationLoginSecretRetryRules := components.CreateHealthCheckAuthenticationLoginSecretRetryRulesBackoff(components.HealthCheckAuthenticationLoginSecretHealthCheckRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckAuthenticationLoginSecretRetryRules.Type {
	case components.HealthCheckAuthenticationLoginSecretRetryRulesTypeNone:
		// healthCheckAuthenticationLoginSecretRetryRules.HealthCheckAuthenticationLoginSecretHealthCheckRetryRulesTypeNone is populated
	case components.HealthCheckAuthenticationLoginSecretRetryRulesTypeStatic:
		// healthCheckAuthenticationLoginSecretRetryRules.HealthCheckAuthenticationLoginSecretHealthCheckRetryRulesTypeStatic is populated
	case components.HealthCheckAuthenticationLoginSecretRetryRulesTypeBackoff:
		// healthCheckAuthenticationLoginSecretRetryRules.HealthCheckAuthenticationLoginSecretHealthCheckRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use healthCheckAuthenticationLoginSecretRetryRules.GetUnknownRaw() for raw JSON
}
```
