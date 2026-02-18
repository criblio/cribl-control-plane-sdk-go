# DiffLine


## Supported Types

### DiffLineDelete

```go
diffLine := components.CreateDiffLineDelete(components.DiffLineDelete{/* values here */})
```

### DiffLineInsert

```go
diffLine := components.CreateDiffLineInsert(components.DiffLineInsert{/* values here */})
```

### DiffLineContext

```go
diffLine := components.CreateDiffLineContext(components.DiffLineContext{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch diffLine.Type {
	case components.DiffLineTypeDelete:
		// diffLine.DiffLineDelete is populated
	case components.DiffLineTypeInsert:
		// diffLine.DiffLineInsert is populated
	case components.DiffLineTypeContext:
		// diffLine.DiffLineContext is populated
	default:
		// Unknown type - use diffLine.GetUnknownRaw() for raw JSON
}
```
