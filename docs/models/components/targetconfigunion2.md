# TargetConfigUnion2


## Supported Types

### TargetConfig2

```go
targetConfigUnion2 := components.CreateTargetConfigUnion2TargetConfig2(components.TargetConfig2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch targetConfigUnion2.Type {
	case components.TargetConfigUnion2TypeTargetConfig2:
		// targetConfigUnion2.TargetConfig2 is populated
}
```
