# OutputGrafanaCloud


## Supported Types

### OutputGrafanaCloudGrafanaCloud1

```go
outputGrafanaCloud := operations.CreateOutputGrafanaCloudOutputGrafanaCloudGrafanaCloud1(operations.OutputGrafanaCloudGrafanaCloud1{/* values here */})
```

### OutputGrafanaCloudGrafanaCloud2

```go
outputGrafanaCloud := operations.CreateOutputGrafanaCloudOutputGrafanaCloudGrafanaCloud2(operations.OutputGrafanaCloudGrafanaCloud2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch outputGrafanaCloud.Type {
	case operations.OutputGrafanaCloudTypeOutputGrafanaCloudGrafanaCloud1:
		// outputGrafanaCloud.OutputGrafanaCloudGrafanaCloud1 is populated
	case operations.OutputGrafanaCloudTypeOutputGrafanaCloudGrafanaCloud2:
		// outputGrafanaCloud.OutputGrafanaCloudGrafanaCloud2 is populated
}
```
