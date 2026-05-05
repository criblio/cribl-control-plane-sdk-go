# OutputResponseOutputGrafanaCloudUnion


## Supported Types

### OutputResponseOutputGrafanaCloud1

```go
outputResponseOutputGrafanaCloudUnion := components.CreateOutputResponseOutputGrafanaCloudUnionOutputResponseOutputGrafanaCloud1(components.OutputResponseOutputGrafanaCloud1{/* values here */})
```

### OutputResponseOutputGrafanaCloud2

```go
outputResponseOutputGrafanaCloudUnion := components.CreateOutputResponseOutputGrafanaCloudUnionOutputResponseOutputGrafanaCloud2(components.OutputResponseOutputGrafanaCloud2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputResponseOutputGrafanaCloudUnion.Type {
	case components.OutputResponseOutputGrafanaCloudUnionTypeOutputResponseOutputGrafanaCloud1:
		// outputResponseOutputGrafanaCloudUnion.OutputResponseOutputGrafanaCloud1 is populated
	case components.OutputResponseOutputGrafanaCloudUnionTypeOutputResponseOutputGrafanaCloud2:
		// outputResponseOutputGrafanaCloudUnion.OutputResponseOutputGrafanaCloud2 is populated
}
```
