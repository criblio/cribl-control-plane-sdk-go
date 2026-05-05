# InputGrafanaInputUnion


## Supported Types

### InputGrafanaGrafanaInput1

```go
inputGrafanaInputUnion := components.CreateInputGrafanaInputUnionInputGrafanaGrafanaInput1(components.InputGrafanaGrafanaInput1{/* values here */})
```

### InputGrafanaGrafanaInput2

```go
inputGrafanaInputUnion := components.CreateInputGrafanaInputUnionInputGrafanaGrafanaInput2(components.InputGrafanaGrafanaInput2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputGrafanaInputUnion.Type {
	case components.InputGrafanaInputUnionTypeInputGrafanaGrafanaInput1:
		// inputGrafanaInputUnion.InputGrafanaGrafanaInput1 is populated
	case components.InputGrafanaInputUnionTypeInputGrafanaGrafanaInput2:
		// inputGrafanaInputUnion.InputGrafanaGrafanaInput2 is populated
}
```
