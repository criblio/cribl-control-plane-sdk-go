# RestAuthenticationOauthRetryRules


## Supported Types

### RestAuthenticationOauthRestRetryRulesTypeNone

```go
restAuthenticationOauthRetryRules := components.CreateRestAuthenticationOauthRetryRulesNone(components.RestAuthenticationOauthRestRetryRulesTypeNone{/* values here */})
```

### RestAuthenticationOauthRestRetryRulesTypeStatic

```go
restAuthenticationOauthRetryRules := components.CreateRestAuthenticationOauthRetryRulesStatic(components.RestAuthenticationOauthRestRetryRulesTypeStatic{/* values here */})
```

### RestAuthenticationOauthRestRetryRulesTypeBackoff

```go
restAuthenticationOauthRetryRules := components.CreateRestAuthenticationOauthRetryRulesBackoff(components.RestAuthenticationOauthRestRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationOauthRetryRules.Type {
	case components.RestAuthenticationOauthRetryRulesTypeNone:
		// restAuthenticationOauthRetryRules.RestAuthenticationOauthRestRetryRulesTypeNone is populated
	case components.RestAuthenticationOauthRetryRulesTypeStatic:
		// restAuthenticationOauthRetryRules.RestAuthenticationOauthRestRetryRulesTypeStatic is populated
	case components.RestAuthenticationOauthRetryRulesTypeBackoff:
		// restAuthenticationOauthRetryRules.RestAuthenticationOauthRestRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use restAuthenticationOauthRetryRules.GetUnknownRaw() for raw JSON
}
```
