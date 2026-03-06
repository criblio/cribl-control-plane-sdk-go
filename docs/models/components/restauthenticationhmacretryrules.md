# RestAuthenticationHmacRetryRules


## Supported Types

### RestAuthenticationHmacRestRetryRulesTypeNone

```go
restAuthenticationHmacRetryRules := components.CreateRestAuthenticationHmacRetryRulesNone(components.RestAuthenticationHmacRestRetryRulesTypeNone{/* values here */})
```

### RestAuthenticationHmacRestRetryRulesTypeStatic

```go
restAuthenticationHmacRetryRules := components.CreateRestAuthenticationHmacRetryRulesStatic(components.RestAuthenticationHmacRestRetryRulesTypeStatic{/* values here */})
```

### RestAuthenticationHmacRestRetryRulesTypeBackoff

```go
restAuthenticationHmacRetryRules := components.CreateRestAuthenticationHmacRetryRulesBackoff(components.RestAuthenticationHmacRestRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationHmacRetryRules.Type {
	case components.RestAuthenticationHmacRetryRulesTypeNone:
		// restAuthenticationHmacRetryRules.RestAuthenticationHmacRestRetryRulesTypeNone is populated
	case components.RestAuthenticationHmacRetryRulesTypeStatic:
		// restAuthenticationHmacRetryRules.RestAuthenticationHmacRestRetryRulesTypeStatic is populated
	case components.RestAuthenticationHmacRetryRulesTypeBackoff:
		// restAuthenticationHmacRetryRules.RestAuthenticationHmacRestRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use restAuthenticationHmacRetryRules.GetUnknownRaw() for raw JSON
}
```
