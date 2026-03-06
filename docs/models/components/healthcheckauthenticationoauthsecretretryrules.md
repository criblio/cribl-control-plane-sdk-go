# HealthCheckAuthenticationOauthSecretRetryRules


## Supported Types

### HealthCheckAuthenticationOauthSecretHealthCheckRetryRulesTypeNone

```go
healthCheckAuthenticationOauthSecretRetryRules := components.CreateHealthCheckAuthenticationOauthSecretRetryRulesNone(components.HealthCheckAuthenticationOauthSecretHealthCheckRetryRulesTypeNone{/* values here */})
```

### HealthCheckAuthenticationOauthSecretHealthCheckRetryRulesTypeStatic

```go
healthCheckAuthenticationOauthSecretRetryRules := components.CreateHealthCheckAuthenticationOauthSecretRetryRulesStatic(components.HealthCheckAuthenticationOauthSecretHealthCheckRetryRulesTypeStatic{/* values here */})
```

### HealthCheckAuthenticationOauthSecretHealthCheckRetryRulesTypeBackoff

```go
healthCheckAuthenticationOauthSecretRetryRules := components.CreateHealthCheckAuthenticationOauthSecretRetryRulesBackoff(components.HealthCheckAuthenticationOauthSecretHealthCheckRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch healthCheckAuthenticationOauthSecretRetryRules.Type {
	case components.HealthCheckAuthenticationOauthSecretRetryRulesTypeNone:
		// healthCheckAuthenticationOauthSecretRetryRules.HealthCheckAuthenticationOauthSecretHealthCheckRetryRulesTypeNone is populated
	case components.HealthCheckAuthenticationOauthSecretRetryRulesTypeStatic:
		// healthCheckAuthenticationOauthSecretRetryRules.HealthCheckAuthenticationOauthSecretHealthCheckRetryRulesTypeStatic is populated
	case components.HealthCheckAuthenticationOauthSecretRetryRulesTypeBackoff:
		// healthCheckAuthenticationOauthSecretRetryRules.HealthCheckAuthenticationOauthSecretHealthCheckRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use healthCheckAuthenticationOauthSecretRetryRules.GetUnknownRaw() for raw JSON
}
```
