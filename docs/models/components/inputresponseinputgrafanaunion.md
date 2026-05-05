# InputResponseInputGrafanaUnion


## Supported Types

### InputResponseInputGrafana1

```go
inputResponseInputGrafanaUnion := components.CreateInputResponseInputGrafanaUnionInputResponseInputGrafana1(components.InputResponseInputGrafana1{/* values here */})
```

### InputResponseInputGrafana2

```go
inputResponseInputGrafanaUnion := components.CreateInputResponseInputGrafanaUnionInputResponseInputGrafana2(components.InputResponseInputGrafana2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputResponseInputGrafanaUnion.Type {
	case components.InputResponseInputGrafanaUnionTypeInputResponseInputGrafana1:
		// inputResponseInputGrafanaUnion.InputResponseInputGrafana1 is populated
	case components.InputResponseInputGrafanaUnionTypeInputResponseInputGrafana2:
		// inputResponseInputGrafanaUnion.InputResponseInputGrafana2 is populated
}
```
