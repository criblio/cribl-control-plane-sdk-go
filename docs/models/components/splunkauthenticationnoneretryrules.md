# SplunkAuthenticationNoneRetryRules


## Supported Types

### SplunkAuthenticationNoneSplunkRetryRulesTypeNone

```go
splunkAuthenticationNoneRetryRules := components.CreateSplunkAuthenticationNoneRetryRulesNone(components.SplunkAuthenticationNoneSplunkRetryRulesTypeNone{/* values here */})
```

### SplunkAuthenticationNoneSplunkRetryRulesTypeStatic

```go
splunkAuthenticationNoneRetryRules := components.CreateSplunkAuthenticationNoneRetryRulesStatic(components.SplunkAuthenticationNoneSplunkRetryRulesTypeStatic{/* values here */})
```

### SplunkAuthenticationNoneSplunkRetryRulesTypeBackoff

```go
splunkAuthenticationNoneRetryRules := components.CreateSplunkAuthenticationNoneRetryRulesBackoff(components.SplunkAuthenticationNoneSplunkRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch splunkAuthenticationNoneRetryRules.Type {
	case components.SplunkAuthenticationNoneRetryRulesTypeNone:
		// splunkAuthenticationNoneRetryRules.SplunkAuthenticationNoneSplunkRetryRulesTypeNone is populated
	case components.SplunkAuthenticationNoneRetryRulesTypeStatic:
		// splunkAuthenticationNoneRetryRules.SplunkAuthenticationNoneSplunkRetryRulesTypeStatic is populated
	case components.SplunkAuthenticationNoneRetryRulesTypeBackoff:
		// splunkAuthenticationNoneRetryRules.SplunkAuthenticationNoneSplunkRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use splunkAuthenticationNoneRetryRules.GetUnknownRaw() for raw JSON
}
```
