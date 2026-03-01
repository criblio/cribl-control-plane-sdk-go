# RestCollectMethodOtherRetryRules


## Supported Types

### RestCollectMethodOtherRestRetryRulesTypeNone

```go
restCollectMethodOtherRetryRules := components.CreateRestCollectMethodOtherRetryRulesNone(components.RestCollectMethodOtherRestRetryRulesTypeNone{/* values here */})
```

### RestCollectMethodOtherRestRetryRulesTypeStatic

```go
restCollectMethodOtherRetryRules := components.CreateRestCollectMethodOtherRetryRulesStatic(components.RestCollectMethodOtherRestRetryRulesTypeStatic{/* values here */})
```

### RestCollectMethodOtherRestRetryRulesTypeBackoff

```go
restCollectMethodOtherRetryRules := components.CreateRestCollectMethodOtherRetryRulesBackoff(components.RestCollectMethodOtherRestRetryRulesTypeBackoff{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodOtherRetryRules.Type {
	case components.RestCollectMethodOtherRetryRulesTypeNone:
		// restCollectMethodOtherRetryRules.RestCollectMethodOtherRestRetryRulesTypeNone is populated
	case components.RestCollectMethodOtherRetryRulesTypeStatic:
		// restCollectMethodOtherRetryRules.RestCollectMethodOtherRestRetryRulesTypeStatic is populated
	case components.RestCollectMethodOtherRetryRulesTypeBackoff:
		// restCollectMethodOtherRetryRules.RestCollectMethodOtherRestRetryRulesTypeBackoff is populated
	default:
		// Unknown type - use restCollectMethodOtherRetryRules.GetUnknownRaw() for raw JSON
}
```
