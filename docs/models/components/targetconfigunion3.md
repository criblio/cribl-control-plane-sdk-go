# TargetConfigUnion3


## Supported Types

### TargetConfig3

```go
targetConfigUnion3 := components.CreateTargetConfigUnion3TargetConfig3(components.TargetConfig3{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch targetConfigUnion3.Type {
	case components.TargetConfigUnion3TypeTargetConfig3:
		// targetConfigUnion3.TargetConfig3 is populated
}
```
