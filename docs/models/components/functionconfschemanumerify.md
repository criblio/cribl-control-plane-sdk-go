# FunctionConfSchemaNumerify


## Supported Types

### NumerifyFormatFix

```go
functionConfSchemaNumerify := components.CreateFunctionConfSchemaNumerifyFix(components.NumerifyFormatFix{/* values here */})
```

### NumerifyFormatNone

```go
functionConfSchemaNumerify := components.CreateFunctionConfSchemaNumerifyNone(components.NumerifyFormatNone{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch functionConfSchemaNumerify.Type {
	case components.FunctionConfSchemaNumerifyTypeFix:
		// functionConfSchemaNumerify.NumerifyFormatFix is populated
	case components.FunctionConfSchemaNumerifyTypeNone:
		// functionConfSchemaNumerify.NumerifyFormatNone is populated
	default:
		// Unknown type - use functionConfSchemaNumerify.GetUnknownRaw() for raw JSON
}
```
