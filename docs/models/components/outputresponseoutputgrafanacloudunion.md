# OutputResponseOutputGrafanaCloudUnion


## Supported Types

### OutputResponseOutputGrafanaCloudGrafanaCloud1

```go
outputResponseOutputGrafanaCloudUnion := components.CreateOutputResponseOutputGrafanaCloudUnionOutputResponseOutputGrafanaCloudGrafanaCloud1(components.OutputResponseOutputGrafanaCloudGrafanaCloud1{/* values here */})
```

### OutputResponseOutputGrafanaCloudGrafanaCloud2

```go
outputResponseOutputGrafanaCloudUnion := components.CreateOutputResponseOutputGrafanaCloudUnionOutputResponseOutputGrafanaCloudGrafanaCloud2(components.OutputResponseOutputGrafanaCloudGrafanaCloud2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputResponseOutputGrafanaCloudUnion.Type {
	case components.OutputResponseOutputGrafanaCloudUnionTypeOutputResponseOutputGrafanaCloudGrafanaCloud1:
		// outputResponseOutputGrafanaCloudUnion.OutputResponseOutputGrafanaCloudGrafanaCloud1 is populated
	case components.OutputResponseOutputGrafanaCloudUnionTypeOutputResponseOutputGrafanaCloudGrafanaCloud2:
		// outputResponseOutputGrafanaCloudUnion.OutputResponseOutputGrafanaCloudGrafanaCloud2 is populated
}
```
