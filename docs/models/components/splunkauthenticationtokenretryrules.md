# SplunkAuthenticationTokenRetryRules


## Supported Types

### SplunkAuthenticationTokenSplunkRetryRulesTypeNone

```go
splunkAuthenticationTokenRetryRules := components.CreateSplunkAuthenticationTokenRetryRulesNone(components.SplunkAuthenticationTokenSplunkRetryRulesTypeNone{/* values here */})
```

### SplunkAuthenticationTokenSplunkRetryRulesTypeStatic

```go
splunkAuthenticationTokenRetryRules := components.CreateSplunkAuthenticationTokenRetryRulesStatic(components.SplunkAuthenticationTokenSplunkRetryRulesTypeStatic{/* values here */})
```

### SplunkAuthenticationTokenSplunkRetryRulesTypeBackoff

```go
splunkAuthenticationTokenRetryRules := components.CreateSplunkAuthenticationTokenRetryRulesBackoff(components.SplunkAuthenticationTokenSplunkRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch splunkAuthenticationTokenRetryRules.Type {
	case components.SplunkAuthenticationTokenRetryRulesTypeNone:
		// splunkAuthenticationTokenRetryRules.SplunkAuthenticationTokenSplunkRetryRulesTypeNone is populated
	case components.SplunkAuthenticationTokenRetryRulesTypeStatic:
		// splunkAuthenticationTokenRetryRules.SplunkAuthenticationTokenSplunkRetryRulesTypeStatic is populated
	case components.SplunkAuthenticationTokenRetryRulesTypeBackoff:
		// splunkAuthenticationTokenRetryRules.SplunkAuthenticationTokenSplunkRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use splunkAuthenticationTokenRetryRules.GetUnknownRaw() for raw JSON
}
```
