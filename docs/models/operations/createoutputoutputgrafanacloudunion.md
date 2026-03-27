# CreateOutputOutputGrafanaCloudUnion


## Supported Types

### CreateOutputOutputGrafanaCloudGrafanaCloud1

```go
createOutputOutputGrafanaCloudUnion := operations.CreateCreateOutputOutputGrafanaCloudUnionCreateOutputOutputGrafanaCloudGrafanaCloud1(operations.CreateOutputOutputGrafanaCloudGrafanaCloud1{/* values here */})
```

### CreateOutputOutputGrafanaCloudGrafanaCloud2

```go
createOutputOutputGrafanaCloudUnion := operations.CreateCreateOutputOutputGrafanaCloudUnionCreateOutputOutputGrafanaCloudGrafanaCloud2(operations.CreateOutputOutputGrafanaCloudGrafanaCloud2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createOutputOutputGrafanaCloudUnion.Type {
	case operations.CreateOutputOutputGrafanaCloudUnionTypeCreateOutputOutputGrafanaCloudGrafanaCloud1:
		// createOutputOutputGrafanaCloudUnion.CreateOutputOutputGrafanaCloudGrafanaCloud1 is populated
	case operations.CreateOutputOutputGrafanaCloudUnionTypeCreateOutputOutputGrafanaCloudGrafanaCloud2:
		// createOutputOutputGrafanaCloudUnion.CreateOutputOutputGrafanaCloudGrafanaCloud2 is populated
}
```
