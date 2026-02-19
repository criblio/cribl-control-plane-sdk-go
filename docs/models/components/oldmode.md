# OldMode


## Supported Types

### 

```go
oldMode := components.CreateOldModeStr(string{/* values here */})
```

### 

```go
oldMode := components.CreateOldModeArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch oldMode.Type {
	case components.OldModeTypeStr:
		// oldMode.Str is populated
	case components.OldModeTypeArrayOfStr:
		// oldMode.ArrayOfStr is populated
}
```
