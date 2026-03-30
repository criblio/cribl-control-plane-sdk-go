# TargetConfigUnion1


## Supported Types

### TargetConfig1

```go
targetConfigUnion1 := components.CreateTargetConfigUnion1TargetConfig1(components.TargetConfig1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch targetConfigUnion1.Type {
	case components.TargetConfigUnion1TypeTargetConfig1:
		// targetConfigUnion1.TargetConfig1 is populated
}
```
