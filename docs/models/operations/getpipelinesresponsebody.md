# GetPipelinesResponseBody

List of Pipeline objects.


## Supported Types

### CountedPipeline

```go
getPipelinesResponseBody := operations.CreateGetPipelinesResponseBodyCountedPipeline(components.CountedPipeline{/* values here */})
```

### PaginatedPipeline

```go
getPipelinesResponseBody := operations.CreateGetPipelinesResponseBodyPaginatedPipeline(components.PaginatedPipeline{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getPipelinesResponseBody.Type {
	case operations.GetPipelinesResponseBodyTypeCountedPipeline:
		// getPipelinesResponseBody.CountedPipeline is populated
	case operations.GetPipelinesResponseBodyTypePaginatedPipeline:
		// getPipelinesResponseBody.PaginatedPipeline is populated
}
```
