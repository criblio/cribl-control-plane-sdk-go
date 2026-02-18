# CreateOutputSystemByPackOutputGrafanaCloudUnion


## Supported Types

### CreateOutputSystemByPackOutputGrafanaCloudGrafanaCloud1

```go
createOutputSystemByPackOutputGrafanaCloudUnion := operations.CreateCreateOutputSystemByPackOutputGrafanaCloudUnionCreateOutputSystemByPackOutputGrafanaCloudGrafanaCloud1(operations.CreateOutputSystemByPackOutputGrafanaCloudGrafanaCloud1{/* values here */})
```

### CreateOutputSystemByPackOutputGrafanaCloudGrafanaCloud2

```go
createOutputSystemByPackOutputGrafanaCloudUnion := operations.CreateCreateOutputSystemByPackOutputGrafanaCloudUnionCreateOutputSystemByPackOutputGrafanaCloudGrafanaCloud2(operations.CreateOutputSystemByPackOutputGrafanaCloudGrafanaCloud2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createOutputSystemByPackOutputGrafanaCloudUnion.Type {
	case operations.CreateOutputSystemByPackOutputGrafanaCloudUnionTypeCreateOutputSystemByPackOutputGrafanaCloudGrafanaCloud1:
		// createOutputSystemByPackOutputGrafanaCloudUnion.CreateOutputSystemByPackOutputGrafanaCloudGrafanaCloud1 is populated
	case operations.CreateOutputSystemByPackOutputGrafanaCloudUnionTypeCreateOutputSystemByPackOutputGrafanaCloudGrafanaCloud2:
		// createOutputSystemByPackOutputGrafanaCloudUnion.CreateOutputSystemByPackOutputGrafanaCloudGrafanaCloud2 is populated
}
```
