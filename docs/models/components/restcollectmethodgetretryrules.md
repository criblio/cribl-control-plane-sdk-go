# RestCollectMethodGetRetryRules


## Supported Types

### RestCollectMethodGetRestRetryRulesTypeNone

```go
restCollectMethodGetRetryRules := components.CreateRestCollectMethodGetRetryRulesNone(components.RestCollectMethodGetRestRetryRulesTypeNone{/* values here */})
```

### RestCollectMethodGetRestRetryRulesTypeStatic

```go
restCollectMethodGetRetryRules := components.CreateRestCollectMethodGetRetryRulesStatic(components.RestCollectMethodGetRestRetryRulesTypeStatic{/* values here */})
```

### RestCollectMethodGetRestRetryRulesTypeBackoff

```go
restCollectMethodGetRetryRules := components.CreateRestCollectMethodGetRetryRulesBackoff(components.RestCollectMethodGetRestRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodGetRetryRules.Type {
	case components.RestCollectMethodGetRetryRulesTypeNone:
		// restCollectMethodGetRetryRules.RestCollectMethodGetRestRetryRulesTypeNone is populated
	case components.RestCollectMethodGetRetryRulesTypeStatic:
		// restCollectMethodGetRetryRules.RestCollectMethodGetRestRetryRulesTypeStatic is populated
	case components.RestCollectMethodGetRetryRulesTypeBackoff:
		// restCollectMethodGetRetryRules.RestCollectMethodGetRestRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use restCollectMethodGetRetryRules.GetUnknownRaw() for raw JSON
}
```
