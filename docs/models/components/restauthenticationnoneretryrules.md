# RestAuthenticationNoneRetryRules


## Supported Types

### RestAuthenticationNoneRestRetryRulesTypeNone

```go
restAuthenticationNoneRetryRules := components.CreateRestAuthenticationNoneRetryRulesNone(components.RestAuthenticationNoneRestRetryRulesTypeNone{/* values here */})
```

### RestAuthenticationNoneRestRetryRulesTypeStatic

```go
restAuthenticationNoneRetryRules := components.CreateRestAuthenticationNoneRetryRulesStatic(components.RestAuthenticationNoneRestRetryRulesTypeStatic{/* values here */})
```

### RestAuthenticationNoneRestRetryRulesTypeBackoff

```go
restAuthenticationNoneRetryRules := components.CreateRestAuthenticationNoneRetryRulesBackoff(components.RestAuthenticationNoneRestRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationNoneRetryRules.Type {
	case components.RestAuthenticationNoneRetryRulesTypeNone:
		// restAuthenticationNoneRetryRules.RestAuthenticationNoneRestRetryRulesTypeNone is populated
	case components.RestAuthenticationNoneRetryRulesTypeStatic:
		// restAuthenticationNoneRetryRules.RestAuthenticationNoneRestRetryRulesTypeStatic is populated
	case components.RestAuthenticationNoneRetryRulesTypeBackoff:
		// restAuthenticationNoneRetryRules.RestAuthenticationNoneRestRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use restAuthenticationNoneRetryRules.GetUnknownRaw() for raw JSON
}
```
