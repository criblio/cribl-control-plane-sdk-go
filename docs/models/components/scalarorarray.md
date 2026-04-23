# ScalarOrArray


## Supported Types

### 

```go
scalarOrArray := components.CreateScalarOrArrayStr(string{/* values here */})
```

### 

```go
scalarOrArray := components.CreateScalarOrArrayArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch scalarOrArray.Type {
	case components.ScalarOrArrayTypeStr:
		// scalarOrArray.Str is populated
	case components.ScalarOrArrayTypeArrayOfStr:
		// scalarOrArray.ArrayOfStr is populated
}
```
