# CreateInputInputGrafanaUnion


## Supported Types

### CreateInputInputGrafanaGrafana1

```go
createInputInputGrafanaUnion := operations.CreateCreateInputInputGrafanaUnionCreateInputInputGrafanaGrafana1(operations.CreateInputInputGrafanaGrafana1{/* values here */})
```

### CreateInputInputGrafanaGrafana2

```go
createInputInputGrafanaUnion := operations.CreateCreateInputInputGrafanaUnionCreateInputInputGrafanaGrafana2(operations.CreateInputInputGrafanaGrafana2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createInputInputGrafanaUnion.Type {
	case operations.CreateInputInputGrafanaUnionTypeCreateInputInputGrafanaGrafana1:
		// createInputInputGrafanaUnion.CreateInputInputGrafanaGrafana1 is populated
	case operations.CreateInputInputGrafanaUnionTypeCreateInputInputGrafanaGrafana2:
		// createInputInputGrafanaUnion.CreateInputInputGrafanaGrafana2 is populated
}
```
