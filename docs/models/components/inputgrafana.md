# InputGrafana


## Supported Types

### InputGrafanaGrafana1

```go
inputGrafana := components.CreateInputGrafanaInputGrafanaGrafana1(components.InputGrafanaGrafana1{/* values here */})
```

### InputGrafanaGrafana2

```go
inputGrafana := components.CreateInputGrafanaInputGrafanaGrafana2(components.InputGrafanaGrafana2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch inputGrafana.Type {
	case components.InputGrafanaTypeInputGrafanaGrafana1:
		// inputGrafana.InputGrafanaGrafana1 is populated
	case components.InputGrafanaTypeInputGrafanaGrafana2:
		// inputGrafana.InputGrafanaGrafana2 is populated
}
```
