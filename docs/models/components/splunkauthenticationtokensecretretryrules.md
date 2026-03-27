# SplunkAuthenticationTokenSecretRetryRules


## Supported Types

### SplunkAuthenticationTokenSecretSplunkRetryRulesTypeNone

```go
splunkAuthenticationTokenSecretRetryRules := components.CreateSplunkAuthenticationTokenSecretRetryRulesNone(components.SplunkAuthenticationTokenSecretSplunkRetryRulesTypeNone{/* values here */})
```

### SplunkAuthenticationTokenSecretSplunkRetryRulesTypeStatic

```go
splunkAuthenticationTokenSecretRetryRules := components.CreateSplunkAuthenticationTokenSecretRetryRulesStatic(components.SplunkAuthenticationTokenSecretSplunkRetryRulesTypeStatic{/* values here */})
```

### SplunkAuthenticationTokenSecretSplunkRetryRulesTypeBackoff

```go
splunkAuthenticationTokenSecretRetryRules := components.CreateSplunkAuthenticationTokenSecretRetryRulesBackoff(components.SplunkAuthenticationTokenSecretSplunkRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch splunkAuthenticationTokenSecretRetryRules.Type {
	case components.SplunkAuthenticationTokenSecretRetryRulesTypeNone:
		// splunkAuthenticationTokenSecretRetryRules.SplunkAuthenticationTokenSecretSplunkRetryRulesTypeNone is populated
	case components.SplunkAuthenticationTokenSecretRetryRulesTypeStatic:
		// splunkAuthenticationTokenSecretRetryRules.SplunkAuthenticationTokenSecretSplunkRetryRulesTypeStatic is populated
	case components.SplunkAuthenticationTokenSecretRetryRulesTypeBackoff:
		// splunkAuthenticationTokenSecretRetryRules.SplunkAuthenticationTokenSecretSplunkRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use splunkAuthenticationTokenSecretRetryRules.GetUnknownRaw() for raw JSON
}
```
