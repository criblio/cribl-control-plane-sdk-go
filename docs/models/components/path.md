# Path


## Supported Types

### 

```go
path := components.CreatePathStr(string{/* values here */})
```

### 

```go
path := components.CreatePathNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch path.Type {
	case components.PathTypeStr:
		// path.Str is populated
	case components.PathTypeNumber:
		// path.Number is populated
}
```
