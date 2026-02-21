# SplunkAuthenticationLoginSecretRetryRules


## Supported Types

### SplunkAuthenticationLoginSecretSplunkRetryRulesTypeNone

```go
splunkAuthenticationLoginSecretRetryRules := components.CreateSplunkAuthenticationLoginSecretRetryRulesNone(components.SplunkAuthenticationLoginSecretSplunkRetryRulesTypeNone{/* values here */})
```

### SplunkAuthenticationLoginSecretSplunkRetryRulesTypeStatic

```go
splunkAuthenticationLoginSecretRetryRules := components.CreateSplunkAuthenticationLoginSecretRetryRulesStatic(components.SplunkAuthenticationLoginSecretSplunkRetryRulesTypeStatic{/* values here */})
```

### SplunkAuthenticationLoginSecretSplunkRetryRulesTypeBackoff

```go
splunkAuthenticationLoginSecretRetryRules := components.CreateSplunkAuthenticationLoginSecretRetryRulesBackoff(components.SplunkAuthenticationLoginSecretSplunkRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch splunkAuthenticationLoginSecretRetryRules.Type {
	case components.SplunkAuthenticationLoginSecretRetryRulesTypeNone:
		// splunkAuthenticationLoginSecretRetryRules.SplunkAuthenticationLoginSecretSplunkRetryRulesTypeNone is populated
	case components.SplunkAuthenticationLoginSecretRetryRulesTypeStatic:
		// splunkAuthenticationLoginSecretRetryRules.SplunkAuthenticationLoginSecretSplunkRetryRulesTypeStatic is populated
	case components.SplunkAuthenticationLoginSecretRetryRulesTypeBackoff:
		// splunkAuthenticationLoginSecretRetryRules.SplunkAuthenticationLoginSecretSplunkRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use splunkAuthenticationLoginSecretRetryRules.GetUnknownRaw() for raw JSON
}
```
