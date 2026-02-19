# HealthCheckAuthenticationOauthRetryRules


## Supported Types

### HealthCheckAuthenticationOauthHealthCheckRetryRulesTypeNone

```go
healthCheckAuthenticationOauthRetryRules := components.CreateHealthCheckAuthenticationOauthRetryRulesNone(components.HealthCheckAuthenticationOauthHealthCheckRetryRulesTypeNone{/* values here */})
```

### HealthCheckAuthenticationOauthHealthCheckRetryRulesTypeStatic

```go
healthCheckAuthenticationOauthRetryRules := components.CreateHealthCheckAuthenticationOauthRetryRulesStatic(components.HealthCheckAuthenticationOauthHealthCheckRetryRulesTypeStatic{/* values here */})
```

### HealthCheckAuthenticationOauthHealthCheckRetryRulesTypeBackoff

```go
healthCheckAuthenticationOauthRetryRules := components.CreateHealthCheckAuthenticationOauthRetryRulesBackoff(components.HealthCheckAuthenticationOauthHealthCheckRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckAuthenticationOauthRetryRules.Type {
	case components.HealthCheckAuthenticationOauthRetryRulesTypeNone:
		// healthCheckAuthenticationOauthRetryRules.HealthCheckAuthenticationOauthHealthCheckRetryRulesTypeNone is populated
	case components.HealthCheckAuthenticationOauthRetryRulesTypeStatic:
		// healthCheckAuthenticationOauthRetryRules.HealthCheckAuthenticationOauthHealthCheckRetryRulesTypeStatic is populated
	case components.HealthCheckAuthenticationOauthRetryRulesTypeBackoff:
		// healthCheckAuthenticationOauthRetryRules.HealthCheckAuthenticationOauthHealthCheckRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use healthCheckAuthenticationOauthRetryRules.GetUnknownRaw() for raw JSON
}
```
