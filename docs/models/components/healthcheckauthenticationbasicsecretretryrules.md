# HealthCheckAuthenticationBasicSecretRetryRules


## Supported Types

### HealthCheckAuthenticationBasicSecretHealthCheckRetryRulesTypeNone

```go
healthCheckAuthenticationBasicSecretRetryRules := components.CreateHealthCheckAuthenticationBasicSecretRetryRulesNone(components.HealthCheckAuthenticationBasicSecretHealthCheckRetryRulesTypeNone{/* values here */})
```

### HealthCheckAuthenticationBasicSecretHealthCheckRetryRulesTypeStatic

```go
healthCheckAuthenticationBasicSecretRetryRules := components.CreateHealthCheckAuthenticationBasicSecretRetryRulesStatic(components.HealthCheckAuthenticationBasicSecretHealthCheckRetryRulesTypeStatic{/* values here */})
```

### HealthCheckAuthenticationBasicSecretHealthCheckRetryRulesTypeBackoff

```go
healthCheckAuthenticationBasicSecretRetryRules := components.CreateHealthCheckAuthenticationBasicSecretRetryRulesBackoff(components.HealthCheckAuthenticationBasicSecretHealthCheckRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckAuthenticationBasicSecretRetryRules.Type {
	case components.HealthCheckAuthenticationBasicSecretRetryRulesTypeNone:
		// healthCheckAuthenticationBasicSecretRetryRules.HealthCheckAuthenticationBasicSecretHealthCheckRetryRulesTypeNone is populated
	case components.HealthCheckAuthenticationBasicSecretRetryRulesTypeStatic:
		// healthCheckAuthenticationBasicSecretRetryRules.HealthCheckAuthenticationBasicSecretHealthCheckRetryRulesTypeStatic is populated
	case components.HealthCheckAuthenticationBasicSecretRetryRulesTypeBackoff:
		// healthCheckAuthenticationBasicSecretRetryRules.HealthCheckAuthenticationBasicSecretHealthCheckRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use healthCheckAuthenticationBasicSecretRetryRules.GetUnknownRaw() for raw JSON
}
```
