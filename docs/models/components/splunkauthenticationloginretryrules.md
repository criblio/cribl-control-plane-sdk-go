# SplunkAuthenticationLoginRetryRules


## Supported Types

### SplunkAuthenticationLoginSplunkRetryRulesTypeNone

```go
splunkAuthenticationLoginRetryRules := components.CreateSplunkAuthenticationLoginRetryRulesNone(components.SplunkAuthenticationLoginSplunkRetryRulesTypeNone{/* values here */})
```

### SplunkAuthenticationLoginSplunkRetryRulesTypeStatic

```go
splunkAuthenticationLoginRetryRules := components.CreateSplunkAuthenticationLoginRetryRulesStatic(components.SplunkAuthenticationLoginSplunkRetryRulesTypeStatic{/* values here */})
```

### SplunkAuthenticationLoginSplunkRetryRulesTypeBackoff

```go
splunkAuthenticationLoginRetryRules := components.CreateSplunkAuthenticationLoginRetryRulesBackoff(components.SplunkAuthenticationLoginSplunkRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch splunkAuthenticationLoginRetryRules.Type {
	case components.SplunkAuthenticationLoginRetryRulesTypeNone:
		// splunkAuthenticationLoginRetryRules.SplunkAuthenticationLoginSplunkRetryRulesTypeNone is populated
	case components.SplunkAuthenticationLoginRetryRulesTypeStatic:
		// splunkAuthenticationLoginRetryRules.SplunkAuthenticationLoginSplunkRetryRulesTypeStatic is populated
	case components.SplunkAuthenticationLoginRetryRulesTypeBackoff:
		// splunkAuthenticationLoginRetryRules.SplunkAuthenticationLoginSplunkRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use splunkAuthenticationLoginRetryRules.GetUnknownRaw() for raw JSON
}
```
