# CreateInputSystemByPackInputGrafanaUnion


## Supported Types

### CreateInputSystemByPackInputGrafanaGrafana1

```go
createInputSystemByPackInputGrafanaUnion := operations.CreateCreateInputSystemByPackInputGrafanaUnionCreateInputSystemByPackInputGrafanaGrafana1(operations.CreateInputSystemByPackInputGrafanaGrafana1{/* values here */})
```

### CreateInputSystemByPackInputGrafanaGrafana2

```go
createInputSystemByPackInputGrafanaUnion := operations.CreateCreateInputSystemByPackInputGrafanaUnionCreateInputSystemByPackInputGrafanaGrafana2(operations.CreateInputSystemByPackInputGrafanaGrafana2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createInputSystemByPackInputGrafanaUnion.Type {
	case operations.CreateInputSystemByPackInputGrafanaUnionTypeCreateInputSystemByPackInputGrafanaGrafana1:
		// createInputSystemByPackInputGrafanaUnion.CreateInputSystemByPackInputGrafanaGrafana1 is populated
	case operations.CreateInputSystemByPackInputGrafanaUnionTypeCreateInputSystemByPackInputGrafanaGrafana2:
		// createInputSystemByPackInputGrafanaUnion.CreateInputSystemByPackInputGrafanaGrafana2 is populated
}
```
