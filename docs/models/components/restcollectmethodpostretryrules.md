# RestCollectMethodPostRetryRules


## Supported Types

### RestCollectMethodPostRestRetryRulesTypeNone

```go
restCollectMethodPostRetryRules := components.CreateRestCollectMethodPostRetryRulesNone(components.RestCollectMethodPostRestRetryRulesTypeNone{/* values here */})
```

### RestCollectMethodPostRestRetryRulesTypeStatic

```go
restCollectMethodPostRetryRules := components.CreateRestCollectMethodPostRetryRulesStatic(components.RestCollectMethodPostRestRetryRulesTypeStatic{/* values here */})
```

### RestCollectMethodPostRestRetryRulesTypeBackoff

```go
restCollectMethodPostRetryRules := components.CreateRestCollectMethodPostRetryRulesBackoff(components.RestCollectMethodPostRestRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodPostRetryRules.Type {
	case components.RestCollectMethodPostRetryRulesTypeNone:
		// restCollectMethodPostRetryRules.RestCollectMethodPostRestRetryRulesTypeNone is populated
	case components.RestCollectMethodPostRetryRulesTypeStatic:
		// restCollectMethodPostRetryRules.RestCollectMethodPostRestRetryRulesTypeStatic is populated
	case components.RestCollectMethodPostRetryRulesTypeBackoff:
		// restCollectMethodPostRetryRules.RestCollectMethodPostRestRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use restCollectMethodPostRetryRules.GetUnknownRaw() for raw JSON
}
```
