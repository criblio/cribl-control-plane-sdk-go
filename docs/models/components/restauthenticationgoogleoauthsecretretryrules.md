# RestAuthenticationGoogleOauthSecretRetryRules


## Supported Types

### RestAuthenticationGoogleOauthSecretRestRetryRulesTypeNone

```go
restAuthenticationGoogleOauthSecretRetryRules := components.CreateRestAuthenticationGoogleOauthSecretRetryRulesNone(components.RestAuthenticationGoogleOauthSecretRestRetryRulesTypeNone{/* values here */})
```

### RestAuthenticationGoogleOauthSecretRestRetryRulesTypeStatic

```go
restAuthenticationGoogleOauthSecretRetryRules := components.CreateRestAuthenticationGoogleOauthSecretRetryRulesStatic(components.RestAuthenticationGoogleOauthSecretRestRetryRulesTypeStatic{/* values here */})
```

### RestAuthenticationGoogleOauthSecretRestRetryRulesTypeBackoff

```go
restAuthenticationGoogleOauthSecretRetryRules := components.CreateRestAuthenticationGoogleOauthSecretRetryRulesBackoff(components.RestAuthenticationGoogleOauthSecretRestRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationGoogleOauthSecretRetryRules.Type {
	case components.RestAuthenticationGoogleOauthSecretRetryRulesTypeNone:
		// restAuthenticationGoogleOauthSecretRetryRules.RestAuthenticationGoogleOauthSecretRestRetryRulesTypeNone is populated
	case components.RestAuthenticationGoogleOauthSecretRetryRulesTypeStatic:
		// restAuthenticationGoogleOauthSecretRetryRules.RestAuthenticationGoogleOauthSecretRestRetryRulesTypeStatic is populated
	case components.RestAuthenticationGoogleOauthSecretRetryRulesTypeBackoff:
		// restAuthenticationGoogleOauthSecretRetryRules.RestAuthenticationGoogleOauthSecretRestRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use restAuthenticationGoogleOauthSecretRetryRules.GetUnknownRaw() for raw JSON
}
```
