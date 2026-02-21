# RestAuthenticationLoginSecretRetryRules


## Supported Types

### RestAuthenticationLoginSecretRestRetryRulesTypeNone

```go
restAuthenticationLoginSecretRetryRules := components.CreateRestAuthenticationLoginSecretRetryRulesNone(components.RestAuthenticationLoginSecretRestRetryRulesTypeNone{/* values here */})
```

### RestAuthenticationLoginSecretRestRetryRulesTypeStatic

```go
restAuthenticationLoginSecretRetryRules := components.CreateRestAuthenticationLoginSecretRetryRulesStatic(components.RestAuthenticationLoginSecretRestRetryRulesTypeStatic{/* values here */})
```

### RestAuthenticationLoginSecretRestRetryRulesTypeBackoff

```go
restAuthenticationLoginSecretRetryRules := components.CreateRestAuthenticationLoginSecretRetryRulesBackoff(components.RestAuthenticationLoginSecretRestRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationLoginSecretRetryRules.Type {
	case components.RestAuthenticationLoginSecretRetryRulesTypeNone:
		// restAuthenticationLoginSecretRetryRules.RestAuthenticationLoginSecretRestRetryRulesTypeNone is populated
	case components.RestAuthenticationLoginSecretRetryRulesTypeStatic:
		// restAuthenticationLoginSecretRetryRules.RestAuthenticationLoginSecretRestRetryRulesTypeStatic is populated
	case components.RestAuthenticationLoginSecretRetryRulesTypeBackoff:
		// restAuthenticationLoginSecretRetryRules.RestAuthenticationLoginSecretRestRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use restAuthenticationLoginSecretRetryRules.GetUnknownRaw() for raw JSON
}
```
