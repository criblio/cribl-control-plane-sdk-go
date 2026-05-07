# InputResponseInputGrafanaUnion


## Supported Types

### InputResponseInputGrafanaGrafana1

```go
inputResponseInputGrafanaUnion := components.CreateInputResponseInputGrafanaUnionInputResponseInputGrafanaGrafana1(components.InputResponseInputGrafanaGrafana1{/* values here */})
```

### InputResponseInputGrafanaGrafana2

```go
inputResponseInputGrafanaUnion := components.CreateInputResponseInputGrafanaUnionInputResponseInputGrafanaGrafana2(components.InputResponseInputGrafanaGrafana2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputResponseInputGrafanaUnion.Type {
	case components.InputResponseInputGrafanaUnionTypeInputResponseInputGrafanaGrafana1:
		// inputResponseInputGrafanaUnion.InputResponseInputGrafanaGrafana1 is populated
	case components.InputResponseInputGrafanaUnionTypeInputResponseInputGrafanaGrafana2:
		// inputResponseInputGrafanaUnion.InputResponseInputGrafanaGrafana2 is populated
}
```
