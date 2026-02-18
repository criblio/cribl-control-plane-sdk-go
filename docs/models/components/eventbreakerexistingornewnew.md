# EventBreakerExistingOrNewNew


## Supported Types

### EventBreakerExistingOrNewNewRuleTypeRegex

```go
eventBreakerExistingOrNewNew := components.CreateEventBreakerExistingOrNewNewRegex(components.EventBreakerExistingOrNewNewRuleTypeRegex{/* values here */})
```

### EventBreakerExistingOrNewNewRuleTypeJSON

```go
eventBreakerExistingOrNewNew := components.CreateEventBreakerExistingOrNewNewJSON(components.EventBreakerExistingOrNewNewRuleTypeJSON{/* values here */})
```

### EventBreakerExistingOrNewNewRuleTypeJSONArray

```go
eventBreakerExistingOrNewNew := components.CreateEventBreakerExistingOrNewNewJSONArray(components.EventBreakerExistingOrNewNewRuleTypeJSONArray{/* values here */})
```

### EventBreakerExistingOrNewNewRuleTypeHeader

```go
eventBreakerExistingOrNewNew := components.CreateEventBreakerExistingOrNewNewHeader(components.EventBreakerExistingOrNewNewRuleTypeHeader{/* values here */})
```

### EventBreakerExistingOrNewNewRuleTypeCsv

```go
eventBreakerExistingOrNewNew := components.CreateEventBreakerExistingOrNewNewCsv(components.EventBreakerExistingOrNewNewRuleTypeCsv{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventBreakerExistingOrNewNew.Type {
	case components.EventBreakerExistingOrNewNewTypeRegex:
		// eventBreakerExistingOrNewNew.EventBreakerExistingOrNewNewRuleTypeRegex is populated
	case components.EventBreakerExistingOrNewNewTypeJSON:
		// eventBreakerExistingOrNewNew.EventBreakerExistingOrNewNewRuleTypeJSON is populated
	case components.EventBreakerExistingOrNewNewTypeJSONArray:
		// eventBreakerExistingOrNewNew.EventBreakerExistingOrNewNewRuleTypeJSONArray is populated
	case components.EventBreakerExistingOrNewNewTypeHeader:
		// eventBreakerExistingOrNewNew.EventBreakerExistingOrNewNewRuleTypeHeader is populated
	case components.EventBreakerExistingOrNewNewTypeCsv:
		// eventBreakerExistingOrNewNew.EventBreakerExistingOrNewNewRuleTypeCsv is populated
	default:
		// Unknown type - use eventBreakerExistingOrNewNew.GetUnknownRaw() for raw JSON
}
```
