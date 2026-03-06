# LookupDbLookupFalse


## Supported Types

### LookupDbLookupFalseMatchModeCidr

```go
lookupDbLookupFalse := components.CreateLookupDbLookupFalseCidr(components.LookupDbLookupFalseMatchModeCidr{/* values here */})
```

### LookupDbLookupFalseMatchModeExact

```go
lookupDbLookupFalse := components.CreateLookupDbLookupFalseExact(components.LookupDbLookupFalseMatchModeExact{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch lookupDbLookupFalse.Type {
	case components.LookupDbLookupFalseTypeCidr:
		// lookupDbLookupFalse.LookupDbLookupFalseMatchModeCidr is populated
	case components.LookupDbLookupFalseTypeExact:
		// lookupDbLookupFalse.LookupDbLookupFalseMatchModeExact is populated
	default:
		// Unknown type - use lookupDbLookupFalse.GetUnknownRaw() for raw JSON
}
```
