# RestAuthenticationBasicRetryRules


## Supported Types

### RestAuthenticationBasicRestRetryRulesTypeNone

```go
restAuthenticationBasicRetryRules := components.CreateRestAuthenticationBasicRetryRulesNone(components.RestAuthenticationBasicRestRetryRulesTypeNone{/* values here */})
```

### RestAuthenticationBasicRestRetryRulesTypeStatic

```go
restAuthenticationBasicRetryRules := components.CreateRestAuthenticationBasicRetryRulesStatic(components.RestAuthenticationBasicRestRetryRulesTypeStatic{/* values here */})
```

### RestAuthenticationBasicRestRetryRulesTypeBackoff

```go
restAuthenticationBasicRetryRules := components.CreateRestAuthenticationBasicRetryRulesBackoff(components.RestAuthenticationBasicRestRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationBasicRetryRules.Type {
	case components.RestAuthenticationBasicRetryRulesTypeNone:
		// restAuthenticationBasicRetryRules.RestAuthenticationBasicRestRetryRulesTypeNone is populated
	case components.RestAuthenticationBasicRetryRulesTypeStatic:
		// restAuthenticationBasicRetryRules.RestAuthenticationBasicRestRetryRulesTypeStatic is populated
	case components.RestAuthenticationBasicRetryRulesTypeBackoff:
		// restAuthenticationBasicRetryRules.RestAuthenticationBasicRestRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use restAuthenticationBasicRetryRules.GetUnknownRaw() for raw JSON
}
```
