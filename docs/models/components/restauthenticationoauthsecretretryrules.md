# RestAuthenticationOauthSecretRetryRules


## Supported Types

### RestAuthenticationOauthSecretRestRetryRulesTypeNone

```go
restAuthenticationOauthSecretRetryRules := components.CreateRestAuthenticationOauthSecretRetryRulesNone(components.RestAuthenticationOauthSecretRestRetryRulesTypeNone{/* values here */})
```

### RestAuthenticationOauthSecretRestRetryRulesTypeStatic

```go
restAuthenticationOauthSecretRetryRules := components.CreateRestAuthenticationOauthSecretRetryRulesStatic(components.RestAuthenticationOauthSecretRestRetryRulesTypeStatic{/* values here */})
```

### RestAuthenticationOauthSecretRestRetryRulesTypeBackoff

```go
restAuthenticationOauthSecretRetryRules := components.CreateRestAuthenticationOauthSecretRetryRulesBackoff(components.RestAuthenticationOauthSecretRestRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationOauthSecretRetryRules.Type {
	case components.RestAuthenticationOauthSecretRetryRulesTypeNone:
		// restAuthenticationOauthSecretRetryRules.RestAuthenticationOauthSecretRestRetryRulesTypeNone is populated
	case components.RestAuthenticationOauthSecretRetryRulesTypeStatic:
		// restAuthenticationOauthSecretRetryRules.RestAuthenticationOauthSecretRestRetryRulesTypeStatic is populated
	case components.RestAuthenticationOauthSecretRetryRulesTypeBackoff:
		// restAuthenticationOauthSecretRetryRules.RestAuthenticationOauthSecretRestRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use restAuthenticationOauthSecretRetryRules.GetUnknownRaw() for raw JSON
}
```
