# RestCollectMethodPostWithBodyRetryRules


## Supported Types

### RestCollectMethodPostWithBodyRestRetryRulesTypeNone

```go
restCollectMethodPostWithBodyRetryRules := components.CreateRestCollectMethodPostWithBodyRetryRulesNone(components.RestCollectMethodPostWithBodyRestRetryRulesTypeNone{/* values here */})
```

### RestCollectMethodPostWithBodyRestRetryRulesTypeStatic

```go
restCollectMethodPostWithBodyRetryRules := components.CreateRestCollectMethodPostWithBodyRetryRulesStatic(components.RestCollectMethodPostWithBodyRestRetryRulesTypeStatic{/* values here */})
```

### RestCollectMethodPostWithBodyRestRetryRulesTypeBackoff

```go
restCollectMethodPostWithBodyRetryRules := components.CreateRestCollectMethodPostWithBodyRetryRulesBackoff(components.RestCollectMethodPostWithBodyRestRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodPostWithBodyRetryRules.Type {
	case components.RestCollectMethodPostWithBodyRetryRulesTypeNone:
		// restCollectMethodPostWithBodyRetryRules.RestCollectMethodPostWithBodyRestRetryRulesTypeNone is populated
	case components.RestCollectMethodPostWithBodyRetryRulesTypeStatic:
		// restCollectMethodPostWithBodyRetryRules.RestCollectMethodPostWithBodyRestRetryRulesTypeStatic is populated
	case components.RestCollectMethodPostWithBodyRetryRulesTypeBackoff:
		// restCollectMethodPostWithBodyRetryRules.RestCollectMethodPostWithBodyRestRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use restCollectMethodPostWithBodyRetryRules.GetUnknownRaw() for raw JSON
}
```
