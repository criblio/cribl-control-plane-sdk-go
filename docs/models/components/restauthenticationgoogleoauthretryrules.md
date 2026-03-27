# RestAuthenticationGoogleOauthRetryRules


## Supported Types

### RestAuthenticationGoogleOauthRestRetryRulesTypeNone

```go
restAuthenticationGoogleOauthRetryRules := components.CreateRestAuthenticationGoogleOauthRetryRulesNone(components.RestAuthenticationGoogleOauthRestRetryRulesTypeNone{/* values here */})
```

### RestAuthenticationGoogleOauthRestRetryRulesTypeStatic

```go
restAuthenticationGoogleOauthRetryRules := components.CreateRestAuthenticationGoogleOauthRetryRulesStatic(components.RestAuthenticationGoogleOauthRestRetryRulesTypeStatic{/* values here */})
```

### RestAuthenticationGoogleOauthRestRetryRulesTypeBackoff

```go
restAuthenticationGoogleOauthRetryRules := components.CreateRestAuthenticationGoogleOauthRetryRulesBackoff(components.RestAuthenticationGoogleOauthRestRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationGoogleOauthRetryRules.Type {
	case components.RestAuthenticationGoogleOauthRetryRulesTypeNone:
		// restAuthenticationGoogleOauthRetryRules.RestAuthenticationGoogleOauthRestRetryRulesTypeNone is populated
	case components.RestAuthenticationGoogleOauthRetryRulesTypeStatic:
		// restAuthenticationGoogleOauthRetryRules.RestAuthenticationGoogleOauthRestRetryRulesTypeStatic is populated
	case components.RestAuthenticationGoogleOauthRetryRulesTypeBackoff:
		// restAuthenticationGoogleOauthRetryRules.RestAuthenticationGoogleOauthRestRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use restAuthenticationGoogleOauthRetryRules.GetUnknownRaw() for raw JSON
}
```
