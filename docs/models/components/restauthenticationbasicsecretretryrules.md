# RestAuthenticationBasicSecretRetryRules


## Supported Types

### RestAuthenticationBasicSecretRestRetryRulesTypeNone

```go
restAuthenticationBasicSecretRetryRules := components.CreateRestAuthenticationBasicSecretRetryRulesNone(components.RestAuthenticationBasicSecretRestRetryRulesTypeNone{/* values here */})
```

### RestAuthenticationBasicSecretRestRetryRulesTypeStatic

```go
restAuthenticationBasicSecretRetryRules := components.CreateRestAuthenticationBasicSecretRetryRulesStatic(components.RestAuthenticationBasicSecretRestRetryRulesTypeStatic{/* values here */})
```

### RestAuthenticationBasicSecretRestRetryRulesTypeBackoff

```go
restAuthenticationBasicSecretRetryRules := components.CreateRestAuthenticationBasicSecretRetryRulesBackoff(components.RestAuthenticationBasicSecretRestRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationBasicSecretRetryRules.Type {
	case components.RestAuthenticationBasicSecretRetryRulesTypeNone:
		// restAuthenticationBasicSecretRetryRules.RestAuthenticationBasicSecretRestRetryRulesTypeNone is populated
	case components.RestAuthenticationBasicSecretRetryRulesTypeStatic:
		// restAuthenticationBasicSecretRetryRules.RestAuthenticationBasicSecretRestRetryRulesTypeStatic is populated
	case components.RestAuthenticationBasicSecretRetryRulesTypeBackoff:
		// restAuthenticationBasicSecretRetryRules.RestAuthenticationBasicSecretRestRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use restAuthenticationBasicSecretRetryRules.GetUnknownRaw() for raw JSON
}
```
