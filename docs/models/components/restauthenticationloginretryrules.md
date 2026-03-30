# RestAuthenticationLoginRetryRules


## Supported Types

### RestAuthenticationLoginRestRetryRulesTypeNone

```go
restAuthenticationLoginRetryRules := components.CreateRestAuthenticationLoginRetryRulesNone(components.RestAuthenticationLoginRestRetryRulesTypeNone{/* values here */})
```

### RestAuthenticationLoginRestRetryRulesTypeStatic

```go
restAuthenticationLoginRetryRules := components.CreateRestAuthenticationLoginRetryRulesStatic(components.RestAuthenticationLoginRestRetryRulesTypeStatic{/* values here */})
```

### RestAuthenticationLoginRestRetryRulesTypeBackoff

```go
restAuthenticationLoginRetryRules := components.CreateRestAuthenticationLoginRetryRulesBackoff(components.RestAuthenticationLoginRestRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restAuthenticationLoginRetryRules.Type {
	case components.RestAuthenticationLoginRetryRulesTypeNone:
		// restAuthenticationLoginRetryRules.RestAuthenticationLoginRestRetryRulesTypeNone is populated
	case components.RestAuthenticationLoginRetryRulesTypeStatic:
		// restAuthenticationLoginRetryRules.RestAuthenticationLoginRestRetryRulesTypeStatic is populated
	case components.RestAuthenticationLoginRetryRulesTypeBackoff:
		// restAuthenticationLoginRetryRules.RestAuthenticationLoginRestRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use restAuthenticationLoginRetryRules.GetUnknownRaw() for raw JSON
}
```
