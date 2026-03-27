# SplunkAuthenticationBasicSecretRetryRules


## Supported Types

### SplunkAuthenticationBasicSecretSplunkRetryRulesTypeNone

```go
splunkAuthenticationBasicSecretRetryRules := components.CreateSplunkAuthenticationBasicSecretRetryRulesNone(components.SplunkAuthenticationBasicSecretSplunkRetryRulesTypeNone{/* values here */})
```

### SplunkAuthenticationBasicSecretSplunkRetryRulesTypeStatic

```go
splunkAuthenticationBasicSecretRetryRules := components.CreateSplunkAuthenticationBasicSecretRetryRulesStatic(components.SplunkAuthenticationBasicSecretSplunkRetryRulesTypeStatic{/* values here */})
```

### SplunkAuthenticationBasicSecretSplunkRetryRulesTypeBackoff

```go
splunkAuthenticationBasicSecretRetryRules := components.CreateSplunkAuthenticationBasicSecretRetryRulesBackoff(components.SplunkAuthenticationBasicSecretSplunkRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch splunkAuthenticationBasicSecretRetryRules.Type {
	case components.SplunkAuthenticationBasicSecretRetryRulesTypeNone:
		// splunkAuthenticationBasicSecretRetryRules.SplunkAuthenticationBasicSecretSplunkRetryRulesTypeNone is populated
	case components.SplunkAuthenticationBasicSecretRetryRulesTypeStatic:
		// splunkAuthenticationBasicSecretRetryRules.SplunkAuthenticationBasicSecretSplunkRetryRulesTypeStatic is populated
	case components.SplunkAuthenticationBasicSecretRetryRulesTypeBackoff:
		// splunkAuthenticationBasicSecretRetryRules.SplunkAuthenticationBasicSecretSplunkRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use splunkAuthenticationBasicSecretRetryRules.GetUnknownRaw() for raw JSON
}
```
