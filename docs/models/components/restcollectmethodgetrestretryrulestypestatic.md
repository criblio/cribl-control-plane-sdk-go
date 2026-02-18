# RestCollectMethodGetRestRetryRulesTypeStatic


## Supported Types

### RestCollectMethodGetRestRetryRulesTypeStaticEnableHeaderFalse

```go
restCollectMethodGetRestRetryRulesTypeStatic := components.CreateRestCollectMethodGetRestRetryRulesTypeStaticRestCollectMethodGetRestRetryRulesTypeStaticEnableHeaderFalse(components.RestCollectMethodGetRestRetryRulesTypeStaticEnableHeaderFalse{/* values here */})
```

### RestCollectMethodGetRestRetryRulesTypeStaticEnableHeaderTrue

```go
restCollectMethodGetRestRetryRulesTypeStatic := components.CreateRestCollectMethodGetRestRetryRulesTypeStaticRestCollectMethodGetRestRetryRulesTypeStaticEnableHeaderTrue(components.RestCollectMethodGetRestRetryRulesTypeStaticEnableHeaderTrue{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch restCollectMethodGetRestRetryRulesTypeStatic.Type {
	case components.RestCollectMethodGetRestRetryRulesTypeStaticTypeRestCollectMethodGetRestRetryRulesTypeStaticEnableHeaderFalse:
		// restCollectMethodGetRestRetryRulesTypeStatic.RestCollectMethodGetRestRetryRulesTypeStaticEnableHeaderFalse is populated
	case components.RestCollectMethodGetRestRetryRulesTypeStaticTypeRestCollectMethodGetRestRetryRulesTypeStaticEnableHeaderTrue:
		// restCollectMethodGetRestRetryRulesTypeStatic.RestCollectMethodGetRestRetryRulesTypeStaticEnableHeaderTrue is populated
}
```
