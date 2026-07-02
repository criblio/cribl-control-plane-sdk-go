# GetPipelinesByPackResponseBody

List of Pipeline objects.


## Supported Types

### CountedPipeline

```go
getPipelinesByPackResponseBody := operations.CreateGetPipelinesByPackResponseBodyCountedPipeline(components.CountedPipeline{/* values here */})
```

### PaginatedPipeline

```go
getPipelinesByPackResponseBody := operations.CreateGetPipelinesByPackResponseBodyPaginatedPipeline(components.PaginatedPipeline{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getPipelinesByPackResponseBody.Type {
	case operations.GetPipelinesByPackResponseBodyTypeCountedPipeline:
		// getPipelinesByPackResponseBody.CountedPipeline is populated
	case operations.GetPipelinesByPackResponseBodyTypePaginatedPipeline:
		// getPipelinesByPackResponseBody.PaginatedPipeline is populated
}
```
