# SplunkAuthenticationBasicRetryRules


## Supported Types

### SplunkAuthenticationBasicSplunkRetryRulesTypeNone

```go
splunkAuthenticationBasicRetryRules := components.CreateSplunkAuthenticationBasicRetryRulesNone(components.SplunkAuthenticationBasicSplunkRetryRulesTypeNone{/* values here */})
```

### SplunkAuthenticationBasicSplunkRetryRulesTypeStatic

```go
splunkAuthenticationBasicRetryRules := components.CreateSplunkAuthenticationBasicRetryRulesStatic(components.SplunkAuthenticationBasicSplunkRetryRulesTypeStatic{/* values here */})
```

### SplunkAuthenticationBasicSplunkRetryRulesTypeBackoff

```go
splunkAuthenticationBasicRetryRules := components.CreateSplunkAuthenticationBasicRetryRulesBackoff(components.SplunkAuthenticationBasicSplunkRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch splunkAuthenticationBasicRetryRules.Type {
	case components.SplunkAuthenticationBasicRetryRulesTypeNone:
		// splunkAuthenticationBasicRetryRules.SplunkAuthenticationBasicSplunkRetryRulesTypeNone is populated
	case components.SplunkAuthenticationBasicRetryRulesTypeStatic:
		// splunkAuthenticationBasicRetryRules.SplunkAuthenticationBasicSplunkRetryRulesTypeStatic is populated
	case components.SplunkAuthenticationBasicRetryRulesTypeBackoff:
		// splunkAuthenticationBasicRetryRules.SplunkAuthenticationBasicSplunkRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use splunkAuthenticationBasicRetryRules.GetUnknownRaw() for raw JSON
}
```
