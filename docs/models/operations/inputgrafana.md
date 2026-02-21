# InputGrafana


## Supported Types

### InputGrafanaGrafana1

```go
inputGrafana := operations.CreateInputGrafanaInputGrafanaGrafana1(operations.InputGrafanaGrafana1{/* values here */})
```

### InputGrafanaGrafana2

```go
inputGrafana := operations.CreateInputGrafanaInputGrafanaGrafana2(operations.InputGrafanaGrafana2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputGrafana.Type {
	case operations.InputGrafanaTypeInputGrafanaGrafana1:
		// inputGrafana.InputGrafanaGrafana1 is populated
	case operations.InputGrafanaTypeInputGrafanaGrafana2:
		// inputGrafana.InputGrafanaGrafana2 is populated
}
```
