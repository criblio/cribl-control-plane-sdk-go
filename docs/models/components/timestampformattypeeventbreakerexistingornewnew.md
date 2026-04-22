# TimestampFormatTypeEventBreakerExistingOrNewNew


## Supported Types

### EventBreakerExistingOrNewNewTimestampTypeAuto

```go
timestampFormatTypeEventBreakerExistingOrNewNew := components.CreateTimestampFormatTypeEventBreakerExistingOrNewNewAuto(components.EventBreakerExistingOrNewNewTimestampTypeAuto{/* values here */})
```

### EventBreakerExistingOrNewNewTimestampTypeFormat

```go
timestampFormatTypeEventBreakerExistingOrNewNew := components.CreateTimestampFormatTypeEventBreakerExistingOrNewNewFormat(components.EventBreakerExistingOrNewNewTimestampTypeFormat{/* values here */})
```

### EventBreakerExistingOrNewNewTimestampTypeCurrent

```go
timestampFormatTypeEventBreakerExistingOrNewNew := components.CreateTimestampFormatTypeEventBreakerExistingOrNewNewCurrent(components.EventBreakerExistingOrNewNewTimestampTypeCurrent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch timestampFormatTypeEventBreakerExistingOrNewNew.Type {
	case components.TimestampFormatTypeEventBreakerExistingOrNewNewTypeAuto:
		// timestampFormatTypeEventBreakerExistingOrNewNew.EventBreakerExistingOrNewNewTimestampTypeAuto is populated
	case components.TimestampFormatTypeEventBreakerExistingOrNewNewTypeFormat:
		// timestampFormatTypeEventBreakerExistingOrNewNew.EventBreakerExistingOrNewNewTimestampTypeFormat is populated
	case components.TimestampFormatTypeEventBreakerExistingOrNewNewTypeCurrent:
		// timestampFormatTypeEventBreakerExistingOrNewNew.EventBreakerExistingOrNewNewTimestampTypeCurrent is populated
	default:
		// Unknown type - use timestampFormatTypeEventBreakerExistingOrNewNew.GetUnknownRaw() for raw JSON
}
```
